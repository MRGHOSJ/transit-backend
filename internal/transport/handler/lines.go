package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"transit-backend/internal/transport/service"

	"github.com/gorilla/mux"
)

type LinesHandler struct {
	service *service.LinesService
}

func NewLinesHandler(service *service.LinesService) *LinesHandler {
	return &LinesHandler{service: service}
}

func (h *LinesHandler) response(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *LinesHandler) errorResponse(w http.ResponseWriter, status int, message string) {
	h.response(w, status, map[string]string{"error": message})
}

func (h *LinesHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	lines := h.service.GetAllLines()
	h.response(w, http.StatusOK, lines)
}

func (h *LinesHandler) GetByType(w http.ResponseWriter, r *http.Request) {
	t := mux.Vars(r)["type"]

	lines, err := h.service.GetLinesByType(t)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	h.response(w, http.StatusOK, lines)
}

func (h *LinesHandler) GetLineByType(w http.ResponseWriter, r *http.Request) {
	t := mux.Vars(r)["type"]
	l := mux.Vars(r)["line"]
	f := mux.Vars(r)["flag"]

	line, err := h.service.GetLine(t, l)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if f == "" {
		h.response(w, http.StatusOK, line)
		return
	}

	switch strings.ToLower(f) {
	case "departure":
		h.response(w, http.StatusOK, line.Departure)
	case "arrival":
		h.response(w, http.StatusOK, line.Arrival)
	case "stations":
		h.response(w, http.StatusOK, line.Stations)
	case "schedule":
		h.response(w, http.StatusOK, line.Schedule)
	default:
		h.errorResponse(w, http.StatusBadRequest,
			"invalid flag (use: departure, arrival, stations, schedule)")
	}
}

func (h *LinesHandler) GetClosestStation(w http.ResponseWriter, r *http.Request) {

	latStr := r.URL.Query().Get("lat")
	lngStr := r.URL.Query().Get("lng")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, "invalid latitude")
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, "invalid longitude")
		return
	}

	line, station, distance, err := h.service.FindClosestStation(lat, lng)
	if err != nil {
		h.errorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	response := map[string]interface{}{
		"line":        line,
		"station":     station,
		"distance_km": distance,
	}

	h.response(w, http.StatusOK, response)
}
