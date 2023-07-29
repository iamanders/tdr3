package app

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/iamanders/tdr3/internal/data"
)

// GET /status
func (app *App) handleStatus() http.HandlerFunc {

	type response struct {
		HasRunningTime bool       `json:"hasRunningTime"`
		Time           *data.Time `json:"time"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var response response

		time, err := data.TimeFindRunning()
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
		}

		if time == nil {
			response.HasRunningTime = false
		} else {
			response.HasRunningTime = true
			response.Time = time
		}

		app.respondJson(w, r, response, http.StatusOK)
	}
}

// GET /status/stop
func (app *App) handleStatusStop() http.HandlerFunc {

	type request struct {
		StoppedAtString *string `json:"stoppedAt" validate:"omitempty,min=5"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		// Decode
		var request request
		err := app.decodeJson(w, r, &request)
		if err != nil {
			app.ErrorLog.Println(err)
		}

		// Validate
		validate := validator.New()
		err = validate.Struct(request)
		if err != nil {
			validationErrors := app.extractValidationErrors(err.(validator.ValidationErrors), request)
			app.respondJson(w, r, validationErrors, http.StatusUnprocessableEntity)
			return
		}

		stoppedAt, err := time.Parse("2006-01-02 15:04:05", *request.StoppedAtString)
		if err != nil {
			app.respondJson(w, r, err, http.StatusBadRequest)
			return
		}

		// Stop all running times
		err = data.TimeStopAll(stoppedAt)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		app.respondJson(w, r, "ok", http.StatusOK)
	}
}
