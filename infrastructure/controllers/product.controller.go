package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"product/infrastructure/responses"
	models "product/model"
	"product/repository/broker"
	"product/repository/db"
	searchrepo "product/repository/search"
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
		result, err := db.InsertOne(ctx, newProduct, "products")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}
		newProduct.Id = result
		searchrepo.Create(ctx, result, newProduct, "products")
		createdProduct, _ := json.Marshal(newProduct)
		broker.Publish("Product.Created", createdProduct)
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

		err := db.FindById(ctx, productId, "products", &product)
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

func GetProducts() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		limit, _ := strconv.Atoi(query.Get("limit"))
		page, _ := strconv.Atoi(query.Get("page"))
		search := query.Get("search")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var result []models.Product
		if search != "" {
			searchrepo.Search("products", search, &result, limit, (page-1)*limit)
		} else {
			err := db.Find(ctx, "products", nil, (page-1)*limit, limit, &result)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
				return
			}
		}
		rw.WriteHeader(http.StatusOK)
		response := responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}
