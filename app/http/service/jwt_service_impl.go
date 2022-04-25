package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/helper"
)

type (
	JWTServiceImpl struct {
		Header      Header
		Payload     Payload
		BasePayload BasePayload
		Secret      string
		Token       string
	}

	Header struct {
		Algorithm string `json:"alg"`
		Type      string `json:"typ"`
	}

	Payload map[string]interface{}

	BasePayload struct {
		IAT int64  `json:"iat"`
		SUB string `json:"sub"`
		EXP int64  `json:"exp"`
	}
)

func (jwt *JWTServiceImpl) GetToken() string {
	return jwt.Token
}

func (jwt *JWTServiceImpl) SetSecret(secret string) *JWTServiceImpl {
	jwt.Secret = secret
	return jwt
}

func (jwt *JWTServiceImpl) SetHeader(header Header) *JWTServiceImpl {
	jwt.Header = header
	return jwt
}

func (jwt *JWTServiceImpl) SetPayload(payload Payload) *JWTServiceImpl {
	jwt.Payload = payload
	return jwt
}

func (jwt *JWTServiceImpl) SetSub(sub string) *JWTServiceImpl {
	jwt.BasePayload.SUB = sub

	return jwt
}

func (jwt *JWTServiceImpl) SetExpired(unixExpDate int64) *JWTServiceImpl {
	jwt.BasePayload.EXP = unixExpDate
	return jwt
}

func (jwt *JWTServiceImpl) GenerateToken() (*JWTServiceImpl, error) {
	var theBasePayload Payload
	headerJson, err := json.Marshal(jwt.Header)
	if err != nil {
		return jwt, err
	}

	basePayloadJson, err := json.Marshal(jwt.BasePayload)
	if err != nil {
		return jwt, err
	}

	err = json.Unmarshal([]byte(basePayloadJson), &theBasePayload)
	if err != nil {
		return jwt, err
	}

	mergedPayload := helper.MergeMaps(jwt.Payload, theBasePayload)

	payloadJson, err := json.Marshal(mergedPayload)
	if err != nil {
		return jwt, err
	}

	base64Header := helper.Base64UrlEncoding(string(headerJson))
	base64Payload := helper.Base64UrlEncoding(string(payloadJson))

	jwt.Token = generateSignature(base64Header, base64Payload, jwt.Secret)

	return jwt, nil
}

func (jwt *JWTServiceImpl) ValidateToken(fullToken string) error {
	var payload BasePayload

	separateBearerToken := strings.Split(fullToken, " ")

	if len(separateBearerToken) != 2 {
		return errors.New("invalid token")
	}

	token := separateBearerToken[1]

	// Validate token (must be header.payload.signature) type, check by length
	split := strings.Split(token, ".")

	if len(split) != 3 {
		return errors.New("invalid token")
	}

	// Validate token signature
	sigtoken := generateSignature(split[0], split[1], jwt.Secret)
	if sigtoken != token {
		return errors.New("invalid token")
	}

	// Token is valid at this step
	// Now check the token expired date
	bytePayload, err := helper.Base64UrlDecoding(split[1])

	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytePayload, &payload); err != nil {
		return err
	}

	if payload.EXP < time.Now().Unix() {
		return errors.New("token expired")
	}

	return nil
}

func generateSignature(base64Header string, base64Payload string, secret string) string {
	hashSignature := hmac.New(sha256.New, []byte(secret))
	hashSignature.Write([]byte(base64Header + "." + base64Payload))

	signature := helper.Base64UrlEncoding(string(hashSignature.Sum(nil)))

	return base64Header + "." + base64Payload + "." + signature
}
