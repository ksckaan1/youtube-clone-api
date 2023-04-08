package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

func main() {
	var (
		dbConn *gorm.DB
		err    error
	)

	if dbConn, err = newDBConnection(); err != nil {
		log.Fatalln(err)
	}

	var v1Handlers *V1Handlers

	if v1Handlers, err = registerV1Handlers(dbConn); err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	linkV1Routes(r, v1Handlers)

	PORT := "3000"

	if p, ok := os.LookupEnv("PORT"); ok {
		PORT = p
	}

	log.Printf("api started successfully on '%s' port\n", PORT)
	if err = http.ListenAndServe(fmt.Sprintf(":%s", PORT), r); err != nil {
		log.Fatalln(err)
	}
}

func newDBConnection() (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	if db, err = gorm.Open(sqlite.Open("youtube-clone.db"), &gorm.Config{}); err != nil {
		return nil, fmt.Errorf("new db conn / sqlite open : %w", err)
	}

	if err = db.AutoMigrate(&dbmodels.User{}); err != nil {
		return nil, fmt.Errorf("new db conn / auto migrate : %w", err)
	}

	return db, nil
}
