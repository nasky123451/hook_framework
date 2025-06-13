package hooks

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

type HookManager struct {
	mu        sync.RWMutex
	hooks     map[string][]HookHandler
	stats     map[string]*HookStats
	statsVer  StatsVersion
	pool      *ants.Pool
	plugins   []Plugin
	hookRoles map[string][]string
}

type HookOptions struct {
	Priority int
	Roles    []string
	Metadata map[string]interface{}
}

func NewHookManager() *HookManager {
	pool, _ := ants.NewPool(100)
	return &HookManager{
		hooks:     make(map[string][]HookHandler),
		stats:     make(map[string]*HookStats),
		statsVer:  StatsVersion{Version: 0, Timestamp: time.Now()},
		pool:      pool,
		plugins:   make([]Plugin, 0),
		hookRoles: make(map[string][]string),
	}
}

func (hm *HookManager) AddHook(name string, handler HookHandler) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	// 只註冊短名（移除插件名稱前綴）
	if dotIdx := strings.LastIndex(name, "."); dotIdx != -1 {
		name = name[dotIdx+1:]
	}
	hm.hooks[name] = append(hm.hooks[name], handler)
}

func (hm *HookManager) RegisterHook(name string, priority int, handler HookHandlerFunc) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	// 包成 HookHandler 實體
	wrapped := &BaseHookHandler{
		name:     name,
		priority: priority,
		handler:  handler,
	}

	// 只保留短名（去掉 Plugin 名稱前綴）
	if dotIdx := strings.LastIndex(name, "."); dotIdx != -1 {
		name = name[dotIdx+1:]
	}

	hm.hooks[name] = append(hm.hooks[name], wrapped)

	// 依照 Priority 排序
	sorted := hm.hooks[name]
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Priority() < sorted[j].Priority()
	})
	hm.hooks[name] = sorted
}

// RegisterHookWithOptions 更彈性的註冊方式
func (hm *HookManager) RegisterHookWithOptions(name string, opt HookOptions, handler HookHandlerFunc) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	wrapped := &BaseHookHandler{
		name:     name,
		priority: opt.Priority,
		handler:  handler,
		roles:    opt.Roles,
		metadata: opt.Metadata,
	}

	// 短名註冊（去掉 Plugin 前綴）
	if dotIdx := strings.LastIndex(name, "."); dotIdx != -1 {
		name = name[dotIdx+1:]
	}

	hm.hooks[name] = append(hm.hooks[name], wrapped)

	sort.Slice(hm.hooks[name], func(i, j int) bool {
		return hm.hooks[name][i].Priority() < hm.hooks[name][j].Priority()
	})
}

func (hm *HookManager) GetRegisteredHooks() map[string][]string {
	hm.mu.RLock()
	defer hm.mu.RUnlock()

	result := make(map[string][]string)
	for name, handlers := range hm.hooks {
		var handlerNames []string
		for _, h := range handlers {
			handlerNames = append(handlerNames, h.Name())
		}
		result[name] = handlerNames
	}
	return result
}

// 簡單角色權限檢查範例 (可整合外部 IAM)
func (hm *HookManager) CheckPermission(hookName string, ctx *HookContext) bool {
	role := ctx.GetUserData("role")
	hooks, ok := hm.hooks[hookName]
	if !ok {
		// 如果無明確設定，視為允許
		return true
	}
	for _, h := range hooks {
		for _, r := range h.Roles() {
			if r == role {
				return true
			}
		}
	}
	ctx.AddError(fmt.Errorf("permission denied: role '%v' cannot execute hook '%s'", role, hookName))
	return false
}

func (hm *HookManager) GenerateHookDocs(path string) error {
	allHooks := GetAllRegisteredHooks()
	lines := []string{
		"# Hook Documentation",
		"",
		fmt.Sprintf("Generated at: %s", time.Now().Format(time.RFC3339)),
		"",
	}

	for _, h := range allHooks {
		lines = append(lines,
			fmt.Sprintf("## %s", h.HookName),
			fmt.Sprintf("- 📄 Description: %s", h.Description),
			fmt.Sprintf("- 🔗 Registered From: %s", h.RegisteredFrom),
		)

		if len(h.Roles) > 0 {
			lines = append(lines, "- 👥 Allowed Roles:")
			for _, r := range h.Roles {
				lines = append(lines, fmt.Sprintf("  - %s", r))
			}
		}

		if len(h.ParamHints) > 0 {
			lines = append(lines, "- 🎯 Expected Parameters:")
			for _, p := range h.ParamHints {
				lines = append(lines, fmt.Sprintf("  - %s", p))
			}
		}

		lines = append(lines, "") // spacing
	}

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}
