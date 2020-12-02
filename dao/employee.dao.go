package dao

import (
	"project-employee/model"
	"project-employee/pojo"
)

type EmployeeDao interface {
	AddEmployee(data *model.Employee) error
	EditEmployee(data *model.Employee) error
	DeleteEmployee(id string) error
	GetEmployees() ([]model.Employee, error)
	GetEmployeeById(id string) (model.Employee, error)
	GetEmployeeByPerson(id string) (model.Employee, error)
	GetManager() ([]pojo.PojoTempEmployee, error)
	GetStaffByUnit(unitId string) ([]pojo.PojoTempEmployee, error)
	GetAllByUnit(unitId string) ([]pojo.PojoTempEmployee, error)
}
