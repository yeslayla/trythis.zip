package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yeslayla/trythis.zip/api/api"
)

type App struct{}

// GetRouter creates a new mux router
func (app *App) GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.HandleError(w, r, &api.MethodNotAllowedError{
			Err: errors.New("method not allowed"),
		})
	})
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.HandleError(w, r, &api.PageNotFoundError{
			Err: errors.New("page not found"),
		})
	})

	api.HandlerWithOptions(app, api.GorillaServerOptions{
		BaseRouter:       router,
		ErrorHandlerFunc: app.HandleError,
	})
	return router
}

func (app *App) HandleError(w http.ResponseWriter, r *http.Request, err error) {

	errorCode := http.StatusInternalServerError
	if _, ok := err.(*api.InvalidParamFormatError); ok {
		errorCode = http.StatusBadRequest
	} else if _, ok := err.(*api.TooManyValuesForParamError); ok {
		errorCode = http.StatusBadRequest
	} else if _, ok := err.(*api.RequiredHeaderError); ok {
		errorCode = http.StatusBadRequest
	} else if _, ok := err.(*api.RequiredParamError); ok {
		errorCode = http.StatusBadRequest
	} else if _, ok := err.(*api.UnmarshallingParamError); ok {
		errorCode = http.StatusBadRequest
	} else if _, ok := err.(*api.UnescapedCookieParamError); ok {
		errorCode = http.StatusBadRequest
	} else if _, ok := err.(*api.PageNotFoundError); ok {
		errorCode = http.StatusNotFound
	} else if _, ok := err.(*api.MethodNotAllowedError); ok {
		errorCode = http.StatusMethodNotAllowed
	}

	response := api.ErrorResponse{
		Code:    int32(errorCode),
		Message: err.Error(),
	}

	w.WriteHeader(errorCode)
	data, _ := json.Marshal(response)
	_, _ = w.Write(data)
}

// GetStatus returns 200 when operational
func (app *App) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	response := api.ErrorResponse{
		Code:    http.StatusOK,
		Message: "OK",
	}

	data, _ := json.Marshal(response)
	_, _ = w.Write(data)
}
