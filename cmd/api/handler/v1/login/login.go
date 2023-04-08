package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"youtube-clone/pkg/validator"
	"youtube-clone/service/login"
)

var _ Interface = (*Login)(nil)

type Login struct {
	loginService login.Interface
}

func New(l login.Interface) (*Login, error) {
	if l == nil {
		return nil, fmt.Errorf("new login handler / login service : %w", ErrServiceShouldBeSetUp)
	}

	return &Login{
		loginService: l,
	}, nil
}

func (l *Login) Login(w http.ResponseWriter, r *http.Request) {
	var (
		icBody EmailAndPasswordLogin
		err    error
	)

	if err = json.NewDecoder(r.Body).Decode(&icBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error": err.Error(),
		})

		return
	}

	var token string

	if token, err = l.loginService.Login(r.Context(), icBody.Email, icBody.Password); err != nil {
		if errors.Is(err, validator.ErrInvalidEmail) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": errors.Unwrap(err).Error(),
			})

			return
		}

		if errors.Is(err, validator.ErrInvalidRange) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "password must be chars between 6 and 32",
			})

			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, login.ErrPasswordMatch) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": ErrWrongEmailOrPassword.Error(),
			})

			return
		}

		// Hatanın tipini bilmiyorsam 500 olarak döndürürüm
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"error": err.Error(),
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"token": token,
	})
}
