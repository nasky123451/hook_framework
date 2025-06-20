package framework

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/internal/plugins"
	"hook_framework/internal/utils"
	"log"
)

func InitializeFramework() *ClientInputProcessor {
	printer := utils.NewPrinter()
	env := hooks.NewHookEnvironment()

	pluginManager := hooks.NewPluginManager()

	initializeHooksAndPlugins(pluginManager, env)

	if len(hooks.GetAllHookNames()) == 0 {
		panic("No hooks registered. Ensure plugins are correctly registering hooks.")
	}

	env.HookGraph.AddChain("create_account", "notify_account_created", "create_jira_task")

	processor := NewClientInputProcessor(env, printer)

	err := hooks.InitPermissions("./permissions.json")
	if err != nil {
		log.Fatalf("初始化權限失敗: %v", err)
	}

	if err := processor.Env.HookManager.GenerateHookDocs("hook_docs.md"); err != nil {
		panic(fmt.Errorf("failed to generate hook docs: %w", err))
	}

	return processor
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
