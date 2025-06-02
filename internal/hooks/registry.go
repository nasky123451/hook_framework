package hooks

import (
	"fmt"
	"log"
	"sync"
)

// RegisterHook 將動態定義的 Hook 註冊至 HookManager 中
//
// 參數：
// - hookManager: 目標 Hook 管理器，不能為 nil
// - name: Hook 名稱，必須已在 hookNameRegistry 註冊過
// - priority: Hook 優先級，數字越大優先級越高
// - filter: Hook 的執行過濾函式，決定是否執行此 Hook
// - run: Hook 具體執行邏輯
func RegisterHook(
	hookManager *HookManager,
	name string,
	priority int,
	filter func(ctx *HookContext) bool,
	run func(ctx *HookContext) HookResult,
) {
	if hookManager == nil {
		panic("HookManager is nil. Ensure it is properly initialized.")
	}

	normalizedName, exists := GetHookName(name)
	if !exists || normalizedName == "" {
		log.Printf("[RegisterHook] Hook name '%s' is invalid or not registered.", name)
		panic(fmt.Sprintf("Hook name '%s' is not registered or invalid.", name))
	}

	log.Printf("[RegisterHook] Registering hook '%s' (normalized: '%s').", name, normalizedName)

	hookManager.AddHook(normalizedName, &dynamicHookHandler{
		name:     normalizedName,
		priority: priority,
		filter:   filter,
		run:      run,
	})

	// Debug: 列出目前已註冊的 hooks
	log.Printf("[RegisterHook] Current registered hooks: %+v", hookManager.GetRegisteredHooks())
}

// dynamicHookHandler 是一個動態實現的 HookHandler，封裝了執行優先級與邏輯
type dynamicHookHandler struct {
	name     string
	priority int
	filter   func(ctx *HookContext) bool
	run      func(ctx *HookContext) HookResult
}

func (h *dynamicHookHandler) Name() string {
	return h.name
}

func (h *dynamicHookHandler) Priority() int {
	return h.priority
}

func (h *dynamicHookHandler) Filter(ctx *HookContext) bool {
	if h.filter == nil {
		// 預設允許執行
		return true
	}
	return h.filter(ctx)
}

func (h *dynamicHookHandler) Run(ctx *HookContext) HookResult {
	if h.run != nil {
		return h.run(ctx)
	}
	return HookResult{}
}

// hookNameRegistry 用於管理所有已註冊的 Hook 名稱與其對應的標準化名稱
var (
	hookNameRegistryMu sync.RWMutex
	hookNameRegistry   = make(map[string]string)
)

// GetHookName 根據原始 Hook 名稱，回傳標準化的名稱（格式：pluginName.HookName）及是否存在
func GetHookName(name string) (string, bool) {
	hookNameRegistryMu.RLock()
	defer hookNameRegistryMu.RUnlock()

	normalizedName, exists := hookNameRegistry[name]
	return normalizedName, exists
}

// GetAllHookNames 返回所有已註冊的 hookName 的副本，避免外部直接改動內部資料結構
func GetAllHookNames() map[string]string {
	hookNameRegistryMu.RLock()
	defer hookNameRegistryMu.RUnlock()

	copied := make(map[string]string, len(hookNameRegistry))
	for k, v := range hookNameRegistry {
		copied[k] = v
	}
	return copied
}

// InitializeHookNames 以 map 形式的 hookConfigs 初始化 hookNameRegistry
//
// 參數 hookConfigs 格式:
//
//	{
//	  "PluginA": ["HookX", "HookY"],
//	  "PluginB": ["HookZ"],
//	}
//
// 對應 hookNameRegistry 會存儲為: {"HookX": "PluginA.HookX", "HookY": "PluginA.HookY", "HookZ": "PluginB.HookZ"}
func InitializeHookNames(hookConfigs map[string][]string) {
	hookNameRegistryMu.Lock()
	defer hookNameRegistryMu.Unlock()

	for pluginName, hooks := range hookConfigs {
		for _, hook := range hooks {
			hookNameRegistry[hook] = fmt.Sprintf("%s.%s", pluginName, hook)
		}
	}
}
