package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
)

type Claims struct {
	UID        string `json:"uid"`
	UserName   string `json:"userName"`
	Phone      string `json:"phone"`
	RoleID     string `json:"roleId"`
	RoleName   string `json:"roleName"`
	State      int32  `json:"state"`
	BufferTime int32  `json:"bufferTime"`
	jwt.StandardClaims
}

// Create 生成token
func Create(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(consts.Conf.Jwt.SigningKey))
}

// Parse 解析token
func Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(consts.Conf.Jwt.SigningKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
