package models

type Maker struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Name      string `json:"name"`
}

type ListMakers struct {
	Makers []*Maker `json:"Makers"`
}
