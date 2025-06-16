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
	Priority       int
	Permissions    string
	Metadata       map[string]interface{}
	Handler        HookHandlerFunc
}

type HookBuilders []*HookBuilder

var registeredMetadata = make([]HookMetadata, 0)

func New(name string) *HookBuilder {
	h := &HookBuilder{
		HookName: name,
		Metadata: make(map[string]interface{}),
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
	b.Priority = p
	return b
}

func (b *HookBuilder) AllowPermissions(permission string) *HookBuilder {
	b.Permissions = permission
	return b
}

func (b *HookBuilder) WithMetadata(key string, val interface{}) *HookBuilder {
	b.Metadata[key] = val
	return b
}

func (b *HookBuilder) Handle(fn HookHandlerFunc) *HookBuilder {
	b.Handler = fn
	return b
}

func (bs *HookBuilders) RegisterHookDefinitions(hm *HookManager, pluginName string) {
	for _, b := range *bs {
		New(b.HookName).
			WithDescription(b.Description).
			WithParamHints(b.ParamHints...).
			WithPriority(b.Priority).
			AllowPermissions(b.Permissions).
			WithMetadata("plugin", pluginName).
			Handle(b.Handler).
			RegisterTo(hm)
	}
}

func (b *HookBuilder) RegisterTo(hm *HookManager) {
	// 記錄來源位置
	if pc, file, line, ok := runtime.Caller(1); ok {
		funcName := runtime.FuncForPC(pc).Name()
		b.RegisteredFrom = fmt.Sprintf("%s:%d (%s)", filepath.Base(file), line, filepath.Base(funcName))
	}

	registeredMetadata = append(registeredMetadata, HookMetadata{
		Name:        b.HookName,
		Description: b.Description,
		ParamHints:  b.ParamHints,
		Permissions: b.Permissions,
	})

	hm.RegisterHookWithOptions(b.HookName, HookOptions{
		Priority:    b.Priority,
		Permissions: b.Permissions,
		Metadata:    b.Metadata,
	}, b.Handler)

}

func GetFormattedHookDescriptions() []string {
	var lines []string
	for _, meta := range registeredMetadata {
		// 動態構造 context
		contextParts := ""
		for i, param := range meta.ParamHints {
			if i > 0 {
				contextParts += ", "
			}
			contextParts += fmt.Sprintf("\"%s\": \"xxx\"", param)
		}

		permissions := meta.Permissions

		lines = append(lines, fmt.Sprintf(
			`🔹 [%s] %s
  - Description: %s
  - Permissions: %v
  - Example: {Input: "%s", Permissions: "%s", Context: map[string]interface{}{%s}}`,
			meta.Plugin, meta.Name,
			meta.Description,
			meta.Permissions,
			meta.Name, permissions, contextParts,
		))
	}
	return lines
}
