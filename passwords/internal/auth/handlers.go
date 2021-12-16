package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/fleischgewehr/crypto-labs/passwords/internal/app"
	"github.com/fleischgewehr/crypto-labs/passwords/internal/models"
)

type registrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func CreateUser(app *app.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		registrationReq := &registrationRequest{}
		json.NewDecoder(r.Body).Decode(registrationReq)

		if err := validatePassword(registrationReq.Password); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Validation errror: %q", err.Error())
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

		phone, _ := EncryptPhone(registrationReq.Phone)
		user.Phone = phone.Cipher
		user.PhoneSalt = phone.Salt
		address, _ := EncryptAddress(registrationReq.Address)
		user.Address = address.Cipher
		user.AddressSalt = address.Salt

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
			time.Sleep(2 * time.Second)
			// if err := HandleInvalidPassword(app, user.Username); err != nil {
			// 	fmt.Fprintf(w, "Error: %q", err.Error())
			// 	return
			// }
			if user.InvalidLoginCount > 40 {
				user.BlockAccount(r.Context(), app)
				fmt.Fprintf(w, "Your account has been blocked")
				return
			}
			user.UpsertInvalidLoginCount(r.Context(), app)
			fmt.Fprintf(w, "Invalid login or password")
		}
	}
}

type profileResponse struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func GetProfile(app *app.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid ID")
			return
		}
		user := &models.User{ID: id}
		if err := user.GetById(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, err.Error())
			return
		}

		encryptedPhone := &CookedCipher{Cipher: user.Phone, Salt: user.PhoneSalt}
		phone, err := DecryptPhone(encryptedPhone)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
		}
		encryptedAddress := &CookedCipher{Cipher: user.Address, Salt: user.AddressSalt}
		address, err := DecryptAddress(encryptedAddress)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
		}
		resp := &profileResponse{Username: user.Username, Phone: phone, Address: address}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
