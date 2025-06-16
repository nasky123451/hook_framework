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
	hookDefs := hooks.HookBuilders{
		{HookName: "book_room",
			Description: "Handles room booking and prevents conflicts in scheduling",
			ParamHints:  []string{"room", "time"},
			Permissions: "",
			Priority:    10,
			Handler:     handleBookRoom,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleBookRoom(ctx *hooks.HookContext) hooks.HookResult {
	room, _ := ctx.GetEnvString("room")
	time, _ := ctx.GetEnvString("time")

	fmt.Println("[BookingLockPlugin] Booking room:", room, "for time:", time)

	return ctx.SuccessWithMessage("Room %s booked for %s", room, time)
}

func init() {
	hooks.RegisterPluginType("BookingLockPlugin", func() hooks.Plugin {
		return &BookingLockPlugin{}
	})
}
