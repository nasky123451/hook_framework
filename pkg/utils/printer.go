package utils

import (
	"fmt"
	"hook_framework/internal/hooks"
)

type Printer struct{}

func NewPrinter() *Printer {
	return &Printer{}
}

func (p *Printer) PrintSection(title string) {
	fmt.Printf("\n========== %s ==========\n", title)
	fmt.Println("----------------------------------------")
}

func (p *Printer) PrintMessage(message string) {
	fmt.Println(message)
}

func (p *Printer) PrintError(err error) {
	fmt.Printf("[Error] %v\n", err)
}

func (p *Printer) PrintKeyValue(key string, value interface{}) {
	fmt.Printf("%s: %v\n", key, value)
}

func (p *Printer) PrintDivider() {
	fmt.Println("----------------------------------------")
}

func (p *Printer) PrintHookStats(stats map[string]*hooks.HookStats) {
	fmt.Println("Hook Execution Statistics:")
	for hookName, stat := range stats {
		p.PrintKeyValue("Hook", hookName)
		p.PrintKeyValue("  Execution Count", stat.ExecutionCount)
		p.PrintKeyValue("  Total Duration", stat.TotalDuration)
		if stat.LastExecutionResult.Name != "" {
			p.PrintKeyValue("  Last Execution Result", fmt.Sprintf(
				"Name: %s, Success: %t, Error: %v, Duration: %v",
				stat.LastExecutionResult.Name,
				stat.LastExecutionResult.Success,
				stat.LastExecutionResult.Error,
				stat.LastExecutionResult.Duration,
			))
		}
	}
	p.PrintDivider()
}
