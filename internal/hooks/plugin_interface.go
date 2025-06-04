package hooks

type PluginWrapper struct {
	Impl Plugin
	Meta map[string]interface{}
}

type ParserRegistrable interface {
	RegisterParsers()
}

func (pw *PluginWrapper) RegisterParsers() {
	if registrable, ok := pw.Impl.(ParserRegistrable); ok {
		registrable.RegisterParsers()
	}
}

func (pw *PluginWrapper) Name() string {
	if pw.Impl != nil {
		return pw.Impl.Name()
	}
	return ""
}

func (pw *PluginWrapper) GetHookNames() []string {
	if pw.Impl != nil {
		return pw.Impl.GetHookNames()
	}
	return nil
}

func (pw *PluginWrapper) RegisterHooks(hm *HookManager) {
	if pw.Impl != nil {
		pw.Impl.RegisterHooks(hm)
	}
}
