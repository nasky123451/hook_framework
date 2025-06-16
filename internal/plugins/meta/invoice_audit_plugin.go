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
	hookDefs := hooks.HookBuilders{
		{HookName: "create_invoice",
			Description: "Handles invoice creation and audits the invoice details",
			ParamHints:  []string{"invoice_no", "amount"},
			Permissions: "",
			Priority:    10,
			Handler:     handleCreateInvoice,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleCreateInvoice(ctx *hooks.HookContext) hooks.HookResult {
	invoiceNo, _ := ctx.GetEnvString("invoice_no")
	amount, _ := ctx.GetEnvString("amount")

	fmt.Printf("[InvoiceAuditPlugin] Auditing invoice: %s with amount: %s\n", invoiceNo, amount)
	// 實際的審核邏輯可以在這裡實現

	return ctx.SuccessWithMessage("Invoice %s audited successfully with amount %s", invoiceNo, amount)
}

func init() {
	hooks.RegisterPluginType("InvoiceAuditPlugin", func() hooks.Plugin {
		return &InvoiceAuditPlugin{}
	})
}
