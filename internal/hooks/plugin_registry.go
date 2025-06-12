package hooks

import (
	"fmt"
	"sync"
)

// Plugin 定義插件接口
type Plugin interface {
	Name() string
	GetHookNames() []string        // 確保插件返回正確的 Hook 名稱
	RegisterHooks(hm *HookManager) // Change HookManager type to *common.HookManager
}

// --------------------
// 全局插件註冊表與鎖
// --------------------

var (
	mu             sync.RWMutex
	pluginRegistry = make(map[string]func() Plugin)
)

// RegisterPluginType 註冊插件類型
func RegisterPluginType(name string, factory func() Plugin) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := pluginRegistry[name]; exists {
		panic(fmt.Sprintf("Plugin type '%s' is already registered.", name))
	}
	pluginRegistry[name] = factory
}

// GetRegisteredPluginTypes 返回所有已註冊的插件類型
func GetRegisteredPluginTypes() []Plugin {
	mu.Lock()
	defer mu.Unlock()

	plugins := make([]Plugin, 0, len(pluginRegistry))
	for _, factory := range pluginRegistry {
		plugins = append(plugins, factory())
	}
	return plugins
}

func LoadPluginsFromRegistry() ([]Plugin, error) {
	mu.RLock()
	factories := make([]func() Plugin, 0, len(pluginRegistry))
	for _, factory := range pluginRegistry {
		factories = append(factories, factory)
	}
	mu.RUnlock()

	var plugins []Plugin
	for _, factory := range factories {
		plugin := factory()
		if plugin == nil {
			return nil, fmt.Errorf("failed to create plugin from factory")
		}
		plugins = append(plugins, plugin)
	}
	return plugins, nil
}
