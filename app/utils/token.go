package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

const TokenExpireDuration = 24 * time.Hour

var mySigningKey = []byte(os.Getenv("secret"))

type JwtCustomClaims struct {
	UId     int
	IsAdmin bool
	jwt.RegisteredClaims
}

func GetToken(UId int, IsAdmin bool) (string, error) {
	claims := JwtCustomClaims{
		UId:     UId,
		IsAdmin: IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Kyaru",                                                 //签发者
			Subject:   strconv.Itoa(UId),                                       //签发对象
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

func ParseToken(tokenString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
