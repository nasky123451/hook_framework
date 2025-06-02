package hooks

import (
	"strings"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

type HookManager struct {
	mu       sync.RWMutex
	hooks    map[string][]HookHandler
	stats    map[string]*HookStats
	statsVer StatsVersion
	pool     *ants.Pool

	plugins []Plugin
}

func NewHookManager() *HookManager {
	pool, _ := ants.NewPool(100)
	return &HookManager{
		hooks: make(map[string][]HookHandler),
		stats: make(map[string]*HookStats),
		statsVer: StatsVersion{
			Version:   0,
			Timestamp: time.Now(),
		},
		pool:    pool,
		plugins: make([]Plugin, 0),
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
