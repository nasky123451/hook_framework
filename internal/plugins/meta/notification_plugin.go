package meta

import (
	"fmt"
	"log"
	"regexp"

	"hook_framework/internal/hooks"
	"hook_framework/pkg/nlp"
	"hook_framework/pkg/utils"
)

type NotificationPlugin struct {
	AddHook     string
	DisableHook string
}

func (p *NotificationPlugin) Name() string {
	return "NotificationPlugin"
}

func (p *NotificationPlugin) GetHookNames() []string {
	return []string{p.AddHook, p.DisableHook}
}

func (p *NotificationPlugin) RegisterHooks(hookManager *hooks.HookManager) {
	if p.AddHook == "" || p.DisableHook == "" {
		log.Fatal("[NotificationPlugin] One or more hook names are not initialized.")
	}

	utils.RegisterDynamicHook(hookManager, p.AddHook, 40, "editor", func(ctx *hooks.HookContext) hooks.HookResult {

		email := ctx.GetString("email")
		message := ctx.GetString("message")

		if email == "" {
			return hooks.HookResult{Error: fmt.Errorf("email is missing, please provide a valid email address")}
		}
		if message == "" {
			return hooks.HookResult{Error: fmt.Errorf("message is missing, please provide a notification message")}
		}

		SendNotification(email, message)
		ctx.Set("notification_added", true)

		return hooks.HookResult{}
	})

	utils.RegisterDynamicHook(hookManager, p.DisableHook, 40, "admin", func(ctx *hooks.HookContext) hooks.HookResult {
		ctx.Set("notifications_disabled", true)
		return hooks.HookResult{}
	})
}

func (p *NotificationPlugin) RegisterParsers(nlpEngine *nlp.NLP) {
	if nlpEngine == nil {
		log.Fatal("[NotificationPlugin] NLP engine is nil.")
	}

	// NLP：禁用通知
	nlpEngine.RegisterParser(func(input string) (nlp.Intent, bool) {
		if input == "禁用通知" {
			return nlp.Intent{Action: "disable_notifications"}, true
		}
		return nlp.Intent{}, false
	})

	// NLP：傳送通知內容為 ... 傳送給 ...
	nlpEngine.RegisterParser(func(input string) (nlp.Intent, bool) {
		pattern := regexp.MustCompile(`傳送通知內容為\s+(.*?)\s+傳送給\s+(\S+@\S+\.\S+)`)
		matches := pattern.FindStringSubmatch(input)
		if len(matches) == 3 {
			return nlp.Intent{
				Action: "add_notification",
				Params: map[string]string{
					"message": matches[1],
					"email":   matches[2],
				},
			}, true
		}
		return nlp.Intent{}, false
	})
}

// SendNotification 發送通知並輸出通知內容
func SendNotification(email string, message string) {
	fmt.Printf("[Notification] Sending notification to '%s' with message: \"%s\"\n", email, message)
}

func init() {
	hooks.RegisterPluginType("NotificationPlugin", func() hooks.Plugin {
		return &NotificationPlugin{
			AddHook:     "add_notification",
			DisableHook: "disable_notifications",
		}
	})
}
