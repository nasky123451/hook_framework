package hooks

type HookEnvironment struct {
	HookManager *HookManager
	Context     *HookContext
}

func NewHookEnvironment(triggeredBy string, origin string) *HookEnvironment {
	hookManager := NewHookManager()
	return &HookEnvironment{
		HookManager: hookManager,
		Context:     NewHookContext(triggeredBy, map[string]interface{}{"origin": origin}),
	}
}
