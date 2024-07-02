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
	missingFeedbackID = "Missing Feedback ID"
	invalidFeedbackID = "Invalid Feedback ID"
)

type FeedbackHandler struct {
	FeedbackService service.FeedbackServiceInterface
}

func NewFeedbackHandler(FeedbackService service.FeedbackServiceInterface) *FeedbackHandler {
	return &FeedbackHandler{
		FeedbackService: FeedbackService,
	}
}

func (h *FeedbackHandler) RegisterRoutes(router chi.Router) {
	const basePath = "/Feedbacks"
	router.Get(basePath, h.listFeedbacks)
	router.Get(basePath+"/{id}", h.getFeedback)
	router.Post(basePath, h.createFeedback)
	router.Put(basePath+"/{id}", h.updateFeedback)
	router.Delete(basePath+"/{id}", h.deleteFeedback)
}

// createFeedback godoc
// @Summary Create a new Feedback
// @Description Create a new Feedback
// @Tags Feedbacks
// @Accept json
// @Produce json
// @Param input body dto.CreateFeedbackInput true "Feedback data"
// @Success 201 {object} entity.Feedback
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /Feedbacks [post]
// @Security ApiKeyAuth
func (h *FeedbackHandler) createFeedback(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateFeedbackInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error reading request body", err)
		return
	}

	if err := input.FromJSON(body); err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	Feedback, err := h.FeedbackService.Create(r.Context(), &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error creating Feedback", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusCreated, Feedback)
}

// listFeedbacks godoc
// @Summary List Feedbacks
// @Description List Feedbacks
// @Tags Feedbacks
// @Produce json
// @Success 200 {array} entity.Feedback
// @Failure 500 {object} errors.HTTPError
// @Router /Feedbacks [get]
// @Security ApiKeyAuth
func (h *FeedbackHandler) listFeedbacks(w http.ResponseWriter, r *http.Request) {
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
	Feedbacks, err := h.FeedbackService.List(r.Context(), pageInt, limitInt)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error listing Feedbacks", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, Feedbacks)
}

// getFeedback godoc
// @Summary Get a Feedback
// @Description Get a Feedback
// @Tags Feedbacks
// @Produce json
// @Param id path string true "Feedback ID"
// @Success 200 {object} entity.Feedback
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /Feedbacks/{id} [get]
// @Security ApiKeyAuth
func (h *FeedbackHandler) getFeedback(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingFeedbackID, nil)
		return
	}

	FeedbackID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidFeedbackID, err)
		return
	}

	Feedback, err := h.FeedbackService.Get(r.Context(), FeedbackID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusNotFound, "Feedback not found", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, Feedback)
}

// updateFeedback godoc
// @Summary Update a Feedback
// @Description Update a Feedback
// @Tags Feedbacks
// @Accept json
// @Produce json
// @Param id path string true "Feedback ID"
// @Param input body dto.UpdateFeedbackInput true "Feedback data"
// @Success 200 {object} entity.Feedback
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /Feedbacks/{id} [put]
// @Security ApiKeyAuth
func (h *FeedbackHandler) updateFeedback(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateFeedbackInput

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
		errors.NewHTTPError(w, http.StatusBadRequest, missingFeedbackID, nil)
		return
	}

	feedbackID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidFeedbackID, err)
		return
	}

	input.ID = feedbackID

	feedback, err := h.FeedbackService.Update(r.Context(), feedbackID, &input)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error updating Feedback", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusOK, feedback)
}

// deleteFeedback godoc
// @Summary Delete a Feedback
// @Description Delete a Feedback
// @Tags Feedbacks
// @Accept json
// @Produce json
// @Param id path string true "Feedback ID"
// @Success 204
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /Feedbacks/{id} [delete]
// @Security ApiKeyAuth
func (h *FeedbackHandler) deleteFeedback(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		errors.NewHTTPError(w, http.StatusBadRequest, missingFeedbackID, nil)
		return
	}

	FeedbackID, err := uuid.Parse(id)
	if err != nil {
		errors.NewHTTPError(w, http.StatusBadRequest, invalidFeedbackID, err)
		return
	}

	err = h.FeedbackService.Delete(r.Context(), FeedbackID)
	if err != nil {
		errors.NewHTTPError(w, http.StatusInternalServerError, "Error deleting Feedback", err)
		return
	}

	dto.RespondWithJSON(w, http.StatusNoContent, nil)
}
