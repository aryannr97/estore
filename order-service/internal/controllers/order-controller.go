package controllers

import (
	"log"
	"net/http"

	"github.com/aryannr97/order-service/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// IOrderController defines all product specific controller layer actions
type IOrderController interface {
	// GetProducts(*gin.Context)
	// GetProduct(*gin.Context)
	// ReserveProducts(*gin.Context)
	CreateOrder(c *gin.Context)
}

// OrderController handles all user related requests
type OrderController struct {
	DB *sqlx.DB
}

// GetProducts fetch list of products list available
// func (p *OrderController) GetProducts(c *gin.Context) {
// 	products := []models.ProductDto{}
// 	fmt.Println(p.DB.Select(&products, "SELECT * FROM products"))
// 	c.JSON(200, products)
// }

// CreateOrder creates order for user
func (p *OrderController) CreateOrder(c *gin.Context) {
	var requestPayload models.CreateOrderRequest

	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")

		return
	}

	_, err := p.DB.NamedExec("INSERT INTO orders(userid,checkoutid,txid,detail) VALUES(:userid,:checkoutid,:txid,:detail)", &requestPayload)
	if err != nil {
		log.Printf("%+v", requestPayload)
		c.JSON(500, err)

		return
	}

	c.JSON(201, nil)
}
