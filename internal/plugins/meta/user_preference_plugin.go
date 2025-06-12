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
	hooks.New("set_user_pref").
		WithDescription("Handles setting user preferences like theme, language, etc.").
		WithParamHints("theme").
		WithPriority(10).
		AllowRoles("user").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			theme, _ := ctx.GetEnvString("theme")

			fmt.Println("[UserPreferencePlugin] Setting user preference for theme:", theme)
			// todo: 實際的用戶偏好設置邏輯可以在這裡實現

			return ctx.SuccessWithMessage("User preference updated successfully")
		}).RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("UserPreferencePlugin", func() hooks.Plugin {
		return &UserPreferencePlugin{}
	})
}
