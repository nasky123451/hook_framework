package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type UserManagementPlugin struct{}

func (p *UserManagementPlugin) Name() string {
	return "UserManagementPlugin"
}

func (p *UserManagementPlugin) GetHookNames() []string {
	return []string{"create_user", "update_user", "delete_user"}
}

func (p *UserManagementPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	handlers := map[string]func(ctx *hooks.HookContext) hooks.HookResult{
		"create_user": func(ctx *hooks.HookContext) hooks.HookResult {
			username, _ := ctx.GetEnvData("username").(string)
			email, _ := ctx.GetEnvData("email").(string)
			message := fmt.Sprintf("[UserManagementPlugin] Created user: %s <%s>", username, email)
			ctx.SetEnvData("action_message", message)
			fmt.Println(message)
			return hooks.HookResult{Success: true}
		},
		"update_user": func(ctx *hooks.HookContext) hooks.HookResult {
			username, _ := ctx.GetEnvData("username").(string)
			newEmail, _ := ctx.GetEnvData("new_email").(string)
			message := fmt.Sprintf("[UserManagementPlugin] Updated user %s to new email: %s", username, newEmail)
			ctx.SetEnvData("action_message", message)
			fmt.Println(message)
			return hooks.HookResult{Success: true}
		},
		"delete_user": func(ctx *hooks.HookContext) hooks.HookResult {
			username, _ := ctx.GetEnvData("username").(string)
			message := fmt.Sprintf("[UserManagementPlugin] Deleted user: %s", username)
			ctx.SetEnvData("action_message", message)
			fmt.Println(message)
			return hooks.HookResult{Success: true}
		},
	}

	for hookName, handler := range handlers {
		utils.RegisterDynamicHook(hookManager, hookName, 10, "admin", handler)
	}
}

func init() {
	hooks.RegisterPluginType("UserManagementPlugin", func() hooks.Plugin {
		return &UserManagementPlugin{}
	})
}
