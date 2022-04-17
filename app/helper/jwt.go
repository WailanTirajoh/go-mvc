package helper

import (
	"encoding/json"
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
