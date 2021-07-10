package models

// ProductDto ...
type ProductDto struct {
	ID        int64 `json:"id" db:"id"`
	UserID    int64 `json:"userid" db:"userid"`
	ProductID int64 `json:"productid" db:"productid"`
	Qty       int64 `json:"quantity" db:"quantity"`
	Price     int64 `json:"price" db:"price"`
}

// AddtoCartRequest ...
type AddtoCartRequest struct {
	UserID    int64 `json:"userid" db:"userid"`
	ProductID int64 `json:"productid" db:"productid"`
	Qty       int64 `json:"quantity" db:"quantity"`
	Price     int64 `json:"price" db:"price"`
}

// ProductResponse ...
type ProductResponse struct {
	ID     int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Seller string `json:"seller" db:"seller"`
	Qty    int64  `json:"quantity" db:"quantity"`
	Price  int64  `json:"price" db:"price"`
}

// ReserveProducts ...
type ReserveProducts struct {
	CheckoutID string    `json:"checkoutid"`
	Products   []Product `json:"products"`
}

// Product ...
type Product struct {
	ProductID int64 `json:"productid"`
	Qty       int64 `json:"quantity"`
}

// CheckoutRequest ...
type CheckoutRequest struct {
	Products []ProductDto `json:"products"`
}

// CreateOrderRequest ...
type CreateOrderRequest struct {
	UserID     int64  `json:"userid"`
	CheckoutID string `json:"checkoutid"`
	TxID       string `json:"txid"`
	Detail     string `json:"detail"`
}
