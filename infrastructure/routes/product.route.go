package routes

import (
	"github.com/gorilla/mux"

	"product/infrastructure/controllers"
)

func ProductRoute(router *mux.Router) {
	router.HandleFunc("/products", controllers.CreateProduct()).Methods("POST")
	router.HandleFunc("/products/{productId}", controllers.GetAProduct()).Methods("GET")
	router.HandleFunc("/products", controllers.GetProducts()).Methods("GET")
	router.HandleFunc("/products/{productId}", controllers.UpdateAProduct()).Methods("PUT")
	router.HandleFunc("/products/{productId}", controllers.DeleteAProduct()).Methods("DELETE")
}
