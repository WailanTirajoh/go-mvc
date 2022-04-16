package helper

import (
	"encoding/base64"
	"strings"
)

func Base64StdEncoding(data string) string {
	base64str := base64.StdEncoding.EncodeToString([]byte(data))
	return cleanUp(base64str)
}

func cleanUp(s string) string {
	s = strings.ReplaceAll(s, "+", "-")
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "=", "")
	return s
}
