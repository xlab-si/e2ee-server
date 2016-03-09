package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetVersionCheckRoute(router)
	router = SetAuthenticationRoutes(router)
	router = SetAccountRoutes(router)
	router = SetContainerRoutes(router)
	router = SetPeerRoutes(router)
	router = SetNotificationRoutes(router)
	return router
}
