package validate

import (
	"github.com/labstack/echo"
	"github.com/go-playground/validator/v10"
)

// CustomValidator
type CustomValidator struct {
    validator *validator.Validate
}

// NewValidator
func NewValidator() echo.Validator {
    return &CustomValidator{validator: validator.New()}
}

// Validate validate
func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}