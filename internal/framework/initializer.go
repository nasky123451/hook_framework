package framework

import (
	"hook_framework/internal/hooks"
	"hook_framework/internal/plugins"
	"hook_framework/pkg/utils"
	"log"
)

func InitializeFramework() (*ClientInputProcessor, *utils.Printer, []hooks.Plugin) {
	printer := utils.NewPrinter()
	env := hooks.NewHookEnvironment("system", "main")

	pluginManager := hooks.NewPluginManager()

	initializeHooksAndPlugins(pluginManager, env)

	if len(hooks.GetAllHookNames()) == 0 {
		panic("No hooks registered. Ensure plugins are correctly registering hooks.")
	}

	log.Println("[InitializeFramework] Registering operation handlers.")
	registerOperationHandlers(env)

	processor := NewClientInputProcessor(env, printer)

	return processor, printer, hooks.GetRegisteredPluginTypes()
}

func initializeHooksAndPlugins(pm *hooks.PluginManager, env *hooks.HookEnvironment) {
	// 初始化 Hook 註冊表與 Hook 名稱
	hooks.InitializePluginRegistry()
	hookConfigs := plugins.GetAllHookConfigs()
	hooks.InitializeHookNames(hookConfigs)

	for _, plugin := range hooks.GetRegisteredPluginTypes() {
		if registrable, ok := plugin.(hooks.ParserRegistrable); ok {
			registrable.RegisterParsers()
		}
		pm.RegisterPlugin(plugin)
	}

	pm.InitializePlugins(env.Context, env.HookManager)
}

func registerOperationHandlers(env *hooks.HookEnvironment) {
	hookConfigs := plugins.GetAllHookConfigs()
	hookNames := hooks.GetAllHookNames()

	utils.RegisterAllHandlers(func(name string, handler func(ctx interface{}, params map[string]interface{})) {
		hooks.RegisterOperationHandler(name, func(ctx *hooks.HookContext, params map[string]interface{}) {
			handler(ctx, params)
		})
	}, hookNames, hookConfigs)
}
