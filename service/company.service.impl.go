package service

import (
	"project-employee/config"
	"project-employee/dao"
	"project-employee/model"
)

type CompanyServiceImpl struct{}

var companyServiceImpl = CompanyServiceImpl{}
var companyDao dao.CompanyDao = dao.CompanyDaoImpl{}

func (CompanyServiceImpl) AddCompany(data *model.Company) (e error) {
	defer config.CatchError(&e)
	companyServiceImpl.validasiIdNull(*data)
	return companyDao.AddCompany(data)
}

func (CompanyServiceImpl) EditCompany(data *model.Company) (e error) {
	defer config.CatchError(&e)
	companyServiceImpl.validasiIdExist(*data)
	return companyDao.EditCompany(data)
}

func (CompanyServiceImpl) DeleteCompany(id string) (e error) {
	defer config.CatchError(&e)
	return companyDao.DeleteCompany(id)
}

func (CompanyServiceImpl) GetCompanies() (data []model.Company, e error) {
	defer config.CatchError(&e)
	result, err := companyDao.GetCompanys()
	if err == nil {
		return result, nil
	}
	return []model.Company{}, err
}

func (CompanyServiceImpl) GetCompanyById(id string) (data model.Company, e error) {
	defer config.CatchError(&e)
	result, err := companyDao.GetCompanyById(id)
	if err == nil {
		return result, nil
	}
	return model.Company{}, err
}

// ========================== validasi ================================ //

func (CompanyServiceImpl) validasiIdNull(data model.Company) {

	if data.BaseModel.Id != "" {
		panic("Id Must Be Null")
	}

}

func (CompanyServiceImpl) validasiIdNotNull(data model.Company) {

	if data.BaseModel.Id == "" {
		panic("Id Must Be Not Null")
	}
}

func (CompanyServiceImpl) validasiIdExist(data model.Company) {
	_, err := companyServiceImpl.GetCompanyById(data.BaseModel.Id)
	if err != nil {
		panic("Id not exist")
	}

}
