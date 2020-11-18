package model

type Unit struct {
	BaseModel
	Code string `gorm:"column:code;unique" json:"code"`
	Name string `gorm:"column:name" json:"name"`
}

func (Unit) TableName() string {
	return "cor_units"
}
