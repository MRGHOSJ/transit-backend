package model

type PathNode struct {
	Station    Station
	CameFrom   *PathNode
	Cost       float64    // g score - actual cost from start
	Heuristic  float64    // h score - estimated cost to goal
	Connection Connection // how we got to this node
}

func (pn PathNode) TotalCost() float64 {
	return pn.Cost + pn.Heuristic // f score
}

type TransportPreferences struct {
	AllowedModes    map[TransportType]bool
	ModePenalties   map[TransportType]float64 // additional time penalty per mode
	TransferPenalty float64                   // penalty when changing transport modes
}
