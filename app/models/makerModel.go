package models

type Maker struct {
	ID        int     `json:"id"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	NameTh    string  `json:"nameTh"`
	NameEn    string  `json:"nameEn"`
}

type ListMakers struct {
	Makers []*Maker `json:"Makers"`
}
