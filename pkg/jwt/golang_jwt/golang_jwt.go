package golang_jwt

import (
	"errors"
	"fmt"
	golang_jwt "github.com/golang-jwt/jwt/v5"
	"time"
	"youtube-clone/pkg/jwt"
)

var _ jwt.Interface = (*JWT)(nil)

type JWT struct {
	secret string
	exp    int64
}

func New(secret string, exp int64) *JWT {
	return &JWT{
		secret: secret,
		exp:    exp,
	}
}

func (j *JWT) Generate(id uint) (string, error) {
	claim := golang_jwt.RegisteredClaims{
		Issuer:    "Company1",
		Subject:   "Authorization",
		ExpiresAt: golang_jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.exp))),
		ID:        fmt.Sprintf("%d", id),
	}

	var (
		token  *golang_jwt.Token
		signed string
		err    error
	)

	token = golang_jwt.NewWithClaims(golang_jwt.SigningMethodHS256, &claim)

	if signed, err = token.SignedString([]byte(j.secret)); err != nil {
		return "", fmt.Errorf("jwt generate: %w", err)
	}

	return signed, nil
}

func (j *JWT) Parse(token string) (*golang_jwt.RegisteredClaims, error) {
	var (
		tkn *golang_jwt.Token
		err error
	)

	if tkn, err = golang_jwt.ParseWithClaims(token, &golang_jwt.RegisteredClaims{}, func(token *golang_jwt.Token) (any, error) {
		return []byte(j.secret), nil
	}); err != nil || !tkn.Valid {
		return nil, fmt.Errorf("jwt parse: %w", err)
	}

	var (
		claim *golang_jwt.RegisteredClaims
		ok    bool
	)

	if claim, ok = tkn.Claims.(*golang_jwt.RegisteredClaims); !ok {
		return nil, errors.New("parse claims")
	}

	return claim, nil
}
