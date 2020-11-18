package service

import (
	"project-employee/config"
	"project-employee/dao"
	"project-employee/model"
)

type PositionServiceImpl struct{}

var positionServiceImpl = PositionServiceImpl{}
var positionDao dao.PositionDao = dao.PositionDaoImpl{}

func (PositionServiceImpl) AddPosition(data *model.Position) (e error) {
	defer config.CatchError(&e)
	positionServiceImpl.validasiIdNull(*data)
	return positionDao.AddPosition(data)
}

func (PositionServiceImpl) EditPosition(data *model.Position) (e error) {
	defer config.CatchError(&e)
	positionServiceImpl.validasiIdExist(*data)
	return positionDao.EditPosition(data)
}

func (PositionServiceImpl) DeletePosition(id string) (e error) {
	defer config.CatchError(&e)
	return positionDao.DeletePosition(id)
}

func (PositionServiceImpl) GetPositions() (data []model.Position, e error) {
	defer config.CatchError(&e)
	result, err := positionDao.GetPositions()
	if err == nil {
		return result, nil
	}
	return []model.Position{}, err
}

func (PositionServiceImpl) GetPositionById(id string) (data model.Position, e error) {
	defer config.CatchError(&e)
	result, err := positionDao.GetPositionById(id)
	if err == nil {
		return result, nil
	}
	return model.Position{}, err
}

// ========================== validasi ================================ //

func (PositionServiceImpl) validasiIdNull(data model.Position) {

	if data.BaseModel.Id != "" {
		panic("Id Must Be Null")
	}

}

func (PositionServiceImpl) validasiIdNotNull(data model.Position) {

	if data.BaseModel.Id == "" {
		panic("Id Must Be Not Null")
	}
}

func (PositionServiceImpl) validasiIdExist(data model.Position) {
	_, err := positionServiceImpl.GetPositionById(data.BaseModel.Id)
	if err != nil {
		panic("Id not exist")
	}

}
