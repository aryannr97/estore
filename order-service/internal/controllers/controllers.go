package controllers

// IControllers encapsulates all controller actions
type IControllers interface {
	IOrderController
}

// OrderControllers encapsulates all the controllers
type OrderControllers struct {
	OrderController IOrderController
}
