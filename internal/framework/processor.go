package framework

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/internal/utils"
	"time"
)

type ClientInput struct {
	Input   string
	Role    string
	Context map[string]interface{}
}

type ClientInputProcessor struct {
	Env      *hooks.HookEnvironment
	Printer  *utils.Printer
	Metadata []hooks.HookMetadata // 用於存儲 Hook 元數據
}

func NewClientInputProcessor(env *hooks.HookEnvironment, printer *utils.Printer) *ClientInputProcessor {
	return &ClientInputProcessor{
		Env:     env,
		Printer: printer,
	}
}

func (p *ClientInputProcessor) Process(clientInput ClientInput) {
	p.Printer.PrintSection(fmt.Sprintf("Simulating Input: %s (Role: %s)", clientInput.Input, clientInput.Role))

	context := hooks.NewHookContext("system", map[string]interface{}{"origin": "main"})
	context.Reset()

	context.SetUserData("role", clientInput.Role)
	context.SetEnvData("role", clientInput.Role)

	for k, v := range clientInput.Context {
		context.SetEnvData(k, v)
	}

	if p.Env.HookManager != nil {
		if err := p.Env.HookManager.Execute(clientInput.Input, context, false); err != nil {
			p.Printer.PrintError(fmt.Errorf("Hook execution error: %w", err))
		}
	}

	p.PrintResult(context)
	p.Env.Contexts = append(p.Env.Contexts, context)
}

func (p *ClientInputProcessor) ProcessWithGraph(clientInput ClientInput) {
	p.Printer.PrintSection(fmt.Sprintf("Simulating Input: %s (Role: %s)", clientInput.Input, clientInput.Role))

	context := hooks.NewHookContext("system", map[string]interface{}{"origin": "main"})
	context.Reset()

	context.SetUserData("role", clientInput.Role)
	context.SetEnvData("role", clientInput.Role)

	for k, v := range clientInput.Context {
		context.SetEnvData(k, v)
	}

	if p.Env.HookManager != nil {
		if err := p.Env.HookGraph.Execute(clientInput.Input, context); err != nil {
			p.Printer.PrintError(fmt.Errorf("HookGraph execution error: %w", err))
		}
	}

	p.PrintResult(context)
	p.Env.Contexts = append(p.Env.Contexts, context)
}

func (p *ClientInputProcessor) PrintResult(ctx *hooks.HookContext) {
	if len(ctx.Results) > 0 {
		p.Printer.PrintMessage("[Result] Execution Results:")
		for i, result := range ctx.Results {
			p.Printer.PrintMessage(fmt.Sprintf("  #%d: %v", i+1, result))
		}
	}

	if len(ctx.Errors) > 0 {
		p.Printer.PrintMessage("[Error] Errors Occurred:")
		for i, err := range ctx.Errors {
			p.Printer.PrintError(fmt.Errorf("  #%d: %v", i+1, err))
		}
	}
}

func (p *ClientInputProcessor) PrintStats() {
	// 獲取 Hook Stats
	commonStats, version := p.Env.HookManager.GetStats()

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
	p.Printer.PrintMessage(fmt.Sprintf("Stats Version: %d", version.Version))
	p.Printer.PrintMessage(fmt.Sprintf("Stats Timestamp: %s", version.Timestamp.Format(time.RFC3339)))
	fmt.Println("Hook Execution Statistics:")
	for hookName, stat := range hookStats {
		p.Printer.PrintKeyValue("Hook", hookName)
		p.Printer.PrintKeyValue("  Execution Count", stat.ExecutionCount)
		p.Printer.PrintKeyValue("  Total Duration", stat.TotalDuration)
		if stat.LastExecutionResult.Name != "" {
			p.Printer.PrintKeyValue("  Last Execution Result", fmt.Sprintf(
				"Name: %s, Success: %t, Error: %v, Duration: %v",
				stat.LastExecutionResult.Name,
				stat.LastExecutionResult.Success,
				stat.LastExecutionResult.Error,
				stat.LastExecutionResult.Duration,
			))
		}
	}
	p.Printer.PrintDivider()

	// 輸出 HookContext 詳細結果
	hooks.PrintExecutionSummary(p.Printer, p.Env, hooks.ReportOptions{
		OnlyFailed:      false,
		FilterHookNames: []string{}, // 空代表全部
		OutputMarkdown:  "hook_execution_report.md",
	})
}
