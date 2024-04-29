package util

import (
	"crypto/rand"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey []byte

// 生成一个安全的随机密钥
func generateJWTKey() []byte {
	key := make([]byte, 32) // 32字节密钥
	_, err := rand.Read(key)
	if err != nil {
		// 处理错误
	}
	return key
}

type TokenClaims struct {
	Id       int64
	Username string
	jwt.StandardClaims
}

func GenerateToken(id int64, username string) (token string, err error) {
	// 在使用GenerateToken函数之前生成jwtKey
	jwtKey = generateJWTKey()
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	claims := TokenClaims{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ByteDance-Tiny-Douyin",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
}

func ParseToken(token string) (*TokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*TokenClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
