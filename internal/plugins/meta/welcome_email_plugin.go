// welcome_email_plugin.go
package meta

import (
	"fmt"

	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type WelcomeEmailPlugin struct{}

func (p *WelcomeEmailPlugin) Name() string {
	return "WelcomeEmailPlugin"
}

func (p *WelcomeEmailPlugin) GetHookNames() []string {
	return []string{"create_account"}
}

func (p *WelcomeEmailPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "create_account", 10, "admin", func(ctx *hooks.HookContext) hooks.HookResult {
		email, ok := ctx.GetEnvData("email").(string)
		if !ok || email == "" {
			return hooks.HookResult{Error: fmt.Errorf("email is required to send welcome email")}
		}

		message := fmt.Sprintf("Welcome email sent to %s", email)
		ctx.SetEnvData("approval_message", message)
		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("WelcomeEmailPlugin", func() hooks.Plugin {
		return &WelcomeEmailPlugin{}
	})
}
