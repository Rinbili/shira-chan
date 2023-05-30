package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

const TokenExpireDuration = 24 * time.Hour

var mySigningKey = []byte(os.Getenv("secret"))

type JwtCustomClaims struct {
	UId   int
	Level string
	UName string
	jwt.RegisteredClaims
}

func GetToken(UId int, UName string, Level string) (string, error) {
	claims := JwtCustomClaims{
		UId:   UId,
		UName: UName,
		Level: Level,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Kyaru",                                                 //签发者
			Subject:   UName,                                                   //签发对象
			Audience:  jwt.ClaimStrings{"web"},                                 //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), //过期日期
			NotBefore: jwt.NewNumericDate(time.Now()),                          //启用日期
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          //签发时间
			ID:        uuid.NewString(),                                        //wt ID
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySigningKey)
	return token, err
}

func TokenParse(tokenString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
