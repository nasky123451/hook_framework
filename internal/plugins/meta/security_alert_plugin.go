package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type SecurityAlertPlugin struct{}

func (p *SecurityAlertPlugin) Name() string {
	return "SecurityAlertPlugin"
}

func (p *SecurityAlertPlugin) GetHookNames() []string {
	return []string{"login_failure_alert"}
}

func (p *SecurityAlertPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "login_failure_alert", 10, "security", func(ctx *hooks.HookContext) hooks.HookResult {
		ip, _ := ctx.GetEnvData("ip").(string)

		message := fmt.Sprintf("Security alert triggered from IP %s", ip)
		ctx.SetEnvData("approval_message", message)
		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("SecurityAlertPlugin", func() hooks.Plugin {
		return &SecurityAlertPlugin{}
	})
}
