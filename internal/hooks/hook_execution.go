package hooks

import (
	"fmt"
	"log"
	"time"
)

func (hm *HookManager) Execute(name string, ctx *HookContext, async bool) {
	hm.mu.RLock()
	hooks := hm.hooks[name]
	hm.mu.RUnlock()

	if len(hooks) == 0 {
		log.Printf("[Hook] No handlers registered for hook '%s'", name)
		return
	}

	start := time.Now()
	log.Printf("[Hook] Executing hook '%s' with %d handlers", name, len(hooks))

	if async {
		hm.executeAsync(hooks, ctx)
	} else {
		hm.executeSync(hooks, ctx)
	}

	duration := time.Since(start)
	hm.updateStats(name, duration, HookResult{
		StopExecution: ctx.IsStopped(),
		Error:         nil,
	})
	log.Printf("[Hook] Finished hook '%s' in %v", name, duration)
}

func (hm *HookManager) executeSync(hooks []HookHandler, ctx *HookContext) {
	for _, handler := range hooks {
		if !handler.Filter(ctx) {
			log.Printf("[Hook] Handler '%s' skipped due to filter", handler.Name())
			continue
		}
		defer recoverHook(ctx, handler.Name())
		res := handler.Run(ctx)
		if res.Error != nil {
			ctx.AddError(res.Error)
		}
		if res.StopExecution {
			log.Printf("[Hook] Execution stopped by hook '%s'", handler.Name())
			ctx.Stop()
			break
		}
	}
}

func (hm *HookManager) executeAsync(hooks []HookHandler, ctx *HookContext) {
	for _, handler := range hooks {
		if !handler.Filter(ctx) {
			log.Printf("[Hook] Handler '%s' skipped due to filter", handler.Name())
			continue
		}
		h := handler
		_ = hm.pool.Submit(func() {
			defer recoverHook(ctx, h.Name())
			res := h.Run(ctx)
			if res.Error != nil {
				ctx.AddError(res.Error)
			}
			if res.StopExecution {
				log.Printf("[Hook] Execution stopped by hook '%s'", h.Name())
				ctx.Stop()
			}
		})
	}
}

func recoverHook(ctx *HookContext, name string) {
	if r := recover(); r != nil {
		err := fmt.Errorf("hook '%s' panic: %v", name, r)
		ctx.AddError(err)
	}
}
