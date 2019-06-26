package models

type Maker struct {
	ID            int     `json:"id"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
	NameTh        string  `json:"nameTh"`
	NameEn        string  `json:"nameEn"`
	LoCode        string  `json:"loCode"`
	LoCodeCountry string  `json:"LoCodeCountry"`
	FullCountry   string  `json:"FullCountry"`
	TruckAmount   int     `json:"TruckAmount"`
	Func1         string  `json:"Func1"`
	Func2         string  `json:"Func2"`
	Func3         string  `json:"Func3"`
	Func4         string  `json:"Func4"`
	Func5         string  `json:"Func5"`
	Func6         string  `json:"Func6"`
	Func7         string  `json:"Func7"`
	Func8         string  `json:"Func8"`
	TruckType     string  `json:"TruckType"`
}

type ListMakers struct {
	Makers []*Maker `json:"Makers"`
}
