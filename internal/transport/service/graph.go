package service

import (
	"strings"
	"sync"
	"transit-backend/internal/transport/components"
	"transit-backend/internal/transport/model"
)

var (
	graphInstance *Graph
	graphOnce     sync.Once
)

type Graph struct {
	Stations map[string]*model.Station
	mu       sync.RWMutex
}

func GetGraph(lines []model.Line) *Graph {
	graphOnce.Do(func() {
		graphInstance = NewGraph(lines)
	})
	return graphInstance
}

func NewGraph(lines []model.Line) *Graph {
	graph := &Graph{
		Stations: make(map[string]*model.Station),
	}

	for _, line := range lines {
		lineType := model.TransportType(strings.ToLower(line.Type))
		for i := 0; i < len(line.Stations); i++ {
			station := line.Stations[i]
			stationKey := station.Key()
			_, exists := graph.Stations[stationKey]

			if !exists {
				newStation := station
				newStation.Type = lineType
				newStation.Line = line.Line
				newStation.Connections = []model.Connection{}
				graph.Stations[stationKey] = &newStation
			}

			if i > 0 {
				prevStation := line.Stations[i-1]
				distance := components.HaversineDistance(
					station.Lat, station.Lng,
					prevStation.Lat, prevStation.Lng,
				)
				duration := calculateDuration(distance, lineType)

				graph.addConnection(station, prevStation, distance, duration, lineType, line.Line)
				graph.addConnection(prevStation, station, distance, duration, lineType, line.Line)
			}

		}
	}

	graph.addWalkingConnections(0.5)

	return graph
}

func calculateDuration(distance float64, transportType model.TransportType) float64 {
	type transportMode struct {
		speedKmH    float64 // Speed in km/h
		fixedDelayM float64 // Fixed delay in minutes (boarding, waiting, etc.)
	}

	modes := map[model.TransportType]transportMode{
		model.Metro: {
			speedKmH:    30.0,
			fixedDelayM: 2.0,
		},
		model.Bus: {
			speedKmH:    20.0,
			fixedDelayM: 5.0,
		},
		model.Walk: {
			speedKmH:    5.0, // More realistic walking speed
			fixedDelayM: 0.0,
		},
	}

	mode, exists := modes[transportType]
	if !exists {
		mode = transportMode{
			speedKmH:    10.0,
			fixedDelayM: 0.0,
		}
	}

	// Calculate travel time in minutes (distance in km, speed in km/h)
	travelTime := (distance / mode.speedKmH) * 60

	// Add fixed delay for this transport type
	totalDuration := travelTime + mode.fixedDelayM

	return totalDuration
}

func (g *Graph) addConnection(from, to model.Station, distance, duration float64,
	connType model.TransportType, line string) {

	g.mu.Lock()
	defer g.mu.Unlock()

	fromKey := from.Key()
	toKey := to.Key()

	// Check if connection already exists
	for _, conn := range from.Connections {
		if conn.StationKey == toKey && conn.Type == connType {
			return
		}
	}

	g.Stations[fromKey].Connections = append(g.Stations[fromKey].Connections, model.Connection{
		StationKey: toKey,
		Distance:   distance,
		Duration:   duration,
		Type:       connType,
		Line:       line,
	})
}

func (g *Graph) addWalkingConnections(maxDistance float64) {
	// This is a simplified approach - in production you'd use spatial indexing
	for _, s1 := range g.Stations {
		for _, s2 := range g.Stations {
			if s1.Key() == s2.Key() {
				continue
			}

			distance := components.HaversineDistance(s1.Lat, s1.Lng, s2.Lat, s2.Lng)
			if distance <= maxDistance {
				duration := calculateDuration(distance, model.Walk)
				g.addConnection(*s1, *s2, distance, duration, model.Walk, "")
			}
		}
	}
}
