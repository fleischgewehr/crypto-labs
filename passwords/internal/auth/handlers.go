package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/fleischgewehr/crypto-labs/passwords/internal/app"
	"github.com/fleischgewehr/crypto-labs/passwords/internal/models"
)

type registrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(app *app.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		registrationReq := &registrationRequest{}
		json.NewDecoder(r.Body).Decode(registrationReq)

		user := &models.User{Username: registrationReq.Username}
		cookedPassword, err := HashPassword(registrationReq.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "An error occurred while creating the user")
			return
		}
		user.PasswordHash = cookedPassword.Hash
		user.PasswordSalt = cookedPassword.ArgonSalt

		if err := user.Create(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "An error occurred while creating the user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(user)
		w.Write(resp)
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
			fmt.Fprintf(w, "User not found")
			return
		}

		stored := &CookedPassword{Hash: user.PasswordHash, ArgonSalt: user.PasswordSalt}
		if CheckPassword(loginReq.Password, stored) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Logged in")
		} else {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Invalid login or password")
		}
	}
}
