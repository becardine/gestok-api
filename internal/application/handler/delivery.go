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
	missingDeliveryID = "Missing Delivery ID"
	invalidDeliveryID = "Invalid Delivery ID"
)

type DeliveryHandler struct {
	DeliveryService service.DeliveryServiceInterface
}

func NewDeliveryHandler(DeliveryService service.DeliveryServiceInterface) *DeliveryHandler {
	return &DeliveryHandler{
		DeliveryService: DeliveryService,
	}
}

func (h *DeliveryHandler) RegisterRoutes(router chi.Router) {
	const basePath = "/deliverys"
	router.Get(basePath, h.listDeliverys)
	router.Get(basePath+"/{id}", h.getDelivery)
	router.Post(basePath, h.createDelivery)
	router.Put(basePath+"/{id}", h.updateDelivery)
	router.Delete(basePath+"/{id}", h.deleteDelivery)
}

// createDelivery godoc
// @Summary Create a new Delivery
// @Description Create a new Delivery
// @Tags Deliverys
// @Accept json
// @Produce json
// @Param input body dto.CreateDeliveryInput true "Delivery data"
// @Success 201 {object} entity.Delivery
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /Deliverys [post]
// @Security ApiKeyAuth
func (h *DeliveryHandler) createDelivery(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateDeliveryInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	Delivery, err := h.DeliveryService.Create(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error creating Delivery", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusCreated, Delivery)
}

// listDeliverys godoc
// @Summary List Deliverys
// @Description List Deliverys
// @Tags Deliverys
// @Produce json
// @Success 200 {array} entity.Delivery
// @Failure 500 {object} errors.HTTPError
// @Router /Deliverys [get]
// @Security ApiKeyAuth
func (h *DeliveryHandler) listDeliverys(w http.ResponseWriter, r *http.Request) {
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
	Deliverys, err := h.DeliveryService.List(r.Context(), pageInt, limitInt)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error listing Deliverys", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, Deliverys)
}

// getDelivery godoc
// @Summary Get a Delivery
// @Description Get a Delivery
// @Tags Deliverys
// @Produce json
// @Param id path string true "Delivery ID"
// @Success 200 {object} entity.Delivery
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /Deliverys/{id} [get]
// @Security ApiKeyAuth
func (h *DeliveryHandler) getDelivery(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingDeliveryID, nil)
		return
	}

	DeliveryID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidDeliveryID, err)
		return
	}

	Delivery, err := h.DeliveryService.Get(r.Context(), DeliveryID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusNotFound, "Delivery not found", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, Delivery)
}

// updateDelivery godoc
// @Summary Update a Delivery
// @Description Update a Delivery
// @Tags Deliverys
// @Accept json
// @Produce json
// @Param id path string true "Delivery ID"
// @Param input body dto.UpdateDeliveryInput true "Delivery data"
// @Success 200 {object} entity.Delivery
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /Deliverys/{id} [put]
// @Security ApiKeyAuth
func (h *DeliveryHandler) updateDelivery(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateDeliveryInput

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
		errors.NewHTTPError(w, http.StatusBadRequest, missingDeliveryID, nil)
		return
	}

	deliveryID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidDeliveryID, err)
		return
	}

	input.ID = deliveryID

	delivery, err := h.DeliveryService.Update(r.Context(), deliveryID, &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error updating Delivery", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, delivery)
}

// deleteDelivery godoc
// @Summary Delete a Delivery
// @Description Delete a Delivery
// @Tags Deliverys
// @Accept json
// @Produce json
// @Param id path string true "Delivery ID"
// @Success 204
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /Deliverys/{id} [delete]
// @Security ApiKeyAuth
func (h *DeliveryHandler) deleteDelivery(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingDeliveryID, nil)
		return
	}

	DeliveryID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidDeliveryID, err)
		return
	}

	err = h.DeliveryService.Delete(r.Context(), DeliveryID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error deleting Delivery", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusNoContent, nil)
}
