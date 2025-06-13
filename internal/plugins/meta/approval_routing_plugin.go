package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
)

type ApprovalRoutingPlugin struct{}

func (p *ApprovalRoutingPlugin) Name() string {
	return "ApprovalRoutingPlugin"
}

func (p *ApprovalRoutingPlugin) GetHookNames() []string {
	return []string{"submit_report"}
}

func (p *ApprovalRoutingPlugin) RegisterHooks(hm *hooks.HookManager) {
	hookDefs := hooks.HookBuilders{
		{
			HookName:    "submit_report",
			Description: "Handles submission of reports and routes them for approval",
			ParamHints:  []string{"doc_type"},
			Roles:       []string{"auditor"},
			Priority:    10,
			Handler:     handleSubmitReport,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleSubmitReport(ctx *hooks.HookContext) hooks.HookResult {
	docType, _ := ctx.GetEnvString("doc_type")
	fmt.Println("[ApprovalRoutingPlugin] Processing document of type:", docType)
	// 實際處理邏輯
	return ctx.SuccessWithMessage("Submitted and routed document of type: %s", docType)
}

func init() {
	hooks.RegisterPluginType("ApprovalRoutingPlugin", func() hooks.Plugin {
		return &ApprovalRoutingPlugin{}
	})
}
