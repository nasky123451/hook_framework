package hooks

import "time"

type HookHandler interface {
	Name() string
	Permissions() string
	Priority() int
	Filter(ctx *HookContext) bool
	Run(ctx *HookContext) HookResult
}

type HookResult struct {
	Permissions   string
	Name          string
	Success       bool
	Duration      time.Duration
	StopExecution bool
	Error         error
	Message       string
	DateTime      time.Time
}

type HookHandlerFunc func(ctx *HookContext) HookResult
