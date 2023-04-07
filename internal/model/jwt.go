package model

import "github.com/golang-jwt/jwt"

type Claims struct {
	UID        string `json:"uid"`
	UserName   string `json:"userName"`
	Phone      string `json:"phone"`
	RoleID     string `json:"roleId"`
	RoleName   string `json:"roleName"`
	State      int    `json:"state"`
	BufferTime int64  `json:"bufferTime"`
	jwt.StandardClaims
}
