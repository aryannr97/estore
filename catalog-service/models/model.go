package models

// ProductDto ...
type ProductDto struct {
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

// ReserveDto ...
type ReserveDto struct {
	ID         int64  `db:"id"`
	CheckoutID string `db:"checkoutid"`
	ProductID  int64  `db:"productid"`
	Qty        int64  `db:"quantity"`
}
