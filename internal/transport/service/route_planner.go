// service/route_planner.go
package service

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"transit-backend/internal/transport/components"
	"transit-backend/internal/transport/model"
)

type RoutePlanner struct {
	graph *Graph
}

func NewRoutePlanner(graph *Graph) *RoutePlanner {
	return &RoutePlanner{graph: graph}
}

func (rp *RoutePlanner) getPreferences(mode string) model.TransportPreferences {
	prefs := model.TransportPreferences{
		AllowedModes: make(map[model.TransportType]bool),
		ModePenalties: map[model.TransportType]float64{
			model.Metro: 0,
			model.Bus:   2, // Prefer metro over bus by adding small penalty
			model.Walk:  5, // Prefer public transport over walking
		},
		TransferPenalty: 5, // 5 minute penalty for transfers
	}
	comfortPrefs := model.TransportPreferences{
		AllowedModes: make(map[model.TransportType]bool),
		ModePenalties: map[model.TransportType]float64{
			model.Walk:  5.0, // discourage long walks
			model.Bus:   1.0,
			model.Metro: 1.0,
		},
		TransferPenalty: 10.0, // strongly penalize line changes
	}

	fastPrefs := model.TransportPreferences{
		AllowedModes: make(map[model.TransportType]bool),
		ModePenalties: map[model.TransportType]float64{
			model.Walk:  1.0,
			model.Bus:   1.0,
			model.Metro: 1.0,
		},
		TransferPenalty: 1.0,
	}

	switch mode {
	case "fast":
		return fastPrefs
	case "comfort":
		return comfortPrefs
	default:
		return prefs
	}
}

func (rp *RoutePlanner) FindRoutes(fromLat, fromLng, toLat, toLng float64, modes []model.TransportType, prefsParam string) ([]model.RouteOption, error) {

	// Convert modes to preferences
	prefs := rp.getPreferences(prefsParam)

	for _, mode := range modes {
		prefs.AllowedModes[mode] = true
	}

	// Find nearest accessible stations
	startStations := rp.findNearestAccessPoints(fromLat, fromLng, modes)
	endStations := rp.findNearestAccessPoints(toLat, toLng, modes)

	var options []model.RouteOption

	for _, start := range startStations {
		for _, end := range endStations {
			// Find path using A*
			var path []model.Connection
			var totalDuration float64
			var err error

			// Select preferences based on user input
			path, totalDuration, err = rp.graph.FindPathAStar(
				start.StationKey,
				end.StationKey,
				prefs,
			)

			if err != nil {
				continue
			}

			// Build route segments
			segments := make([]model.RouteSegment, 0)
			currentPos := model.Station{Lat: fromLat, Lng: fromLng}
			totalDistance := 0.0

			// Add initial walking segment if needed
			if currentPos.Key() != start.StationKey {

				parts := strings.Split(start.StationKey, ",")
				if len(parts) != 2 {
					panic("Invalid coordinate format")
				}

				startLat, err := strconv.ParseFloat(parts[0], 64)
				if err != nil {
					panic("Invalid latitude")
				}

				startLng, err := strconv.ParseFloat(parts[1], 64)
				if err != nil {
					panic("Invalid longitude")
				}

				walkDist := components.HaversineDistance(
					currentPos.Lat, currentPos.Lng,
					startLat, startLng,
				)
				walkDur := calculateDuration(walkDist, model.Walk)

				segments = append(segments, model.RouteSegment{
					From:         currentPos,
					To:           model.Station{Lat: startLat, Lng: startLng},
					Mode:         model.Walk,
					Duration:     walkDur,
					Distance:     walkDist,
					Instructions: "Walk to " + rp.graph.Stations[start.StationKey].Name,
				})
				totalDistance += walkDist
				currentPos = model.Station{Lat: startLat, Lng: startLng}
			}

			// Add path segments
			for _, conn := range path {
				toStation := rp.graph.Stations[conn.StationKey]
				segments = append(segments, model.RouteSegment{
					From:         currentPos,
					To:           *toStation,
					Mode:         conn.Type,
					Line:         conn.Line,
					Duration:     conn.Duration,
					Distance:     conn.Distance,
					Instructions: rp.getInstructions(conn, currentPos, *toStation),
				})
				totalDistance += conn.Distance
				currentPos = *toStation
			}

			// Add final walking segment if needed
			if !currentPos.IsSame(model.Station{Lat: toLat, Lng: toLng}) {
				walkDist := components.HaversineDistance(
					currentPos.Lat, currentPos.Lng,
					toLat, toLng,
				)
				walkDur := calculateDuration(walkDist, model.Walk)

				segments = append(segments, model.RouteSegment{
					From:         currentPos,
					To:           model.Station{Lat: toLat, Lng: toLng},
					Mode:         model.Walk,
					Duration:     walkDur,
					Distance:     walkDist,
					Instructions: "Walk to destination",
				})
				totalDistance += walkDist
			}

			options = append(options, model.RouteOption{
				Segments:      segments,
				TotalDuration: totalDuration,
				TotalDistance: totalDistance,
				Transfers:     rp.countTransfers(segments),
			})
		}
	}

	// Sort options by total duration
	sort.Slice(options, func(i, j int) bool {
		return options[i].TotalDuration < options[j].TotalDuration
	})

	// Return top 3 options
	if len(options) > 3 {
		return options[:3], nil
	}
	return options, nil
}

// service/route_planner.go
func (rp *RoutePlanner) findNearestAccessPoints(lat, lng float64, modes []model.TransportType) []model.Connection {
	// Convert modes to a set for quick lookup
	modeSet := make(map[model.TransportType]bool)
	for _, mode := range modes {
		modeSet[mode] = true
	}

	var results []model.Connection
	minDistance := math.MaxFloat64

	// Find closest stations for each allowed mode
	for _, station := range rp.graph.Stations {
		// Skip stations that don't match our allowed modes
		if !modeSet[station.Type] {
			continue
		}

		distance := components.HaversineDistance(lat, lng, station.Lat, station.Lng)

		// Keep track of the closest stations (within 1km or other threshold)
		if distance < 1.0 { // 1km radius
			duration := calculateDuration(distance, model.Walk)

			results = append(results, model.Connection{
				StationKey: station.Key(),
				Distance:   distance,
				Duration:   duration,
				Type:       model.Walk, // The connection TO the station is always walking
			})

			// Track absolute minimum distance
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	// If we found nothing within 1km, return the absolute closest station
	if len(results) == 0 {
		for _, station := range rp.graph.Stations {
			if !modeSet[station.Type] {
				continue
			}

			distance := components.HaversineDistance(lat, lng, station.Lat, station.Lng)
			if distance < minDistance {
				minDistance = distance
				duration := calculateDuration(distance, model.Walk)

				// Replace results with just this one station
				results = []model.Connection{{
					StationKey: station.Key(),
					Distance:   distance,
					Duration:   duration,
					Type:       model.Walk,
				}}
			}
		}
	}

	// Sort by distance and return top 3 closest stations
	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})

	if len(results) > 3 {
		return results[:3]
	}
	return results
}

func (rp *RoutePlanner) countTransfers(segments []model.RouteSegment) int {
	transfers := 0
	for i := 1; i < len(segments); i++ {
		if segments[i].Mode != segments[i-1].Mode {
			transfers++
		}
	}
	return transfers
}

func (rp *RoutePlanner) getInstructions(conn model.Connection, from, to model.Station) string {
	switch conn.Type {
	case model.Metro:
		return "Take " + conn.Line + " metro from " + from.Name + " to " + to.Name
	case model.Bus:
		return "Take " + conn.Line + " bus from " + from.Name + " to " + to.Name
	case model.Walk:
		return "Walk from " + from.Name + " to " + to.Name
	default:
		return "Continue to " + to.Name
	}
}
