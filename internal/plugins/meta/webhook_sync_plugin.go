package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type WebhookSyncPlugin struct{}

func (p *WebhookSyncPlugin) Name() string {
	return "WebhookSyncPlugin"
}

func (p *WebhookSyncPlugin) GetHookNames() []string {
	return []string{"webhook_sync"}
}

func (p *WebhookSyncPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "webhook_sync", 10, "integration", func(ctx *hooks.HookContext) hooks.HookResult {
		source := ctx.GetEnvData("source")

		message := fmt.Sprintf("[WebhookSyncPlugin] Received webhook from %v.\n", source)
		ctx.SetEnvData("approval_message", message)
		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("WebhookSyncPlugin", func() hooks.Plugin {
		return &WebhookSyncPlugin{}
	})
}
