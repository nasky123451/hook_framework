package meta

import (
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
	hookDefs := hooks.HookBuilders{
		{HookName: "update_username",
			Description: "Handles updating a user's username",
			ParamHints:  []string{"username"},
			Permissions: "",
			Priority:    10,
			Handler:     handleUpdateUsername,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleUpdateUsername(ctx *hooks.HookContext) hooks.HookResult {
	username, _ := ctx.GetEnvString("username")

	// 實際的用戶名更新邏輯可以在這裡實現
	// 例如，更新數據庫中的用戶名

	return ctx.SuccessWithMessage("Username updated to %s", username)
}

func init() {
	hooks.RegisterPluginType("UpdateUsernamePlugin", func() hooks.Plugin {
		return &UpdateUsernamePlugin{}
	})
}
