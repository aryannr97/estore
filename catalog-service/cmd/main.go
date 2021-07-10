package main

import (
	"log"

	"github.com/aryannr97/catalog-service/internal/controllers"
	datastore "github.com/aryannr97/catalog-service/internal/datasource/mysql"

	"github.com/aryannr97/catalog-service/internal/router"
)

func main() {
	db := datastore.Init()
	log.Print(db)
	ctr := controllers.CatalogControllers{
		ProductController: &controllers.ProductController{
			DB: db,
		},
	}
	router.Init(ctr)
}
