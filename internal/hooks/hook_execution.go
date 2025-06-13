package hooks

import (
	"fmt"
	"sort"
	"time"
)

func (hm *HookManager) Execute(name string, ctx *HookContext, async bool) error {

	// 權限檢查（基於 ctx 的 role）
	if !hm.CheckPermission(name, ctx) {
		return fmt.Errorf("permission denied for hook %s and role %s", name, ctx.GetUserData("role"))
	}

	result := hm.ExecuteHookByName(name, ctx)
	ctx.AddExecutionLog(result) // 紀錄流程與結果

	if result.StopExecution {
		fmt.Printf("[HookGraph] Hook %s 停止後續流程\n", name)
		return nil
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
	finalResult.Role = ctx.EnvData["role"].(string)
	finalResult.Name = name
	finalResult.Success = true

	start := time.Now()
	finalResult.DateTime = start
	for _, handler := range sorted {
		if !handler.Filter(ctx) {
			continue
		}
		result := handler.Run(ctx)
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
	finalResult.Duration = time.Since(start)

	duration := time.Since(start)
	hm.updateStats(name, duration, HookResult{
		StopExecution: ctx.IsStopped(),
		Error:         nil,
	})

	return finalResult
}
