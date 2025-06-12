package framework

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/internal/utils"
	"time"
)

func PrintStats(env *hooks.HookEnvironment, printer *utils.Printer) {
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
	fmt.Println("Hook Execution Statistics:")
	for hookName, stat := range hookStats {
		printer.PrintKeyValue("Hook", hookName)
		printer.PrintKeyValue("  Execution Count", stat.ExecutionCount)
		printer.PrintKeyValue("  Total Duration", stat.TotalDuration)
		if stat.LastExecutionResult.Name != "" {
			printer.PrintKeyValue("  Last Execution Result", fmt.Sprintf(
				"Name: %s, Success: %t, Error: %v, Duration: %v",
				stat.LastExecutionResult.Name,
				stat.LastExecutionResult.Success,
				stat.LastExecutionResult.Error,
				stat.LastExecutionResult.Duration,
			))
		}
	}
	printer.PrintDivider()

	// 輸出 HookContext 詳細結果
	hooks.PrintExecutionSummary(printer, env, hooks.ReportOptions{
		OnlyFailed:      false,
		FilterHookNames: []string{}, // 空代表全部
		OutputMarkdown:  "hook_execution_report.md",
	})
}
