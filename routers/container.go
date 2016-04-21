package routers

import (
	"github.com/xlab-si/e2ee-server/controllers"
	"github.com/xlab-si/e2ee-server/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetContainerRoutes(router *mux.Router, authenticationRequired bool) *mux.Router {
	if (authenticationRequired) {
		router.Handle("/container/{containerNameHmac}",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.ContainerGet),
			)).Methods("GET")
		router.Handle("/container/{containerNameHmac}",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.ContainerCreate),
			)).Methods("PUT")
		router.Handle("/container/{containerNameHmac}",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.ContainerDelete),
			)).Methods("DELETE")
		router.Handle("/container/record",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.ContainerRecordCreate),
			)).Methods("POST")
		router.Handle("/container/share",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.ContainerShare),
			)).Methods("POST")
		router.Handle("/container/unshare",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.ContainerUnshare),
			)).Methods("POST")
	} else {
		router.Handle("/container/{containerNameHmac}",
			negroni.New(
				negroni.HandlerFunc(controllers.ContainerGet),
			)).Methods("GET")
		router.Handle("/container/{containerNameHmac}",
			negroni.New(
				negroni.HandlerFunc(controllers.ContainerCreate),
			)).Methods("PUT")
		router.Handle("/container/{containerNameHmac}",
			negroni.New(
				negroni.HandlerFunc(controllers.ContainerDelete),
			)).Methods("DELETE")
		router.Handle("/container/record",
			negroni.New(
				negroni.HandlerFunc(controllers.ContainerRecordCreate),
			)).Methods("POST")
		router.Handle("/container/share",
			negroni.New(
				negroni.HandlerFunc(controllers.ContainerShare),
			)).Methods("POST")
		router.Handle("/container/unshare",
			negroni.New(
				negroni.HandlerFunc(controllers.ContainerUnshare),
			)).Methods("POST")

	}

 	return router
}
