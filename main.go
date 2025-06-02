package main

import (
	"hook_framework/internal/framework"
)

func main() {
	// 初始化框架
	processor, printer, initializedPlugins := framework.InitializeFramework()

	// 模擬多個客戶端輸入
	clientInputs := []struct {
		Input string
		Role  string
	}{
		{"修改 email 為 example@example.com", "admin"},            // 測試修改 email
		{"更新用戶名為 John Doe", "user"},                            // 測試更新用戶名
		{"啟用兩步驗證", "admin"},                                    // 測試啟用兩步驗證
		{"刪除關鍵數據", "admin"},                                    // 測試刪除操作
		{"刪除關鍵數據", "user"},                                     // 測試無權限刪除
		{"檢查系統狀態", "viewer"},                                   // 測試檢查系統狀態
		{"啟用夜間模式", "editor"},                                   // 測試啟用夜間模式
		{"傳送通知內容為 Hello World 傳送給 user@example.com", "editor"}, // 測試帶有 email 的通知
		{"傳送通知內容為 Hello World", "admin"},                       // 測試 admin 無視規則
		{"禁用通知", "admin"},                                      // 測試禁用通知
		{"未知操作測試", "guest"},                                    // 測試未知操作
	}

	for _, clientInput := range clientInputs {
		processor.Process(clientInput)
	}

	// 輸出 Hook Stats
	framework.PrintStats(processor.Env, printer, initializedPlugins)
}
