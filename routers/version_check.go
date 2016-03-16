package routers

import (
	"github.com/mancabizjak/e2ee-server/controllers"
	"github.com/mancabizjak/e2ee-server/core/authentication"
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
