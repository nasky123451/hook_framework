package hooks

import "log"

func (hm *HookManager) LoadAndInitPlugins(ctx *HookContext, plugins []Plugin) {
	for _, plugin := range plugins {
		hm.plugins = append(hm.plugins, plugin)
		plugin.RegisterHooks(hm)
		log.Printf("[Plugin] Loaded plugin: %T", plugin)
	}
}

func (hm *HookManager) GetPlugins() []Plugin {
	return hm.plugins
}

var pluginFactories = map[string]func() Plugin{}

func RegisterPluginFactory(name string, factory func() Plugin) {
	pluginFactories[name] = factory
}

func CreatePluginsFromConfig(configs []map[string]interface{}) []Plugin {
	var plugins []Plugin
	for _, conf := range configs {
		pluginType, ok := conf["type"].(string)
		if !ok || pluginType == "" {
			continue
		}
		factory, ok := GetPluginFactoryByName(pluginType)
		if !ok {
			log.Printf("[Plugin] No factory found for type: %s", pluginType)
			continue
		}
		instance := factory()
		if instance == nil {
			log.Printf("[Plugin] Factory returned nil for type: %s", pluginType)
			continue
		}
		plugins = append(plugins, instance)
	}
	return plugins
}
