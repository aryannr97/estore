package controllers

// IControllers encapsulates all controller actions
type IControllers interface {
	IProductController
}

// CatalogControllers encapsulates all the controllers
type CatalogControllers struct {
	ProductController IProductController
}
