package middleware

import (
	"errors"
	"strings"

	"github.com/AlejandroAldana99/mvp_api/constants"
	er "github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/libs"
	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/go-playground/validator"

	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get(constants.GenericHeader)
		if len(header) == 0 {
			c.Set("role", constants.GenericName)
			c.Set("userid", "")
		} else {
			jwt := strings.Split(header, " ")[1]
			role, id, _ := libs.DecodeJWT(jwt)
			c.Set("role", role)
			c.Set("userid", id)
		}
		return next(c)
	}
}

// ParamsValidator :
func ParamsValidatorID(next echo.HandlerFunc) echo.HandlerFunc {
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

func ParamsOrderValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {
		id := strings.ToLower(c.QueryParam("orderid"))
		sts := strings.ToLower(c.QueryParam("status"))

		if id == "undefined" || id == "null" || id == "" || sts == "undefined" || sts == "null" || sts == "" {
			logger.Error("middleware", "ParamsValidator", "Instance param cannot be null")
			return er.HandleServiceError(errors.New("invalid parameters"))
		}

		e = next(c)
		return e
	}
}

func ParamsLoginValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {
		us := strings.ToLower(c.QueryParam("user"))
		pass := strings.ToLower(c.QueryParam("password"))

		if us == "undefined" || us == "null" || us == "" || pass == "undefined" || pass == "null" || pass == "" {
			logger.Error("middleware", "ParamsValidator", "Instance param cannot be null")
			return er.HandleServiceError(errors.New("invalid parameters"))
		}

		e = next(c)
		return e
	}

}

//ValidateInput :
func ValidateOrderBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {
		dto, err := orderBodyToStruct(c)
		if err != nil {
			logger.Error("middleware", "ValidateBody", err.Error())
			return er.HandleServiceError(errors.New("invalid body"))
		}
		errValidation := validateOrderModel(dto)
		if errValidation != nil {
			return er.HandleServiceError(errValidation)
		}

		c.Set("dto", dto)
		return next(c)
	}
}

func ValidateUserBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {
		dto, err := userBodyToStruct(c)
		if err != nil {
			logger.Error("middleware", "ValidateBody", err.Error())
			return er.HandleServiceError(errors.New("invalid body"))
		}
		errValidation := validateUserModel(dto)
		if errValidation != nil {
			return er.HandleServiceError(errValidation)
		}

		c.Set("dto", dto)
		return next(c)
	}
}

//SendgridPayloadToStruct :
func orderBodyToStruct(c echo.Context) (models.OrderData, error) {
	dto := new(models.OrderData)
	err := c.Bind(dto)
	return *dto, err
}

//ValidateInputModelSendgrid :
func validateOrderModel(dto models.OrderData) error {
	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return err
	}
	return nil
}

func userBodyToStruct(c echo.Context) (models.UserData, error) {
	dto := new(models.UserData)
	err := c.Bind(dto)
	return *dto, err
}

//ValidateInputModelSendgrid :
func validateUserModel(dto models.UserData) error {
	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return err
	}
	return nil
}
