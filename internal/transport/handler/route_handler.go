// handler/route_handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"transit-backend/internal/transport/model"
	"transit-backend/internal/transport/service"
)

type RouteHandler struct {
	routePlanner *service.RoutePlanner
}

func NewRouteHandler(planner *service.RoutePlanner) *RouteHandler {
	return &RouteHandler{routePlanner: planner}
}

func (h *RouteHandler) GetRoutes(w http.ResponseWriter, r *http.Request) {
	// Parse from coordinates
	from := r.URL.Query().Get("from")
	fromParts := strings.Split(from, ",")
	if len(fromParts) != 2 {
		h.errorResponse(w, http.StatusBadRequest, "invalid 'from' parameter")
		return
	}
	fromLat, err := strconv.ParseFloat(fromParts[0], 64)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, "invalid latitude in 'from'")
		return
	}
	fromLng, err := strconv.ParseFloat(fromParts[1], 64)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, "invalid longitude in 'from'")
		return
	}

	// Parse to coordinates
	to := r.URL.Query().Get("to")
	toParts := strings.Split(to, ",")
	if len(toParts) != 2 {
		h.errorResponse(w, http.StatusBadRequest, "invalid 'to' parameter")
		return
	}
	toLat, err := strconv.ParseFloat(toParts[0], 64)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, "invalid latitude in 'to'")
		return
	}
	toLng, err := strconv.ParseFloat(toParts[1], 64)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, "invalid longitude in 'to'")
		return
	}

	// Parse transport modes
	modesParam := r.URL.Query().Get("modes")
	if modesParam == "" {
		modesParam = "metro,bus,walk" // Default modes
	}
	modeStrings := strings.Split(modesParam, ",")
	var modes []model.TransportType
	for _, m := range modeStrings {
		switch strings.ToLower(m) {
		case "metro":
			modes = append(modes, model.Metro)
		case "bus":
			modes = append(modes, model.Bus)
		case "walk":
			modes = append(modes, model.Walk)
		default:
			h.errorResponse(w, http.StatusBadRequest, "invalid transport mode: "+m)
			return
		}
	}

	prefsParam := r.URL.Query().Get("prefs")
	if prefsParam == "" {
		prefsParam = "fast" // Default modes
	} else if prefsParam != "fast" && prefsParam != "comfort" {

		h.errorResponse(w, http.StatusBadRequest, "invalid transport preferance: "+prefsParam+" should be 'fast' or 'comfort'")
		return
	}

	// Get route options
	options, err := h.routePlanner.FindRoutes(fromLat, fromLng, toLat, toLng, modes, prefsParam)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(options) == 0 {
		h.errorResponse(w, http.StatusNotFound, "no routes found")
		return
	}

	h.response(w, http.StatusOK, map[string]interface{}{
		"routes": options,
	})
}

func (h *RouteHandler) response(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *RouteHandler) errorResponse(w http.ResponseWriter, status int, message string) {
	h.response(w, status, map[string]string{"error": message})
}
