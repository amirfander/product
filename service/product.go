package service

import (
	"context"
	"encoding/json"

	"product/model"
	"product/repository/broker"
	"product/repository/db"
	searchrepo "product/repository/search"
)

func Create(ctx context.Context, product model.Product) (string, error) {
	result, err := db.InsertOne(ctx, product, "products")
	if err != nil {
		return "", err
	}
	newProduct := product
	newProduct.Id = result
	searchrepo.Create(ctx, result, newProduct, "products")
	createdProduct, _ := json.Marshal(newProduct)
	broker.Publish("Product.Created", createdProduct)
	return result, nil
}

func FindById(ctx context.Context, productId string) (model.Product, error) {
	var product model.Product
	err := db.FindById(ctx, productId, "products", &product)
	return product, err
}

func Find(ctx context.Context, filter interface{}, search string, limit int, skip int) ([]model.Product, error) {
	var result []model.Product
	if search != "" {
		searchrepo.Search("products", search, &result, limit, skip)
	} else {
		err := db.Find(ctx, "products", nil, skip, limit, &result)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

func UpdateById(ctx context.Context, id string, product model.Product) error {
	err := db.UpdateById(ctx, "products", id, product)
	if err != nil {
		return err
	}
	searchrepo.UpdateById("products", id, product)
	return nil
}

func DeleteById(ctx context.Context, id string) error {
	err := db.DeleteById(ctx, "products", id)
	if err != nil {
		return err
	}
	searchrepo.DeleteById("products", id)
	return nil
}
