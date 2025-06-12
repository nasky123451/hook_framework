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
	// 模擬發送 Email
	hooks.New("notify_account_created").
		WithDescription("Handles sending a welcome email when an account is created").
		WithParamHints("email").
		WithPriority(10).
		AllowRoles("admin", "system").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			email, _ := ctx.GetEnvString("email")

			fmt.Println("[NotifyPlugin] Sending welcome email to:", email)
			// todo: 實際的發送 Email 邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Welcome email sent to %s", email)
		}).RegisterTo(hm)

	// 模擬建立 Jira 任務
	hooks.New("create_jira_task").
		WithDescription("Handles creating a Jira task for account setup").
		WithParamHints("task_details").
		WithPriority(10).
		AllowRoles("admin").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			fmt.Println("[NotifyPlugin] Creating Jira task for account setup...")

			//todo: 實際的 Jira 任務建立邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Jira task created for account setup")
		}).RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("NotifyPlugin", func() hooks.Plugin {
		return &NotifyPlugin{}
	})
}
