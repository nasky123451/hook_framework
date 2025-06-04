package main

import (
	"hook_framework/internal/framework"
	"log"
)

func main() {
	// 初始化框架
	processor, printer, initializedPlugins := framework.InitializeFramework()

	clientInputs := []framework.ClientInput{
		{Input: "create_account", Role: "admin", Context: map[string]interface{}{"email": "new_user@example.com", "env": "web"}},
		{Input: "login_failure_alert", Role: "security", Context: map[string]interface{}{"ip": "192.168.0.1"}},
		{Input: "book_room", Role: "employee", Context: map[string]interface{}{"room": "A101", "time": "2025-06-05 10:00", "device": "mobile"}},
		{Input: "submit_report", Role: "auditor", Context: map[string]interface{}{"doc_type": "financial"}},
		{Input: "webhook_sync", Role: "integration", Context: map[string]interface{}{"source": "GitHub"}},
		{Input: "system_monitor", Role: "devops", Context: map[string]interface{}{"server": "prod-api-1"}},
		{Input: "create_invoice", Role: "finance", Context: map[string]interface{}{"invoice_no": "INV-123456", "amount": 12900}},
		{Input: "set_user_pref", Role: "user", Context: map[string]interface{}{"theme": "dark_mode"}},
		{Input: "set_language", Role: "user", Context: map[string]interface{}{"language": "zh-TW"}},
		{Input: "subscription_reminder", Role: "subscriber", Context: map[string]interface{}{"user_id": 142}},
		{Input: "create_user", Role: "admin", Context: map[string]interface{}{"username": "alice", "email": "alice@example.com"}},
		{Input: "update_user", Role: "admin", Context: map[string]interface{}{"username": "alice", "new_email": "alice.new@example.com"}},
		{Input: "delete_user", Role: "user", Context: map[string]interface{}{"username": "alice"}},
	}

	for _, input := range clientInputs {
		log.Printf("=== 測試輸入：%s (角色：%s) ===", input.Input, input.Role)
		processor.Process(input)
	}

	// 輸出 Hook Stats
	framework.PrintStats(processor.Env, printer, initializedPlugins)
}
