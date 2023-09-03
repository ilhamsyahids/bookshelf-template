package utils

func NewSuccessResp(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"ok":   true,
		"data": data,
	}
}
