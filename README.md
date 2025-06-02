# Hook Framework

This project is a modular hook framework designed to dynamically manage and execute hooks. It includes components for plugin management, NLP-based input processing, and hook execution.

## Folder Structure

```
hook_framework/
├── internal/
│   ├── common/
│   ├── framework/
│   ├── hooks/
│   ├── plugins/
│   │   ├── meta/
├── pkg/
│   ├── utils/
│   ├── nlp/
├── main.go
└── README.md
```

## Overview of Components

### `internal/framework/`
Handles the core framework logic:
- **Processor**: Processes client input and executes hooks.
- **Initializer**: Sets up the framework, plugins, and hooks.
- **Cleanup**: Outputs hook statistics and handles resource cleanup.
- **Utils**: Provides utility functions for framework operations.

### `internal/hooks/`
Manages hooks and their execution:
- **Registry**: Manages hook names and their registration.
- **Dispatcher**: Dispatches input to appropriate handlers and hooks.
- **Environment**: Sets up the hook environment.
- **Plugin Manager**: Manages plugins and their hooks.
- **Plugin Registry**: Registers and retrieves plugin types.
- **Plugin Interface**: Defines the plugin interface and metadata.

### `internal/plugins/meta/`
Contains individual plugins:
- **UpdateUsernamePlugin**: Updates usernames dynamically.
- **UpdateEmailPlugin**: Updates email addresses dynamically.
- **TwoFactorPlugin**: Enables two-factor authentication.
- **NotificationPlugin**: Manages notifications.
- **DeleteProtectionPlugin**: Protects critical data from deletion.
- **Hook Config**: Provides hook configurations for plugins.

### `pkg/utils/`
Provides utility functions:
- **Config Loader**: Loads plugin configuration files.
- **Operations**: Registers and manages operation handlers.
- **Simulations**: Simulates hook execution scenarios.

### `pkg/nlp/`
Handles NLP-based input parsing:
- **NLP Engine**: Parses client input to determine actions and parameters.

### `main.go`
Entry point of the application. Initializes the framework and processes client inputs.

---

## Usage

### Initialization
The framework is initialized using the `InitializeFramework` function in `internal/framework/initializer.go`. This function:
1. Sets up the hook environment.
2. Loads plugin configurations from `plugin_config.json`.
3. Registers plugins and their hooks.
4. Initializes the NLP engine for input parsing.

### Processing Client Input
Client input is processed using the `ClientInputProcessor` in `internal/framework/processor.go`. This processor:
1. Parses the input using the NLP engine.
2. Dispatches the parsed action to the appropriate handler or hook.
3. Handles errors and stops execution if necessary.

### Hook Execution
Hooks are dynamically registered and executed using the `RegisterHook` and `DispatchInput` functions in `internal/hooks/registry.go` and `internal/hooks/dispatcher.go`.

### Plugins
Plugins are modular components that extend the framework's functionality. Each plugin:
1. Implements the `hooks.Plugin` interface.
2. Registers hooks dynamically using `RegisterDynamicHook`.
3. Optionally registers parsers for NLP-based input handling.

### Simulations
Simulations for operations like save and delete are provided in `pkg/utils/simulations.go`. These simulate hook execution scenarios for testing.

---

## Example Configuration (`plugin_config.json`)

```json
[
  {
    "name": "UpdateUsernamePlugin",
    "priority": 15,
    "enabled": true,
    "hooks": ["update_username"]
  },
  {
    "name": "UpdateEmailPlugin",
    "priority": 15,
    "enabled": true,
    "hooks": ["update_email"]
  },
  {
    "name": "TwoFactorPlugin",
    "priority": 25,
    "enabled": true,
    "hooks": ["enable_two_factor"]
  },
  {
    "name": "NotificationPlugin",
    "priority": 40,
    "enabled": true,
    "hooks": ["add_notification", "disable_notifications"]
  },
  {
    "name": "DeleteProtectionPlugin",
    "priority": 50,
    "enabled": true,
    "hooks": ["before_delete"]
  }
]
```

---

## Data Structures

### `hooks.Plugin`
```go
type Plugin interface {
    Name() string
    GetHookNames() []string
    RegisterHooks(hm *HookManager)
}
```

### `nlp.Intent`
```go
type Intent struct {
    Action string
    Params map[string]string
}
```

### `hooks.HookResult`
```go
type HookResult struct {
    StopExecution bool
    Error         error
}
```

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

- Ensure the `plugin_config.json` file is present in the root directory for plugin configuration.
- Plugins must implement the `hooks.Plugin` interface to be compatible with the framework.
- The framework dynamically loads plugins and hooks based on the configuration file and registered plugins.

---

## License

This project is licensed under the MIT License.# hook_framework
