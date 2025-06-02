package framework

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/internal/plugins/meta"
	"hook_framework/pkg/nlp"
	"hook_framework/pkg/utils"
	"log"
)

func InitializeFramework() (*ClientInputProcessor, *utils.Printer, []hooks.Plugin) {
	printer := utils.NewPrinter()
	env := hooks.NewHookEnvironment("system", "main")

	pluginManager := hooks.NewPluginManager()

	// 初始化 NLP 引擎
	nlpEngine := nlp.NewNLP()

	// 初始化插件註冊表與 Hook 註冊機制
	hooks.InitializePluginRegistry()
	hookConfigs := meta.GetAllHookConfigs() // 使用 meta.GetAllHookConfigs
	hooks.InitializeHookNames(hookConfigs)

	// 讀取插件設定檔
	configPath := "./plugin_config.json"
	configData, err := utils.LoadPluginConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load plugin configuration: %v", err))
	}

	switch configData.(type) {
	case map[string]interface{}, []interface{}:
		log.Printf("[InitializeFramework] Plugin configuration loaded: %T", configData)
	default:
		panic(fmt.Sprintf("Unsupported plugin configuration format: %T", configData))
	}

	// 將 configData 轉換成 plugin 配置陣列
	var pluginConfigs []map[string]interface{}
	switch v := configData.(type) {
	case []interface{}:
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				pluginConfigs = append(pluginConfigs, m)
			}
		}
	case map[string]interface{}:
		pluginConfigs = append(pluginConfigs, v)
	default:
		panic(fmt.Sprintf("Unsupported plugin configuration format: %T", configData))
	}

	// 從 config 建立 plugin 實例
	pluginsList := hooks.CreatePluginsFromConfig(pluginConfigs)

	// 若 plugin 支援 parser，註冊至 NLP engine
	for _, plugin := range pluginsList {
		if registrable, ok := plugin.(hooks.ParserRegistrable); ok {
			log.Printf("[RegisterParsers] Plugin %T is ParserRegistrable, registering...", plugin)
			registrable.RegisterParsers(nlpEngine)
		} else {
			log.Printf("[RegisterParsers] Plugin %T is NOT ParserRegistrable", plugin)
		}
		pluginManager.RegisterPlugin(plugin)
	}

	// 初始化 plugin 並註冊 hooks
	pluginManager.InitializePlugins(env.Context, env.HookManager)

	if len(hooks.GetAllHookNames()) == 0 {
		panic("No hooks registered. Ensure plugins are correctly registering hooks.")
	}

	// 註冊所有 NLP operation handler
	log.Println("[InitializeFramework] Registering operation handlers.")
	utils.RegisterAllHandlers(func(name string, handler func(ctx interface{}, params map[string]interface{})) {
		hooks.RegisterOperationHandler(name, func(ctx *hooks.HookContext, params map[string]interface{}) {
			handler(ctx, params)
		})
	}, hooks.GetAllHookNames(), hookConfigs)

	// 初始化 ClientInputProcessor 並指派 HookManager
	processor := NewClientInputProcessor(env, nlpEngine, printer)
	processor.Env.HookManager = env.HookManager

	return processor, printer, pluginsList
}
