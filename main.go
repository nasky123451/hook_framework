package main

import (
	"fmt"
	flowtest "hook_framework/flow_test"
	"hook_framework/internal/framework"
	"hook_framework/internal/hooks"
	"log"
)

func main() {
	// 初始化框架
	processor := framework.InitializeFramework()

	for _, line := range hooks.GetFormattedHookDescriptions() {
		fmt.Println(line)
	}

	clientInputs := []framework.ClientInput{
		{Input: "create_account", Permissions: "security", Context: map[string]interface{}{"email": "new_user@example.com"}},
		{Input: "login_failure_alert", Permissions: "security", Context: map[string]interface{}{"ip": "192.168.0.1"}},
		{Input: "book_room", Permissions: "security", Context: map[string]interface{}{"room": "A101", "time": "2025-06-05 10:00", "device": "mobile"}},
		{Input: "submit_report", Permissions: "security", Context: map[string]interface{}{"doc_type": "financial"}},
		{Input: "webhook_sync", Permissions: "admin", Context: map[string]interface{}{"source": "GitHub"}},
		{Input: "system_monitor", Permissions: "admin", Context: map[string]interface{}{"server": "prod-api-1"}},
		{Input: "create_invoice", Permissions: "admin", Context: map[string]interface{}{"invoice_no": "INV-123456", "amount": 12900}},
		{Input: "set_user_pref", Permissions: "admin", Context: map[string]interface{}{"theme": "dark_mode"}},
		{Input: "set_language", Permissions: "admin", Context: map[string]interface{}{"language": "zh-TW"}},
		{Input: "subscription_reminder", Permissions: "admin", Context: map[string]interface{}{"user_id": 142}},
		{Input: "create_user", Permissions: "admin", Context: map[string]interface{}{"username": "alice", "email": "alice@example.com"}},
		{Input: "update_user", Permissions: "security", Context: map[string]interface{}{"username": "alice", "new_email": "alice.new@example.com"}},
		{Input: "delete_user", Permissions: "security", Context: map[string]interface{}{"username": "alice"}},
	}

	for _, input := range clientInputs {
		processor.Process(input)
	}

	clientInputs = []framework.ClientInput{
		{Input: "create_account", Permissions: "admin", Context: map[string]interface{}{"email": "new_user2@example.com"}},
		{Input: "create_account", Permissions: "admin", Context: map[string]interface{}{"email": "user2@example.com"}},
	}

	for _, input := range clientInputs {
		processor.ProcessWithGraph(input)
	}

	err := flowtest.RunCreateAccountFlow("new_user@example.com")
	if err != nil {
		log.Fatal("流程執行失敗:", err)
	}

	// 輸出 Hook Stats
	processor.PrintStats()
}
