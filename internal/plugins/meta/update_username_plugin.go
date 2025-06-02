package meta

import (
	"fmt"
	"log"
	"strings"

	"hook_framework/internal/hooks"
	"hook_framework/pkg/nlp"
	"hook_framework/pkg/utils"
)

type UpdateUsernamePlugin struct {
	HookName string
}

func (p *UpdateUsernamePlugin) Name() string {
	return "UpdateUsernamePlugin"
}

func (p *UpdateUsernamePlugin) GetHookNames() []string {
	return []string{p.HookName}
}

func (p *UpdateUsernamePlugin) RegisterHooks(hookManager *hooks.HookManager) {
	if p.HookName == "" {
		log.Fatal("[UpdateUsernamePlugin] HookName is not set.")
	}

	utils.RegisterDynamicHook(hookManager, p.HookName, 15, "admin", func(ctx *hooks.HookContext) hooks.HookResult {
		text := ctx.GetString("input_text")
		username := nlp.ExtractName(text)
		if username == "" {
			return hooks.HookResult{Error: fmt.Errorf("no valid username found in input text")}
		}
		ctx.Set("username", username)
		return hooks.HookResult{}
	})
}

func (p *UpdateUsernamePlugin) RegisterParsers(nlpEngine *nlp.NLP) {
	if nlpEngine == nil {
		log.Fatal("[UpdateUsernamePlugin] NLP engine is nil.")
	}

	nlpEngine.RegisterParser(func(input string) (nlp.Intent, bool) {
		if strings.HasPrefix(input, "更新用戶名為 ") {
			username := strings.TrimPrefix(input, "更新用戶名為 ")
			return nlp.Intent{
				Action: "update_username",
				Params: map[string]string{"username": username},
			}, true
		}
		return nlp.Intent{}, false
	})
}

func init() {
	hooks.RegisterPluginType("UpdateUsernamePlugin", func() hooks.Plugin {
		return &UpdateUsernamePlugin{
			HookName: "update_username",
		}
	})
}
