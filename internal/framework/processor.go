package framework

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/internal/utils"
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

func (p *ClientInputProcessor) Process(clientInput ClientInput, h *hooks.HookManager) {
	p.Printer.PrintSection(fmt.Sprintf("Simulating Input: %s (Role: %s)", clientInput.Input, clientInput.Role))

	context := hooks.NewHookContext("system", map[string]interface{}{"origin": "main"})
	context.Reset()

	context.SetUserData("role", clientInput.Role)
	context.SetEnvData("role", clientInput.Role)

	for k, v := range clientInput.Context {
		context.SetEnvData(k, v)
	}

	if p.Env.HookManager != nil {
		if err := h.Execute(clientInput.Input, context, false); err != nil {
			p.Printer.PrintError(fmt.Errorf("Hook execution error: %w", err))
		}
	}

	p.PrintResult(context)
	p.Env.Contexts = append(p.Env.Contexts, context)
}

func (p *ClientInputProcessor) ProcessWithGraph(clientInput ClientInput, hg *hooks.HookGraph) {
	p.Printer.PrintSection(fmt.Sprintf("Simulating Input: %s (Role: %s)", clientInput.Input, clientInput.Role))

	context := hooks.NewHookContext("system", map[string]interface{}{"origin": "main"})
	context.Reset()

	context.SetUserData("role", clientInput.Role)
	context.SetEnvData("role", clientInput.Role)

	for k, v := range clientInput.Context {
		context.SetEnvData(k, v)
	}

	if p.Env.HookManager != nil {
		if err := hg.Execute(clientInput.Input, context); err != nil {
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
