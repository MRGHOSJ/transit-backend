package model

import "fmt"

type Station struct {
	Lat         float64       `json:"lat"`
	Lng         float64       `json:"lng"`
	Name        string        `json:"name,omitempty"`        // Optional if available
	Type        TransportType `json:"type,omitempty"`        // metro/bus
	Line        string        `json:"line,omitempty"`        // Line identifier
	Connections []Connection  `json:"connections,omitempty"` // Added this field

}

func (s Station) Key() string {
	return fmt.Sprintf("%.6f,%.6f", s.Lat, s.Lng)
}

func (s Station) IsSame(other Station) bool {
	return s.Lat == other.Lat && s.Lng == other.Lng
}
