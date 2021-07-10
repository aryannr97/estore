package externalcaller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/aryannr97/cart-service/models"
)

// IExternalCaller ...
type IExternalCaller interface {
	GetProduct(string) (models.ProductResponse, error)
	ReserveProducts(models.ReserveProducts) error
	UpdateInventory(models.ReserveProducts) error
	CreateOrder(models.CreateOrderRequest) error
}

// ExternalCaller ...
type ExternalCaller struct {
	CartServiceURL  string
	OrderServiceURL string
}

// GetProduct ...
func (ex *ExternalCaller) GetProduct(productid string) (models.ProductResponse, error) {
	var responseBody models.ProductResponse
	url := ex.CartServiceURL + "/internal/product/" + productid

	req, _ := http.NewRequest("GET", url, nil)

	c := http.Client{}

	res, err := c.Do(req)

	if err != nil {
		return models.ProductResponse{}, err
	}

	resBody, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(resBody, &responseBody)

	return responseBody, err
}

// ReserveProducts ...
func (ex *ExternalCaller) ReserveProducts(request models.ReserveProducts) error {
	url := ex.CartServiceURL + "/internal/product/reserve"

	reqBytes, _ := json.Marshal(request)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))

	c := http.Client{}

	res, err := c.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		return err
	}

	return nil
}

// UpdateInventory ...
func (ex *ExternalCaller) UpdateInventory(request models.ReserveProducts) error {
	url := ex.CartServiceURL + "/internal/updateinventory"

	reqBytes, _ := json.Marshal(request)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(reqBytes))

	c := http.Client{}

	res, err := c.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		return err
	}

	return nil
}

// CreateOrder ...
func (ex *ExternalCaller) CreateOrder(request models.CreateOrderRequest) error {
	url := ex.OrderServiceURL + "/internal/create"

	reqBytes, _ := json.Marshal(request)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))

	c := http.Client{}

	res, err := c.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		return errors.New("Error placing order")
	}

	return nil
}
