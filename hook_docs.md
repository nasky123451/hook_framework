# Hook Documentation

Generated at: 2025-06-13T15:39:32+08:00

## book_room
- 📄 Description: Handles room booking and prevents conflicts in scheduling
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
  - employee
- 🎯 Expected Parameters:
  - room
  - time

## create_invoice
- 📄 Description: Handles invoice creation and audits the invoice details
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
  - finance
- 🎯 Expected Parameters:
  - invoice_no
  - amount

## set_language
- 📄 Description: Handles setting the language for the user based on their preferences
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
  - user
- 🎯 Expected Parameters:
  - language

## notify_account_created
- 📄 Description: Handles sending a welcome email when an account is created
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
  - system
- 🎯 Expected Parameters:
  - email

## create_jira_task
- 📄 Description: Handles creating a Jira task for account setup
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - task_details

## system_monitor
- 📄 Description: Handles system monitoring alerts
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - devops
- 🎯 Expected Parameters:
  - server

## update_username
- 📄 Description: Handles updating a user's username
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - username

## create_user
- 📄 Description: Creates a new user with a username and email
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - username
  - email

## update_user
- 📄 Description: Updates a user's email address
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - username
  - new_email

## delete_user
- 📄 Description: Deletes a user by username
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - username

## set_user_pref
- 📄 Description: Handles setting user preferences like theme, language, etc.
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - user
- 🎯 Expected Parameters:
  - theme

## create_account
- 📄 Description: Handles sending a welcome email when a new account is created
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - email

## submit_report
- 📄 Description: Handles submission of reports and routes them for approval
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - auditor
- 🎯 Expected Parameters:
  - doc_type

## login_failure_alert
- 📄 Description: Handles security alerts for login failures
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - security
- 🎯 Expected Parameters:
  - ip

## subscription_reminder
- 📄 Description: Handles sending subscription reminders to users
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - subscriber
- 🎯 Expected Parameters:
  - user_id

## webhook_sync
- 📄 Description: Handles synchronization of webhooks from external sources
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - integration
- 🎯 Expected Parameters:
  - source
