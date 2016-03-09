package routers

import (
	"e2ee/controllers"
	"e2ee/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAccountRoutes(router *mux.Router) *mux.Router {
	router.Handle("/accountexists",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.AccountExists),
		)).Methods("GET")
	router.Handle("/account",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.AccountGet),
		)).Methods("GET")
	router.Handle("/account",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.AccountCreate),
		)).Methods("POST")

	return router
}
