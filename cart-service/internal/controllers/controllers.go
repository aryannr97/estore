package controllers

// IControllers encapsulates all controller actions
type IControllers interface {
	ICartController
}

// CatalogControllers encapsulates all the controllers
type CatalogControllers struct {
	CartController ICartController
}
