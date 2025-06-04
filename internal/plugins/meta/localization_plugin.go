package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type LocalizationPlugin struct{}

func (p *LocalizationPlugin) Name() string {
	return "LocalizationPlugin"
}

func (p *LocalizationPlugin) GetHookNames() []string {
	return []string{"set_language"}
}

func (p *LocalizationPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "set_language", 10, "user", func(ctx *hooks.HookContext) hooks.HookResult {
		language := ctx.GetEnvData("language")

		message := fmt.Sprintf("[LocalizationPlugin] Language set to %v.\n", language)
		ctx.SetEnvData("approval_message", message)
		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("LocalizationPlugin", func() hooks.Plugin {
		return &LocalizationPlugin{}
	})
}
