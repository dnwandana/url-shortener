package model

// SuccessResponse is a JSON struct representing a successful request
type SuccessResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

// ErrorResponse is a JSON struct representing a failed request
type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Error      interface{} `json:"data"`
}
