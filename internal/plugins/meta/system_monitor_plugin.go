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
	hookDefs := hooks.HookBuilders{
		{HookName: "system_monitor",
			Description: "Handles system monitoring alerts",
			ParamHints:  []string{"server"},
			Permissions: "",
			Priority:    10,
			Handler:     handleSystemMonitor,
		},
	}

	hookDefs.RegisterHookDefinitions(hm, p.Name())
}

func handleSystemMonitor(ctx *hooks.HookContext) hooks.HookResult {
	server, _ := ctx.GetEnvData("server")

	// 實際的監控邏輯可以在這裡實現

	return ctx.SuccessWithMessage("Monitoring alert on server: %s", server)
}

func init() {
	hooks.RegisterPluginType("SystemMonitorPlugin", func() hooks.Plugin {
		return &SystemMonitorPlugin{}
	})
}
