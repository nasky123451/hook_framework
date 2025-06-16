// File: internal/hooks/context_manager.go
package hooks

import (
	"sync"
	"time"
	"unsafe"
)

// HookContextWrapper 包裹 HookContext 並記錄 metadata
// 以利清除策略進行
type HookContextWrapper struct {
	Context      *HookContext
	Success      bool
	LastAccessed time.Time
	AccessCount  int
	WrittenCount int // 被寫入次數（用於活躍度統計）
}

type HookContextManager struct {
	mu         sync.Mutex
	contexts   map[string][]*HookContextWrapper // hookName -> list
	maxTotal   int
	minPerHook int
	evictStats struct {
		removed int
	}
}

func NewHookContextManager(maxTotal, minPerHook int) *HookContextManager {
	return &HookContextManager{
		contexts:   make(map[string][]*HookContextWrapper),
		maxTotal:   maxTotal,
		minPerHook: minPerHook,
	}
}

func (m *HookContextManager) Add(ctx *HookContext, name string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, log := range ctx.GetExecutionLog() {
		wrapper := &HookContextWrapper{
			Context:      ctx,
			Success:      log.Success,
			LastAccessed: time.Now(),
			AccessCount:  1,
			WrittenCount: 1,
		}

		// 若已有相同內容 context，則視為活躍度累加（不重複儲存）
		existing := m.contexts[name]
		for _, w := range existing {
			if w.Context == ctx {
				w.WrittenCount++
				w.LastAccessed = time.Now()
				return
			}
		}

		m.contexts[name] = append(m.contexts[name], wrapper)
		m.cleanupIfNeeded()
	}

}

// GetRecent returns the recent HookContexts for a hook
func (m *HookContextManager) GetRecent(hookName string) []*HookContextWrapper {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, ctx := range m.contexts[hookName] {
		ctx.AccessCount++
		ctx.LastAccessed = time.Now()
	}
	return m.contexts[hookName]
}

func (m *HookContextManager) cleanupIfNeeded() {
	total := 0
	for _, list := range m.contexts {
		total += len(list)
	}
	if total <= m.maxTotal {
		return
	}

	// 優先清除寫入次數高且成功的 context（活躍度高）
	for hook, list := range m.contexts {
		if len(list) <= m.minPerHook {
			continue
		}
		newList := []*HookContextWrapper{}
		trueSeen, falseSeen := false, false

		// 預先統計成功/失敗各保留是否達標
		for _, ctx := range list {
			if ctx.Success {
				trueSeen = true
			} else {
				falseSeen = true
			}
		}

		for _, ctx := range list {
			if ctx.WrittenCount > 3 && ctx.Success && trueSeen && falseSeen {
				m.evictStats.removed++
				continue // 清除高頻成功紀錄
			}
			newList = append(newList, ctx)
		}
		m.contexts[hook] = newList
	}
}

// Stats 回傳目前保留的總量與清除數
func (m *HookContextManager) Stats() (total int, removed int, approxMemUsage int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	total = 0
	approxMemUsage = 0
	for _, list := range m.contexts {
		total += len(list)
		for _, w := range list {
			approxMemUsage += int64(unsafe.Sizeof(*w))
		}
	}
	return total, m.evictStats.removed, approxMemUsage
}
