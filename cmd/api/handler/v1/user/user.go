package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"youtube-clone/pkg/contextmanager"
	"youtube-clone/pkg/validator"
	"youtube-clone/service/user"
)

var _ Interface = (*User)(nil)

type User struct {
	userService user.Interface
}

func New(us user.Interface) (*User, error) {
	if us == nil {
		return nil, fmt.Errorf("new user handler / user : %w", ErrServiceShouldBeSetUp)
	}

	return &User{userService: us}, nil
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		icBody user.CreateUserModel
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

	if icBody.Password != icBody.ConfirmPassword {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error": "passwords must be equal",
		})

		return
	}

	if err = u.userService.CreateUser(r.Context(), &icBody); err != nil {
		if errors.Is(err, validator.ErrInvalidEmail) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": validator.ErrInvalidEmail.Error(),
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

		if strings.Contains(err.Error(), "email") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "email already using",
			})

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"error": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (u *User) Profile(w http.ResponseWriter, r *http.Request) {
	var (
		cm  *contextmanager.ContextManager
		err error
	)

	if cm, err = contextmanager.New(r.Context()); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"error": err.Error(),
		})

		return
	}

	var profileResp *user.GetUserModel

	if profileResp, err = u.userService.GetUserByID(r.Context(), cm.GetUserID()); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]any{
				"error": err.Error(),
			})

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"error": err.Error(),
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profileResp)
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")

	var (
		userID uint64
		err    error
	)

	if userID, err = strconv.ParseUint(userIDStr, 10, 64); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error": ErrInvalidParam.Error(),
		})

		return
	}

	var userResp *user.GetUserModel

	if userResp, err = u.userService.GetUserByID(r.Context(), uint(userID)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]any{
				"error": gorm.ErrRecordNotFound.Error(),
			})

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"error": err.Error(),
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResp)
}
