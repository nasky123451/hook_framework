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
	hookDefs := hooks.HookBuilders{
		{HookName: "webhook_sync",
			Description: "Handles synchronization of webhooks from external sources",
			ParamHints:  []string{"source"},
			Roles:       []string{"integration"},
			Priority:    10,
			Handler:     handleWebhookSync,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleWebhookSync(ctx *hooks.HookContext) hooks.HookResult {
	source, _ := ctx.GetEnvData("source")

	fmt.Printf("[WebhookSyncPlugin] Syncing webhooks from source: %s\n", source)
	// 實際的 webhook 同步邏輯可以在這裡實現

	return ctx.SuccessWithMessage("Webhook sync completed successfully")
}

func init() {
	hooks.RegisterPluginType("WebhookSyncPlugin", func() hooks.Plugin {
		return &WebhookSyncPlugin{}
	})
}
