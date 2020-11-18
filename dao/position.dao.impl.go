package dao

import (
	"project-employee/config"
	"project-employee/model"
)

type PositionDaoImpl struct{}

func (PositionDaoImpl) AddPosition(data *model.Position) (e error) {
	defer config.CatchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (PositionDaoImpl) EditPosition(data *model.Position) (e error) {
	defer config.CatchError(&e)
	result := g.Save(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (PositionDaoImpl) DeletePosition(id string) (e error) {
	defer config.CatchError(&e)
	var Position model.Position
	result := g.Delete(&Position, id)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (PositionDaoImpl) GetPositions() (data []model.Position, e error) {
	defer config.CatchError(&e)
	var Positions []model.Position
	result := g.Find(&Positions)
	if result.Error == nil {
		return Positions, nil
	}
	return []model.Position{}, result.Error
}

func (PositionDaoImpl) GetPositionById(id string) (data model.Position, e error) {
	defer config.CatchError(&e)
	Position := model.Position{BaseModel: model.BaseModel{Id: id}}
	result := g.First(&Position)
	if result.Error == nil {
		return Position, nil
	}
	return model.Position{}, result.Error
}
