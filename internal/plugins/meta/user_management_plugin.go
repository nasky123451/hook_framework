package meta

import (
	"hook_framework/internal/hooks"
)

type UserManagementPlugin struct{}

func (p *UserManagementPlugin) Name() string {
	return "UserManagementPlugin"
}

func (p *UserManagementPlugin) GetHookNames() []string {
	return []string{"create_user", "update_user", "delete_user"}
}

func (p *UserManagementPlugin) RegisterHooks(hm *hooks.HookManager) {
	handlers := map[string]func(ctx *hooks.HookContext) hooks.HookResult{
		"create_user": func(ctx *hooks.HookContext) hooks.HookResult {
			username, _ := ctx.GetEnvString("username")
			email, _ := ctx.GetEnvString("email")

			return ctx.SuccessWithMessage("User %s created with email %s", username, email)
		},
		"update_user": func(ctx *hooks.HookContext) hooks.HookResult {
			username, _ := ctx.GetEnvString("username")
			newEmail, _ := ctx.GetEnvString("new_email")

			return ctx.SuccessWithMessage("User %s updated to new email %s", username, newEmail)
		},
		"delete_user": func(ctx *hooks.HookContext) hooks.HookResult {
			username, _ := ctx.GetEnvString("username")

			return ctx.SuccessWithMessage("User %s deleted successfully", username)
		},
	}

	for hookName, handler := range handlers {
		hooks.New(hookName).
			WithDescription("Handles user management operations: "+hookName).
			WithParamHints("username", "email", "new_email").
			WithPriority(10).
			AllowRoles("admin").
			Handle(handler).RegisterTo(hm)
	}
}

func init() {
	hooks.RegisterPluginType("UserManagementPlugin", func() hooks.Plugin {
		return &UserManagementPlugin{}
	})
}
