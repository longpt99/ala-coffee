package validate

import (
	"ecommerce/errors"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	v = validator.New()
)

func RegisterValidation() error {
	err := v.RegisterValidation("date_string", func(fl validator.FieldLevel) bool {
		dob := fl.Field().String()

		if !fl.Field().IsZero() {
			// Regular expression pattern for "dd/mm/yyyy" format
			pattern := `^(0[1-9]|[1-2][0-9]|3[0-1])/(0[1-9]|1[0-2])/\d{4}$`
			match, _ := regexp.MatchString(pattern, dob)

			return match
		}

		return true
	})

	if err != nil {
		return err
	}

	return nil
}

func getFieldName(req interface{}, namespace string) string {
	t := reflect.TypeOf(req)
	nsSplit := strings.Split(namespace, ".")[1:]
	fieldNames := make([]string, 0)

	var embeddedType reflect.StructField

	for i, n := range nsSplit {
		var embeddedField reflect.StructField
		if i == 0 {
			embeddedField, _ = t.Elem().FieldByName(n)
			embeddedType = embeddedField
		} else {
			embeddedField, _ = embeddedType.Type.FieldByName(n)
		}

		jsonTag := embeddedField.Tag.Get("json")
		fieldNames = append(fieldNames, jsonTag)
	}

	return strings.Join(fieldNames, ".")
}

// Valid validates the given struct.
func Valid(dst interface{}) error {
	err := v.Struct(dst)
	if err == nil {
		return nil
	}

	userFacingErrors := make(errors.M)

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := getFieldName(dst, err.Namespace())

		switch err.Tag() {
		case "required":
			userFacingErrors[fieldName] = "This field is required."
		case "min":
			if err.Type().Kind() == reflect.String {
				userFacingErrors[fieldName] =
					fmt.Sprintf("This field must be at least %s characters long.", err.Param())
			} else {
				userFacingErrors[fieldName] =
					fmt.Sprintf("This value does not meet the minimum of %s.", err.Param())
			}
		case "max":
			if err.Type().Kind() == reflect.String {
				userFacingErrors[fieldName] =
					fmt.Sprintf("This field must be less than %s characters long.", err.Param())
			} else {
				userFacingErrors[fieldName] =
					fmt.Sprintf("This value exceeds the maximum of %s.", err.Param())
			}
		case "email":
			userFacingErrors[fieldName] = "This isn't a valid email."
		case "date_string":
			userFacingErrors[fieldName] = "Wrong format type (dd/mm/yyyy)."
		default:
			userFacingErrors[fieldName] = "Got some errors"
		}
	}

	return errors.E(errors.Op("payload.Valid"), err, userFacingErrors, http.StatusBadRequest)
}

// ReadValid is equivalent to calling Read followed by Valid.
func ReadValid(dst interface{}, r *http.Request) error {
	op := errors.Op("utils.validate.ReadValid")

	if err := Read(dst, r); err != nil {
		return errors.E(op, err)
	}

	if err := Valid(dst); err != nil {
		return errors.E(op, err)
	}

	return nil
}

// Read unmarshals the payload from the incoming request to the given struct pointer.
func Read(dst interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return errors.E(errors.Op("utils.validate.Read"), http.StatusBadRequest, err,
			map[string]string{"message": "Could not decode request body"})
	}

	return nil
}

func lowerFirstLetter(s string) string {
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}

	if s[len(s)-2:] == "ID" {
		s = s[:len(s)-2] + "Id"
	}

	return s
}
