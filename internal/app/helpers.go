package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// Write JSON from data
func (app *App) respondJson(w http.ResponseWriter, r *http.Request, data interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		app.ErrorLog.Println("cannot encode json", err)
		return err
	}

	return nil
}

// Decode JSON to struct
func (app *App) decodeJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 * 5 // 5MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

// Get validation error messages as a map from validator.ValidationErrors errors
func (app *App) extractValidationErrors(errors validator.ValidationErrors, model interface{}) map[string]string {
	validationErrorsMap := map[string]string{}

	for _, validationError := range errors {

		// Get JSON field from model struct
		val := reflect.ValueOf(model)
		field, _ := val.Type().FieldByName(validationError.Field())
		jsonField := field.Tag.Get("json")
		if jsonField == "" {
			jsonField = validationError.Field()
		}

		// Construct a nice error message
		message := fmt.Sprintf("Invalid value for %s (%s)", validationError.Field(), validationError.Tag())
		switch validationError.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", validationError.Field())
		case "gte":
		case "min":
			message = fmt.Sprintf("%s is too short/small", validationError.Field())
		case "lte":
		case "max":
			message = fmt.Sprintf("%s is too long/big", validationError.Field())
		}

		validationErrorsMap[jsonField] = message
	}

	return validationErrorsMap
}
