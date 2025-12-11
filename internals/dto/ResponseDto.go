package dto

type ResponseDto struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	ResultObj interface{} `json:"resultObj"`
}
