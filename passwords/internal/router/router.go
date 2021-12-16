package router

import (
	"github.com/julienschmidt/httprouter"

	"github.com/fleischgewehr/crypto-labs/passwords/internal/app"
	"github.com/fleischgewehr/crypto-labs/passwords/internal/auth"
)

func Get(app *app.Application) *httprouter.Router {
	mux := httprouter.New()

	mux.POST("/users", auth.CreateUser(app))
	mux.POST("/users/login", auth.Login(app))
	mux.GET("/users/:id", auth.GetProfile(app))

	return mux
}
