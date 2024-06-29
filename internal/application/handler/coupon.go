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

const (
	missingCouponID = "Missing coupon ID"
	invalidCouponID = "Invalid coupon ID"
)

type CouponHandler struct {
	couponService service.CouponServiceInterface
}

func NewCouponHandler(couponService service.CouponServiceInterface) *CouponHandler {
	return &CouponHandler{
		couponService: couponService,
	}
}

func (h *CouponHandler) RegisterRoutes(router chi.Router) {
	const basePath = "/coupons"
	router.Get(basePath, h.listCoupons)
	router.Get(basePath+"/{id}", h.getCoupon)
	router.Post(basePath, h.createCoupon)
	router.Put(basePath+"/{id}", h.updateCoupon)
	router.Delete(basePath+"/{id}", h.deleteCoupon)
}

// createCoupon godoc
// @Summary Create a new coupon
// @Description Create a new coupon
// @Tags coupons
// @Accept json
// @Produce json
// @Param input body dto.CreateCouponInput true "Coupon data"
// @Success 201 {object} entity.Coupon
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /coupons [post]
// @Security ApiKeyAuth
func (h *CouponHandler) createCoupon(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateCouponInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	coupon, err := h.couponService.Create(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error creating coupon", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusCreated, coupon)
}

// listCoupons godoc
// @Summary List coupons
// @Description List coupons
// @Tags coupons
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} []entity.Coupon
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /coupons [get]
func (h *CouponHandler) listCoupons(w http.ResponseWriter, r *http.Request) {
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

	coupons, err := h.couponService.List(r.Context(), pageInt, limitInt)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error listing coupons", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, coupons)
}

// getCoupon godoc
// @Summary Get a coupon by ID
// @Description Get a coupon by ID
// @Tags coupons
// @Accept json
// @Produce json
// @Param id path string true "Coupon ID"
// @Success 200 {object} entity.Coupon
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /coupons/{id} [get]
func (h *CouponHandler) getCoupon(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingCouponID, nil)
		return
	}
	couponID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidCouponID, err)
		return
	}

	coupon, err := h.couponService.Get(r.Context(), couponID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error getting coupon", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, coupon)
}

// updateCoupon godoc
// @Summary Update a coupon by ID
// @Description Update a coupon by ID
// @Tags coupons
// @Accept json
// @Produce json
// @Param id path string true "Coupon ID"
// @Param input body dto.UpdateCouponInput true "Coupon data"
// @Success 200 {object} entity.Coupon
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /coupons/{id} [put]
// @Security ApiKeyAuth
func (h *CouponHandler) updateCoupon(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateCouponInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingCouponID, nil)
		return
	}
	_, err = uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidCouponID, err)
		return
	}

	err = h.couponService.Update(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error updating coupon", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, input)
}

// deleteCoupon godoc
// @Summary Delete a coupon by ID
// @Description Delete a coupon by ID
// @Tags coupons
// @Accept json
// @Produce json
// @Param id path string true "Coupon ID"
// @Success 204 "No Content"
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /coupons/{id} [delete]
// @Security ApiKeyAuth
func (h *CouponHandler) deleteCoupon(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingCouponID, nil)
		return
	}
	couponID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidCouponID, err)
		return
	}

	err = h.couponService.Delete(r.Context(), couponID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error deleting coupon", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
