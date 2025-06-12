package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
)

type LocalizationPlugin struct{}

func (p *LocalizationPlugin) Name() string {
	return "LocalizationPlugin"
}

func (p *LocalizationPlugin) GetHookNames() []string {
	return []string{"set_language"}
}

func (p *LocalizationPlugin) RegisterHooks(hm *hooks.HookManager) {
	hooks.New("set_language").
		WithDescription("Handles setting the language for the user based on their preferences").
		WithParamHints("language").
		WithPriority(10).
		AllowRoles("admin", "user").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			language, _ := ctx.GetEnvData("language")

			fmt.Println("[LocalizationPlugin] Setting language to:", language)
			// todo: 實際的語言設置邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Language set to %s", language)
		}).RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("LocalizationPlugin", func() hooks.Plugin {
		return &LocalizationPlugin{}
	})
}
