package plugins

import (
	"fmt"
	"log"

	"hook_framework/internal/hooks"
	_ "hook_framework/internal/plugins/meta"
)

type HookConfig struct {
	Name string
}

type PluginConfig struct {
	Name     string
	Priority int
	Enabled  bool
	Hooks    []HookConfig
}

// GetAllHookConfigs 返回所有插件的 Hook 配置
func GetAllHookConfigs() map[string][]string {
	hookConfigs := make(map[string][]string)

	// 動態讀取所有已註冊的插件
	plugins, err := hooks.LoadPluginsFromRegistry() // 使用 hookmanager.LoadPluginsFromRegistry 動態加載插件
	if err != nil {
		panic(fmt.Sprintf("Failed to load plugins: %v", err))
	}

	for _, plugin := range plugins {
		pluginName := plugin.Name()
		hookNames := plugin.GetHookNames()

		if len(hookNames) == 0 {
			log.Printf("[GetAllHookConfigs] Plugin '%s' has no hooks registered. Ensure hooks are properly initialized.", pluginName)
		}

		hookConfigs[pluginName] = hookNames
	}

	return hookConfigs
}
