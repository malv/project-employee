package controller

import (
	"errors"
	"project-employee/config"
	"project-employee/model"
	"project-employee/service"

	"github.com/labstack/echo"
)

var positionService service.PositionService = service.PositionServiceImpl{}

func SetPositionController(e *echo.Group) {
	e.POST("/positions", addPosition)
	e.PUT("/positions", editPosition)
	e.DELETE("/position/:id", deletePosition)
	e.GET("/positions", getPositions)
	e.GET("/position/:id", getPositionById)
}

func addPosition(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := &model.Position{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := positionService.AddPosition(data)
	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func editPosition(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := &model.Position{}

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := positionService.EditPosition(data)
	if err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}

func deletePosition(c echo.Context) (e error) {
	defer config.CatchError(&e)
	id := c.Param("id")

	err := positionService.DeletePosition(id)

	if err == nil {
		return resSuccess(c)
	}
	return resErr(c, err)
}

func getPositions(c echo.Context) (e error) {
	defer config.CatchError(&e)

	result, err := positionService.GetPositions()

	if err == nil {
		if len(result) > 0 {
			return res(c, result)
		}
		return resErr(c, errors.New("data not found"))
	}
	return resErr(c, err)
}

func getPositionById(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")

	result, err := positionService.GetPositionById(id)

	if err == nil {
		return res(c, result)
	}

	return resErr(c, errors.New("id not found"))
}
