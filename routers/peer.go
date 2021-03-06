package routers

import (
	"github.com/xlab-si/e2ee-server/controllers"
	"github.com/xlab-si/e2ee-server/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetPeerRoutes(router *mux.Router, authenticationRequired bool) *mux.Router {
	if (authenticationRequired) {
		router.Handle("/peer/{username}",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.PeerGet),
			)).Methods("GET")
		router.Handle("/peer",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.PeerNotify),
			)).Methods("POST")
	} else {
		router.Handle("/peer/{username}",
			negroni.New(
				negroni.HandlerFunc(controllers.PeerGet),
			)).Methods("GET")
		router.Handle("/peer",
			negroni.New(
				negroni.HandlerFunc(controllers.PeerNotify),
			)).Methods("POST")
	}

	return router
}
