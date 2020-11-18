package model

type Employee struct {
	BaseModel
	CompanyId       string   `gorm:"column:company_id" json:"companyId"`
	Company         Company  `gorm:"foreignKey:CompanyId"`
	PersonId        string   `gorm:"column:person_id" json:"personId"`
	UnitId          string   `gorm:"column:unit_id" json:unitId`
	Unit            Unit     `gorm:"foreignKey:UnitId"`
	PositionId      string   `gorm:"column:position_id" json:positionId`
	Position        Position `gorm:"foreignKey:PositionId"`
	Nik             string   `gorm:"column:nik;unique" json:"nik"`
	HiredDate       string   `gorm:"column:hired_date" json:"hiredDate"`
	TerminationDate string   `gorm:"column:termination_date" json:"terminationDate"`
	IsActive        bool     `gorm:"column:is_active" json:"isActive"`
}

func (Employee) TableName() string {
	return "cor_employees"
}
