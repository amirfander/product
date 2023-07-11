package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"product/config"
	"product/routes"
)

func main() {
	router := mux.NewRouter()

	//run database
	config.ConnectDB()

	config.ConnectCache()

	//routes
	routes.ProductRoute(router) //add this

	log.Fatal(http.ListenAndServe(":6000", router))
}
