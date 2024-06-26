package handler

import (
	"io"
	"net/http"
	"strconv"

	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const missingCategoryID = "Missing category ID"
const invalidCategoryID = "Invalid category ID"

type CategoryHandler struct {
	categoryService service.CategoryServiceInterface
}

func NewCategoryHandler(categoryService service.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) RegisterRoutes(router chi.Router) {
	const basePath = "/categories"
	router.Get(basePath, h.listCategories)
	router.Get(basePath+"/{id}", h.getCategory)
	router.Post(basePath, h.createCategory)
	router.Put(basePath+"/{id}", h.updateCategory)
	router.Delete(basePath+"/{id}", h.deleteCategory)
	router.Get(basePath+"/{id}/products", h.getProducts)
}

// createCategory godoc
// @Summary Create category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body entity.Category true "Category object"
// @Success 201 {object} entity.Category
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /categories [post]
func (h *CategoryHandler) createCategory(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateCategoryInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Error parsing request body", err)
		return
	}

	category, err := h.categoryService.Create(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error creating category", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusCreated, category)
}

// updateCategory godoc
// @Summary Update category
// @Description Update a category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param category body entity.Category true "Category object"
// @Success 204 No Content
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /categories/{id} [put]
func (h *CategoryHandler) updateCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingCategoryID, nil)
		return
	}

	categoryId, err := uuid.Parse(id)
	if err != nil {

		errors.NewHTTPError(w, http.StatusBadRequest, invalidCategoryID, err)
		return
	}

	var input dto.UpdateCategoryInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Error parsing request body", err)
		return
	}

	input.ID = categoryId

	if err := h.categoryService.Update(r.Context(), &input); err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error updating category", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusNoContent, nil)
}

// deleteCategory godoc
// @Summary Delete category
// @Description Delete a category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 204 No Content
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /categories/{id} [delete]
func (h *CategoryHandler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingCategoryID, nil)
		return
	}

	categoryId, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidCategoryID, err)
		return
	}

	if err := h.categoryService.Delete(r.Context(), categoryId); err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error deleting category", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusNoContent, nil)
}

// getProducts godoc
// @Summary Get category products
// @Description Get all products for a category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {array} entity.Product
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /categories/{id}/products [get]
func (h *CategoryHandler) getProducts(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingCategoryID, nil)
		return
	}

	categoryId, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidCategoryID, err)
		return
	}

	products, err := h.categoryService.GetCategoryProducts(r.Context(), categoryId)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error fetching category products", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, products)
}

// getCategory godoc
// @Summary Get category
// @Description Get a category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} entity.Category
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /categories/{id} [get]
func (h *CategoryHandler) getCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingCategoryID, nil)
		return
	}

	categoryId, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidCategoryID, err)
		return
	}

	category, err := h.categoryService.Get(r.Context(), categoryId)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error fetching category", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, category)
}

// listCategories godoc
// @Summary List categories
// @Description List all categories
// @Tags categories
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Limit per page"
// @Success 200 {array} entity.Category
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /categories [get]
func (h *CategoryHandler) listCategories(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	categories, err := h.categoryService.List(r.Context(), pageInt, limitInt)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error fetching categories", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, categories)
}
