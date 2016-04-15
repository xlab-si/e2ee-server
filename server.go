package main

import (
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"github.com/xlab-si/e2ee-server/config"
	"github.com/xlab-si/e2ee-server/core/db"
	"github.com/xlab-si/e2ee-server/routers"
	"github.com/xlab-si/e2ee-server/settings"
	"net/http"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		//AllowedHeaders: []string{"Authorization"}, // doesn't work
		AllowedHeaders: []string{"*"},
		//Debug: true,
	})

	config.Init()
	settings.Init()
	db.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)
	//http.ListenAndServe(":8080", n)
	http.ListenAndServeTLS(":4443", "server.crt", "server.key", n)
}
