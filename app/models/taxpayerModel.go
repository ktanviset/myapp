package models

type Taxpayer struct {
	NID               string `json:"NID"`
	BranchNumber      string `json:"BranchNumber"`
	BranchTitle       string `json:"BranchTitle"`
	BranchName        string `json:"BranchName"`
	BuildingName      string `json:"BuildingName"`
	RoomNumber        string `json:"RoomNumber"`
	FloorNumber       string `json:"FloorNumber"`
	VillageName       string `json:"VillageName"`
	HouseNumber       string `json:"HouseNumber"`
	MooNumber         string `json:"MooNumber"`
	SoiName           string `json:"SoiName"`
	StreetName        string `json:"StreetName"`
	ThumbolName       string `json:"ThumbolName"`
	AmphurName        string `json:"AmphurName"`
	ProvinceName      string `json:"ProvinceName"`
	PostCode          string `json:"PostCode"`
	BusinessFirstDate string `json:"BusinessFirstDate"`
	BranchTypeCode    string `json:"BranchTypeCode"`
	BranchTypeName    string `json:"BranchTypeName"`
	RegisteredCapital string `json:"RegisteredCapital"`
	Status            string `json:"Status"`
}
