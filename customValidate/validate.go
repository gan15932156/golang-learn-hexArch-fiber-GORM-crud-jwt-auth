package customvalidate

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)


func Validate(validate *validator.Validate,data interface{}) []string{

	err := validate.Struct(data)

	if err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			errorsMsg := []string{}
			for _, e := range validateErrs {
				errorMsg := transformErrorMessage(e)
				errorsMsg = append(errorsMsg,errorMsg)
			}
			return errorsMsg
		}

	}

	return nil
}

func transformErrorMessage(e validator.FieldError) string{
	switch e.Tag(){
		case  "required":
			return fmt.Sprintf("%s is required.",e.Field())
		case  "min":
			return fmt.Sprintf("%s must min %s character.",e.Field(),e.Param())
		case  "max":
			return fmt.Sprintf("%s must less than %s character.",e.Field(),e.Param())
		case  "email":
			return fmt.Sprintf("%s must valid email format.",e.Field())
		default:
			return fmt.Sprintf("Not found error message for field %s",e.Field())
	}
}