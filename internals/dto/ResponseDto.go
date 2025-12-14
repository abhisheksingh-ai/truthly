package dto

type ResponseDto[T any] struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ResultObj T      `json:"resultObj"`
	Error     string `json:"error"`
}
