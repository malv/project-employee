package pojo

type PojoGetEmployee struct {
	Id       string       `json:"id"`
	Nik      string       `json:"nik"`
	Name     string       `json:"name"`
	Unit     PojoUnit     `json:"unit"`
	Position PojoPosition `json:"position"`
}
