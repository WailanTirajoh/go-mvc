package helper

import (
	"encoding/json"
	"errors"
	"strings"
)

func GetUserPayload(token string) (map[string]interface{}, error) {
	var payload map[string]interface{}

	split := strings.Split(token, ".")
	bytePayload, err := Base64UrlDecoding(split[1])

	if err != nil {
		return payload, err
	}

	if err := json.Unmarshal(bytePayload, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func GetStrKey(m map[string]interface{}, key string) (string, error) {
	value, exists := m[key]
	if !exists {
		return "", errors.New(key + " is required")
	}
	strVal, ok := value.(string)
	if !ok {
		return "", errors.New("error on " + key)
	}

	return strVal, nil
}
