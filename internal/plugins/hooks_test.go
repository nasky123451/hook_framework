package plugins_test

import (
	"errors"
	"hook_framework/internal/hooks"
	"testing"
)

type SampleHandler struct{}

func (h *SampleHandler) Name() string {
	return "sample_handler"
}
func (h *SampleHandler) Priority() int {
	return 10
}
func (h *SampleHandler) Filter(ctx *hooks.HookContext) bool {
	return true
}
func (h *SampleHandler) Run(ctx *hooks.HookContext) hooks.HookResult {
	return hooks.HookResult{Name: h.Name(), Success: true}
}

func TestHookManager_RegisterAndExecute(t *testing.T) {
	hm := hooks.NewHookManager()
	runCounter := 0

	// 用 RegisterHook 註冊兩個 handler (priority 越小越先跑)
	hm.RegisterHook("test_hook", 20, func(ctx *hooks.HookContext) hooks.HookResult {
		runCounter++
		return hooks.HookResult{Name: "handler1", Success: true, Message: "handler1 run"}
	})
	hm.RegisterHook("test_hook", 10, func(ctx *hooks.HookContext) hooks.HookResult {
		runCounter++
		return hooks.HookResult{Name: "handler2", Success: true, Message: "handler2 run"}
	})

	ctx := hooks.NewHookContext("test_hook", nil)
	ctx.SetUserData("permissions", "admin")

	err := hm.Execute("test_hook", ctx, false)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if runCounter != 2 {
		t.Errorf("Expected 2 handlers to run, got %d", runCounter)
	}
}

func TestHookManager_StopExecution(t *testing.T) {
	hm := hooks.NewHookManager()
	runCounter := 0

	// 第一個 handler 會停止後續執行
	hm.RegisterHook("stop_hook", 10, func(ctx *hooks.HookContext) hooks.HookResult {
		runCounter++
		return hooks.HookResult{Name: "stopper", Success: true, StopExecution: true}
	})
	// 第二個 handler 不該被執行
	hm.RegisterHook("stop_hook", 20, func(ctx *hooks.HookContext) hooks.HookResult {
		runCounter++
		return hooks.HookResult{Name: "should_not_run", Success: true}
	})

	ctx := hooks.NewHookContext("stop_hook", nil)
	ctx.SetUserData("permissions", "admin")

	err := hm.Execute("stop_hook", ctx, false)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if runCounter != 1 {
		t.Errorf("Expected only 1 handler to run due to StopExecution, got %d", runCounter)
	}
}

func TestHookManager_HandlerError(t *testing.T) {
	hm := hooks.NewHookManager()

	// Handler 回傳錯誤
	hm.RegisterHook("error_hook", 10, func(ctx *hooks.HookContext) hooks.HookResult {
		return hooks.HookResult{
			Name:    "error_handler",
			Success: false,
			Error:   errors.New("simulated error"),
		}
	})

	ctx := hooks.NewHookContext("error_hook", nil)
	ctx.SetUserData("permissions", "admin")

	err := hm.Execute("error_hook", ctx, false)
	if err == nil {
		t.Error("Expected error from Execute, but got nil")
	}
}

func TestHookManager_GetRegisteredHooks(t *testing.T) {
	hm := hooks.NewHookManager()

	sh := &SampleHandler{}
	hm.RegisterHook("sample_hook", sh.Priority(), func(ctx *hooks.HookContext) hooks.HookResult {
		return sh.Run(ctx)
	})

	all := hm.GetRegisteredHooks()
	handlers, ok := all["sample_hook"]
	if !ok || len(handlers) == 0 {
		t.Fatalf("Expected sample_hook handlers registered, got none")
	}

	// 檢查是否至少有一個 handler 名稱為 sample_handler (或你註冊時 handler 的 Name())
	found := false
	for _, name := range handlers {
		if name == "sample_hook" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected handler named 'sample_hook' but got %v", handlers)
	}
}
