package structs

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserName  string `json:"userName"`
	UserId    int    `json:"userId"`
	UserRoles string `json:"userRoles"`
	jwt.StandardClaims
}
