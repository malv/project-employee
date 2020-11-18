package service

import (
	"project-employee/config"
	"project-employee/dao"
	"project-employee/model"
)

type UnitServiceImpl struct{}

var unitDao dao.UnitDao = dao.UnitDaoImpl{}
var unitServiceImpl = UnitServiceImpl{}

func (UnitServiceImpl) AddUnit(data *model.Unit) (e error) {
	defer config.CatchError(&e)
	unitServiceImpl.validasiIdNull(*data)
	return unitDao.AddUnit(data)
}

func (UnitServiceImpl) EditUnit(data *model.Unit) (e error) {
	defer config.CatchError(&e)
	unitServiceImpl.validasiIdExist(*data)
	return unitDao.EditUnit(data)
}

func (UnitServiceImpl) DeleteUnit(id string) (e error) {
	defer config.CatchError(&e)
	return unitDao.DeleteUnit(id)
}

func (UnitServiceImpl) GetUnits() (data []model.Unit, e error) {
	defer config.CatchError(&e)
	result, err := unitDao.GetUnits()
	if err == nil {
		return result, nil
	}
	return []model.Unit{}, err
}

func (UnitServiceImpl) GetUnitById(id string) (data model.Unit, e error) {
	defer config.CatchError(&e)
	result, err := unitDao.GetUnitById(id)
	if err == nil {
		return result, nil
	}
	return model.Unit{}, err
}

// ========================== validasi ================================ //

func (UnitServiceImpl) validasiIdNull(data model.Unit) {

	if data.BaseModel.Id != "" {
		panic("Id Must Be Null")
	}

}

func (UnitServiceImpl) validasiIdNotNull(data model.Unit) {

	if data.BaseModel.Id == "" {
		panic("Id Must Be Not Null")
	}
}

func (UnitServiceImpl) validasiIdExist(data model.Unit) {
	_, err := unitServiceImpl.GetUnitById(data.BaseModel.Id)
	if err != nil {
		panic("Id not exist")
	}

}
