package errors

import (
	"fmt"
	"net/http"

	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/models"
	"github.com/labstack/echo/v4"
)

const (
	DataSourceException = iota + 1
	InvalidParameters
	InvalidRole
	NonCancellable
	InvalidAssigment
	InvalidStatus
)

// ServiceErrors :
var ServiceErrors map[int]string = map[int]string{
	DataSourceException: "Data source exception",
	InvalidParameters:   "Invalid parameters",
	InvalidAssigment:    "Invalid Assigment",
}

// NewAPIErrorResponse :
func NewAPIErrorResponse(errors ...models.ErrorResponse) models.APIErrorResponse {
	return models.APIErrorResponse{
		Errors: errors,
	}
}

// MapErrorCode :
func MapErrorCode(code int) string {
	return ServiceErrors[code]
}

// ErrorCodeString :
func ErrorCodeString(code int) string {
	return fmt.Sprintf("CDS-%d", code)
}

func HandleServiceError(err error) error {
	var (
		status, code int
	)
	switch err.Error() {
	case "invalid parameters":
		status = http.StatusBadRequest
		code = InvalidParameters
		break
	case "invalid Assigment":
		status = http.StatusUnauthorized
		code = InvalidAssigment
		break
	default:
		status = http.StatusInternalServerError
		code = DataSourceException
	}
	return echo.NewHTTPError(
		status,
		NewAPIErrorResponse(
			models.ErrorResponse{
				Code:    ErrorCodeString(code),
				Message: MapErrorCode(code),
			},
		))
}
