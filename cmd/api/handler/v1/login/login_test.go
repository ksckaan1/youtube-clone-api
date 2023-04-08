package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"testing"
	"youtube-clone/pkg/repository/gormadp"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
	"youtube-clone/service/login"
)

const (
	jwtSecret  = "verysecretkey"
	jwtExpHour = int64(24)
)

func TestLogin_Login(t *testing.T) {
	db := gormadp.NewTestDBConnection(t)
	repo := gormadp.NewRepository(db)
	loginService := login.New(repo, jwtSecret, jwtExpHour)

	var (
		loginHnd *Login
		err      error
	)

	if loginHnd, err = New(loginService); err != nil {
		t.Error(err)
	}

	r := chi.NewRouter()

	r.Post("/api/v1/auth/login", loginHnd.Login)

	server := httptest.NewServer(r)

	t.Cleanup(func() {
		server.Close()
	})

	t.Run("CreateUser", func(t *testing.T) {
		testUser := dbmodels.User{
			Email:    "me@kaanksc.com",
			Password: "$2a$10$XQSil8wLLOjbhsuxnhsBdOH6Z8dXszXA9b2ELKLZeuN.DG13aTHSi", // asdf1234
			FullName: "Kaan Kuscu",
		}

		if err = db.Create(&testUser).Error; err != nil {
			t.Error(err)
		}
	})

	loginTests := []struct {
		description  string
		email        string
		password     string
		expectedCode int
	}{
		// invalid
		{
			description:  "WrongEmailAndPassword",
			email:        "abc@gmail.com",
			password:     "asdasdasd",
			expectedCode: http.StatusBadRequest,
		},
		// valid
		{
			description:  "LoginSuccess",
			email:        "me@kaanksc.com",
			password:     "asdf1234",
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range loginTests {
		t.Run(fmt.Sprintf("Login_%s", tt.description), func(t *testing.T) {
			reqBody := EmailAndPasswordLogin{
				Email:    tt.email,
				Password: tt.password,
			}

			var body []byte

			if body, err = json.Marshal(reqBody); err != nil {
				t.Error(err)
			}

			client := &http.Client{}

			var req *http.Request

			if req, err = http.NewRequest("POST", server.URL+"/api/v1/auth/login", bytes.NewBuffer(body)); err != nil {
				t.Error(err)
			}

			req.Header.Set("Content-Type", "application/json")

			var resp *http.Response

			if resp, err = client.Do(req); err != nil {
				t.Error(err)
			}

			if tt.expectedCode != resp.StatusCode {
				t.Errorf("wrong status code : %d", resp.StatusCode)
			}
		})
	}
}
