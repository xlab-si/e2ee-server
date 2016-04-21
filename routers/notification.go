package routers

import (
	"github.com/xlab-si/e2ee-server/controllers"
	"github.com/xlab-si/e2ee-server/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetNotificationRoutes(router *mux.Router, authenticationRequired bool) *mux.Router {
	if (authenticationRequired) {
		router.Handle("/messages",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.NotificationsGet),
			)).Methods("GET")
		router.Handle("/messages",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.NotificationsDelete),
			)).Methods("DELETE")
	} else {
		router.Handle("/messages",
			negroni.New(
				negroni.HandlerFunc(controllers.NotificationsGet),
			)).Methods("GET")
		router.Handle("/messages",
			negroni.New(
				negroni.HandlerFunc(controllers.NotificationsDelete),
			)).Methods("DELETE")
	}
	
	return router
}
