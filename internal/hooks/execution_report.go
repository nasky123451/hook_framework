package hooks

import (
	"fmt"
	"hook_framework/internal/utils"
	"os"
	"strings"
	"time"

	// 替換成你實際的 utils 路徑

	"github.com/fatih/color"
)

// ReportOptions 控制顯示選項
type ReportOptions struct {
	OnlyFailed      bool     // 僅顯示失敗
	FilterHookNames []string // 指定要顯示的 Hook 名稱（空表示全顯）
	OutputMarkdown  string   // 若指定檔名，則輸出 Markdown 檔案
}

// PrintExecutionSummary 輸出 Hook 執行摘要（可搭配條件與檔案輸出）
func PrintExecutionSummary(printer *utils.Printer, env *HookEnvironment, opts ReportOptions) {
	hookLogsByName := make(map[string][]*HookResult)

	// 分群
	for _, wrappers := range env.ContextManager.contexts {
		for _, wrapper := range wrappers {
			logs := wrapper.Context.GetExecutionLog()
			for _, log := range logs {
				if opts.OnlyFailed && log.Success {
					continue
				}
				if len(opts.FilterHookNames) > 0 && !contains(opts.FilterHookNames, log.Name) {
					continue
				}
				// 注意這裡是取址 &log，要先轉成變數避免 range 變動造成所有指標相同
				logCopy := log
				hookLogsByName[log.Name] = append(hookLogsByName[log.Name], &logCopy)
			}
		}
	}

	// 終端輸出
	printer.PrintMessage("🎯 Hook Execution Summary")
	for name, logs := range hookLogsByName {
		blue := color.New(color.FgHiCyan).SprintFunc()
		printer.PrintMessage(fmt.Sprintf("\n%s %s", blue("🔧 Hook:"), name))
		printer.PrintMessage(strings.Repeat("=", 40))
		for _, log := range logs {
			t := log.DateTime.Format("2006-01-02 15:04:05")
			permissions := log.Permissions
			colorizer := color.New(color.FgGreen)
			if !log.Success {
				colorizer = color.New(color.FgRed)
			}
			printer.PrintMessage(colorizer.Sprintf("  [%s] %s | Success: %v | Duration: %v", t, permissions, log.Success, log.Duration))
			printer.PrintMessage("    Message:" + log.Message)
			if log.Error != nil {
				printer.PrintMessage("    Error  :" + log.Error.Error())
			}
			printer.PrintMessage("    --------")
		}
	}

	// 可選：輸出 Markdown
	if opts.OutputMarkdown != "" {
		var lines []string
		lines = append(lines, "# Hook Execution Report", fmt.Sprintf("> Generated at %s", time.Now().Format(time.RFC3339)), "")
		for name, logs := range hookLogsByName {
			lines = append(lines, fmt.Sprintf("## %s", name), "")
			for _, log := range logs {
				lines = append(lines,
					fmt.Sprintf("- 🕒 Time: `%s`", log.DateTime.Format(time.RFC3339)),
					fmt.Sprintf("- ✅ Success: `%v`", log.Success),
					fmt.Sprintf("- ⏱ Duration: `%v`", log.Duration),
					fmt.Sprintf("- 👤 Permissions: `%s`", log.Permissions),
					fmt.Sprintf("- 💬 Message: %s", log.Message),
				)
				if log.Error != nil {
					lines = append(lines, fmt.Sprintf("- ❌ Error: `%v`", log.Error.Error()))
				}
				lines = append(lines, "")
			}
		}
		_ = os.WriteFile(opts.OutputMarkdown, []byte(strings.Join(lines, "\n")), 0644)
	}
}

// contains 檢查 string slice 中是否包含某值
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
