package hooks

import "time"

type HookHandler interface {
	Name() string
	Priority() int
	Filter(ctx *HookContext) bool
	Run(ctx *HookContext) HookResult
}

type HookResult struct {
	Name          string
	Success       bool
	Duration      time.Duration
	StopExecution bool
	Error         error
	Message       string
}

type HookHandlerFunc func(ctx *HookContext) HookResult
