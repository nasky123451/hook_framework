package hooks

import (
	"log"
	"strings"
)

type OperationHandler func(ctx *HookContext, params map[string]interface{})

var operationRegistry = make(map[string]OperationHandler)

// RegisterOperationHandler 註冊操作處理器
func RegisterOperationHandler(action string, handler OperationHandler) {
	// Simplify action name for registration
	simplifiedAction := simplifyActionName(action)
	operationRegistry[simplifiedAction] = handler
}

// DispatchInput 根據解析結果執行對應的操作處理器
func DispatchInput(action string, params map[string]interface{}, ctx *HookContext, hookManager *HookManager) {
	if params == nil {
		params = make(map[string]interface{})
	}
	params["action"] = action

	// 如果 action 是 unknown，直接停止後續動作
	if action == "unknown" {
		log.Printf("[DispatchInput] Action '%s' is unknown. Stopping further execution.", action)
		ctx.Stop()
		return
	}

	handleCommonParameters(params, ctx)

	// Simplify action name for lookup
	simplifiedAction := simplifyActionName(action)

	handler, exists := operationRegistry[simplifiedAction]
	if !exists {
		log.Printf("[DispatchInput] No handler found for action: %s (simplified: %s)", action, simplifiedAction)
		return
	}

	handler(ctx, params)

	// 確保只執行一次對應的 Hook
	if ctx.IsStopped() {
		log.Printf("[DispatchInput] Execution stopped for action: %s", action)
		return
	}

	// 使用 GetHookName 獲取標準化名稱
	hookName, exists := GetHookName(action)
	if !exists {
		log.Printf("[DispatchInput] No HookName found for action: %s (simplified: %s)", action, simplifiedAction)
		return
	}

	normalizedName := simplifyActionName(hookName)
	if hookManager != nil {
		hookManager.Execute(normalizedName, ctx, false)
	}
}

// handleCommonParameters 處理通用參數檢查和上下文設置
func handleCommonParameters(params map[string]interface{}, ctx *HookContext) {
	for key, value := range params {
		// 忽略 nil 值
		if value == nil {
			continue
		}
		ctx.Set(key, value)
	}
}

// simplifyActionName 移除插件名稱前綴，返回簡化的操作名稱
func simplifyActionName(action string) string {
	parts := strings.Split(action, ".")
	if len(parts) > 1 {
		return parts[1] // 返回操作名稱部分
	}
	return action // 如果沒有前綴，直接返回原始名稱
}
