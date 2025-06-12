package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
)

type WebhookSyncPlugin struct{}

func (p *WebhookSyncPlugin) Name() string {
	return "WebhookSyncPlugin"
}

func (p *WebhookSyncPlugin) GetHookNames() []string {
	return []string{"webhook_sync"}
}

func (p *WebhookSyncPlugin) RegisterHooks(hm *hooks.HookManager) {
	hooks.New("webhook_sync").
		WithDescription("Handles synchronization of webhooks from external sources").
		WithParamHints("source").
		WithPriority(10).
		AllowRoles("integration").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			source, _ := ctx.GetEnvData("source")

			fmt.Printf("[WebhookSyncPlugin] Syncing webhooks from source: %s\n", source)
			//todo: 實際的 webhook 同步邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Webhook sync completed successfully")
		}).RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("WebhookSyncPlugin", func() hooks.Plugin {
		return &WebhookSyncPlugin{}
	})
}
