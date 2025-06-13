package framework

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/internal/plugins"
	"hook_framework/internal/utils"
)

func InitializeFramework() (*ClientInputProcessor, *utils.Printer, *hooks.HookManager, *hooks.HookGraph) {
	printer := utils.NewPrinter()
	env := hooks.NewHookEnvironment()

	pluginManager := hooks.NewPluginManager()

	initializeHooksAndPlugins(pluginManager, env)

	if len(hooks.GetAllHookNames()) == 0 {
		panic("No hooks registered. Ensure plugins are correctly registering hooks.")
	}

	// 建立 HookGraph 並定義流程鏈結
	hg := hooks.NewHookGraph(env.HookManager)

	hg.AddChain("create_account", "notify_account_created", "create_jira_task")

	processor := NewClientInputProcessor(env, printer)

	if err := processor.Env.HookManager.GenerateHookDocs("hook_docs.md"); err != nil {
		panic(fmt.Errorf("failed to generate hook docs: %w", err))
	}

	return processor, printer, processor.Env.HookManager, hg
}
func initializeHooksAndPlugins(pm *hooks.PluginManager, env *hooks.HookEnvironment) {
	// 初始化 Hook 註冊表與 Hook 名稱
	hookConfigs := plugins.GetAllHookConfigs()
	hooks.InitializeHookNames(hookConfigs)

	for _, plugin := range hooks.GetRegisteredPluginTypes() {
		if registrable, ok := plugin.(hooks.ParserRegistrable); ok {
			registrable.RegisterParsers()
		}
		pm.RegisterPlugin(plugin)
	}

	pm.InitializePlugins(env.HookManager)
}
