package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type UserPreferencePlugin struct{}

func (p *UserPreferencePlugin) Name() string {
	return "UserPreferencePlugin"
}

func (p *UserPreferencePlugin) GetHookNames() []string {
	return []string{"set_user_pref"}
}

func (p *UserPreferencePlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "set_user_pref", 10, "user", func(ctx *hooks.HookContext) hooks.HookResult {
		theme, _ := ctx.GetEnvData("theme").(string)

		message := fmt.Sprintf("[UserPreferencePlugin] User preference updated: theme = %s.", theme)
		ctx.SetEnvData("approval_message", message)

		fmt.Println(message) // 可保留日誌輸出

		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("UserPreferencePlugin", func() hooks.Plugin {
		return &UserPreferencePlugin{}
	})
}
