package hooks

type HookEnvironment struct {
	HookManager    *HookManager
	HookGraph      *HookGraph
	ContextManager *HookContextManager
}

func NewHookEnvironment() *HookEnvironment {
	hookManager := NewHookManager()
	return &HookEnvironment{
		HookManager:    hookManager,
		HookGraph:      NewHookGraph(hookManager),
		ContextManager: NewHookContextManager(512, 1),
	}
}
