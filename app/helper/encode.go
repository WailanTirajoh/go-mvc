package helper

import (
	"encoding/base64"
	"strings"
)

func Base64StdEncoding(s string) string {
	base64str := base64.StdEncoding.EncodeToString([]byte(s))
	return cleanUp(base64str)
}

func Base64StdDecoding(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func cleanUp(s string) string {
	s = strings.ReplaceAll(s, "+", "-")
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "=", "")
	return s
}
