package middleware

import "net/http"

type Interface interface {
	Auth(next http.Handler) http.Handler
	AuthOptional(next http.Handler) http.Handler
}
