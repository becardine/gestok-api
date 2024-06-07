package handler

import (
	"net/http"

	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {
}

func (h *ProductHandler) Routes() chi.Router {
	r := chi.NewRouter()

	// Defina as rotas aqui (ex: r.Get("/", h.listProducts))

	return r
}
