package model

type TypeIndex map[string][]Line

type Schedule struct {
	Frequency              string `json:"frequency"`
	LastTrip               string `json:"last_trip"`
	FirstTrip              string `json:"first_trip"`
	FirstTripToMainStation string `json:"first_trip_to_main_station"`
	LastTripToMainStation  string `json:"last_trip_to_main_station"`
}

type Line struct {
	Type      string    `json:"type"`
	Line      string    `json:"line"`
	Departure string    `json:"departure"`
	Arrival   string    `json:"arrival"`
	Stations  []Station `json:"stations"`
	Schedule  Schedule  `json:"schedule"`
}

type Data struct {
	Lines []Line `json:"lines"`
}
