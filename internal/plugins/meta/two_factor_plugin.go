package meta

import (
	"fmt"
	"log"

	"hook_framework/internal/hooks"
	"hook_framework/pkg/nlp"
	"hook_framework/pkg/utils"
)

type TwoFactorPlugin struct {
	HookName string
}

func (p *TwoFactorPlugin) Name() string {
	return "TwoFactorPlugin"
}

func (p *TwoFactorPlugin) GetHookNames() []string {
	return []string{p.HookName}
}

func (p *TwoFactorPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	if p.HookName == "" {
		log.Fatal("[TwoFactorPlugin] HookName is not set.")
	}

	utils.RegisterDynamicHook(hookManager, p.HookName, 25, "admin", func(ctx *hooks.HookContext) hooks.HookResult {
		user, ok := ctx.Get("user").(map[string]string)
		if !ok || user["email"] == "" {
			return hooks.HookResult{Error: fmt.Errorf("no valid user found to enable two-factor authentication")}
		}
		ctx.Set("two_factor_auth_enabled", true)

		return hooks.HookResult{}
	})
}

func (p *TwoFactorPlugin) RegisterParsers(nlpEngine *nlp.NLP) {
	if nlpEngine == nil {
		log.Fatal("[TwoFactorPlugin] NLP engine is nil.")
	}

	nlpEngine.RegisterParser(func(input string) (nlp.Intent, bool) {
		if input == "啟用兩步驗證" {
			return nlp.Intent{
				Action: "enable_two_factor",
				Params: nil,
			}, true
		}
		return nlp.Intent{}, false
	})
}

func init() {
	hooks.RegisterPluginType("TwoFactorPlugin", func() hooks.Plugin {
		return &TwoFactorPlugin{
			HookName: "enable_two_factor",
		}
	})
}
