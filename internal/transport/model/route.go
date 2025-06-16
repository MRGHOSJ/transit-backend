// model/route.go
package model

type RouteOption struct {
	Segments      []RouteSegment `json:"segments"`
	TotalDuration float64        `json:"total_duration"`
	TotalDistance float64        `json:"total_distance"`
	Transfers     int            `json:"transfers"`
}

type RouteSegment struct {
	From         Station       `json:"from"`
	To           Station       `json:"to"`
	Mode         TransportType `json:"mode"`
	Line         string        `json:"line,omitempty"`
	Duration     float64       `json:"duration"`
	Distance     float64       `json:"distance"`
	Instructions string        `json:"instructions"`
}
