package framework

import (
	"fmt"
	"hook_framework/internal/hooks"
	"time"
)

func PrintStats(env *hooks.HookEnvironment, printer *Printer) {
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

	// 輸出 HookContext 詳細結果
	// printer.PrintMessage("HookContext Executions:")
	// for _, ctx := range env.HookManager.GetContexts() {
	// 	role, _ := ctx.GetEnvString("role")
	// 	for _, log := range ctx.GetExecutionLog() {
	// 		printer.PrintMessage(fmt.Sprintf("  DateTime      : %s", log.DateTime.Format(time.RFC3339)))
	// 		printer.PrintMessage(fmt.Sprintf("  Input         : %s", log.Name))
	// 		printer.PrintMessage(fmt.Sprintf("  Role          : %s", role))
	// 		printer.PrintMessage(fmt.Sprintf("  Success       : %v", log.Success))
	// 		printer.PrintMessage(fmt.Sprintf("  Duration (µs) : %d", log.Duration))
	// 		printer.PrintMessage(fmt.Sprintf("  StopExecution : %v", log.StopExecution))
	// 		printer.PrintMessage(fmt.Sprintf("  Error         : %v", log.Error))
	// 		printer.PrintMessage(fmt.Sprintf("  Message       : %s", log.Message))
	// 		printer.PrintMessage("---------------------------------------")
	// 	}
	// }
}
