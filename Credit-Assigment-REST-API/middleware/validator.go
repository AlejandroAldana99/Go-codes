package middleware

import (
	"errors"
	"strings"

	er "github.com/AlejandroAldana99/Credit-Assigment-REST-API/errors"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/models"
	"github.com/go-playground/validator"

	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/libs/logger"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

// ParamsValidator :
func ValidatorParams(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {
		id := strings.TrimSpace(c.Param("id"))

		if id == "undefined" || id == "null" || id == "" {
			logger.Error("middleware", "ParamsValidator", "Instance param cannot be null")
			return er.HandleServiceError(errors.New("invalid parameters"))
		}

		e = next(c)
		return e
	}
}

func ValidateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {
		dto, err := bodyToStruct(c)
		if err != nil {
			logger.Error("middleware", "ValidateBody", err.Error())
			return er.HandleServiceError(errors.New("invalid body"))
		}
		errValidation := validateModel(dto)
		if errValidation != nil {
			return er.HandleServiceError(errValidation)
		}

		c.Set("dto", dto)
		return next(c)
	}
}

func bodyToStruct(c echo.Context) (models.AssigmentData, error) {
	dto := new(models.AssigmentData)
	err := c.Bind(dto)
	return *dto, err
}

func validateModel(dto models.AssigmentData) error {
	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return err
	}
	return nil
}
