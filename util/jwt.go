package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var mySigningKey = []byte("a343be788860d10ef767c078884cccaf")

func CreateJwtToken(id string) (string ,error){
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour*24).Unix(),
		Issuer:    "neko",
		Subject: id,

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err!=nil {
		return "",err
	}
	return ss,nil
}

func ValidateJwtToken(tokenString string)( *jwt.Token,error ){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if token.Valid {
		return token,nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("token 无效")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("token 过期")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	return nil,err
}