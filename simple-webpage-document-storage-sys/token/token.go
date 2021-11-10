package token

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt"
	"simple-webpage-document-storage-sys/logging"
	"time"
)

const signingKey = "simple"

type MyClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

// GenerateToken generates a token containing the user id and the expiration data; by default the token is valid for an hour
func GenerateToken(uid string) (string, error) {
	expire := time.Now().Add(time.Hour)
	claims := MyClaims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

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
	logging.ConditionalLogError(err)
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims.Uid, true
	}
	return "", false
}
