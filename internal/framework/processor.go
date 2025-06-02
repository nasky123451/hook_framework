package framework

import (
	"fmt"
	"hook_framework/internal/hooks"
	"hook_framework/pkg/nlp"
	"hook_framework/pkg/utils"
)

type ClientInputProcessor struct {
	Env       *hooks.HookEnvironment
	NLPEngine *nlp.NLP
	Printer   *utils.Printer
}

func NewClientInputProcessor(env *hooks.HookEnvironment, nlpEngine *nlp.NLP, printer *utils.Printer) *ClientInputProcessor {
	return &ClientInputProcessor{
		Env:       env,
		NLPEngine: nlpEngine,
		Printer:   printer,
	}
}

func (p *ClientInputProcessor) Process(clientInput struct{ Input, Role string }) {
	p.Printer.PrintSection(fmt.Sprintf("Simulating Input: %s (Role: %s)", clientInput.Input, clientInput.Role))

	// 重置上下文的錯誤和停止狀態
	p.Env.Context.Reset()

	// 設置角色到上下文
	p.Env.Context.Set("role", clientInput.Role)

	// 解析輸入
	intent := p.NLPEngine.ParseInput(clientInput.Input)
	if intent.Action == "" {
		p.Printer.PrintMessage("[Process] No action detected from input.")
		return
	}

	params := utils.ConvertParamsToInterface(intent.Params)

	// 添加輸入文本到參數
	params["input_text"] = clientInput.Input

	// 通過 DispatchInput 處理操作
	hooks.DispatchInput(intent.Action, params, p.Env.Context, p.Env.HookManager)

	// 檢查是否有錯誤
	if len(p.Env.Context.Errors) > 0 {
		p.Printer.PrintMessage("[Process] Errors occurred during processing:")
		for _, err := range p.Env.Context.Errors {
			p.Printer.PrintError(err)
		}
	}
}
