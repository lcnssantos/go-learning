package provider

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"main/shared"
	"time"
)

type JwtProviderImpl struct {
}

func NewJwtProviderImpl() *JwtProviderImpl {
	return &JwtProviderImpl{}
}

func (this *JwtProviderImpl) Encode(id string, kind string, expirationTime int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  id,
		"kind": kind,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Second * time.Duration(expirationTime)).Unix(),
	})

	return token.SignedString([]byte(shared.GetEnvironmentConfiguration().JWT_SECRET))
}

func (this *JwtProviderImpl) Decode(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(shared.GetEnvironmentConfiguration().JWT_SECRET), nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return fmt.Sprintf("%v", claims["uid"]), fmt.Sprintf("%v", claims["kind"]), err
	} else {
		return "", "", err
	}
}
