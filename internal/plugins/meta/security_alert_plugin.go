package meta

import (
	"hook_framework/internal/hooks"
)

type SecurityAlertPlugin struct{}

func (p *SecurityAlertPlugin) Name() string {
	return "SecurityAlertPlugin"
}

func (p *SecurityAlertPlugin) GetHookNames() []string {
	return []string{"login_failure_alert"}
}

func (p *SecurityAlertPlugin) RegisterHooks(hm *hooks.HookManager) {
	hookDefs := hooks.HookBuilders{
		{HookName: "login_failure_alert",
			Description: "Handles security alerts for login failures",
			ParamHints:  []string{"ip"},
			Permissions: "",
			Priority:    10,
			Handler:     handleLoginFailureAlert,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleLoginFailureAlert(ctx *hooks.HookContext) hooks.HookResult {
	ip, _ := ctx.GetEnvString("ip")

	// 實際的安全警報邏輯可以在這裡實現

	return ctx.SuccessWithMessage("Security alert triggered from IP %s", ip)
}

func init() {
	hooks.RegisterPluginType("SecurityAlertPlugin", func() hooks.Plugin {
		return &SecurityAlertPlugin{}
	})
}
