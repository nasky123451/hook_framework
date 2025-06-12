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
	hooks.New("login_failure_alert").
		WithDescription("Handles security alerts for login failures").
		WithParamHints("ip").
		WithPriority(10).
		AllowRoles("security").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			ip, _ := ctx.GetEnvString("ip")

			// todo: 實際的安全警報邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Security alert triggered from IP %s", ip)
		}).RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("SecurityAlertPlugin", func() hooks.Plugin {
		return &SecurityAlertPlugin{}
	})
}
