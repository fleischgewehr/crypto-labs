package auth

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/fleischgewehr/crypto-labs/passwords/internal/app"
)

func CreateUser(app *app.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Yo")
	}
}
