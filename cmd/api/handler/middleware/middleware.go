package middleware

import (
	"youtube-clone/pkg/jwt"
	"youtube-clone/pkg/jwt/golang_jwt"
)

var _ Interface = (*MiddleWare)(nil)

type MiddleWare struct {
	jwt jwt.Interface
}

func New(secret string, exp int64) (*MiddleWare, error) {
	j := golang_jwt.New(secret, exp)

	return &MiddleWare{jwt: j}, nil
}
