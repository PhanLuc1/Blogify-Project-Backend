package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignedDetails struct {
	UserId int
	jwt.RegisteredClaims
}

var SECRET_KEY = os.Getenv("SECRET_LOVE")

func TokenGeneration(userid int) (signedtoken string, err error) {
	claims := &SignedDetails{
		UserId: userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 100)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, err
}
func ValidateToken(signedtoken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "The Token is invalid"
		return
	}
	expireAt := claims.ExpiresAt.Time.Unix()
	if expireAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}
	return claims, msg
}
