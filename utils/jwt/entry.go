package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayLoad struct {
	NickName string `json:"nickName"`
	RoleId   uint   `json:"roleId"`
	UserId   uint   `json:"userId"`
	// Password string `json:"-"`
}

var Secret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
