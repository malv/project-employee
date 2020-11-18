package dao

import "project-employee/model"

type CompanyDao interface {
	AddCompany(data *model.Company) error
	EditCompany(data *model.Company) error
	DeleteCompany(id string) error
	GetCompanys() ([]model.Company, error)
	GetCompanyById(id string) (model.Company, error)
}
