package handler

import (
	"encoding/json"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"

	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/errors"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) RegisterRoutes(router chi.Router) {
	router.Get("/products", h.listProducts)
	router.Get("/products/{id}", h.getProduct)
	router.Post("/products", h.createProduct)
	router.Put("/products/{id}", h.updateProduct)
	router.Delete("/products/{id}", h.deleteProduct)
}

// createProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param input body service.CreateProductInput true "Product data"
// @Success 201 {object} entity.Product
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /products [post]
// @Security ApiKeyAuth
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

	config.GetLogger("handler").Infof("Input: %+v", input)

	newProduct, err := h.productService.CreateProduct(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Failed to create product", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newProduct)
	if err != nil {
		return
	}
}

// updateProduct godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param input body service.UpdateProductInput true "Product data"
// @Success 200 {object} entity.Product
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) updateProduct(w http.ResponseWriter, r *http.Request) {
	var input service.UpdateProductInput

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Error parsing request body", err)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, "Missing brand ID", nil)
		return
	}

	productID, err := common.NewIDFromString(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid product ID", err)
		return
	}

	input.ID = productID

	err = h.productService.UpdateProduct(r.Context(), productID, &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Failed to update product", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(input)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// deleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 "No Content"
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, "Missing product ID", nil)
		return
	}

	productID, err := common.NewIDFromString(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid product ID", err)
		return
	}

	err = h.productService.DeleteProduct(r.Context(), productID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Failed to delete product", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// listProducts godoc
// @Summary List all products
// @Description List all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} entity.Product
// @Failure 500 {object} errors.HTTPError
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) listProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.ListProducts(r.Context())
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Failed to list products", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		return
	}
}

// getProduct godoc
// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) getProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, "Missing product ID", nil)
		return
	}

	productID, err := common.NewIDFromString(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid brand ID", err)
		return
	}

	product, err := h.productService.GetProduct(r.Context(), productID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Failed to get product", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		return
	}
}
