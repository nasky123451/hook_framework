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
	hookDefs := hooks.HookBuilders{
		{HookName: "create_user",
			Description: "Creates a new user with a username and email",
			ParamHints:  []string{"username", "email"},
			Roles:       []string{"admin"},
			Priority:    10,
			Handler:     handleCreateUser,
		},
		{HookName: "update_user",
			Description: "Updates a user's email address",
			ParamHints:  []string{"username", "new_email"},
			Roles:       []string{"admin"},
			Priority:    10,
			Handler:     handleUpdateUser,
		},
		{HookName: "delete_user",
			Description: "Deletes a user by username",
			ParamHints:  []string{"username"},
			Roles:       []string{"admin"},
			Priority:    10,
			Handler:     handleDeleteUser,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleCreateUser(ctx *hooks.HookContext) hooks.HookResult {
	username, _ := ctx.GetEnvString("username")
	email, _ := ctx.GetEnvString("email")

	// 實際的用戶創建邏輯可以在這裡實現
	// 例如，將用戶信息存儲到數據庫中

	return ctx.SuccessWithMessage("User %s created with email %s", username, email)
}

func handleUpdateUser(ctx *hooks.HookContext) hooks.HookResult {
	username, _ := ctx.GetEnvString("username")
	newEmail, _ := ctx.GetEnvString("new_email")

	// 實際的用戶更新邏輯可以在這裡實現
	// 例如，更新數據庫中的用戶電子郵件地址

	return ctx.SuccessWithMessage("User %s updated to new email %s", username, newEmail)
}

func handleDeleteUser(ctx *hooks.HookContext) hooks.HookResult {
	username, _ := ctx.GetEnvString("username")
	// 實際的用戶刪除邏輯可以在這裡實現
	// 例如，從數據庫中刪除用戶
	return ctx.SuccessWithMessage("User %s deleted successfully", username)
}

func init() {
	hooks.RegisterPluginType("UserManagementPlugin", func() hooks.Plugin {
		return &UserManagementPlugin{}
	})
}
