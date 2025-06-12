package hooks

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type HookBuilder struct {
	HookName       string
	Description    string
	ParamHints     []string
	RegisteredFrom string
	priority       int
	roles          []string
	metadata       map[string]interface{}
	handler        HookHandlerFunc
}

func New(name string) *HookBuilder {
	h := &HookBuilder{
		HookName: name,
		metadata: make(map[string]interface{}),
	}
	registerGlobalHook(h)
	return h
}

func (h *HookBuilder) WithDescription(desc string) *HookBuilder {
	h.Description = desc
	return h
}

func (h *HookBuilder) WithParamHints(params ...string) *HookBuilder {
	h.ParamHints = params
	return h
}

func (b *HookBuilder) WithPriority(p int) *HookBuilder {
	b.priority = p
	return b
}

func (b *HookBuilder) AllowRoles(roles ...string) *HookBuilder {
	b.roles = append(b.roles, roles...)
	return b
}

func (b *HookBuilder) WithMetadata(key string, val interface{}) *HookBuilder {
	b.metadata[key] = val
	return b
}

func (b *HookBuilder) Handle(fn HookHandlerFunc) *HookBuilder {
	b.handler = fn
	return b
}

func (b *HookBuilder) RegisterTo(hm *HookManager) {
	// 記錄來源位置
	if pc, file, line, ok := runtime.Caller(1); ok {
		funcName := runtime.FuncForPC(pc).Name()
		b.RegisteredFrom = fmt.Sprintf("%s:%d (%s)", filepath.Base(file), line, filepath.Base(funcName))
	}

	hm.RegisterHookWithOptions(b.HookName, HookOptions{
		Priority: b.priority,
		Roles:    b.roles,
		Metadata: b.metadata,
	}, b.handler)
}
