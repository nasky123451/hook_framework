package hooks

type HookEnvironment struct {
	HookManager *HookManager
	Contexts    []*HookContext
}

func NewHookEnvironment() *HookEnvironment {
	hookManager := NewHookManager()
	return &HookEnvironment{
		HookManager: hookManager,
	}
}
