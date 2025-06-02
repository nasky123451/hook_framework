package meta

import (
	"fmt"
	"log"

	"hook_framework/internal/hooks"
	"hook_framework/pkg/nlp"
	"hook_framework/pkg/utils"
)

type DeleteProtectionPlugin struct {
	BeforeDeleteHook string
}

func (p *DeleteProtectionPlugin) Name() string {
	return "DeleteProtectionPlugin"
}

func (p *DeleteProtectionPlugin) GetHookNames() []string {
	return []string{p.BeforeDeleteHook}
}

func (p *DeleteProtectionPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	if p.BeforeDeleteHook == "" {
		log.Fatal("[DeleteProtectionPlugin] Hook name 'BeforeDeleteHook' is not initialized.")
	}

	utils.RegisterDynamicHook(hookManager, p.BeforeDeleteHook, 50, "admin", func(ctx *hooks.HookContext) hooks.HookResult {
		fmt.Println("[BeforeDeleteHook] Checking if data is critical...")

		if ctx.Get("data") == "critical" {
			fmt.Println("[BeforeDeleteHook] Deletion blocked: Critical data detected.")
			return hooks.HookResult{
				StopExecution: true,
				Error:         fmt.Errorf("cannot delete critical data"),
			}
		}

		fmt.Println("[BeforeDeleteHook] Data is safe to delete.")
		return hooks.HookResult{}
	})
}

func (p *DeleteProtectionPlugin) RegisterParsers(nlpEngine *nlp.NLP) {
	if nlpEngine == nil {
		log.Fatal("[DeleteProtectionPlugin] NLP engine is nil.")
	}

	nlpEngine.RegisterParser(func(input string) (nlp.Intent, bool) {
		if input == "刪除關鍵數據" {
			return nlp.Intent{
				Action: "delete_protection",
			}, true
		}
		return nlp.Intent{}, false
	})
}

func init() {
	hooks.RegisterPluginType("DeleteProtectionPlugin", func() hooks.Plugin {
		return &DeleteProtectionPlugin{
			BeforeDeleteHook: "before_delete",
		}
	})
}
