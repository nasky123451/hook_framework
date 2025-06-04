package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type SubscriptionReminderPlugin struct{}

func (p *SubscriptionReminderPlugin) Name() string {
	return "SubscriptionReminderPlugin"
}

func (p *SubscriptionReminderPlugin) GetHookNames() []string {
	return []string{"subscription_reminder"}
}

func (p *SubscriptionReminderPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "subscription_reminder", 10, "subscriber", func(ctx *hooks.HookContext) hooks.HookResult {
		userID := ctx.GetEnvData("user_id")

		message := fmt.Sprintf("[SubscriptionReminderPlugin] Subscription reminder triggered for user_id = %v.\n", userID)
		ctx.SetEnvData("approval_message", message)
		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("SubscriptionReminderPlugin", func() hooks.Plugin {
		return &SubscriptionReminderPlugin{}
	})
}
