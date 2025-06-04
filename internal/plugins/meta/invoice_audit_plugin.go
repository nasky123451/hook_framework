package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type InvoiceAuditPlugin struct{}

func (p *InvoiceAuditPlugin) Name() string {
	return "InvoiceAuditPlugin"
}

func (p *InvoiceAuditPlugin) GetHookNames() []string {
	return []string{"create_invoice"}
}

func (p *InvoiceAuditPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "create_invoice", 10, "finance", func(ctx *hooks.HookContext) hooks.HookResult {
		invoiceNo := ctx.GetEnvData("invoice_no")
		amount := ctx.GetEnvData("amount")

		message := fmt.Sprintf("[InvoiceAuditPlugin] Auditing invoice %v with amount %v.\n", invoiceNo, amount)
		ctx.SetEnvData("approval_message", message)
		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("InvoiceAuditPlugin", func() hooks.Plugin {
		return &InvoiceAuditPlugin{}
	})
}
