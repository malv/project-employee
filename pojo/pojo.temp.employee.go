package pojo

type PojoTempEmployee struct {
	Id       string       `json:id`
	Nik      string       `json:nik`
	PersonId string       `json:personId`
	Unit     PojoUnit     `json:unit`
	Position PojoPosition `json:position`
}
