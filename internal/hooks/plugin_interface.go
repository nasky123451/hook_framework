package hooks

import (
	"hook_framework/pkg/nlp"
)

type PluginWrapper struct {
	Impl Plugin
	Meta map[string]interface{}
}

type Wrapper interface {
	Unwrap() Plugin
}

type ParserRegistrable interface {
	RegisterParsers(*nlp.NLP)
}

func (pw *PluginWrapper) RegisterParsers(nlpEngine *nlp.NLP) {
	if registrable, ok := pw.Impl.(ParserRegistrable); ok {
		registrable.RegisterParsers(nlpEngine)
	}
}

func (w *PluginWrapper) Unwrap() Plugin {
	return w.Impl
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
