package router

import (
	"log"

	"github.com/aryannr97/cart-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

// Init starts http server
func Init(ctr controllers.CatalogControllers) {
	routerEngine := gin.Default()

	cartGroup := routerEngine.Group("/services/cart/api/1.0/")

	cartGroup.PUT("/add", ctr.CartController.AddtoCart)
	cartGroup.GET("/user/:userid", ctr.CartController.GetCart)
	cartGroup.POST("/checkout", ctr.CartController.Checkout)

	err := routerEngine.Run(":8081")
	if err != nil {
		log.Fatalf("Catalog service start up failed due to %v", err)
	}
}
