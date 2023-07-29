package app

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/iamanders/tdr3/internal/data"
)

// GET /times?startDate=X&stopDate=X
func (app *App) handleTimeIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentDate := time.Now()
		startDate, _ := parseDateToTime(r.URL.Query().Get("startDate"), time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.UTC().Location()))
		stopDate, _ := parseDateToTime(r.URL.Query().Get("stopDate"), time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 23, 59, 59, 0, currentDate.UTC().Location()))

		times, err := data.TimeFindBetween(startDate, stopDate)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
		}

		app.respondJson(w, r, times, http.StatusOK)
	}
}

// POST /times
func (app *App) handleTimeCreate() http.HandlerFunc {

	type request struct {
		ClientString    string  `json:"client" validate:"required,min=1,max=100"`
		ProjectString   string  `json:"project" validate:"required,min=1,max=100"`
		StartedAtString string  `json:"startedAt" validate:"required"`        // TODO: Better validations!
		StoppedAtString *string `json:"stoppedAt" validate:"omitempty,min=5"` // TODO: Better validations!
		Code            *string `json:"code"`
		Comment         *string `json:"comment"`
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

		// Insert/get client
		client, err := data.ClientCreate(request.ClientString, nil)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		// Insert/get project
		project, err := data.ProjectCreate(request.ProjectString, nil)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		// Start/stop date
		startedAt, err := time.Parse("2006-01-02 15:04:05", request.StartedAtString)
		if err != nil {
			app.respondJson(w, r, err, http.StatusBadRequest)
			return
		}
		var stoppedAt *time.Time
		if request.StoppedAtString != nil {
			nullableStoppedAt, err := time.Parse("2006-01-02 15:04:05", *request.StoppedAtString)
			if err != nil {
				app.respondJson(w, r, err, http.StatusBadRequest)
				return
			}
			stoppedAt = &nullableStoppedAt
		}

		time, err := data.TimeCreate(client.ID, project.ID, startedAt, stoppedAt, request.Code, request.Comment)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		// Repsonse
		app.respondJson(w, r, time, http.StatusOK)
	}
}

// GET /times/:id
func (app *App) handleTimeGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time, err := data.TimeFind(chi.URLParam(r, "id"))
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		if time == nil {
			app.respondJson(w, r, err, http.StatusNotFound)
			return
		}

		app.respondJson(w, r, time, http.StatusOK)
	}
}

// PUT /times/:id
func (app *App) updateTime() http.HandlerFunc {

	type request struct {
		ClientString    string  `json:"client" validate:"required,min=1,max=100"`
		ProjectString   string  `json:"project" validate:"required,min=1,max=100"`
		StartedAtString string  `json:"startedAt" validate:"required"`  // TODO: Better validations!
		StoppedAtString *string `json:"stoppedAt" validate:"omitempty"` // TODO: Better validations!
		Code            *string `json:"code"`
		Comment         *string `json:"comment"`
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

		// Find time
		timeToUpdate, err := data.TimeFind(chi.URLParam(r, "id"))
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}
		if timeToUpdate == nil {
			app.respondJson(w, r, err, http.StatusNotFound)
			return
		}

		// Insert/get client
		client, err := data.ClientCreate(request.ClientString, nil)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		// Insert/get project
		project, err := data.ProjectCreate(request.ProjectString, nil)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		// Update
		startedAt, err := time.Parse("2006-01-02 15:04:05", request.StartedAtString)
		if err != nil {
			app.respondJson(w, r, err, http.StatusBadRequest)
			return
		}
		var stoppedAt *time.Time
		if request.StoppedAtString != nil {
			nullableStoppedAt, err := time.Parse("2006-01-02 15:04:05", *request.StoppedAtString)
			if err != nil {
				app.respondJson(w, r, err, http.StatusBadRequest)
				return
			}
			stoppedAt = &nullableStoppedAt
		}

		updatedTime, err := data.TimeUpdate(timeToUpdate.ID, client.ID, project.ID, startedAt, stoppedAt, request.Code, request.Comment)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		app.respondJson(w, r, updatedTime, http.StatusOK)
	}
}

// DELETE /times/:id
func (app *App) deleteTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time, err := data.TimeFind(chi.URLParam(r, "id"))
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}
		if time == nil {
			app.respondJson(w, r, "not found", http.StatusNotFound)
			return
		}

		err = data.TimeDelete(time.ID)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
		}

		app.respondJson(w, r, true, http.StatusOK)
	}
}

// GET /times/search?client=X&project=X&code=X&comment=X&startDate=X&stopDate=X
func (app *App) handleTimeSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		currentDate := time.Now()
		startDate, _ := parseDateToTime(r.URL.Query().Get("startDate"), time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.UTC().Location()))
		stopDate, _ := parseDateToTime(r.URL.Query().Get("stopDate"), time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 23, 59, 59, 0, currentDate.UTC().Location()))

		times, err := data.TimeSearch(
			r.URL.Query().Get("client"),
			r.URL.Query().Get("project"),
			r.URL.Query().Get("code"),
			r.URL.Query().Get("comment"),
			startDate,
			stopDate,
		)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
		}

		app.respondJson(w, r, times, http.StatusOK)
	}
}
