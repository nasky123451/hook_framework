package framework

import (
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
	nlpEngine := nlp.NewNLP()

	initializeHooksAndPlugins(pluginManager, nlpEngine, env)

	if len(hooks.GetAllHookNames()) == 0 {
		panic("No hooks registered. Ensure plugins are correctly registering hooks.")
	}

	log.Println("[InitializeFramework] Registering operation handlers.")
	registerOperationHandlers(nlpEngine, env)

	processor := NewClientInputProcessor(env, nlpEngine, printer)

	return processor, printer, hooks.GetRegisteredPluginTypes()
}

func initializeHooksAndPlugins(pm *hooks.PluginManager, nlpEngine *nlp.NLP, env *hooks.HookEnvironment) {
	// 初始化 Hook 註冊表與 Hook 名稱
	hooks.InitializePluginRegistry()
	hookConfigs := meta.GetAllHookConfigs()
	hooks.InitializeHookNames(hookConfigs)

	for _, plugin := range hooks.GetRegisteredPluginTypes() {
		if registrable, ok := plugin.(hooks.ParserRegistrable); ok {
			registrable.RegisterParsers(nlpEngine)
		}
		pm.RegisterPlugin(plugin)
	}

	pm.InitializePlugins(env.Context, env.HookManager)
}

func registerOperationHandlers(nlpEngine *nlp.NLP, env *hooks.HookEnvironment) {
	hookConfigs := meta.GetAllHookConfigs()
	hookNames := hooks.GetAllHookNames()

	utils.RegisterAllHandlers(func(name string, handler func(ctx interface{}, params map[string]interface{})) {
		hooks.RegisterOperationHandler(name, func(ctx *hooks.HookContext, params map[string]interface{}) {
			handler(ctx, params)
		})
	}, hookNames, hookConfigs)
}
