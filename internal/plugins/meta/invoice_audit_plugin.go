package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
)

type InvoiceAuditPlugin struct{}

func (p *InvoiceAuditPlugin) Name() string {
	return "InvoiceAuditPlugin"
}

func (p *InvoiceAuditPlugin) GetHookNames() []string {
	return []string{"create_invoice"}
}

func (p *InvoiceAuditPlugin) RegisterHooks(hm *hooks.HookManager) {
	hooks.New("create_invoice").
		WithDescription("Handles invoice creation and audits the invoice details").
		WithParamHints("invoice_no", "amount").
		WithPriority(10).
		AllowRoles("admin", "finance").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			invoiceNo, _ := ctx.GetEnvData("invoice_no")
			amount, _ := ctx.GetEnvData("amount")

			fmt.Printf("[InvoiceAuditPlugin] Auditing invoice: %s with amount: %s\n", invoiceNo, amount)
			// todo: 實際的審核邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Invoice %s audited successfully with amount %s", invoiceNo, amount)
		}).RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("InvoiceAuditPlugin", func() hooks.Plugin {
		return &InvoiceAuditPlugin{}
	})
}
