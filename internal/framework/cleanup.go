package framework

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
	"time"
)

func PrintStats(env *hooks.HookEnvironment, printer *utils.Printer, initializedPlugins []hooks.Plugin) {
	// 獲取 Hook Stats
	commonStats, version := env.HookManager.GetStats()

	// 將 common.HookStats 轉換為 hooks.HookStats
	hookStats := make(map[string]*hooks.HookStats)
	for name, stat := range commonStats {
		hookStats[name] = &hooks.HookStats{
			ExecutionCount: stat.ExecutionCount,
			TotalDuration:  stat.TotalDuration,
			LastExecutionResult: hooks.HookResult{ // 使用 hooks.HookResult
				StopExecution: stat.LastExecutionResult.StopExecution, // 修正字段名稱
				Error:         stat.LastExecutionResult.Error,         // 修正字段名稱
			},
		}
	}

	// 輸出統計數據
	printer.PrintMessage(fmt.Sprintf("Stats Version: %d", version.Version))
	printer.PrintMessage(fmt.Sprintf("Stats Timestamp: %s", version.Timestamp.Format(time.RFC3339)))
	printer.PrintHookStats(hookStats)
}
