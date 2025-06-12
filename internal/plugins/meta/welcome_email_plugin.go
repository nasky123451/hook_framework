// welcome_email_plugin.go
package meta

import (
	"fmt"

	"hook_framework/internal/hooks"
)

type WelcomeEmailPlugin struct{}

func (p *WelcomeEmailPlugin) Name() string {
	return "WelcomeEmailPlugin"
}

func (p *WelcomeEmailPlugin) GetHookNames() []string {
	return []string{"create_account"}
}

func (p *WelcomeEmailPlugin) RegisterHooks(hm *hooks.HookManager) {
	hooks.New("create_account").
		WithDescription("Handles sending a welcome email when a new account is created").
		WithParamHints("email").
		WithPriority(10).
		AllowRoles("admin").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			email, ok := ctx.GetEnvData("email")
			if !ok || email == "" {
				return hooks.HookResult{Error: fmt.Errorf("email is required to send welcome email")}
			}

			fmt.Println("[WelcomeEmailPlugin] Sending welcome email to:", email)
			// todo: 實際的發送郵件邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Welcome email sent to %s", email)
		}).RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("WelcomeEmailPlugin", func() hooks.Plugin {
		return &WelcomeEmailPlugin{}
	})
}
