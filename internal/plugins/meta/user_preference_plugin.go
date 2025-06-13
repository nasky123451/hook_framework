package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
)

type UserPreferencePlugin struct{}

func (p *UserPreferencePlugin) Name() string {
	return "UserPreferencePlugin"
}

func (p *UserPreferencePlugin) GetHookNames() []string {
	return []string{"set_user_pref"}
}

func (p *UserPreferencePlugin) RegisterHooks(hm *hooks.HookManager) {
	hookDefs := hooks.HookBuilders{
		{HookName: "set_user_pref",
			Description: "Handles setting user preferences like theme, language, etc.",
			ParamHints:  []string{"theme"},
			Roles:       []string{"user"},
			Priority:    10,
			Handler:     handleSetUserPreference,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())

}

func handleSetUserPreference(ctx *hooks.HookContext) hooks.HookResult {
	theme, _ := ctx.GetEnvString("theme")

	fmt.Println("[UserPreferencePlugin] Setting user preference for theme:", theme)
	// 實際的用戶偏好設置邏輯可以在這裡實現

	return ctx.SuccessWithMessage("User preference updated successfully")
}

func init() {
	hooks.RegisterPluginType("UserPreferencePlugin", func() hooks.Plugin {
		return &UserPreferencePlugin{}
	})
}
