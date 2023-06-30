package models


// ResponseModel struct to describe the BaseResponse model
type ResponseModel[T any] struct {
	Error      bool   `json:"error"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Payload    T    `json:"payload"`
}
