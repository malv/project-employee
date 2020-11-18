package service

import "project-employee/model"

type CompanyService interface {
	AddCompany(data *model.Company) error
	EditCompany(data *model.Company) error
	DeleteCompany(id string) error
	GetCompanies() ([]model.Company, error)
	GetCompanyById(id string) (model.Company, error)
}
