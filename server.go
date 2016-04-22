package main

import (
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"github.com/xlab-si/e2ee-server/config"
	"github.com/xlab-si/e2ee-server/core/db"
	"github.com/xlab-si/e2ee-server/routers"
	"net/http"
	"fmt"
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
	db.Init()
	router := routers.InitRoutes(true)
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)

	http_conf := viper.GetStringMap("https")
	port := fmt.Sprintf(":%s", http_conf["port"])
	key := fmt.Sprintf("%s/%s", http_conf["path"], http_conf["key"])
	cert := fmt.Sprintf("%s/%s", http_conf["path"], http_conf["cert"])
	
	err := http.ListenAndServeTLS(port, cert, key, n)
	fmt.Println(err)
}
