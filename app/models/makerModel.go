package models

type Maker struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
}

type ListMakers struct {
	Makers []*Maker `json:"Makers"`
}
