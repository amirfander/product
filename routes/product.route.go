package routes

import (
	"github.com/gorilla/mux"

	"product/controllers"
)

func ProductRoute(router *mux.Router) {
	router.HandleFunc("/products", controllers.CreateProduct()).Methods("POST")
	router.HandleFunc("/products/{productId}", controllers.GetAProduct()).Methods("GET")
	router.HandleFunc("/products", controllers.GetProducts()).Methods("GET")

	// router.HandleFunc("/user/{userId}", controllers.EditAUser()).Methods("PUT")
	// router.HandleFunc("/user/{userId}", controllers.DeleteAUser()).Methods("DELETE")
	// router.HandleFunc("/users", controllers.GetAllUser()).Methods("GET")
}
