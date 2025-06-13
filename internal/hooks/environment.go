package hooks

type HookEnvironment struct {
	HookManager *HookManager
	HookGraph   *HookGraph
	Contexts    []*HookContext
}

func NewHookEnvironment() *HookEnvironment {
	hookManager := NewHookManager()
	return &HookEnvironment{
		HookManager: hookManager,
		HookGraph:   NewHookGraph(hookManager),
	}
}
