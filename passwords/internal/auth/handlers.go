package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/fleischgewehr/crypto-labs/passwords/internal/app"
	"github.com/fleischgewehr/crypto-labs/passwords/internal/models"
)

func CreateUser(app *app.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		user := &models.User{}
		json.NewDecoder(r.Body).Decode(user)

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
