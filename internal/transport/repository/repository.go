package repository

import (
	"errors"
	"fmt"
	"strings"
	"transit-backend/internal/transport/model"
)

type transportRepository struct {
	data      *model.Data
	typeIndex model.TypeIndex
}

func NewTransportRepository(data *model.Data) TransportRepository {
	return &transportRepository{
		data:      data,
		typeIndex: buildTypeIndex(data.Lines),
	}
}

func buildTypeIndex(lines []model.Line) model.TypeIndex {
	index := make(model.TypeIndex)
	for _, line := range lines {
		key := strings.ToLower(line.Type)
		index[key] = append(index[key], line)
	}
	return index
}

// GetAllLines implements TransportRepository.
func (repo *transportRepository) GetAllLines() []model.Line {
	return repo.data.Lines
}

// GetLine implements TransportRepository.
func (repo *transportRepository) GetLine(t string, l string) (*model.Line, error) {
	t = strings.ToLower(t)
	l = strings.ToLower(l)

	lines, ok := repo.typeIndex[t]

	if !ok {
		return nil, fmt.Errorf("invalid transport type: %s", t)
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("no lines found for type %s", t)
	}

	var filteredLine *model.Line

	for _, line := range lines {
		if strings.EqualFold(l, line.Line) {
			filteredLine = &line
			break
		}
	}

	if filteredLine == nil {
		return nil, fmt.Errorf("line %s not found for type %s", l, t)
	}

	return filteredLine, nil
}

func (repo *transportRepository) GetLinesByType(t string) ([]model.Line, error) {
	t = strings.ToLower(t)
	lines, ok := repo.typeIndex[t]

	if !ok {
		return nil, errors.New("invalid transport type")
	}

	if len(lines) == 0 {
		return nil, errors.New("no lines found")
	}

	return lines, nil
}
