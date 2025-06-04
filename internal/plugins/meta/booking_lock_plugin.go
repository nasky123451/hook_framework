package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type BookingLockPlugin struct{}

func (p *BookingLockPlugin) Name() string {
	return "BookingLockPlugin"
}

func (p *BookingLockPlugin) GetHookNames() []string {
	return []string{"book_room"}
}

func (p *BookingLockPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "book_room", 10, "employee", func(ctx *hooks.HookContext) hooks.HookResult {
		room, _ := ctx.GetEnvData("room").(string)
		time, _ := ctx.GetEnvData("time").(string)

		message := fmt.Sprintf("Room %s booked for %s", room, time)
		ctx.SetEnvData("approval_message", message)
		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("BookingLockPlugin", func() hooks.Plugin {
		return &BookingLockPlugin{}
	})
}
