package hooks

import (
	"time"
)

type HookStats struct {
	ExecutionCount      int
	TotalDuration       time.Duration
	LastExecutionResult HookResult
}

type StatsVersion struct {
	Version   int
	Timestamp time.Time
}

func (hm *HookManager) updateStats(name string, duration time.Duration, result HookResult) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	if stat, exists := hm.stats[name]; exists {
		stat.ExecutionCount++
		stat.TotalDuration += duration
		stat.LastExecutionResult = result
	} else {
		hm.stats[name] = &HookStats{
			ExecutionCount:      1,
			TotalDuration:       duration,
			LastExecutionResult: result,
		}
	}
}

func (hm *HookManager) GetStats() (map[string]*HookStats, StatsVersion) {
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	return hm.stats, hm.statsVer
}
