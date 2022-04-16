package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"time"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/helper"
)

type (
	JWTService interface {
		SetPayload(payload map[string]interface{}) *JWTServiceImpl
		SetHeader(header Header) *JWTServiceImpl
		SetSecret(header string) *JWTServiceImpl
		SetExpired(unixExpDate int64) *JWTServiceImpl
		GenerateToken() (*JWTServiceImpl, error)
		GetToken() string
	}

	JWTServiceImpl struct {
		Type      string
		Algorithm string
		Header    Header
		Payload   map[string]interface{}
		Secret    string
		Token     string
		IAT       int64
		SUB       string
		EXP       int64
	}

	Header struct {
		Algorithm string `json:"alg"`
		Type      string `json:"typ"`
	}
)

func NewJWT(payload map[string]interface{}) JWTService {
	return &JWTServiceImpl{
		Header: Header{
			Algorithm: "HS256",
			Type:      "JWT",
		},
		Payload: payload,
		IAT:     time.Now().Unix(),
		EXP:     time.Now().Add(time.Second * 60).Unix(),
	}
}

func (jwt *JWTServiceImpl) GenerateToken() (*JWTServiceImpl, error) {
	if len(jwt.Payload) == 0 {
		return jwt, errors.New("payload is unset")
	}

	jwt.Payload = helper.MergeMaps(jwt.Payload, map[string]interface{}{
		"iat": jwt.IAT,
		"exp": jwt.EXP,
	})

	headerJson, err := json.Marshal(jwt.Header)
	if err != nil {
		return jwt, err
	}

	payloadJson, err := json.Marshal(jwt.Payload)
	if err != nil {
		return jwt, err
	}

	base64UrlHeader := helper.Base64StdEncoding(string(headerJson))
	base64UrlPayload := helper.Base64StdEncoding(string(payloadJson))
	secret := jwt.Secret

	hashSignature := hmac.New(sha256.New, []byte(secret))
	hashSignature.Write([]byte(base64UrlHeader + "." + base64UrlPayload))

	signature := helper.Base64StdEncoding(string(hashSignature.Sum(nil)))

	jwt.Token = base64UrlHeader + "." + base64UrlPayload + "." + signature

	return jwt, nil
}

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

func (jwt *JWTServiceImpl) SetPayload(payload map[string]interface{}) *JWTServiceImpl {
	jwt.Payload = payload
	return jwt
}

func (jwt *JWTServiceImpl) SetExpired(unixExpDate int64) *JWTServiceImpl {
	jwt.EXP = unixExpDate
	return jwt
}
