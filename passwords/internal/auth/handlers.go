package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/validator.v2"

	"github.com/fleischgewehr/crypto-labs/passwords/internal/app"
	"github.com/fleischgewehr/crypto-labs/passwords/internal/models"
)

type registrationRequest struct {
	// Any string from 3 to 16 symbols
	Username string `json:"username" validate:"min=3,max=16"`
	// Any string with len >= 8 and containing at least one digit, lower/upper case char & punctuation
	Password string `json:"password" validate:"min=10,regexp=(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*\W)"`
}

func CreateUser(app *app.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		registrationReq := &registrationRequest{}
		json.NewDecoder(r.Body).Decode(registrationReq)

		if err := validator.Validate(registrationReq); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err.Error())
			fmt.Fprintf(w, "Validation error: %q", err.Error())
			return
		}

		user := &models.User{Username: registrationReq.Username}
		if err := user.GetByUsername(r.Context(), app); err == nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "User with given login already exists")
			return
		}

		cookedPassword, err := HashPassword(registrationReq.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Invalid password: %q", err.Error())
			return
		}
		user.PasswordHash = cookedPassword.Hash
		user.PasswordSalt = cookedPassword.ArgonSalt

		if err := user.Create(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not create user: %q", err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(app *app.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		loginReq := &loginRequest{}
		json.NewDecoder(r.Body).Decode(loginReq)

		user := &models.User{Username: loginReq.Username}
		if err := user.GetByUsername(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Invalid login or password")
			return
		}

		stored := &CookedPassword{Hash: user.PasswordHash, ArgonSalt: user.PasswordSalt}
		if CheckPassword(loginReq.Password, stored) {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Invalid login or password")
		}
	}
}
