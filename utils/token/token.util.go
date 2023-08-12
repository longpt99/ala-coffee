package token

import (
	"ecommerce/errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"ecommerce/configs"
)

var (
	jwtKey = []byte(configs.Env.JWT.SecretKey)
)

type Claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func SignToken(id string) string {
	claims := &Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(configs.Env.JWT.ExpiresTime) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	return tokenString
}

func VerifyToken(signedToken string) (*Claims, error) {
	op := errors.Op("VerifyToken")
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		fmt.Println(err.Error())
		// if err.Error() == fmt.Sprintf("%s: %s", jwt.ErrTokenInvalidClaims.Error(), jwt.ErrTokenExpired.Error()) {
		return nil, errors.E(op, http.StatusUnauthorized, strings.TrimSpace(strings.Split(err.Error(), ":")[1]))
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.E(op, http.StatusBadRequest, "couldn't parse claims")
	}

	return claims, nil
}

func GetPayload(r *http.Request) *Claims {
	claims, ok := r.Context().Value("claims").(*Claims)
	if !ok {
		return nil
	}

	return claims
}
