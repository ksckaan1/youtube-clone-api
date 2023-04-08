package login

import "net/http"

type Interface interface {
	Login(w http.ResponseWriter, r *http.Request)
}
