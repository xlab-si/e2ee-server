package routers

import (
	"github.com/xlab-si/e2ee-server/controllers"
	"github.com/xlab-si/e2ee-server/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAccountRoutes(router *mux.Router, authenticationRequired bool) *mux.Router {
	if (authenticationRequired) {
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
	} else {
		router.Handle("/accountexists",
			negroni.New(
				negroni.HandlerFunc(controllers.AccountExists),
			)).Methods("GET")
		router.Handle("/account",
			negroni.New(
				negroni.HandlerFunc(controllers.AccountGet),
			)).Methods("GET")
		router.Handle("/account",
			negroni.New(
				negroni.HandlerFunc(controllers.AccountCreate),
			)).Methods("POST")

	}

	return router
}
