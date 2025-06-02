package utils

import (
	"fmt"
	"hook_framework/internal/hooks"
	"log"
	"strings"
)

// Operation 定義操作結構
type Operation struct {
	Name    string
	Handler func(ctx interface{}, params map[string]interface{})
}

// RegisterAllHandlers 註冊所有操作處理器
func RegisterAllHandlers(registerFunc func(name string, handler func(ctx interface{}, params map[string]interface{})), hookNames map[string]string, hookConfigs map[string][]string) {
	if hookNames == nil {
		panic("[RegisterAllHandlers] hookNames map is nil. Ensure it is properly initialized.")
	}

	operations := []Operation{}

	for pluginName, hooks := range hookConfigs {
		for _, hook := range hooks {
			normalizedName, exists := hookNames[hook]
			if !exists {
				panic(fmt.Sprintf("[RegisterAllHandlers] Hook name '%s' is not registered in hookNames.", hook))
			}

			// Simplify operation name by removing plugin name prefix
			simplifiedName := simplifyActionName(normalizedName)

			operations = append(operations, Operation{
				Name: simplifiedName,
				Handler: func(ctx interface{}, params map[string]interface{}) {
					fmt.Printf("[Handler] Executing hook '%s' for plugin '%s' with params: %+v\n", simplifiedName, pluginName, params)
				},
			})
		}
	}

	// 註冊所有操作處理器
	for _, op := range operations {
		if op.Name == "" {
			panic(fmt.Sprintf("[RegisterAllHandlers] Operation name is empty for handler: %+v", op))
		}
		registerFunc(op.Name, op.Handler)
		log.Printf("[RegisterAllHandlers] Handler registered for operation: %s", op.Name)
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

// RegisterDynamicHook 動態註冊 Hook
func RegisterDynamicHook(hookManager *hooks.HookManager, hookName string, priority int, role string, runFunc func(ctx *hooks.HookContext) hooks.HookResult) {
	hooks.RegisterHook(hookManager, hookName, priority,
		func(ctx *hooks.HookContext) bool {
			// 過濾條件：檢查角色是否符合
			return ctx.GetString("role") == role || role == ""
		},
		func(ctx *hooks.HookContext) hooks.HookResult {
			return runFunc(ctx)
		},
	)
}
