package hooks

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// HookContext 定義執行 Hook 時的上下文結構
// 分離 UserData 與 EnvData，並加入元數據、錯誤收集與停止控制
type HookContext struct {
	UserData           map[string]interface{} // 使用者相關資料（如角色、user_id 等）
	EnvData            map[string]interface{} // 補充環境上下文（如 IP、裝置、語言等）
	Metadata           HookMetadata           // 執行元資料（觸發來源、時間戳等）
	Errors             []error                // 執行中產生的錯誤列表
	Results            []interface{}          // 執行結果列表
	StopSignal         bool                   // 是否停止後續 Hook 執行
	mu                 sync.RWMutex           // 併發保護
	executionLog       []HookResult           // 執行日誌，記錄每個 Hook 的執行結果
	CurrentHandlerName string
}

// HookMetadata 用於描述 Hook 執行的額外元資訊
type HookMetadata struct {
	TriggeredBy string    // 觸發 Hook 的來源（如 Plugin 名稱）
	Origin      string    // 事件原點（如 API、系統）
	Timestamp   time.Time // 觸發時間

	Name        string   // Hook 的名稱
	Description string   // Hook 的描述
	Permissions string   // 可執行的角色列表（若為 nil 表示允許全部）
	ParamHints  []string // 參數提示（如類型、範例值等）
	Plugin      string   // 所屬插件名稱
}

// NewHookContext 時初始化
func NewHookContext(role string, extra map[string]interface{}) *HookContext {
	return &HookContext{
		UserData:   map[string]interface{}{"role": role},
		EnvData:    make(map[string]interface{}),
		Errors:     make([]error, 0),
		Results:    make([]interface{}, 0),
		StopSignal: false,
		Metadata:   HookMetadata{Timestamp: time.Now()},
	}
}

func (ctx *HookContext) AddExecutionLog(result HookResult) {
	ctx.executionLog = append(ctx.executionLog, result)
}

func (ctx *HookContext) GetExecutionLog() []HookResult {
	return ctx.executionLog
}

func (c *HookContext) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// 先從 UserData 找
	if val, ok := c.UserData[key]; ok {
		return val
	}

	// 再從 EnvData 找
	return c.EnvData[key]
}

func (c *HookContext) Set(key string, val interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 先放到 UserData 裡，這裡可以根據需求調整
	if c.UserData == nil {
		c.UserData = make(map[string]interface{})
	}
	c.UserData[key] = val
}

func (c *HookContext) GetString(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	// 優先從 UserData 取
	if val, ok := c.UserData[key]; ok {
		if s, ok2 := val.(string); ok2 {
			return s
		}
	}
	// 若 UserData 沒有，再從 EnvData 取
	if val, ok := c.EnvData[key]; ok {
		if s, ok2 := val.(string); ok2 {
			return s
		}
	}
	return ""
}

// SetUserData 設定使用者資料鍵值對
func (c *HookContext) SetUserData(key string, val interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.UserData == nil {
		c.UserData = make(map[string]interface{})
	}
	c.UserData[key] = val
}

// GetUserData 取得使用者資料值
func (c *HookContext) GetUserData(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.UserData == nil {
		return nil
	}
	return c.UserData[key]
}

// SetEnvData 設定環境資料鍵值對
func (c *HookContext) SetEnvData(key string, val interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.EnvData == nil {
		c.EnvData = make(map[string]interface{})
	}
	c.EnvData[key] = val
}

// GetEnvData 取得環境資料值
func (c *HookContext) GetEnvData(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.EnvData == nil {
		return nil, false
	}
	val, ok := c.EnvData[key]
	return val, ok
}

// GetEnvString 回傳字串型別的環境變數
func (c *HookContext) GetEnvString(key string) (string, bool) {
	val, ok := c.GetEnvData(key)
	if !ok {
		return "", false
	}
	str, ok := val.(string)
	return str, ok
}

func (c *HookContext) SuccessWithMessage(format string, args ...any) HookResult {
	message := fmt.Sprintf(format, args...)
	c.SetEnvData("approval_message", message)

	return HookResult{
		Success: true,
		Message: message,
	}
}

// AddError 新增錯誤到錯誤列表
func (c *HookContext) AddError(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Errors = append(c.Errors, err)
}

// Reset 清空錯誤列表，重置停止狀態及資料
func (c *HookContext) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Errors = nil
	c.StopSignal = false
	c.UserData = make(map[string]interface{})
	c.EnvData = make(map[string]interface{})
	c.Metadata = HookMetadata{
		Timestamp: time.Now(),
	}
}

// Stop 標記停止執行後續 Hook
func (c *HookContext) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.StopSignal = true
	log.Println("[HookContext] Execution stopped.")
}

// IsStopped 查詢是否已標記停止
func (c *HookContext) IsStopped() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.StopSignal
}

func (ctx *HookContext) Clone() *HookContext {
	ctx.mu.RLock()
	defer ctx.mu.RUnlock()

	// 深拷貝 UserData
	clonedUserData := make(map[string]interface{})
	for k, v := range ctx.UserData {
		clonedUserData[k] = v
	}

	// 深拷貝 EnvData
	clonedEnvData := make(map[string]interface{})
	for k, v := range ctx.EnvData {
		clonedEnvData[k] = v
	}

	// 拷貝 Metadata, Errors, Results 等（依需求深淺拷貝）
	cloned := &HookContext{
		UserData:     clonedUserData,
		EnvData:      clonedEnvData,
		Metadata:     ctx.Metadata, // 若 Metadata 裡有指標欄位需另拷
		Errors:       append([]error{}, ctx.Errors...),
		Results:      append([]interface{}{}, ctx.Results...),
		StopSignal:   ctx.StopSignal,
		executionLog: ctx.executionLog,
	}

	return cloned
}
