package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"product/infrastructure/config"
	"product/infrastructure/routes"
)

func main() {
	router := mux.NewRouter()

	//run database
	config.ConnectDB()

	config.ConnectCache()

	config.ConnectBroker()

	config.ConnectSearch()

	//routes
	routes.ProductRoute(router) //add this

	log.Fatal(http.ListenAndServe(":6000", router))
}
