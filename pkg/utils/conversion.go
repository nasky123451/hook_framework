package utils

// ConvertParamsToInterface 將 map[string]string 類型轉換為 map[string]interface{}
func ConvertParamsToInterface(params map[string]string) map[string]interface{} {
	converted := make(map[string]interface{})
	for key, value := range params {
		converted[key] = value
	}
	return converted
}
