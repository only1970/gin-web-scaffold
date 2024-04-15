package jwt

import (
	"gvd_server/global"

	"github.com/dgrijalva/jwt-go/v4"
)

func ParseToken(token string) (*CustomClaims, error) {
	Token, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		return nil, err
	}
	if !Token.Valid {
		//令牌无效
		return nil, err
	}
	return claims, nil
}
