package config

type JWTConfig struct {
	Lifetime  string
	Secret    string
	Algorithm string
}

func NewJwt() JWTConfig {
	return JWTConfig{
		Lifetime:  GetEnv("JWT_LIFETIME", "1"),
		Secret:    GetEnv("JWT_SECRET", "mysupersecret"),
		Algorithm: GetEnv("JWT_ALGORITHM", "HS256"),
	}
}
