# Hook Documentation

Generated at: 2025-06-13T11:51:25+08:00

## create_invoice
- 📄 Description: Handles invoice creation and audits the invoice details
- 🔗 Registered From: invoice_audit_plugin.go:32 (meta.(*InvoiceAuditPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - admin
  - finance
- 🎯 Expected Parameters:
  - invoice_no
  - amount

## subscription_reminder
- 📄 Description: Handles sending subscription reminders to users
- 🔗 Registered From: subscription_reminder_plugin.go:29 (meta.(*SubscriptionReminderPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - subscriber
- 🎯 Expected Parameters:
  - user_id

## update_username
- 📄 Description: Handles updating a user's username
- 🔗 Registered From: update_username_plugin.go:35 (meta.(*UpdateUsernamePlugin).RegisterHooks)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - username

## create_user
- 📄 Description: Handles user management operations: create_user
- 🔗 Registered From: user_management_plugin.go:44 (meta.(*UserManagementPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - username
  - email
  - new_email

## update_user
- 📄 Description: Handles user management operations: update_user
- 🔗 Registered From: user_management_plugin.go:44 (meta.(*UserManagementPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - username
  - email
  - new_email

## delete_user
- 📄 Description: Handles user management operations: delete_user
- 🔗 Registered From: user_management_plugin.go:44 (meta.(*UserManagementPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - username
  - email
  - new_email

## set_user_pref
- 📄 Description: Handles setting user preferences like theme, language, etc.
- 🔗 Registered From: user_preference_plugin.go:31 (meta.(*UserPreferencePlugin).RegisterHooks)
- 👥 Allowed Roles:
  - user
- 🎯 Expected Parameters:
  - theme

## webhook_sync
- 📄 Description: Handles synchronization of webhooks from external sources
- 🔗 Registered From: webhook_sync_plugin.go:31 (meta.(*WebhookSyncPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - integration
- 🎯 Expected Parameters:
  - source

## create_account
- 📄 Description: Handles sending a welcome email when a new account is created
- 🔗 Registered From: welcome_email_plugin.go:36 (meta.(*WelcomeEmailPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - email

## submit_report
- 📄 Description: Handles submission of reports and routes them for approval
- 🔗 Registered From: approval_routing_plugin.go:32 (meta.(*ApprovalRoutingPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - auditor
- 🎯 Expected Parameters:
  - doc_type

## book_room
- 📄 Description: Handles room booking and prevents conflicts in scheduling
- 🔗 Registered From: booking_lock_plugin.go:33 (meta.(*BookingLockPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - admin
  - employee
- 🎯 Expected Parameters:
  - room
  - time

## set_language
- 📄 Description: Handles setting the language for the user based on their preferences
- 🔗 Registered From: localization_plugin.go:31 (meta.(*LocalizationPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - admin
  - user
- 🎯 Expected Parameters:
  - language

## notify_account_created
- 📄 Description: Handles sending a welcome email when an account is created
- 🔗 Registered From: notify_plugin.go:30 (meta.(*NotifyPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - admin
  - system
- 🎯 Expected Parameters:
  - email

## create_jira_task
- 📄 Description: Handles creating a Jira task for account setup
- 🔗 Registered From: notify_plugin.go:44 (meta.(*NotifyPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - task_details

## login_failure_alert
- 📄 Description: Handles security alerts for login failures
- 🔗 Registered From: security_alert_plugin.go:29 (meta.(*SecurityAlertPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - security
- 🎯 Expected Parameters:
  - ip

## system_monitor
- 📄 Description: Handles system monitoring alerts
- 🔗 Registered From: system_monitor_plugin.go:29 (meta.(*SystemMonitorPlugin).RegisterHooks)
- 👥 Allowed Roles:
  - devops
- 🎯 Expected Parameters:
  - server
