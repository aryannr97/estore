package router

import (
	"log"

	"github.com/aryannr97/order-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

// Init starts http server
func Init(ctr controllers.OrderControllers) {
	routerEngine := gin.Default()

	//orderGroup := routerEngine.Group("/services/order/api/1.0")
	internalGroup := routerEngine.Group("/services/order/api/1.0/internal")

	internalGroup.POST("/create", ctr.OrderController.CreateOrder)

	err := routerEngine.Run(":8082")
	if err != nil {
		log.Fatalf("Catalog service start up failed due to %v", err)
	}
}
