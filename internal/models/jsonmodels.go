package models

type JSONResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
