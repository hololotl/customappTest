package internal

import (
	"CustomappTest/internal/models"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler struct {
	generator Geneator
	logger    *slog.Logger
}

func NewHandler(rtp float64, log *slog.Logger) *Handler {
	g := NewParetoAlgorithm(rtp)
	return &Handler{g, log}
}

func (h *Handler) GetRandom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.logger.Info("GetRandom request method not allowed ", "Request Method", r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "Method not allowed"})
		return
	}
	resp := h.generator.GetMultiplier()
	res := models.Response{resp}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
