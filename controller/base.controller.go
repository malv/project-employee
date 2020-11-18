package controller

import (
	"net/http"
	"project-employee/config"

	"github.com/labstack/echo"
)

func SetInit(e *echo.Group) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rest API started")
	})
}

func res(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func resErr(c echo.Context, err error) error {
	return c.String(http.StatusBadRequest, err.Error())
}

func catchError(e *error) {
	config.CatchError(e)
}

func resSuccess(c echo.Context) error {
	return c.String(http.StatusOK, "Success")
}

func resString(c echo.Context, msg string) error {
	return c.String(http.StatusOK, msg)
}
