package hooks

import (
	"fmt"
	"sort"
	"time"
)

func (hm *HookManager) Execute(name string, ctx *HookContext, async bool) error {
	handlers := hm.GetHookDefinitionByName(name)
	if len(handlers) == 0 {
		return nil
	}

	if !hm.CheckPermissions(ctx, handlers) {
		permRaw := ctx.GetUserData("permissions")
		perm, ok := permRaw.(string)
		if !ok {
			perm = ""
		}
		return fmt.Errorf("permission denied for hook %s and role %v", name, perm)
	}

	result := hm.ExecuteHookByName(name, ctx)
	ctx.AddExecutionLog(result)

	if result.StopExecution {
		fmt.Printf("[HookGraph] Hook %s 停止後續流程\n", name)
	}

	// *** 新增：如果有錯誤，回傳錯誤 ***
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (hm *HookManager) ExecuteHookByName(name string, ctx *HookContext) HookResult {
	hm.mu.RLock()
	handlers := hm.hooks[name]
	hm.mu.RUnlock()

	if len(handlers) == 0 {
		return HookResult{
			Name:    name,
			Success: false,
			Error:   fmt.Errorf("no handlers registered for hook '%s'", name),
		}
	}

	// 優先順序排序：小數值優先執行
	sorted := make([]HookHandler, len(handlers))
	copy(sorted, handlers)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Priority() < sorted[j].Priority()
	})

	var finalResult HookResult
	perm, _ := ctx.EnvData["permissions"].(string)
	finalResult.Permissions = perm
	finalResult.Name = name
	finalResult.Success = true

	start := time.Now()
	finalResult.DateTime = start
	for _, handler := range sorted {
		if !handler.Filter(ctx) {
			continue
		}
		retryCfg := RetryConfig{MaxRetries: 3, Backoff: time.Millisecond * 100}
		result := executeWithRetry(ctx, handler, retryCfg)
		finalResult.Message = result.Message
		if result.Error != nil {
			finalResult.Success = false
			finalResult.Error = result.Error
		}
		if result.StopExecution {
			finalResult.StopExecution = true
			break
		}
	}
	duration := time.Since(start)
	finalResult.Duration = duration

	hm.updateStats(name, duration, HookResult{
		StopExecution: ctx.IsStopped(),
		Error:         nil,
	})

	return finalResult
}
