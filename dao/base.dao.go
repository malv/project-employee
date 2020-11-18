package dao

import (
	"project-employee/config"

	"gorm.io/gorm"
)

var g *gorm.DB

func SetDao(gDB *gorm.DB) {
	g = gDB
}
func catchError(e *error) {
	config.CatchError(e)
}
