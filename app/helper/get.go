package helper

import "errors"

func GetStrKey(m map[string]interface{}, key string) (string, error) {
	value, exists := m[key]
	if !exists {
		return "", errors.New("sub is required")
	}
	strVal, ok := value.(string)
	if !ok {
		return "", errors.New("sub is not string")
	}

	return strVal, nil
}
