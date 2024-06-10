package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/errors"
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
	config.GetLogger("handler").Info("Entrou no handler createProduct")
	var input service.CreateProductInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	config.GetLogger("handler").Infof("Dados recebidos: %+v", input)

	newProduct, err := h.productService.CreateProduct(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Failed to create product", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func (h *ProductHandler) Routes(router chi.Router) {
	config.GetLogger("handler").Info("Entrou no handler Routes")
	router.Post("/", h.createProduct)
}
