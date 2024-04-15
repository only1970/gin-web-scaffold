package jwt

import (
	"gvd_server/global"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// 生成token
func GenToken(user JwtPayLoad) (string, error) {
	Secret = []byte(global.Config.JWT.Secret)
	claims := CustomClaims{
		JwtPayLoad: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(global.Config.JWT.Expires) * time.Hour)),
			Issuer:    global.Config.JWT.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}
