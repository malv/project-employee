package service

import "project-employee/model"

type PositionService interface {
	AddPosition(data *model.Position) error
	EditPosition(data *model.Position) error
	DeletePosition(id string) error
	GetPositions() ([]model.Position, error)
	GetPositionById(id string) (model.Position, error)
}
