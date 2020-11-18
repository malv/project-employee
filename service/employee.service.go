package service

import (
	"project-employee/model"
	"project-employee/pojo"
)

type EmployeeService interface {
	AddEmployee(data *model.Employee) error
	EditEmployee(data *model.Employee) error
	DeleteEmployee(id string) error
	GetEmployees() ([]pojo.PojoGetEmployee, error)
	GetEmployeeById(id string) (pojo.PojoGetEmployee, error)
	GetManager() ([]pojo.PojoGetEmployee, error)
	GetStaffByUnit(unitId string) ([]pojo.PojoGetEmployee, error)
	GetAllByUnit(unitId string) ([]pojo.PojoGetEmployee, error)
}
