package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/dto"
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
	const basePath = "/products"
	router.Get(basePath, h.listProducts)
	router.Get(basePath+"/{id}", h.getProduct)
	router.Post(basePath, h.createProduct)
	router.Put(basePath+"/{id}", h.updateProduct)
	router.Delete(basePath+"/{id}", h.deleteProduct)
}

// createProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param input body dto.CreateProductInput true "Product data"
// @Success 201 {object} entity.Product
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	config.GetLogger("handler").Info("Entrou no handler createProduct")
	var input dto.CreateProductInput

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
// @Param input body dto.UpdateProductInput true "Product data"
// @Success 200 {object} entity.Product
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) updateProduct(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateProductInput

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

	err := h.productService.UpdateProduct(r.Context(), input.ID, &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Failed to update product", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, input)
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

	productID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid product ID", err)
		return
	}

	err = h.productService.DeleteProduct(r.Context(), productID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Failed to delete product", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusNoContent, nil)
}

// listProducts godoc
// @Summary List all products
// @Description List all products
// @Tags products
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Accept json
// @Produce json
// @Success 200 {array} entity.Product
// @Failure 500 {object} errors.HTTPError
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) listProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0 // default to 0 if not provided or invalid
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10 // default to 10 if not provided or invalid
	}

	products, err := h.productService.ListProducts(r.Context(), pageInt, limitInt)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Failed to list products", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, products)
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

	productID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid brand ID", err)
		return
	}

	product, err := h.productService.GetProduct(r.Context(), productID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Failed to get product", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, product)
}
