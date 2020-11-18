package dao

import (
	"project-employee/config"
	"project-employee/model"
)

type UnitDaoImpl struct{}

func (UnitDaoImpl) AddUnit(data *model.Unit) (e error) {
	defer config.CatchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (UnitDaoImpl) EditUnit(data *model.Unit) (e error) {
	defer config.CatchError(&e)
	result := g.Save(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (UnitDaoImpl) DeleteUnit(id string) (e error) {
	defer config.CatchError(&e)
	var unit model.Unit
	result := g.Delete(&unit, id)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (UnitDaoImpl) GetUnits() (data []model.Unit, e error) {
	defer config.CatchError(&e)
	var units []model.Unit
	result := g.Find(&units)
	if result.Error == nil {
		return units, nil
	}
	return []model.Unit{}, result.Error
}

func (UnitDaoImpl) GetUnitById(id string) (data model.Unit, e error) {
	defer config.CatchError(&e)
	unit := model.Unit{BaseModel: model.BaseModel{Id: id}}
	result := g.First(&unit)
	if result.Error == nil {
		return unit, nil
	}
	return model.Unit{}, result.Error
}
