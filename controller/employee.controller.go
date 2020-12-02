package controller

import (
	"errors"
	"project-employee/config"
	"project-employee/model"
	"project-employee/service"

	pb "project-employee/proto/model"

	"github.com/labstack/echo"
)

var employeeService service.EmployeeService = service.EmployeeServiceImpl{}

func SetEmployeeController(e *echo.Group) {
	e.POST("/employees", addEmployee)
	e.PUT("/employees", editEmployee)
	e.DELETE("/employee/:id", deleteEmployee)
	e.GET("/employees", getEmployees)
	e.GET("/employee/:id", getEmployeeById)
	e.GET("/employee/manager", getManager)
	e.GET("/employee/:unitId/staff", getStaffByUnit)
	e.GET("/employee/unit/:unitId", getAllByUnit)
}

func addEmployee(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := &model.Employee{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := employeeService.AddEmployee(data)
	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func editEmployee(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := &model.Employee{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := employeeService.EditEmployee(data)
	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func deleteEmployee(c echo.Context) (e error) {
	defer config.CatchError(&e)
	id := c.Param("id")

	err := employeeService.DeleteEmployee(id)

	if err == nil {
		return resSuccess(c)
	}
	return resErr(c, err)
}

func getEmployees(c echo.Context) (e error) {
	defer config.CatchError(&e)

	// result, err := employeeService.GetEmployees()
	result, err := config.ClientEmployee.GetEmployees(config.CtxEmployee, &pb.Tokens{Token: service.ReqToken})

	if err == nil {
		return res(c, result)
	}
	return resErr(c, err)
}

func getEmployeeById(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	result, err := employeeService.GetEmployeeById(id)

	if err == nil {
		return res(c, result)
	}

	return resErr(c, errors.New("id not found"))
}

func getManager(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := employeeService.GetManager()

	if err == nil {
		return res(c, result)
	}

	return resErr(c, errors.New("manager not found"))
}

func getStaffByUnit(c echo.Context) (e error) {
	defer config.CatchError(&e)

	unitId := c.Param("unitId")

	result, err := employeeService.GetStaffByUnit(unitId)

	if err == nil {
		return res(c, result)
	}

	return resErr(c, errors.New("manager not found"))
}

func getAllByUnit(c echo.Context) (e error) {
	defer config.CatchError(&e)

	unitId := c.Param("unitId")

	result, err := employeeService.GetAllByUnit(unitId)

	if err == nil {
		return res(c, result)
	}

	return resErr(c, errors.New("manager not found"))
}
