package main

import (
	"github.com/aryannr97/cart-service/internal/controllers"
	datastore "github.com/aryannr97/cart-service/internal/datasource/mysql"
	"github.com/aryannr97/cart-service/pkg/externalcaller"

	"github.com/aryannr97/cart-service/internal/router"
)

func main() {
	db := datastore.Init()
	ex := &externalcaller.ExternalCaller{
		CartServiceURL:  "http://catalog-service:8080/services/catalog/api/1.0",
		OrderServiceURL: "http://order-service:8082/services/order/api/1.0",
	}
	ctr := controllers.CatalogControllers{
		CartController: &controllers.CartController{
			DB:             db,
			ExternalCaller: ex,
		},
	}
	router.Init(ctr)
}
