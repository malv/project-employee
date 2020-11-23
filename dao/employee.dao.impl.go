package dao

import (
	"project-employee/config"
	"project-employee/model"
	"project-employee/pojo"
	"strings"
)

type EmployeeDaoImpl struct{}

func (EmployeeDaoImpl) AddEmployee(data *model.Employee) (e error) {
	defer config.CatchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (EmployeeDaoImpl) EditEmployee(data *model.Employee) (e error) {
	defer config.CatchError(&e)
	result := g.Save(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (EmployeeDaoImpl) DeleteEmployee(id string) (e error) {
	defer config.CatchError(&e)
	var employee model.Employee
	result := g.Delete(&employee, id)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (EmployeeDaoImpl) GetEmployees() (data []model.Employee, e error) {
	defer config.CatchError(&e)
	var employees []model.Employee
	result := g.Preload("Company").Preload("Unit").Preload("Position").Find(&employees)
	if result.Error == nil {
		return employees, nil
	}
	return []model.Employee{}, result.Error
}

func (EmployeeDaoImpl) GetEmployeeById(id string) (data model.Employee, e error) {
	defer config.CatchError(&e)
	employee := model.Employee{BaseModel: model.BaseModel{Id: id}}
	result := g.Preload("Company").Preload("Unit").Preload("Position").First(&employee)

	if result.Error == nil {
		return employee, nil
	}
	return model.Employee{}, result.Error
}

func (EmployeeDaoImpl) GetManager() (data []pojo.PojoTempEmployee, e error) {
	defer config.CatchError(&e)
	var sb strings.Builder
	sb.WriteString("SELECT ce.id AS employee_id , ")
	sb.WriteString("ce.nik AS nik, ")
	sb.WriteString("ce.person_id AS person_id, ")
	sb.WriteString("cp.id AS position_id, ")
	sb.WriteString("cp.\"name\" AS position_name, ")
	sb.WriteString("cu.id AS unit_id, ")
	sb.WriteString("cu.name AS unit_name ")
	sb.WriteString("FROM cor_employees ce ")
	sb.WriteString("JOIN cor_positions cp ON cp.id = ce.position_id ")
	sb.WriteString("JOIN cor_units cu ON cu.id = ce.unit_id  ")
	sb.WriteString("WHERE cp.\"level\" = 3")

	rows, err := g.Raw(sb.String()).Rows()

	listManager := []pojo.PojoTempEmployee{}
	pojoUnit := pojo.PojoUnit{}
	pojoPosition := pojo.PojoPosition{}

	for rows.Next() {
		pojoGetManager := pojo.PojoTempEmployee{}
		rows.Scan(&pojoGetManager.Id, &pojoGetManager.Nik, &pojoGetManager.PersonId,
			&pojoPosition.Id, &pojoPosition.Name,
			&pojoUnit.Id, &pojoUnit.Name,
		)
		pojoGetManager.Position = pojoPosition
		pojoGetManager.Unit = pojoUnit

		listManager = append(listManager, pojoGetManager)
	}
	defer rows.Close()

	return listManager, err
}

func (EmployeeDaoImpl) GetStaffByUnit(unitId string) (data []pojo.PojoTempEmployee, e error) {
	defer config.CatchError(&e)
	var sb strings.Builder
	sb.WriteString("SELECT ce.id AS employee_id , ")
	sb.WriteString("ce.nik AS nik, ")
	sb.WriteString("ce.person_id AS person_id, ")
	sb.WriteString("cp.id AS position_id, ")
	sb.WriteString("cp.\"name\" AS position_name, ")
	sb.WriteString("cu.id AS unit_id, ")
	sb.WriteString("cu.name AS unit_name ")
	sb.WriteString("FROM cor_employees ce ")
	sb.WriteString("JOIN cor_positions cp ON cp.id = ce.position_id ")
	sb.WriteString("JOIN cor_units cu ON cu.id = ce.unit_id  ")
	sb.WriteString("WHERE cp.\"level\" < 3")
	sb.WriteString("AND cu.id = ?")

	// var listEmp []model.Employee
	// g.
	// 	Joins("JOIN ").
	// 	Joins("JOIN cor_units cu ON cu.id = cor_employees.unit_id").Select("cu.id,Position.Id").Where("Level < 3 AND cu.id = ? ", unitId).Find(&listEmp)
	// log.Print(listEmp)

	rows, err := g.Raw(sb.String(), unitId).Rows()

	listStaff := []pojo.PojoTempEmployee{}
	pojoUnit := pojo.PojoUnit{}
	pojoPosition := pojo.PojoPosition{}

	for rows.Next() {
		pojoGetStaff := pojo.PojoTempEmployee{}
		rows.Scan(&pojoGetStaff.Id, &pojoGetStaff.Nik, &pojoGetStaff.PersonId,
			&pojoPosition.Id, &pojoPosition.Name,
			&pojoUnit.Id, &pojoUnit.Name,
		)
		pojoGetStaff.Position = pojoPosition
		pojoGetStaff.Unit = pojoUnit

		listStaff = append(listStaff, pojoGetStaff)
	}
	defer rows.Close()
	return listStaff, err
}

func (EmployeeDaoImpl) GetAllByUnit(unitId string) (data []pojo.PojoTempEmployee, e error) {
	defer config.CatchError(&e)
	var sb strings.Builder
	sb.WriteString("SELECT ce.id AS employee_id , ")
	sb.WriteString("ce.nik AS nik, ")
	sb.WriteString("ce.person_id AS person_id, ")
	sb.WriteString("cp.id AS position_id, ")
	sb.WriteString("cp.\"name\" AS position_name, ")
	sb.WriteString("cu.id AS unit_id, ")
	sb.WriteString("cu.name AS unit_name ")
	sb.WriteString("FROM cor_employees ce ")
	sb.WriteString("JOIN cor_positions cp ON cp.id = ce.position_id ")
	sb.WriteString("JOIN cor_units cu ON cu.id = ce.unit_id  ")
	sb.WriteString("WHERE cu.id = ?")

	rows, err := g.Raw(sb.String(), unitId).Rows()

	listManager := []pojo.PojoTempEmployee{}
	pojoUnit := pojo.PojoUnit{}
	pojoPosition := pojo.PojoPosition{}

	for rows.Next() {
		pojoGetManager := pojo.PojoTempEmployee{}
		rows.Scan(&pojoGetManager.Id, &pojoGetManager.Nik, &pojoGetManager.PersonId,
			&pojoPosition.Id, &pojoPosition.Name,
			&pojoUnit.Id, &pojoUnit.Name,
		)
		pojoGetManager.Position = pojoPosition
		pojoGetManager.Unit = pojoUnit

		listManager = append(listManager, pojoGetManager)
	}
	defer rows.Close()
	return listManager, err
}
