package meta

import (
	"hook_framework/internal/hooks"
)

type SystemMonitorPlugin struct{}

func (p *SystemMonitorPlugin) Name() string {
	return "SystemMonitorPlugin"
}

func (p *SystemMonitorPlugin) GetHookNames() []string {
	return []string{"system_monitor"}
}

func (p *SystemMonitorPlugin) RegisterHooks(hm *hooks.HookManager) {
	hooks.New("system_monitor").
		WithDescription("Handles system monitoring alerts").
		WithParamHints("server").
		WithPriority(10).
		AllowRoles("devops").
		Handle(func(ctx *hooks.HookContext) hooks.HookResult {
			server, _ := ctx.GetEnvData("server")

			//todo: 實際的監控邏輯可以在這裡實現

			return ctx.SuccessWithMessage("Monitoring alert on server: %s", server)
		}).RegisterTo(hm)
}

func init() {
	hooks.RegisterPluginType("SystemMonitorPlugin", func() hooks.Plugin {
		return &SystemMonitorPlugin{}
	})
}
