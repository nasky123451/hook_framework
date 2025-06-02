package utils

import (
	"hook_framework/internal/hooks"
)

// GenericHook 通用 Hook 結構，實現 HookHandler 接口
type GenericHook struct {
	HookName     string
	HookPriority int
	Role         string // 重命名字段，避免與方法名稱衝突
	RunFunc      func(ctx *hooks.HookContext) hooks.HookResult
}

// Name 返回 Hook 的名稱
func (h *GenericHook) Name() string {
	return h.HookName
}

// Priority 返回 Hook 的優先級
func (h *GenericHook) Priority() int {
	return h.HookPriority
}

// Filter 返回是否執行 Hook，默認為 true
func (h *GenericHook) Filter(ctx *hooks.HookContext) bool {
	return true
}

// RequiredRole 返回執行 Hook 所需的角色
func (h *GenericHook) RequiredRole() string {
	return h.Role
}

// Run 執行 Hook 的邏輯
func (h *GenericHook) Run(ctx *hooks.HookContext) hooks.HookResult {
	if h.RunFunc != nil {
		return h.RunFunc(ctx)
	}
	return hooks.HookResult{}
}
