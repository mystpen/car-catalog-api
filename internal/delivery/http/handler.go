package http

type Handler struct{
	service CarCatalogService
}

func NewHandler(service CarCatalogService) *Handler{
	return &Handler{service: service}
}