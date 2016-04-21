package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes(authenticationRequired bool) *mux.Router {
	router := mux.NewRouter()
	router = SetVersionCheckRoute(router, authenticationRequired)
	router = SetAccountRoutes(router, authenticationRequired)
	router = SetContainerRoutes(router, authenticationRequired)
	router = SetPeerRoutes(router, authenticationRequired)
	router = SetNotificationRoutes(router, authenticationRequired)
	return router
}
