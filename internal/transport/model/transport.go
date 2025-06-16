package model

type TransportType string

const (
	Metro TransportType = "metro"
	Bus   TransportType = "bus"
	Walk  TransportType = "walk"
)

type Connection struct {
	StationKey string        `json:"station_id"`
	Distance   float64       `json:"distance"`
	Type       TransportType `json:"type"`
	Line       string        `json:"line,omitempty"`
	Duration   float64       `json:"duration"`
}
