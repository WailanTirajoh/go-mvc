package service

import (
	"strconv"
	"time"

	"github.com/WailanTirajoh/go-simple-clean-architecture/config"
)

type JWTService interface {
	// Setters
	SetPayload(payload Payload) *JWTServiceImpl
	SetHeader(header Header) *JWTServiceImpl
	SetSecret(header string) *JWTServiceImpl
	SetExpired(unixExpDate int64) *JWTServiceImpl
	SetSub(sub string) *JWTServiceImpl

	// Getters
	GetToken() string

	GenerateToken() (*JWTServiceImpl, error)
	ValidateToken(token string) error
}

func NewJWT() JWTService {
	jwtConfig := config.NewJwt()

	lifetime, err := strconv.Atoi(jwtConfig.Lifetime)
	if err != nil {
		panic(err)
	}

	return &JWTServiceImpl{
		Header: Header{
			Algorithm: jwtConfig.Algorithm,
			Type:      "JWT",
		},
		BasePayload: BasePayload{
			IAT: time.Now().Unix(),
			EXP: time.Now().Add(time.Second * 60 * time.Duration(lifetime)).Unix(),
		},
		Secret: jwtConfig.Secret,
	}
}
