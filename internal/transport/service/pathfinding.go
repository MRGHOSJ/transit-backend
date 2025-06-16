package service

import (
	"container/heap"
	"fmt"
	"transit-backend/internal/transport/components"
	"transit-backend/internal/transport/model"
)

// Priority queue implementation for A*
type PriorityQueue []*model.PathNode

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].TotalCost() < pq[j].TotalCost() }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*model.PathNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Updated FindPathAStar function with correct heap.Fix usage
func (g *Graph) FindPathAStar(
	startKey, endKey string,
	preferences model.TransportPreferences,
) ([]model.Connection, float64, error) {

	g.mu.RLock()
	defer g.mu.RUnlock()

	// Check if stations exist
	startStation, startExists := g.Stations[startKey]
	endStation, endExists := g.Stations[endKey]
	if !startExists || !endExists {
		return nil, 0, fmt.Errorf("start or end station not found")
	}

	// Initialize open and closed sets
	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)
	closedSet := make(map[string]bool)

	// Track nodes in open set for quick lookup
	openSetNodes := make(map[string]*model.PathNode)

	// Add start node
	startNode := &model.PathNode{
		Station:   *startStation,
		Cost:      0,
		Heuristic: g.heuristic(*startStation, *endStation),
	}
	heap.Push(&openSet, startNode)
	openSetNodes[startStation.Key()] = startNode

	// For reconstructing path
	cameFrom := make(map[string]*model.PathNode)

	for openSet.Len() > 0 {
		// Get node with lowest f-score
		current := heap.Pop(&openSet).(*model.PathNode)
		delete(openSetNodes, current.Station.Key())

		// Check if we've reached the destination
		if current.Station.Key() == endKey {
			return g.reconstructPath(cameFrom, current), current.Cost, nil
		}

		closedSet[current.Station.Key()] = true

		// Explore neighbors
		for _, conn := range current.Station.Connections {
			// Skip if mode not allowed
			if !preferences.AllowedModes[conn.Type] {
				continue
			}

			neighbor, exists := g.Stations[conn.StationKey]
			if !exists {
				continue
			}

			// Skip if already evaluated
			if closedSet[neighbor.Key()] {
				continue
			}

			// Calculate tentative cost
			tentativeCost := current.Cost + conn.Duration

			// Add transfer penalty if mode changed
			if current.Connection.Type != "" && current.Connection.Type != conn.Type {
				tentativeCost += preferences.TransferPenalty
			}

			if current.Connection.Line != "" && current.Connection.Line != conn.Line {
				tentativeCost += preferences.TransferPenalty
			}

			// Add mode-specific penalty
			tentativeCost += preferences.ModePenalties[conn.Type]

			// Check if neighbor is in open set
			existingNode, inOpenSet := openSetNodes[neighbor.Key()]

			if !inOpenSet {
				// New node discovered
				newNode := &model.PathNode{
					Station:    *neighbor,
					CameFrom:   current,
					Cost:       tentativeCost,
					Heuristic:  g.heuristic(*neighbor, *endStation),
					Connection: conn,
				}
				heap.Push(&openSet, newNode)
				openSetNodes[neighbor.Key()] = newNode
				cameFrom[neighbor.Key()] = current
			} else if tentativeCost < existingNode.Cost {
				// Found a better path to this node
				existingNode.Cost = tentativeCost
				existingNode.CameFrom = current
				existingNode.Connection = conn

				// Find the node's index in the heap
				for i, node := range openSet {
					if node.Station.Key() == neighbor.Key() {
						heap.Fix(&openSet, i)
						break
					}
				}
			}
		}
	}

	return nil, 0, fmt.Errorf("no path found")
}

func (g *Graph) heuristic(a, b model.Station) float64 {
	// Use Haversine distance divided by max speed (metro speed) as heuristic
	distance := components.HaversineDistance(a.Lat, a.Lng, b.Lat, b.Lng)
	return distance / 30.0 * 60 // Convert to minutes at 30 km/h
}

func (g *Graph) reconstructPath(cameFrom map[string]*model.PathNode, end *model.PathNode) []model.Connection {
	path := make([]model.Connection, 0)
	current := end

	for current.CameFrom != nil {
		path = append([]model.Connection{current.Connection}, path...)
		current = current.CameFrom
	}

	return path
}
