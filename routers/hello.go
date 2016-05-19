package routers

import (
	"github.com/xlab-si/e2ee-server/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
