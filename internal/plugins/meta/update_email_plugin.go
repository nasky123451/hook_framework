package meta

import (
	"fmt"
	"log"
	"strings"

	"hook_framework/internal/hooks"
	"hook_framework/pkg/nlp"
	"hook_framework/pkg/utils"
)

type UpdateEmailPlugin struct {
	HookName string
}

func (p *UpdateEmailPlugin) Name() string {
	return "UpdateEmailPlugin"
}

func (p *UpdateEmailPlugin) GetHookNames() []string {
	return []string{p.HookName}
}

func (p *UpdateEmailPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	if p.HookName == "" {
		log.Fatal("[UpdateEmailPlugin] HookName is not set.")
	}

	utils.RegisterDynamicHook(hookManager, p.HookName, 15, "admin", func(ctx *hooks.HookContext) hooks.HookResult {
		email := ctx.GetString("email")
		if email == "" {
			return hooks.HookResult{Error: fmt.Errorf("email is missing")}
		}
		return hooks.HookResult{}
	})
}

func (p *UpdateEmailPlugin) RegisterParsers(nlpEngine *nlp.NLP) {
	if nlpEngine == nil {
		log.Fatal("[UpdateEmailPlugin] NLP engine is nil.")
	}

	nlpEngine.RegisterParser(func(input string) (nlp.Intent, bool) {
		if strings.HasPrefix(input, "修改 email 為 ") {
			email := strings.TrimPrefix(input, "修改 email 為 ")
			return nlp.Intent{
				Action: "update_email",
				Params: map[string]string{"email": email},
			}, true
		}
		return nlp.Intent{}, false
	})
}

func init() {
	hooks.RegisterPluginType("UpdateEmailPlugin", func() hooks.Plugin {
		return &UpdateEmailPlugin{
			HookName: "update_email",
		}
	})
}
