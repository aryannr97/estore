package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aryannr97/cart-service/models"
	"github.com/aryannr97/cart-service/pkg/externalcaller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ICartController provides cart operations
type ICartController interface {
	AddtoCart(*gin.Context)
	GetCart(*gin.Context)
	Checkout(c *gin.Context)
}

// CartController handles cart operations
type CartController struct {
	DB             *sqlx.DB
	ExternalCaller externalcaller.IExternalCaller
}

var (
	cartChan     = make(chan int, 10000)
	checkoutChan = make(chan int, 1)
)

// AddtoCart adds products to user cart
func (p *CartController) AddtoCart(c *gin.Context) {
	defer func() {
		<-cartChan
	}()
	cartChan <- 1
	var requestPayload models.AddtoCartRequest
	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")

		return
	}

	var cartItems int64

	productInfo, err := p.ExternalCaller.GetProduct(fmt.Sprint(requestPayload.ProductID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	if productInfo.Qty < requestPayload.Qty {
		c.JSON(http.StatusForbidden, "Requested quantity not in stock")

		return
	}

	p.DB.Get(&cartItems, "SELECT COUNT(*) FROM cart WHERE userid = ?  AND productid = ?",
		requestPayload.UserID, requestPayload.ProductID)

	if cartItems > 0 {
		_, err := p.DB.NamedExec("UPDATE cart SET quantity = :quantity, price = :price WHERE userid = :userid  AND productid = :productid", &requestPayload)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)

			return
		}

		c.JSON(204, nil)

		return
	}
	_, err = p.DB.NamedExec("INSERT INTO cart(userid,productid,quantity,price) VALUES(:userid,:productid,:quantity,:price)", &requestPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(201, nil)
}

// GetCart fetch list of products in cart
func (p *CartController) GetCart(c *gin.Context) {
	userid := c.Param("userid")
	products := []models.ProductDto{}
	fmt.Println(p.DB.Select(&products, "SELECT * FROM cart WHERE userid = ?", userid))
	c.JSON(200, products)
}

// Checkout starts product order placement
func (p *CartController) Checkout(c *gin.Context) {
	var requestPayload models.CheckoutRequest
	var reserveRequest models.ReserveProducts
	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")

		return
	}

	reserveRequest.CheckoutID = uuid.New().String()
	for _, item := range requestPayload.Products {
		reserveRequest.Products = append(reserveRequest.Products, models.Product{
			ProductID: item.ProductID,
			Qty:       item.Qty,
		})
	}

	checkoutChan <- 1

	if err := p.ExternalCaller.ReserveProducts(reserveRequest); err != nil {
		c.JSON(500, err)

		<-checkoutChan
		return
	}

	<-checkoutChan

	/**Perform payment operations*/
	txID := uuid.New().String()
	time.Sleep(10 * time.Second)
	/**/

	pBytes, _ := json.Marshal(requestPayload.Products)

	req := models.CreateOrderRequest{
		UserID:     requestPayload.Products[0].UserID,
		CheckoutID: reserveRequest.CheckoutID,
		TxID:       txID,
		Detail:     string(pBytes),
	}

	if err := p.ExternalCaller.CreateOrder(req); err != nil {
		c.JSON(500, err)

		return
	}

	if _, err := p.DB.Exec("DELETE FROM cart WHERE userid = ?", requestPayload.Products[0].UserID); err != nil {
		c.JSON(500, err)

		return
	}

	if err := p.ExternalCaller.UpdateInventory(reserveRequest); err != nil {
		c.JSON(500, err)

		return
	}

	c.JSON(200, "Order placed")
}
