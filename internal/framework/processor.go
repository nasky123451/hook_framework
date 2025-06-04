package framework

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/utils"
)

type ClientInput struct {
	Input   string
	Role    string
	Context map[string]interface{}
}

type ClientInputProcessor struct {
	Env     *hooks.HookEnvironment
	Printer *utils.Printer
}

func NewClientInputProcessor(env *hooks.HookEnvironment, printer *utils.Printer) *ClientInputProcessor {
	return &ClientInputProcessor{
		Env:     env,
		Printer: printer,
	}
}

func (p *ClientInputProcessor) ProcessWithContext(clientInput ClientInput) {
	p.Printer.PrintSection(fmt.Sprintf("Simulating Input: %s (Role: %s)", clientInput.Input, clientInput.Role))

	p.Env.Context.Reset()

	p.Env.Context.SetUserData("role", clientInput.Role)

	for k, v := range clientInput.Context {
		p.Env.Context.SetEnvData(k, v)
	}

	// 取消 NLP，直接以 clientInput.Input 當作 hookName
	hookName := clientInput.Input

	// params 可以放一些 context 裡的環境變數，或直接空參數
	params := map[string]interface{}{
		"input_text": clientInput.Input,
	}

	hooks.DispatchInput(hookName, params, p.Env.Context, p.Env.HookManager)

	p.PrintResult(p.Env.Context)
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
