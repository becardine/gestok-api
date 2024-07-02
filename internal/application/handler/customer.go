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
	missingCustomerID = "Missing Customer ID"
	invalidCustomerID = "Invalid Customer ID"
)

type CustomerHandler struct {
	customerService service.CustomerServiceInterface
}

func NewCustomerHandler(customerService service.CustomerServiceInterface) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (h *CustomerHandler) RegisterRoutes(router chi.Router) {
	const basePath = "/customers"
	router.Get(basePath, h.listCustomers)
	router.Get(basePath+"/{id}", h.getCustomer)
	router.Post(basePath, h.createCustomer)
	router.Put(basePath+"/{id}", h.updateCustomer)
	router.Delete(basePath+"/{id}", h.deleteCustomer)
}

// createCustomer godoc
// @Summary Create a new customer
// @Description Create a new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param input body dto.CreateCustomerInput true "Customer data"
// @Success 201 {object} entity.Customer
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /customers [post]
// @Security ApiKeyAuth
func (h *CustomerHandler) createCustomer(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateCustomerInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	customer, err := h.customerService.Create(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error creating customer", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusCreated, customer)
}

// listCustomers godoc
// @Summary List customers
// @Description List customers
// @Tags customers
// @Produce json
// @Success 200 {array} entity.Customer
// @Failure 500 {object} errors.HTTPError
// @Router /customers [get]
// @Security ApiKeyAuth
func (h *CustomerHandler) listCustomers(w http.ResponseWriter, r *http.Request) {
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
	customers, err := h.customerService.List(r.Context(), pageInt, limitInt)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error listing customers", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, customers)
}

// getCustomer godoc
// @Summary Get a customer
// @Description Get a customer
// @Tags customers
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} entity.Customer
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /customers/{id} [get]
// @Security ApiKeyAuth
func (h *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingCustomerID, nil)
		return
	}

	customerID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidCustomerID, err)
		return
	}

	customer, err := h.customerService.Get(r.Context(), customerID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusNotFound, "Customer not found", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, customer)
}

// updateCustomer godoc
// @Summary Update a customer
// @Description Update a customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Param input body dto.UpdateCustomerInput true "Customer data"
// @Success 200 {object} entity.Customer
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /customers/{id} [put]
// @Security ApiKeyAuth
func (h *CustomerHandler) updateCustomer(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateCustomerInput

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
		errors.NewHTTPError(w, http.StatusBadRequest, missingCustomerID, nil)
		return
	}

	customerID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidCustomerID, err)
		return
	}

	input.ID = customerID

	err = h.customerService.Update(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error updating customer", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, input)
}

// deleteCustomer godoc
// @Summary Delete a customer
// @Description Delete a customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Success 204
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /customers/{id} [delete]
// @Security ApiKeyAuth
func (h *CustomerHandler) deleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingCustomerID, nil)
		return
	}

	customerID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidCustomerID, err)
		return
	}

	err = h.customerService.Delete(r.Context(), customerID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error deleting customer", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusNoContent, nil)
}
