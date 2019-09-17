package structs

import (
  "github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	Id_user              string              `json:"id_user"`
  User                 string              `json:"user"`
	jwt.StandardClaims
}
