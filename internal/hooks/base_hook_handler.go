package hooks

import "fmt"

type BaseHookHandler struct {
	name     string
	priority int
	handler  HookHandlerFunc

	roles    []string
	metadata map[string]interface{}
}

// Name 回傳 hook 名稱
func (h *BaseHookHandler) Name() string {
	return h.name
}

// Priority 回傳執行優先順序
func (h *BaseHookHandler) Priority() int {
	return h.priority
}

// Roles 回傳可執行的角色（若為 nil 表示允許全部）
func (h *BaseHookHandler) Roles() []string {
	return h.roles
}

// Metadata 回傳自定義資訊
func (h *BaseHookHandler) Metadata() map[string]interface{} {
	return h.metadata
}

// Execute 執行 hook handler
func (h *BaseHookHandler) Execute(ctx *HookContext) HookResult {
	return h.handler(ctx)
}

// Filter 權限篩選（預設允許全部）
func (h *BaseHookHandler) Filter(ctx *HookContext) bool {
	if len(h.roles) == 0 {
		return true // 沒有限制
	}
	userRole, _ := ctx.GetUserData("role").(string)
	for _, role := range h.roles {
		if role == userRole {
			return true
		}
	}
	return false
}

// Run 執行 hook 並自動補 message
func (h *BaseHookHandler) Run(ctx *HookContext) HookResult {
	result := h.handler(ctx)
	if result.Success {
		if _, ok := ctx.GetEnvData("approval_message"); !ok {
			ctx.SetEnvData("approval_message", fmt.Sprintf("Hook %s executed successfully", h.name))
		}
	}
	return result
}
