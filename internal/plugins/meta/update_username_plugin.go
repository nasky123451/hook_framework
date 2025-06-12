package meta

import (
	"fmt"

	"hook_framework/internal/hooks"
)

type UpdateUsernamePlugin struct{}

func (p *UpdateUsernamePlugin) Name() string {
	return "UpdateUsernamePlugin"
}

func (p *UpdateUsernamePlugin) GetHookNames() []string {
	return []string{"update_username"}
}

func (p *UpdateUsernamePlugin) RegisterHooks(hm *hooks.HookManager) {
	hooks.New("update_username").
		WithDescription("Handles updating a user's username").
		WithParamHints("username").
		WithPriority(10).
		AllowRoles("admin").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			// 直接從 ctx 拿要更新的 username
			username, ok := ctx.GetEnvString("username")
			if !ok || username == "" {
				return hooks.HookResult{Error: fmt.Errorf("username is missing or invalid")}
			}

			//todo: 實際的更新邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Username updated to %s", username)
		}).RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("UpdateUsernamePlugin", func() hooks.Plugin {
		return &UpdateUsernamePlugin{}
	})
}
