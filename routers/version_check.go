package routers

import (
	"github.com/xlab-si/e2ee-server/controllers"
	"github.com/xlab-si/e2ee-server/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetVersionCheckRoute(router *mux.Router, authenticationRequired bool) *mux.Router {
	if (authenticationRequired) {
		router.Handle("/versioncheck",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.VersionCheck),
			)).Methods("GET")
	} else {
		router.Handle("/versioncheck",
			negroni.New(
				negroni.HandlerFunc(controllers.VersionCheck),
			)).Methods("GET")
	}

	return router
}
