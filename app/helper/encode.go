package helper

import (
	"encoding/base64"
)

func Base64UrlEncoding(s string) string {
	base64str := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString([]byte(s))
	return base64str
}

func Base64UrlDecoding(s string) ([]byte, error) {
	return base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(s)
}
