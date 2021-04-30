package httpmodels

import (
	"github.com/brianfromlife/golang-ecs/pkg/errors"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ValidateRegisterRequest(c echo.Context) (*RegisterRequest, *errors.ApiError) {

	registerRequest := new(RegisterRequest)
	if err := c.Bind(registerRequest); err != nil {
		return nil, errors.BindError()
	}

	var validationErrors []string

	if len(registerRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 characters")
	}

	if len(registerRequest.Username) < 3 {
		validationErrors = append(validationErrors, "Username must be longer than 2 characters")
	}

	if len(validationErrors) > 0 {
		return nil, errors.ValidationError(validationErrors)
	}

	return registerRequest, nil
}

func ValidateLoginRequest(c echo.Context) (*LoginRequest, *errors.ApiError) {
	loginRequest := new(LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return nil, errors.BindError()
	}

	var validationErrors []string

	if len(loginRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 characters")
	}

	if len(loginRequest.Username) < 3 {
		validationErrors = append(validationErrors, "Username must be longer than 2 characters")
	}

	if len(validationErrors) > 0 {
		return nil, errors.ValidationError(validationErrors)
	}

	return loginRequest, nil
}
