package hooks

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// --------------------
// 全局插件註冊表與鎖
// --------------------

var (
	mu             sync.RWMutex
	pluginRegistry = make(map[string]func() Plugin)
)

// toPascalCase 轉換字串為 PascalCase（底線分隔詞）
func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, "")
}

// normalizePluginName 統一插件名稱格式（可擴充）
func normalizePluginName(name string) string {
	return toPascalCase(name)
}

// --------------------
// 插件註冊相關函式
// --------------------

// RegisterPlugin 註冊插件工廠函式，名稱會被標準化為 PascalCase
// 如果重複註冊會回傳錯誤
func RegisterPlugin(name string, factory func() Plugin) error {
	RegisterPluginType(name, factory)
	return nil
}

// UnregisterPlugin 取消註冊指定名稱插件，名稱會被標準化為 PascalCase
// 找不到插件時回傳錯誤
func UnregisterPlugin(name string) error {
	nName := normalizePluginName(name)

	mu.Lock()
	defer mu.Unlock()

	if _, exists := pluginRegistry[nName]; !exists {
		return fmt.Errorf("plugin '%s' is not registered", nName)
	}
	delete(pluginRegistry, nName)
	return nil
}

// RegisterPluginType 註冊插件類型
func RegisterPluginType(name string, factory func() Plugin) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := pluginRegistry[name]; exists {
		panic(fmt.Sprintf("Plugin type '%s' is already registered.", name))
	}
	pluginRegistry[name] = factory
}

// UnregisterPluginType 根據名稱取消註冊插件類型
func UnregisterPluginType(name string) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := pluginRegistry[name]; !exists {
		panic(fmt.Sprintf("Plugin type '%s' is not registered.", name))
	}
	delete(pluginRegistry, name)
}

// GetRegisteredPluginTypes 返回所有已註冊的插件類型
func GetRegisteredPluginTypes() []Plugin {
	mu.Lock()
	defer mu.Unlock()

	plugins := make([]Plugin, 0, len(pluginRegistry))
	for _, factory := range pluginRegistry {
		plugins = append(plugins, factory())
	}
	return plugins
}

// IsPluginRegistered 判斷插件是否已註冊，大小寫不敏感
func IsPluginRegistered(name string) bool {
	nName := normalizePluginName(name)

	mu.RLock()
	defer mu.RUnlock()

	_, exists := pluginRegistry[nName]
	return exists
}

// --------------------
// 插件獲取相關函式
// --------------------

// GetPluginNames 回傳所有已註冊插件名稱（標準化格式）
func GetPluginNames() []string {
	mu.RLock()
	defer mu.RUnlock()

	names := make([]string, 0, len(pluginRegistry))
	for name := range pluginRegistry {
		names = append(names, name)
	}
	return names
}

// GetPluginByName 根據插件名稱獲取一個新實例（大小寫不敏感）
func GetPluginByName(name string) (Plugin, error) {
	nName := normalizePluginName(name)

	mu.RLock()
	factory, exists := pluginRegistry[nName]
	mu.RUnlock()

	if !exists {
		// 支援大小寫不敏感模糊匹配
		mu.RLock()
		for regName, f := range pluginRegistry {
			if strings.EqualFold(regName, name) {
				factory = f
				exists = true
				break
			}
		}
		mu.RUnlock()

		if !exists {
			return nil, fmt.Errorf("plugin '%s' not found. Available plugins: %v", name, GetPluginNames())
		}
	}

	plugin := factory()
	if plugin == nil {
		return nil, fmt.Errorf("plugin factory for '%s' returned nil", name)
	}
	return plugin, nil
}

// LoadAllPlugins 創建所有已註冊插件實例並返回
func LoadAllPlugins() ([]Plugin, error) {
	mu.RLock()
	factories := make([]func() Plugin, 0, len(pluginRegistry))
	for _, f := range pluginRegistry {
		factories = append(factories, f)
	}
	mu.RUnlock()

	plugins := make([]Plugin, 0, len(factories))
	for _, factory := range factories {
		p := factory()
		if p == nil {
			return nil, fmt.Errorf("plugin factory returned nil")
		}
		plugins = append(plugins, p)
	}
	return plugins, nil
}

// --------------------
// 插件初始化：讀取資料夾檔案名稱並回傳插件名稱列表
// --------------------

// ScanPluginsFromDir 從指定目錄掃描符合條件的插件檔案名稱並回傳標準化的插件名稱
// 注意：此函式不會自動註冊插件，只掃描檔名，真正的註冊需外部呼叫 RegisterPlugin
func ScanPluginsFromDir(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read plugin directory '%s': %w", dir, err)
	}

	var pluginNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if strings.HasSuffix(name, "_plugin.go") {
			pluginName := strings.TrimSuffix(name, "_plugin.go")
			pluginNames = append(pluginNames, normalizePluginName(pluginName)+"Plugin")
		}
	}
	return pluginNames, nil
}

// --------------------
// 範例：初始化插件註冊表 (示範，需由外部呼叫)
// --------------------

// InitializePluginRegistry 從預設資料夾掃描插件檔案名稱，並回傳掃描結果
// 不直接註冊插件，由外部依需求決定如何註冊（例如 import 對應包並呼叫 RegisterPlugin）
func InitializePluginRegistry() ([]string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}
	metaFolder := filepath.Join(workingDir, "internal/plugins/meta")

	pluginNames, err := ScanPluginsFromDir(metaFolder)
	if err != nil {
		return nil, err
	}

	log.Printf("[InitializePluginRegistry] Scanned plugins: %v\n", pluginNames)
	return pluginNames, nil
}

func LoadPluginsFromRegistry() ([]Plugin, error) {
	mu.RLock()
	factories := make([]func() Plugin, 0, len(pluginRegistry))
	for _, factory := range pluginRegistry {
		factories = append(factories, factory)
	}
	mu.RUnlock()

	var plugins []Plugin
	for _, factory := range factories {
		plugin := factory()
		if plugin == nil {
			return nil, fmt.Errorf("failed to create plugin from factory")
		}
		plugins = append(plugins, plugin)
	}
	return plugins, nil
}

// GetPluginFactoryByName 回傳 pluginFactory，若找不到則回傳 nil 和 false
func GetPluginFactoryByName(name string) (func() Plugin, bool) {
	mu.RLock()
	defer mu.RUnlock()

	factory, ok := pluginRegistry[normalizePluginName(name)]
	if ok {
		return factory, true
	}

	// 支援大小寫不敏感模糊匹配
	for k, f := range pluginRegistry {
		if strings.EqualFold(k, name) {
			return f, true
		}
	}
	return nil, false
}
