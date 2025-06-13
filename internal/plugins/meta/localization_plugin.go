package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
)

type LocalizationPlugin struct{}

func (p *LocalizationPlugin) Name() string {
	return "LocalizationPlugin"
}

func (p *LocalizationPlugin) GetHookNames() []string {
	return []string{"set_language"}
}

func (p *LocalizationPlugin) RegisterHooks(hm *hooks.HookManager) {
	hookDefs := hooks.HookBuilders{
		{HookName: "set_language",
			Description: "Handles setting the language for the user based on their preferences",
			ParamHints:  []string{"language"},
			Roles:       []string{"admin", "user"},
			Priority:    10,
			Handler:     handleSetLanguage,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleSetLanguage(ctx *hooks.HookContext) hooks.HookResult {
	language, _ := ctx.GetEnvString("language")

	fmt.Println("[LocalizationPlugin] Setting language to:", language)
	// 實際的語言設置邏輯可以在這裡實現

	return ctx.SuccessWithMessage("Language set to %s", language)
}

func init() {
	hooks.RegisterPluginType("LocalizationPlugin", func() hooks.Plugin {
		return &LocalizationPlugin{}
	})
}
