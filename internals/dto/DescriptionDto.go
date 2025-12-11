package dto

type DescriptionReqDto struct {
	ImageId string `json:"imageId"`
	UserId  string `json:"userId"`

	Description string `json:"description"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
}
