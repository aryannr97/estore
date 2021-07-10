package router

import (
	"log"

	"github.com/aryannr97/catalog-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

// Init starts http server
func Init(ctr controllers.CatalogControllers) {
	routerEngine := gin.Default()

	productGroup := routerEngine.Group("/services/catalog/api/1.0/product")
	internalGroup := routerEngine.Group("/services/catalog/api/1.0/internal")

	productGroup.GET("/list", ctr.ProductController.GetProducts)
	productGroup.GET("/:productid", ctr.ProductController.GetProduct)
	internalGroup.GET("/product/:productid", ctr.ProductController.GetProduct)
	internalGroup.POST("/product/reserve", ctr.ProductController.ReserveProducts)
	internalGroup.PUT("/updateinventory", ctr.ProductController.UpdateInventory)

	err := routerEngine.Run(":8080")
	if err != nil {
		log.Fatalf("Catalog service start up failed due to %v", err)
	}
}
