package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"transit-backend/internal/transport/model"
	"transit-backend/internal/transport/service"
)

type GraphHandler struct {
	graphService *service.Graph
}

func NewGraphHandler(graphService *service.Graph) *GraphHandler {
	return &GraphHandler{graphService: graphService}
}

func (h *GraphHandler) response(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *GraphHandler) errorResponse(w http.ResponseWriter, status int, message string) {
	h.response(w, status, map[string]string{"error": message})
}

func (h *GraphHandler) GetGraphStructure(w http.ResponseWriter, r *http.Request) {
	h.response(w, http.StatusOK, h.graphService.Stations)
}

// GetStationDetails returns details for a specific station
func (h *GraphHandler) GetStationDetails(w http.ResponseWriter, r *http.Request) {

	latStr := r.URL.Query().Get("lat")
	lngStr := r.URL.Query().Get("lng")

	// Convert string coordinates to float64
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, "invalid latitude value")
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, "invalid longitude value")
		return
	}

	stationID := fmt.Sprintf("%.6f,%.6f", lat, lng)

	if stationID == "" {
		h.errorResponse(w, http.StatusBadRequest, "station ID parameter is required")
		return
	}

	station, exists := h.graphService.Stations[stationID]
	if !exists {
		h.errorResponse(w, http.StatusNotFound, "station not found")
		return
	}

	h.response(w, http.StatusOK, station)
}

// GetStationConnections returns all connections for a station
func (h *GraphHandler) GetStationConnections(w http.ResponseWriter, r *http.Request) {

	latStr := r.URL.Query().Get("lat")
	lngStr := r.URL.Query().Get("lng")

	// Convert string coordinates to float64
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, "invalid latitude value")
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, "invalid longitude value")
		return
	}

	stationID := fmt.Sprintf("%.6f,%.6f", lat, lng)

	if stationID == "" {
		h.errorResponse(w, http.StatusBadRequest, "station ID parameter is required")
		return
	}

	station, exists := h.graphService.Stations[stationID]
	if !exists {
		h.errorResponse(w, http.StatusNotFound, "station not found")
		return
	}

	// Get connected stations with their details
	connections := make([]model.Connection, 0, len(station.Connections))
	for _, conn := range station.Connections {
		if connectedStation, exists := h.graphService.Stations[conn.StationKey]; exists {
			connections = append(connections, model.Connection{
				StationKey: connectedStation.Key(),
				Distance:   conn.Distance,
				Duration:   conn.Duration,
				Type:       conn.Type,
				Line:       conn.Line,
			})
		}
	}

	h.response(w, http.StatusOK, connections)
}
