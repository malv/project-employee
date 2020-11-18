package dao

import "project-employee/model"

type UnitDao interface {
	AddUnit(data *model.Unit) error
	EditUnit(data *model.Unit) error
	DeleteUnit(id string) error
	GetUnits() ([]model.Unit, error)
	GetUnitById(id string) (model.Unit, error)
}
