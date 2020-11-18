package controller

import (
	"errors"
	"project-employee/config"
	"project-employee/model"
	"project-employee/service"

	"github.com/labstack/echo"
)

var unitService service.UnitService = service.UnitServiceImpl{}

func SetUnitController(e *echo.Group) {
	e.POST("/units", addUnit)
	e.PUT("/units", editUnit)
	e.DELETE("/unit/:id", deleteUnit)
	e.GET("/units", getUnits)
	e.GET("/unit/:id", getUnitById)
}

func addUnit(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := &model.Unit{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := unitService.AddUnit(data)
	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func editUnit(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := &model.Unit{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := unitService.EditUnit(data)
	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func deleteUnit(c echo.Context) (e error) {
	defer config.CatchError(&e)
	id := c.Param("id")

	err := unitService.DeleteUnit(id)

	if err == nil {
		return resSuccess(c)
	}
	return resErr(c, err)
}

func getUnits(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := unitService.GetUnits()

	if err == nil {
		if len(result) > 0 {
			return res(c, result)
		}
		return resErr(c, errors.New("data not found"))
	}
	return resErr(c, err)
}

func getUnitById(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	result, err := unitService.GetUnitById(id)

	if err == nil {
		return res(c, result)
	}

	return resErr(c, errors.New("id not found"))
}
