package hooks

type PluginHandler func(ctx *HookContext, params map[string]interface{}) HookResult

type DynamicPlugin struct {
	Name        string
	Version     string
	Description string
	Handler     PluginHandler
}

type PluginManager struct {
	plugins []Plugin
}

func NewPluginManager() *PluginManager {
	return &PluginManager{}
}

// RegisterPlugin 註冊插件
func (pm *PluginManager) RegisterPlugin(plugin Plugin) {
	pm.plugins = append(pm.plugins, plugin)
}

// InitializePlugins 初始化所有插件
func (pm *PluginManager) InitializePlugins(ctx *HookContext, hm *HookManager) []Plugin {
	var initializedPlugins []Plugin
	for _, plugin := range pm.plugins {
		plugin.RegisterHooks(hm)
		initializedPlugins = append(initializedPlugins, plugin)
	}
	return initializedPlugins
}
