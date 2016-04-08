package main

import (
	"github.com/xlab-si/e2ee-server/routers"
	"github.com/mancabizjak/e2ee-server/settings"
	"github.com/xlab-si/e2ee-server/core/db"
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	c := cors.New(cors.Options{
                AllowedOrigins:   []string{"*"},
                AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
                AllowCredentials: true,
		//AllowedHeaders: []string{"Authorization"}, // doesn't work
		AllowedHeaders:   []string{"*"},
		//Debug: true,
        })

	settings.Init()
	db.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)
	http.ListenAndServe(":8080", n)
}
