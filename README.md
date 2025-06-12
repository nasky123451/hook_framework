# Hook Framework

This project is a modular hook framework designed to dynamically manage and execute hooks. It includes components for plugin management, NLP-based input processing, and hook execution.

## Folder Structure

```
hook_framework/
├── internal/
│   ├── framework/
│   ├── hooks/
│   ├── plugins/
│   │   ├── meta/
├── flow_test/
├── main.go
├── README.md
├── hook_docs.md
├── go.mod
├── go.sum
└── .gitignore
```

## Overview of Components

### `internal/framework/`
Handles the core framework logic:
- **Processor**: Processes client input and executes hooks.
- **Initializer**: Sets up the framework, plugins, and hooks.
- **Cleanup**: Outputs hook statistics and handles resource cleanup.
- **Printer**: Provides utility functions for printing logs and statistics.

### `internal/hooks/`
Manages hooks and their execution:
- **Registry**: Manages hook names and their registration.
- **Dispatcher**: Dispatches input to appropriate handlers and hooks.
- **Environment**: Sets up the hook environment.
- **Plugin Manager**: Manages plugins and their hooks.
- **Plugin Registry**: Registers and retrieves plugin types.
- **Plugin Interface**: Defines the plugin interface and metadata.
- **Hook Builder**: Provides a builder pattern for defining hooks.
- **Hook Graph**: Manages directed acyclic graphs (DAG) for hook execution.
- **Hook Execution**: Handles the execution of hooks.
- **Hook Stats**: Tracks statistics for hook execution.

### `internal/plugins/meta/`
Contains individual plugins:
- **ApprovalRoutingPlugin**: Routes submitted reports for approval.
- **BookingLockPlugin**: Prevents conflicts in room booking schedules.
- **InvoiceAuditPlugin**: Audits invoice creation details.
- **LocalizationPlugin**: Sets user language preferences.
- **NotifyPlugin**: Handles notifications and Jira task creation.
- **SecurityAlertPlugin**: Manages security alerts for login failures.
- **SubscriptionReminderPlugin**: Sends subscription reminders.
- **SystemMonitorPlugin**: Handles system monitoring alerts.
- **UpdateUsernamePlugin**: Updates usernames dynamically.
- **UserManagementPlugin**: Manages user creation, updates, and deletion.
- **UserPreferencePlugin**: Handles user preferences like themes and languages.
- **WebhookSyncPlugin**: Synchronizes webhooks from external sources.
- **WelcomeEmailPlugin**: Sends welcome emails upon account creation.

### `flow_test/`
Contains flow tests for simulating hook execution scenarios:
- **CreateAccountFlow**: Simulates account creation, notification, and Jira task creation.

### `hook_docs.md`
Auto-generated documentation for all registered hooks.

---

## Usage

### Initialization
The framework is initialized using the `InitializeFramework` function in `internal/framework/initializer.go`. This function:
1. Sets up the hook environment.
2. Loads plugin configurations dynamically.
3. Registers plugins and their hooks.
4. Initializes the HookGraph for DAG-based execution.

### Processing Client Input
Client input is processed using the `ClientInputProcessor` in `internal/framework/processor.go`. This processor:
1. Parses the input and sets up the context.
2. Executes hooks or DAG chains based on the input.
3. Handles errors and stops execution if necessary.

### Hook Execution
Hooks are dynamically registered and executed using the `RegisterHook` and `ExecuteHookByName` functions in `internal/hooks/registry.go` and `internal/hooks/hook_execution.go`.

### Plugins
Plugins are modular components that extend the framework's functionality. Each plugin:
1. Implements the `hooks.Plugin` interface.
2. Registers hooks dynamically using `RegisterHooks`.
3. Optionally registers parsers for NLP-based input handling.

### Flow Tests
Flow tests simulate hook execution scenarios for testing. For example, `flow_test/create_account_flow.go` simulates account creation and related operations.

---

## How to Run

1. Ensure you have Go installed.
2. Navigate to the project directory:
   ```bash
   cd /Users/linjunhan/works/go/src/hook_framework
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

---

## Notes

- Ensure the `hook_docs.md` file is generated for hook documentation.
- Plugins must implement the `hooks.Plugin` interface to be compatible with the framework.
- The framework dynamically loads plugins and hooks based on the configuration file and registered plugins.

---

## License

This project is licensed under the MIT License.
