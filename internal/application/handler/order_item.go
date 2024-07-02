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
	missingOrderItemID = "Missing OrderItem ID"
	invalidOrderItemID = "Invalid OrderItem ID"
)

type orderItemHandler struct {
	orderItemService service.OrderItemServiceInterface
}

func NewOrderItemHandler(orderItemService service.OrderItemServiceInterface) *orderItemHandler {
	return &orderItemHandler{
		orderItemService: orderItemService,
	}
}

func (h *orderItemHandler) RegisterRoutes(router chi.Router) {
	const basePath = "/order-items"
	router.Get(basePath, h.listOrderItems)
	router.Get(basePath+"/{id}", h.getOrderItem)
	router.Post(basePath, h.createOrderItem)
	router.Put(basePath+"/{id}", h.updateOrderItem)
	router.Delete(basePath+"/{id}", h.deleteOrderItem)
}

// createOrderItem godoc
// @Summary Create a new OrderItem
// @Description Create a new OrderItem
// @Tags OrderItems
// @Accept json
// @Produce json
// @Param input body dto.CreateOrderItemInput true "OrderItem data"
// @Success 201 {object} entity.OrderItem
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /order-items [post]
// @Security ApiKeyAuth
func (h *orderItemHandler) createOrderItem(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateOrderItemInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid input", err)
		return
	}

	orderItem, err := h.orderItemService.Create(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error creating OrderItem", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusCreated, orderItem)
}

// listOrderItems godoc
// @Summary List OrderItems
// @Description List OrderItems
// @Tags OrderItems
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Accept json
// @Produce json
// @Success 200 {array} entity.OrderItem
// @Failure 500 {object} errors.HTTPError
// @Router /order-items [get]
// @Security ApiKeyAuth
func (h *orderItemHandler) listOrderItems(w http.ResponseWriter, r *http.Request) {
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
	orderItems, err := h.orderItemService.List(r.Context(), pageInt, limitInt)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error listing OrderItems", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, orderItems)
}

// getOrderItem godoc
// @Summary Get an OrderItem
// @Description Get an OrderItem
// @Tags OrderItems
// @Produce json
// @Param id path string true "OrderItem ID"
// @Success 200 {object} entity.OrderItem
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /order-items/{id} [get]
// @Security ApiKeyAuth
func (h *orderItemHandler) getOrderItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingOrderItemID, nil)
		return
	}

	orderItemID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidOrderItemID, err)
		return
	}

	orderItem, err := h.orderItemService.Get(r.Context(), orderItemID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error getting OrderItem", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, orderItem)
}

// updateOrderItem godoc
// @Summary Update an OrderItem
// @Description Update an OrderItem
// @Tags OrderItems
// @Accept json
// @Produce json
// @Param id path string true "OrderItem ID"
// @Param input body dto.UpdateOrderItemInput true "OrderItem data"
// @Success 200 {object} entity.OrderItem
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Router /order-items/{id} [put]
// @Security ApiKeyAuth
func (h *orderItemHandler) updateOrderItem(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateOrderItemInput
	id := chi.URLParam(r, "id")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid input", err)
		return
	}

	orderItemID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidOrderItemID, err)
		return
	}

	orderItem, err := h.orderItemService.Update(r.Context(), orderItemID, &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error updating OrderItem", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, orderItem)
}

// deleteOrderItem godoc
// @Summary Delete an OrderItem
// @Description Delete an OrderItem
// @Tags OrderItems
// @Produce json
// @Param id path string true "OrderItem ID"
// @Success 204 "No content"
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /order-items/{id} [delete]
// @Security ApiKeyAuth
func (h *orderItemHandler) deleteOrderItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingOrderItemID, nil)
		return
	}

	orderItemID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidOrderItemID, err)
		return
	}

	err = h.orderItemService.Delete(r.Context(), orderItemID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error deleting OrderItem", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
