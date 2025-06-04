package hooks

// PluginMetadata 表示插件的元數據
type PluginMetadata struct {
	Name        string
	Version     string
	Description string
	Priority    int  // 添加 Priority 字段
	Enabled     bool // 添加 Enabled 字段
}

// Plugin 定義插件接口
type Plugin interface {
	Name() string
	GetHookNames() []string        // 確保插件返回正確的 Hook 名稱
	RegisterHooks(hm *HookManager) // Change HookManager type to *common.HookManager
}

// LoadPlugin 根據名稱加載插件
func LoadPlugin(name string) (Plugin, error) {
	return GetPluginByName(name) // 使用 plugin_registry.go 中的 GetPluginByName
}

// GetRegisteredPlugins 返回所有已註冊的插件名稱
func GetRegisteredPlugins() []string {
	pluginTypes := GetRegisteredPluginTypes() // 使用 plugin_registry.go 中的 GetRegisteredPluginTypes
	var pluginNames []string
	for _, plugin := range pluginTypes {
		pluginNames = append(pluginNames, plugin.Name())
	}
	return pluginNames
}
