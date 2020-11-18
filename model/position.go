package model

type Position struct {
	BaseModel
	Code        string `gorm:"column:code;unique" json:"code"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Level       int    `gorm:"column:level" json:"level"`
}

func (Position) TableName() string {
	return "cor_positions"
}
