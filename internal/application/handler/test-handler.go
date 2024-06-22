package handler

import (
	"fmt"
	"net/http"

	"github.com/becardine/gestock-api/config"
	"github.com/go-chi/chi/v5"
)

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Routes(router chi.Router) {
	router.Get("/", h.handleTest)
}

func (h *TestHandler) handleTest(w http.ResponseWriter, r *http.Request) {
	config.GetLogger("handler").Info("Entrou no handler handleTest")
	fmt.Fprintln(w, "Teste OK!")
}
