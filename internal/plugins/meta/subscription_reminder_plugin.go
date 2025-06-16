package meta

import (
	"hook_framework/internal/hooks"
)

type SubscriptionReminderPlugin struct{}

func (p *SubscriptionReminderPlugin) Name() string {
	return "SubscriptionReminderPlugin"
}

func (p *SubscriptionReminderPlugin) GetHookNames() []string {
	return []string{"subscription_reminder"}
}

func (p *SubscriptionReminderPlugin) RegisterHooks(hm *hooks.HookManager) {
	hookDefs := hooks.HookBuilders{
		{HookName: "subscription_reminder",
			Description: "Handles sending subscription reminders to users",
			ParamHints:  []string{"user_id"},
			Permissions: "",
			Priority:    10,
			Handler:     handleSubscriptionReminder,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleSubscriptionReminder(ctx *hooks.HookContext) hooks.HookResult {
	userID, _ := ctx.GetEnvData("user_id")

	// 實際的訂閱提醒邏輯可以在這裡實現

	return ctx.SuccessWithMessage("Subscription reminder sent for user ID %s", userID)
}

func init() {
	hooks.RegisterPluginType("SubscriptionReminderPlugin", func() hooks.Plugin {
		return &SubscriptionReminderPlugin{}
	})
}
