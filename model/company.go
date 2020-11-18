package model

type Company struct {
	BaseModel
	Code             string `gorm:"column:code;unique" json:"code"`
	Name             string `gorm:"column:name" json:"name"`
	Description      string `gorm:"column:description" json:"description"`
	CompanyTaxNumber string `gorm:"column:company_tax_number" json:"companyTaxNumber"`
}

func (Company) TableName() string {
	return "cor_companies"
}
