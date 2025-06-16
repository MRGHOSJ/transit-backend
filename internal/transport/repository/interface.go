package repository

import "transit-backend/internal/transport/model"

type TransportRepository interface {
	GetAllLines() []model.Line
	GetLinesByType(t string) ([]model.Line, error)
	GetLine(t, l string) (*model.Line, error)
}
