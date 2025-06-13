package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
)

type NotifyPlugin struct{}

func (p *NotifyPlugin) Name() string { return "NotifyPlugin" }

func (p *NotifyPlugin) GetHookNames() []string {
	return []string{"notify_account_created", "create_jira_task"}
}

func (p *NotifyPlugin) RegisterHooks(hm *hooks.HookManager) {
	hookDefs := hooks.HookBuilders{
		{
			HookName:    "notify_account_created",
			Description: "Handles sending a welcome email when an account is created",
			ParamHints:  []string{"email"},
			Roles:       []string{"admin", "system"},
			Priority:    10,
			Handler:     handleNotifyAccountCreated,
		},
		{
			HookName:    "create_jira_task",
			Description: "Handles creating a Jira task for account setup",
			ParamHints:  []string{"task_details"},
			Roles:       []string{"admin"},
			Priority:    10,
			Handler:     handleCreateJiraTask,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleNotifyAccountCreated(ctx *hooks.HookContext) hooks.HookResult {
	email, _ := ctx.GetEnvString("email")

	fmt.Println("[NotifyPlugin] Sending welcome email to:", email)
	// 實際的發送 Email 邏輯可以在這裡實現

	return ctx.SuccessWithMessage("Welcome email sent to %s", email)
}

func handleCreateJiraTask(ctx *hooks.HookContext) hooks.HookResult {
	fmt.Println("[NotifyPlugin] Creating Jira task for account setup...")

	// 實際的 Jira 任務建立邏輯可以在這裡實現

	return ctx.SuccessWithMessage("Jira task created for account setup")
}

func init() {
	hooks.RegisterPluginType("NotifyPlugin", func() hooks.Plugin {
		return &NotifyPlugin{}
	})
}
