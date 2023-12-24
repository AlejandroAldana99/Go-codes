package models

// ErrorResponse :
type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// APIErrorResponse :
type APIErrorResponse struct {
	Errors []ErrorResponse `json:"errors"`
}
