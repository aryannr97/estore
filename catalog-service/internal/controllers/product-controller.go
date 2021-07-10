package controllers

import (
	"fmt"
	"net/http"

	"github.com/aryannr97/catalog-service/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// IProductController defines all product specific controller layer actions
type IProductController interface {
	GetProducts(*gin.Context)
	GetProduct(*gin.Context)
	ReserveProducts(*gin.Context)
	UpdateInventory(c *gin.Context)
}

// ProductController handles all user related requests
type ProductController struct {
	DB *sqlx.DB
}

// GetProducts fetch list of products list available
func (p *ProductController) GetProducts(c *gin.Context) {
	products := []models.ProductDto{}
	fmt.Println(p.DB.Select(&products, "SELECT * FROM products"))
	c.JSON(200, products)
}

// GetProduct fetch prodduct details
func (p *ProductController) GetProduct(c *gin.Context) {
	products := models.ProductDto{}
	productid := c.Param("productid")
	fmt.Println(p.DB.Get(&products, "SELECT * FROM products WHERE id = ?", productid))
	c.JSON(200, products)
}

// ReserveProducts marls product for checkout
func (p *ProductController) ReserveProducts(c *gin.Context) {
	var requestPayload models.ReserveProducts

	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")

		return
	}

	tx, _ := p.DB.Beginx()

	for _, item := range requestPayload.Products {
		var qty int
		var reserve int
		itemDto := models.ReserveDto{
			CheckoutID: requestPayload.CheckoutID,
			ProductID:  item.ProductID,
			Qty:        item.Qty,
		}
		p.DB.Get(&qty, "SELECT quantity FROM products WHERE id = ?", item.ProductID)
		p.DB.Get(&reserve, "SELECT SUM(quantity) FROM reserved_products WHERE productid = ?", item.ProductID)

		if item.Qty > int64(qty-reserve) {
			c.JSON(500, fmt.Sprintf("Quantity exceeded for productid:%v", item.ProductID))

			return
		}
		_, err := tx.NamedExec("INSERT INTO reserved_products(checkoutid,productid,quantity) VALUES(:checkoutid,:productid,:quantity)", &itemDto)
		if err != nil {
			tx.Rollback()
			c.JSON(500, err)

			return
		}

	}

	tx.Commit()
	c.JSON(204, nil)
}

// UpdateInventory marls product for checkout
func (p *ProductController) UpdateInventory(c *gin.Context) {
	var requestPayload models.ReserveProducts

	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")

		return
	}

	tx, _ := p.DB.Beginx()

	for _, item := range requestPayload.Products {
		var qty int
		itemDto := models.ReserveDto{
			CheckoutID: requestPayload.CheckoutID,
			ProductID:  item.ProductID,
			Qty:        item.Qty,
		}

		p.DB.Get(&qty, "SELECT SUM(quantity) FROM reserved_products WHERE productid = ? AND checkoutid = ?", itemDto.ProductID, itemDto.CheckoutID)

		_, err := tx.NamedExec("UPDATE products SET quantity = quantity - :quantity", &itemDto)
		if err != nil {
			tx.Rollback()
			c.JSON(500, err)

			return
		}

	}

	_, err := tx.Exec("DELETE FROM reserved_products WHERE checkoutid = ?", requestPayload.CheckoutID)
	if err != nil {
		tx.Rollback()
		c.JSON(500, err)

		return
	}

	tx.Commit()
	c.JSON(204, nil)
}
