package routers

import (
	"e2ee/controllers"
	"e2ee/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetVersionCheckRoute(router *mux.Router) *mux.Router {
	router.Handle("/versioncheck",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.VersionCheck),
		)).Methods("GET")

	return router
}
