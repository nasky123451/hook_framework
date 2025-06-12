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
	hooks.New("subscription_reminder").
		WithDescription("Handles sending subscription reminders to users").
		WithParamHints("user_id").
		WithPriority(10).
		AllowRoles("subscriber").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			userID, _ := ctx.GetEnvData("user_id")

			//todo: 實際的訂閱提醒邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Subscription reminder sent for user ID %s", userID)
		}).RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("SubscriptionReminderPlugin", func() hooks.Plugin {
		return &SubscriptionReminderPlugin{}
	})
}
