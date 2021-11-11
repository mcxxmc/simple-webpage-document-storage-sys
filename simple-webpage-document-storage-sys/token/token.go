package token

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt"
	"simple-webpage-document-storage-sys/logging"
	"time"
)

const signingKey = "simple"
const s1Token = "token"

type MyClaims struct {
	Uid []byte `json:"uid"`
	jwt.StandardClaims
}

// GenerateToken generates a token containing the user id and the expiration data; by default the token is valid for an hour
func GenerateToken(uid string) (string, error) {
	expire := time.Now().Add(time.Hour)
	claims := MyClaims{
		Uid: []byte(uid),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signingKey))
}

// DecodeToken decodes a token
//
// returns a string (uid) and a bool (true if valid);
func DecodeToken(tokenStr string) (string, bool) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		logging.Error(err, logging.S(s1Token, tokenStr))
		return "", false
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return string(claims.Uid), true
	}
	return "", false
}
