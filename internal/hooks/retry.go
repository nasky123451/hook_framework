package hooks

import (
	"log"
	"time"
)

type RetryConfig struct {
	MaxRetries int
	Backoff    time.Duration
}

func executeWithRetry(ctx *HookContext, handler HookHandler, retryCfg RetryConfig) HookResult {
	var result HookResult
	for i := 0; i <= retryCfg.MaxRetries; i++ {
		result = handler.Run(ctx)
		if result.Success || !result.Retryable || result.Error == nil {
			return result
		}
		log.Printf("[Retry] Handler %s failed: %v, retrying (%d/%d)...",
			handler.Name(), result.Error, i+1, retryCfg.MaxRetries)
		time.Sleep(retryCfg.Backoff)
	}
	return result
}
