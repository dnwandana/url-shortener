package model

// constant variable for http status message
const (
	StatusOK                  = "OK"
	StatusCreated             = "CREATED"
	StatusBadRequest          = "BAD_REQUEST"
	StatusInternalServerError = "INTERNAL_SERVER_ERROR"
)

// ResponseMessage struct that returns a success or error message.
type ResponseMessage struct {
	Code    int32  `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ResponseData struct represents JSON response after shortening the URL.
type ResponseData struct {
	Code   int32       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
