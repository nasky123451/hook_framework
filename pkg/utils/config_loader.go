package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadPluginConfig 加載插件配置文件
func LoadPluginConfig(configPath string) (interface{}, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config interface{} // 支持 map 或 array
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	fmt.Printf("[LoadPluginConfig] Loaded plugin configuration: %+v\n", config)
	return config, nil
}
