package controller

import (
	"errors"
	"project-employee/config"
	"project-employee/model"
	"project-employee/service"

	"github.com/labstack/echo"
)

var companyService service.CompanyService = service.CompanyServiceImpl{}

func SetCompanyController(e *echo.Group) {
	e.POST("/companies", addCompany)
	e.PUT("/companies", editCompany)
	e.DELETE("/company/:id", deleteCompany)
	e.GET("/companies", getCompanies)
	e.GET("/company/:id", getCompanyById)
}

func addCompany(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := &model.Company{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := companyService.AddCompany(data)
	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func editCompany(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := &model.Company{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := companyService.EditCompany(data)
	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func deleteCompany(c echo.Context) (e error) {
	defer config.CatchError(&e)
	id := c.Param("id")

	err := companyService.DeleteCompany(id)

	if err == nil {
		return resSuccess(c)
	}
	return resErr(c, err)
}

func getCompanies(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := companyService.GetCompanies()

	if err == nil {
		if len(result) > 0 {
			return res(c, result)
		}
		return resErr(c, errors.New("data not found"))
	}
	return resErr(c, err)
}

func getCompanyById(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	result, err := companyService.GetCompanyById(id)

	if err == nil {
		return res(c, result)
	}

	return resErr(c, errors.New("id not found"))
}
