package handler

import (
	"encoding/json"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/becardine/gestock-api/internal/utils"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

type BrandHandler struct {
	brandService service.BrandService
}

func NewBrandHandler(brandService service.BrandService) *BrandHandler {
	return &BrandHandler{
		brandService: brandService,
	}
}

func (h *BrandHandler) RegisterRoutes(router chi.Router) {
	router.Get("/brands", h.listBrands)
	router.Get("/brands/{id}", h.getBrand)
	router.Post("/brands", h.createBrand)
	router.Put("/brands/{id}", h.updateBrand)
	router.Delete("/brands/{id}", h.deleteBrand)
}

// createBrand godoc
// @Summary Create a new brand
// @Description Create a new brand
// @Tags brands
// @Accept json
// @Produce json
// @Param input body dto.CreateBrandInput true "Brand data"
// @Success 201 {object} entity.Brand
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /brands [post]
// @Security ApiKeyAuth
func (h *BrandHandler) createBrand(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateBrandInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Error parsing request body", err)
		return
	}

	brand, err := h.brandService.Create(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error creating brand", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, brand)
}

// updateBrand godoc
// @Summary Update a brand
// @Description Update a brand
// @Tags brands
// @Accept json
// @Produce json
// @Param id path string true "Brand ID"
// @Param input body dto.UpdateBrandInput true "Brand data"
// @Success 204 "No Content"
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /brands/{id} [put]
// @Security ApiKeyAuth
func (h *BrandHandler) updateBrand(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateBrandInput

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

	commonID, err := common.NewIDFromString(id) // convert string to common.ID
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid brand ID", err)
		return
	}

	input.ID = commonID

	err = h.brandService.Update(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error updating brand", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}

// deleteBrand godoc
// @Summary Delete a brand
// @Description Delete a brand
// @Tags brands
// @Param id path string true "Brand ID"
// @Success 204 "No Content"
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /brands/{id} [delete]
// @Security ApiKeyAuth
func (h *BrandHandler) deleteBrand(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, "Missing brand ID", nil)
		return
	}

	commonID, err := common.NewIDFromString(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid brand ID", err)
		return
	}

	err = h.brandService.Delete(r.Context(), commonID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error deleting brand", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}

// getBrand godoc
// @Summary Get a brand
// @Description Get a brand
// @Tags brands
// @Param id path string true "Brand ID" Format(uuid)
// @Produce json
// @Success 200 {object} entity.Brand
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /brands/{id} [get]
// @Security ApiKeyAuth
func (h *BrandHandler) getBrand(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, "Missing brand ID", nil)
		return
	}

	commonID, err := common.NewIDFromString(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid brand ID", err)
		return
	}

	brand, err := h.brandService.Get(r.Context(), commonID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error fetching brand", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, brand)
}

// listBrands godoc
// @Summary List brands
// @Description List brands
// @Tags brands
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Produce json
// @Success 200 {array} entity.Brand
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /brands [get]
// @Security ApiKeyAuth
func (h *BrandHandler) listBrands(w http.ResponseWriter, r *http.Request) {
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

	brands, err := h.brandService.List(r.Context(), pageInt, limitInt)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error fetching brands", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, brands)
}
