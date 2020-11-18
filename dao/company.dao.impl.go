package dao

import (
	"project-employee/config"
	"project-employee/model"
)

type CompanyDaoImpl struct{}

func (CompanyDaoImpl) AddCompany(data *model.Company) (e error) {
	defer config.CatchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (CompanyDaoImpl) EditCompany(data *model.Company) (e error) {
	defer config.CatchError(&e)
	result := g.Save(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (CompanyDaoImpl) DeleteCompany(id string) (e error) {
	defer config.CatchError(&e)
	var unit model.Company
	result := g.Delete(&unit, id)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (CompanyDaoImpl) GetCompanys() (data []model.Company, e error) {
	defer config.CatchError(&e)
	var units []model.Company
	result := g.Find(&units)
	if result.Error == nil {
		return units, nil
	}
	return []model.Company{}, result.Error
}

func (CompanyDaoImpl) GetCompanyById(id string) (data model.Company, e error) {
	defer config.CatchError(&e)
	unit := model.Company{BaseModel: model.BaseModel{Id: id}}
	result := g.First(&unit)
	if result.Error == nil {
		return unit, nil
	}
	return model.Company{}, result.Error
}
