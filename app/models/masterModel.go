package models

type Master struct {
	ID      int    `json:"id"`
	Val     string `json:"val"`
	Name    string `json:"name"`
	Display string `json:"display"`
}

type ListMasters struct {
	Masters []*Master `json:"Masters"`
}
