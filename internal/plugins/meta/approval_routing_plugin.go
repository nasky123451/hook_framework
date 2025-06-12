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
	hooks.New("submit_report").
		WithDescription("Handles submission of reports and routes them for approval").
		WithParamHints("doc_type").
		WithPriority(10).
		AllowRoles("auditor").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			docType, _ := ctx.GetEnvString("doc_type")

			fmt.Println("[ApprovalRoutingPlugin] Processing document of type:", docType)

			// 實際處理邏輯
			return ctx.SuccessWithMessage("Submitted and routed document of type: %s", docType)
		}).
		RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("ApprovalRoutingPlugin", func() hooks.Plugin {
		return &ApprovalRoutingPlugin{}
	})
}
