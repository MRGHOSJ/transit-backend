package service

import (
	"errors"
	"math"
	"transit-backend/internal/transport/components"
	"transit-backend/internal/transport/model"
	"transit-backend/internal/transport/repository"
)

type LinesService struct {
	repo repository.TransportRepository
}

func NewLinesService(repo repository.TransportRepository) *LinesService {
	return &LinesService{repo: repo}
}

func (s *LinesService) GetAllLines() []model.Line {
	return s.repo.GetAllLines()
}

func (s *LinesService) GetLinesByType(t string) ([]model.Line, error) {
	return s.repo.GetLinesByType(t)
}

func (s *LinesService) GetLine(t string, l string) (*model.Line, error) {
	return s.repo.GetLine(t, l)
}

func (s *LinesService) FindClosestStation(lat, lng float64) (model.Line, model.Station, float64, error) {
	var closestLine model.Line
	var closestStation model.Station
	minDistance := math.MaxFloat64

	for _, line := range s.repo.GetAllLines() {
		for _, station := range line.Stations {
			distance := components.HaversineDistance(lat, lng, station.Lat, station.Lng)
			if distance < minDistance {
				minDistance = distance
				closestStation = station
				closestLine = line
			}
		}
	}

	if minDistance == math.MaxFloat64 {
		return model.Line{}, model.Station{}, 0, errors.New("no stations found")
	}

	return closestLine, closestStation, minDistance, nil
}
