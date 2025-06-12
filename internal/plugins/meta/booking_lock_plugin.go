package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
)

type BookingLockPlugin struct{}

func (p *BookingLockPlugin) Name() string {
	return "BookingLockPlugin"
}

func (p *BookingLockPlugin) GetHookNames() []string {
	return []string{"book_room"}
}

func (p *BookingLockPlugin) RegisterHooks(hm *hooks.HookManager) {
	hooks.New("book_room").
		WithDescription("Handles room booking and prevents conflicts in scheduling").
		WithParamHints("room", "time").
		WithPriority(10).
		AllowRoles("admin", "employee").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			room, _ := ctx.GetEnvString("room")
			time, _ := ctx.GetEnvString("time")

			fmt.Println("[BookingLockPlugin] Booking room:", room, "for time:", time)
			// TODO: 實際的房間預訂邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Room %s booked for %s", room, time)
		}).
		RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("BookingLockPlugin", func() hooks.Plugin {
		return &BookingLockPlugin{}
	})
}
