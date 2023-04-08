package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"youtube-clone/cmd/api/handler/middleware"
	loginhnd "youtube-clone/cmd/api/handler/v1/login"
	userhnd "youtube-clone/cmd/api/handler/v1/user"
	"youtube-clone/pkg/repository/gormadp"
	loginserv "youtube-clone/service/login"
	usersrv "youtube-clone/service/user"
)

const (
	jwtSecret  = "verysecretkey"
	jwtExpHour = int64(24)
)

type V1Handlers struct {
	login loginhnd.Interface
	user  userhnd.Interface
	mw    middleware.Interface
}

func registerV1Handlers(dbConn *gorm.DB) (*V1Handlers, error) {
	var err error

	repo := gormadp.NewRepository(dbConn)

	loginService := loginserv.New(repo, jwtSecret, jwtExpHour)

	var loginHandler *loginhnd.Login

	if loginHandler, err = loginhnd.New(loginService); err != nil {
		return nil, fmt.Errorf("register v1 handlers : %w", err)
	}

	var userService *usersrv.User

	if userService, err = usersrv.New(repo, jwtSecret); err != nil {
		return nil, fmt.Errorf("register v1 handlers : %w", err)
	}

	var userHandler *userhnd.User

	if userHandler, err = userhnd.New(userService); err != nil {
		return nil, fmt.Errorf("register v1 handlers : %w", err)
	}

	var mw *middleware.MiddleWare

	if mw, err = middleware.New(jwtSecret, jwtExpHour); err != nil {
		return nil, fmt.Errorf("register v1 handlers : %w", err)
	}

	return &V1Handlers{
		login: loginHandler,
		user:  userHandler,
		mw:    mw,
	}, nil
}

func linkV1Routes(r *chi.Mux, h *V1Handlers) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("YouTube Clone API is running!"))
	})

	// Login
	r.Post("/api/v1/auth/login", h.login.Login)

	r.Post("/api/v1/user", h.user.CreateUser)

	r.Group(func(r chi.Router) {
		r.Use(h.mw.Auth)
		r.Get("/api/v1/profile", h.user.Profile)
	})

	r.Get("/api/v1/user/{id}", h.user.GetUser)
}
