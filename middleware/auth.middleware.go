package middleware

import (
	"context"
	"ecommerce/errors"

	res "ecommerce/utils/response"
	t "ecommerce/utils/token"

	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		op := errors.Op("Authorization")

		if token == "" {
			res.WriteError(w, r, errors.E(op, http.StatusUnauthorized))
			return
		}

		tokenString := strings.Split(token, " ")
		claims, err := t.VerifyToken(tokenString[1])

		if err != nil {
			res.WriteError(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
