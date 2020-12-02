package service

import (
	"log"
	"project-employee/config"
	"project-employee/dao"
	"project-employee/model"
	"project-employee/pojo"
	"sync"
)

type EmployeeServiceImpl struct{}

var employeeServiceImpl = EmployeeServiceImpl{}
var employeeDao dao.EmployeeDao = dao.EmployeeDaoImpl{}

func (EmployeeServiceImpl) AddEmployee(data *model.Employee) (e error) {
	defer config.CatchError(&e)
	employeeServiceImpl.validasiIdNull(*data)
	employeeServiceImpl.validasiBKIsExist(data)
	return employeeDao.AddEmployee(data)
}

func (EmployeeServiceImpl) EditEmployee(data *model.Employee) (e error) {
	defer config.CatchError(&e)
	employeeServiceImpl.validasiIdExist(*data)
	return employeeDao.EditEmployee(data)
}

func (EmployeeServiceImpl) DeleteEmployee(id string) (e error) {
	defer config.CatchError(&e)
	return employeeDao.DeleteEmployee(id)
}

func (EmployeeServiceImpl) GetEmployees() (data []pojo.PojoGetEmployee, e error) {
	defer config.CatchError(&e)
	var wg sync.WaitGroup
	result, err := employeeDao.GetEmployees()
	if err == nil {
		var listEmployees []pojo.PojoGetEmployee
		for i := 0; i < len(result); i++ {
			wg.Add(1)
			pojoEmp := pojo.PojoGetEmployee{}
			person, err := GetPersonById(result[i].PersonId)
			log.Print(person)
			if err == nil {
				pojoEmp.Id = result[i].Id
				pojoEmp.Nik = result[i].Nik
				pojoEmp.Name = person["first_name"].(string)
				if person["last_name"] != nil {
					pojoEmp.Name = pojoEmp.Name + " " + person["last_name"].(string)
				}
				pojoEmp.Position = pojo.PojoPosition{Id: result[i].Position.Id, Name: result[i].Position.Name}
				pojoEmp.Unit = pojo.PojoUnit{Id: result[i].Unit.Id, Name: result[i].Unit.Name}
				listEmployees = append(listEmployees, pojoEmp)
			}
		}
		log.Print(listEmployees)
		return listEmployees, nil
	}
	return data, err
}

func (EmployeeServiceImpl) GetEmployeeById(id string) (data pojo.PojoGetEmployee, e error) {
	defer config.CatchError(&e)
	result, err := employeeDao.GetEmployeeById(id)
	if err == nil {
		pojoEmp := pojo.PojoGetEmployee{}
		person, err := GetPersonById(result.PersonId)
		if err == nil {
			pojoEmp.Id = result.Id
			pojoEmp.Nik = result.Nik
			pojoEmp.Name = person["first_name"].(string)
			if person["last_name"] != nil {
				pojoEmp.Name = pojoEmp.Name + " " + person["last_name"].(string)
			}
			pojoEmp.Position = pojo.PojoPosition{Id: result.Position.Id, Name: result.Position.Name}
			pojoEmp.Unit = pojo.PojoUnit{Id: result.Unit.Id, Name: result.Unit.Name}
		}
		return pojoEmp, nil
	}
	return data, err
}

func (EmployeeServiceImpl) GetManager() (data []pojo.PojoGetEmployee, e error) {
	result, err := employeeDao.GetManager()
	if err == nil {
		var listPojo []pojo.PojoGetEmployee
		for i := 0; i < len(result); i++ {
			pojoEmp := pojo.PojoGetEmployee{}
			person, err := GetPersonById(result[i].PersonId)
			if err == nil {
				pojoEmp.Id = result[i].Id
				pojoEmp.Nik = result[i].Nik
				pojoEmp.Name = person["first_name"].(string)
				if person["last_name"] != nil {
					pojoEmp.Name = pojoEmp.Name + " " + person["last_name"].(string)
				}
				pojoEmp.Position = result[i].Position
				pojoEmp.Unit = result[i].Unit
				listPojo = append(listPojo, pojoEmp)
			}
		}
		return listPojo, err
	}
	return data, err
}

func (EmployeeServiceImpl) GetStaffByUnit(unitId string) (data []pojo.PojoGetEmployee, e error) {
	result, err := employeeDao.GetStaffByUnit(unitId)
	if err == nil {
		var listPojo []pojo.PojoGetEmployee
		for i := 0; i < len(result); i++ {
			pojoEmp := pojo.PojoGetEmployee{}
			person, err := GetPersonById(result[i].PersonId)
			if err == nil {
				pojoEmp.Id = result[i].Id
				pojoEmp.Nik = result[i].Nik
				pojoEmp.Name = person["first_name"].(string)
				if person["last_name"] != nil {
					pojoEmp.Name = pojoEmp.Name + " " + person["last_name"].(string)
				}
				pojoEmp.Position = result[i].Position
				pojoEmp.Unit = result[i].Unit
				listPojo = append(listPojo, pojoEmp)
			}
		}
		return listPojo, err
	}
	return data, err
}

func (EmployeeServiceImpl) GetAllByUnit(unitId string) (data []pojo.PojoGetEmployee, e error) {
	result, err := employeeDao.GetAllByUnit(unitId)
	if err == nil {
		var listPojo []pojo.PojoGetEmployee
		for i := 0; i < len(result); i++ {
			pojoEmp := pojo.PojoGetEmployee{}
			person, err := GetPersonById(result[i].PersonId)
			if err == nil {
				pojoEmp.Id = result[i].Id
				pojoEmp.Nik = result[i].Nik
				pojoEmp.Name = person["first_name"].(string)
				if person["last_name"] != nil {
					pojoEmp.Name = pojoEmp.Name + " " + person["last_name"].(string)
				}
				pojoEmp.Position = result[i].Position
				pojoEmp.Unit = result[i].Unit
				listPojo = append(listPojo, pojoEmp)
			}
		}
		return listPojo, err
	}
	return data, err
}

// ========================== validasi ================================ //

func (EmployeeServiceImpl) validasiIdNull(data model.Employee) {

	if data.BaseModel.Id != "" {
		panic("Id Must Be Null")
	}

}

func (EmployeeServiceImpl) validasiIdNotNull(data model.Employee) {

	if data.BaseModel.Id == "" {
		panic("Id Must Be Not Null")
	}
}

func (EmployeeServiceImpl) validasiIdExist(data model.Employee) {
	_, err := employeeServiceImpl.GetEmployeeById(data.BaseModel.Id)
	if err != nil {
		panic("Id not exist")
	}

}

func (EmployeeServiceImpl) validasiBKIsExist(data *model.Employee) (e error) {
	defer config.CatchError(&e)
	_, err := employeeDao.GetEmployeeByPerson(data.PersonId)
	if err != nil {
		return err
	}
	return nil
}

func (EmployeeServiceImpl) GetEmployeesWithToken(token string) (data []pojo.PojoGetEmployee, e error) {
	defer config.CatchError(&e)
	var wg sync.WaitGroup
	result, err := employeeDao.GetEmployees()
	if err == nil {
		var listEmployees []pojo.PojoGetEmployee
		for i := 0; i < len(result); i++ {
			wg.Add(1)
			pojoEmp := pojo.PojoGetEmployee{}
			person, err := GetPersonByIdWithToken(result[i].PersonId, token)
			if err == nil {
				pojoEmp.Id = result[i].Id
				pojoEmp.Nik = result[i].Nik
				pojoEmp.Name = person["first_name"].(string)
				if person["last_name"] != nil {
					pojoEmp.Name = pojoEmp.Name + " " + person["last_name"].(string)
				}
				pojoEmp.Position = pojo.PojoPosition{Id: result[i].Position.Id, Name: result[i].Position.Name}
				pojoEmp.Unit = pojo.PojoUnit{Id: result[i].Unit.Id, Name: result[i].Unit.Name}
				listEmployees = append(listEmployees, pojoEmp)
			}
		}
		return listEmployees, nil
	}
	return data, err
}

// func (EmployeeServiceImpl) GetEmployeesFromProto() (data *pb.Employees, e error) {
// 	defer config.CatchError(&e)
// 	res, err := config.ClientEmployee.GetEmployees(config.CtxEmployee, &pb.Empty{Token: ReqToken})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return res, nil
// }
