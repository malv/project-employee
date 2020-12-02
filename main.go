package main

import (
	"project-employee/config"
	"project-employee/controller"
	"project-employee/dao"
	"project-employee/service"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func main() {
	echo := echo.New()
	g := initDb()
	e := echo.Group("/api")

	dao.SetDao(g)
	service.SetService(g)
	config.Connected()
	config.ConnectedEmployee()
	e.Use(service.MiddlewareCredential)

	controller.SetInit(e)
	controller.SetUnitController(e)
	controller.SetPositionController(e)
	controller.SetCompanyController(e)
	controller.SetEmployeeController(e)
	echo.Logger.Fatal(echo.Start(":1234"))
}

func initDb() *gorm.DB {
	g, err := config.Conn()
	if err == nil {
		config.MigrateSchema(g)
		return g
	}
	panic(err)
}
