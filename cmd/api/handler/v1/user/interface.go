package user

import (
	"net/http"
)

type Interface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}
