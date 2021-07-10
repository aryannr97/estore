package main

import (
	"log"

	"github.com/aryannr97/order-service/internal/controllers"
	datastore "github.com/aryannr97/order-service/internal/datasource/mysql"

	"github.com/aryannr97/order-service/internal/router"
)

func main() {
	db := datastore.Init()
	log.Print(db)
	ctr := controllers.OrderControllers{
		OrderController: &controllers.OrderController{
			DB: db,
		},
	}
	router.Init(ctr)
}
