package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/wanchai23chai/go-shared-lib/pkg/errmsg"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateBodyParser(ctx *fiber.Ctx, out interface{}) error {
	body := ctx.Body()
	if len(body) == 0 {
		return nil
	}

	isBodyError := ctx.BodyParser(out)
	if isBodyError != nil {
		if isBodyError.Error() == fiber.ErrUnprocessableEntity.Message {
			return errmsg.ErrInvalidContentType
		}
		err := json.Unmarshal([]byte(ctx.Body()), &out)
		if err != nil {
			typeErr := err.(*json.UnmarshalTypeError)
			errMsg := fmt.Sprintf("Invalid field (%s), it should be %s.", strings.ToLower(typeErr.Field), typeErr.Type)
			return errors.New(errMsg)
		}
	} else {
		validate := validator.New()
		errorstruct := validate.Struct(out)
		if errorstruct != nil {
			for _, err_v := range errorstruct.(validator.ValidationErrors) {
				param := strings.TrimSpace(fmt.Sprintf("%v %s", err_v.Tag(), err_v.Param()))
				errMsg := fmt.Sprintf("Invalid field (%s), it should be %v.", strings.ToLower(err_v.Field()), param)
				return errors.New(errMsg)
			}
		}
	}
	return nil
}
