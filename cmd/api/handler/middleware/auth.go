package middleware

import (
	"context"
	"encoding/json"
	golang_jwt "github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
)

func (mw *MiddleWare) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.ReplaceAll(r.Header.Get("Authorization"), "Bearer ", "")
		var (
			claims *golang_jwt.RegisteredClaims
			err    error
		)

		if claims, err = mw.jwt.Parse(token); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{
				"error": err.Error(),
			})

			return
		}

		var userID uint64

		if userID, err = strconv.ParseUint(claims.ID, 10, 64); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "user id can not converted",
			})

			return
		}

		ctx1 := context.WithValue(r.Context(), "id", uint(userID))
		ctx2 := context.WithValue(ctx1, "is_logged_in", true)

		next.ServeHTTP(w, r.WithContext(ctx2))
	})
}
