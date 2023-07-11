package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"product/dbrepository"
	"product/models"
	"product/responses"
)

func CreateProduct() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var product models.Product
		defer cancel()

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		newProduct := models.Product{
			Title:       product.Title,
			Category:    product.Category,
			Tags:        product.Tags,
			Description: product.Description,
		}
		result, err := dbrepository.InsertOne(ctx, newProduct, "products")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetAProduct() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		productId := params["productId"]
		var product models.Product
		defer cancel()

		err := dbrepository.FindById(ctx, productId, "products", &product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": product}}
		json.NewEncoder(rw).Encode(response)
	}
}
