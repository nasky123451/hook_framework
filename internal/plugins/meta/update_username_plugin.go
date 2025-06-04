package meta

import (
	"fmt"

	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type UpdateUsernamePlugin struct{}

func (p *UpdateUsernamePlugin) Name() string {
	return "UpdateUsernamePlugin"
}

func (p *UpdateUsernamePlugin) GetHookNames() []string {
	return []string{"update_username"}
}

func (p *UpdateUsernamePlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "update_username", 15, "admin", func(ctx *hooks.HookContext) hooks.HookResult {
		// 直接從 ctx 拿要更新的 username
		username, ok := ctx.Get("username").(string)
		if !ok || username == "" {
			return hooks.HookResult{Error: fmt.Errorf("username is missing or invalid")}
		}

		message := fmt.Sprintf("[UpdateUsernamePlugin] Updating username to: %s", username)
		ctx.SetEnvData("approval_message", message)
		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("UpdateUsernamePlugin", func() hooks.Plugin {
		return &UpdateUsernamePlugin{}
	})
}
