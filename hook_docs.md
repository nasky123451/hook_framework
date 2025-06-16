# Hook Documentation

Generated at: 2025-06-16T17:38:33+08:00

## webhook_sync
- 📄 Description: Handles synchronization of webhooks from external sources
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - source

## create_account
- 📄 Description: Handles sending a welcome email when a new account is created
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 👥 Allowed Roles:
  - admin
- 🎯 Expected Parameters:
  - email

## login_failure_alert
- 📄 Description: Handles security alerts for login failures
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - ip

## subscription_reminder
- 📄 Description: Handles sending subscription reminders to users
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - user_id

## update_username
- 📄 Description: Handles updating a user's username
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
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

## notify_account_created
- 📄 Description: Handles sending a welcome email when an account is created
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - email

## create_jira_task
- 📄 Description: Handles creating a Jira task for account setup
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - task_details

## system_monitor
- 📄 Description: Handles system monitoring alerts
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - server

## set_user_pref
- 📄 Description: Handles setting user preferences like theme, language, etc.
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - theme

## submit_report
- 📄 Description: Handles submission of reports and routes them for approval
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - doc_type

## book_room
- 📄 Description: Handles room booking and prevents conflicts in scheduling
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - room
  - time

## create_invoice
- 📄 Description: Handles invoice creation and audits the invoice details
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - invoice_no
  - amount

## set_language
- 📄 Description: Handles setting the language for the user based on their preferences
- 🔗 Registered From: hook_builder.go:72 (hooks.(*HookBuilders).RegisterHookDefinitions)
- 🎯 Expected Parameters:
  - language

## 🔐 Permission Levels (from highest to lowest):
- superadmin
- admin
  - auditor
    - employee
      - user
        - subscriber
  - finance
    - employee
      - user
        - subscriber
  - security
    - employee
      - user
        - subscriber
  - devops
    - integration
      - employee
        - user
          - subscriber
