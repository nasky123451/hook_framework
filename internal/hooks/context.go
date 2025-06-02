package hooks

import (
	"log"
	"sync"
	"time"
)

// HookContext 定義上下文結構
type HookContext struct {
	Data     map[string]interface{}
	Errors   []error
	Metadata HookMetadata
	mu       sync.RWMutex
	_stop    bool // 私有字段，用於標記是否停止執行
}

// HookMetadata 定義上下文的元數據
type HookMetadata struct {
	TriggeredBy string
	Origin      string
	Timestamp   time.Time
}

// NewHookContext 創建新的 HookContext
func NewHookContext(triggeredBy string, origin string) *HookContext {
	return &HookContext{
		Data: make(map[string]interface{}),
		Metadata: HookMetadata{
			TriggeredBy: triggeredBy,
			Origin:      origin,
			Timestamp:   time.Now(),
		},
	}
}

// Set 設置上下文中的鍵值對
func (c *HookContext) Set(key string, val interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.Data == nil {
		c.Data = make(map[string]interface{})
	}
	c.Data[key] = val
}

// Get 獲取上下文中的值
func (c *HookContext) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Data[key]
}

// GetString 獲取上下文中的字符串值
func (c *HookContext) GetString(key string) string {
	if val, ok := c.Get(key).(string); ok {
		return val
	}
	return ""
}

// AddError 添加錯誤到上下文
func (c *HookContext) AddError(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Errors = append(c.Errors, err)
	log.Printf("[HookContext] Error added: %v", err)
}

// Reset 重置上下文的錯誤和停止狀態
func (c *HookContext) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Errors = nil
	c._stop = false
}

// Stop 停止上下文執行
func (c *HookContext) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c._stop = true
	log.Println("[HookContext] Execution stopped.")
}

// IsStopped 檢查上下文是否已停止
func (c *HookContext) IsStopped() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c._stop
}
