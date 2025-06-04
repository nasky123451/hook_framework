package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type ApprovalRoutingPlugin struct{}

func (p *ApprovalRoutingPlugin) Name() string {
	return "ApprovalRoutingPlugin"
}

func (p *ApprovalRoutingPlugin) GetHookNames() []string {
	return []string{"submit_report"}
}

func (p *ApprovalRoutingPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "submit_report", 10, "auditor", func(ctx *hooks.HookContext) hooks.HookResult {
		docType, _ := ctx.GetEnvData("doc_type").(string)

		message := fmt.Sprintf("Submitted and routed document of type: %s", docType)
		ctx.SetEnvData("approval_message", message)
		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("ApprovalRoutingPlugin", func() hooks.Plugin {
		return &ApprovalRoutingPlugin{}
	})
}
