package main

import (
	"baseGo/internal/dto/payload"
	"net/http"
)

const (
	SuccessCode      = "00"
	BadRequestCode   = "01"
	NotFoundCode     = "02"
	UnauthorizedCode = "03"
	SystemErrorCode  = "99"

	SuccessMessage      = "Success"
	SystemErrorMessage  = "System error"
	UnauthorizedMessage = "Unauthorized"

	InvalidPayload = "Invalid payload"
)

func (app *application) logError(err error, resultCode string, message string) {
	app.logger.PrintError(err, payload.BaseResponse{
		ResultCode:    resultCode,
		ResultMessage: message,
		Data:          nil,
	})
}

func (app *application) errorResponse(writer http.ResponseWriter, status int, data payload.BaseResponse) {
	err := app.writeJson(writer, status, data, nil)
	if err != nil {
		app.logError(err, data.ResultCode, data.ResultMessage)
		writer.WriteHeader(status)
	}
}

func (app *application) badRequestErrorResponse(writer http.ResponseWriter, message string) {
	app.errorResponse(writer, http.StatusBadRequest, payload.BaseResponse{
		ResultCode:    BadRequestCode,
		ResultMessage: message,
		Data:          nil,
	})
}

func (app *application) unauthorizedErrorResponse(writer http.ResponseWriter, message string) {
	app.errorResponse(writer, http.StatusUnauthorized, payload.BaseResponse{
		ResultCode:    UnauthorizedCode,
		ResultMessage: message,
		Data:          nil,
	})
}

func (app *application) notFoundErrorResponse(writer http.ResponseWriter, message string) {
	app.errorResponse(writer, http.StatusNotFound, payload.BaseResponse{
		ResultCode:    NotFoundCode,
		ResultMessage: message,
		Data:          nil,
	})
}

func (app *application) serverErrorResponse(writer http.ResponseWriter, err error) {
	app.logError(err, SystemErrorCode, SystemErrorMessage)
	response := payload.BaseResponse{
		ResultCode:    SystemErrorCode,
		ResultMessage: err.Error(),
		Data:          nil,
	}

	app.errorResponse(writer, http.StatusInternalServerError, response)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, errors map[string]string) {
	app.errorResponse(w, http.StatusBadRequest, payload.BaseResponse{
		ResultCode:    BadRequestCode,
		ResultMessage: InvalidPayload,
		Data:          map[string]interface{}{"errors": errors},
	})
}
