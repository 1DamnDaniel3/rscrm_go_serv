package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAdapter struct {
	sectetKey string
	ttl       time.Duration
}

func NewJWTAdapter(secret string, ttl time.Duration) *JWTAdapter {
	return &JWTAdapter{sectetKey: secret, ttl: ttl}
}

func (j *JWTAdapter) Sign(claims map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenClaims := token.Claims.(jwt.MapClaims)

	for k, v := range claims {
		tokenClaims[k] = v
	}
	tokenClaims["exp"] = time.Now().Add(j.ttl).Unix()

	return token.SignedString([]byte(j.sectetKey))
}

func (j *JWTAdapter) Verify(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.sectetKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result := make(map[string]interface{})
		for k, v := range claims {
			result[k] = v
		}
		return result, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
