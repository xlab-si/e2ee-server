package routers

import (
	"e2ee/controllers"
	"e2ee/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetNotificationRoutes(router *mux.Router) *mux.Router {
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
	
	return router
}
