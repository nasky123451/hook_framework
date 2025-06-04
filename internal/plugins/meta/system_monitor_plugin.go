package meta

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type SystemMonitorPlugin struct{}

func (p *SystemMonitorPlugin) Name() string {
	return "SystemMonitorPlugin"
}

func (p *SystemMonitorPlugin) GetHookNames() []string {
	return []string{"system_monitor"}
}

func (p *SystemMonitorPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	utils.RegisterDynamicHook(hookManager, "system_monitor", 10, "devops", func(ctx *hooks.HookContext) hooks.HookResult {
		server := ctx.GetEnvData("server")

		message := fmt.Sprintf("[SystemMonitorPlugin] Monitoring alert on server: %v.\n", server)
		ctx.SetEnvData("approval_message", message)
		return hooks.HookResult{Success: true}
	})
}

func init() {
	hooks.RegisterPluginType("SystemMonitorPlugin", func() hooks.Plugin {
		return &SystemMonitorPlugin{}
	})
}
