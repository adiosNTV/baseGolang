package main

import (
	"baseGo/internal/dto/payload"
	"net/http"
)

func (app *application) routes() http.Handler {
	routes := http.NewServeMux()

	// Root routes
	routes.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		response := payload.BaseResponse{
			ResultCode:    SuccessCode,
			ResultMessage: "Hello world from tech-store 2",
			Data:          nil,
		}

		err := app.writeJson(writer, http.StatusOK, response, nil)
		if err != nil {
			app.serverErrorResponse(writer, err)
		}
	})

	routes.HandleFunc("GET /health-check", func(writer http.ResponseWriter, request *http.Request) {
		err := app.db.Ping()

		response := payload.BaseResponse{
			ResultCode:    SuccessCode,
			ResultMessage: SuccessMessage,
			Data:          nil,
		}

		if err != nil {
			response.ResultCode = SystemErrorCode
			response.ResultMessage = SystemErrorMessage
			app.serverErrorResponse(writer, err)
			return
		}

		err = app.writeJson(writer, 200, response, nil)
		if err != nil {
			app.serverErrorResponse(writer, err)
		}
	})

	//// Product routes
	//routes.HandleFunc("GET /products/{productId}", app.getProductDetailHandler)
	//routes.HandleFunc("GET /products", app.getProductsPaginatedHandler)
	//
	//// Order routes
	//routes.HandleFunc("GET /orders", app.getOrdersPaginated)
	//routes.HandleFunc("GET /orders/{orderId}", app.getOrderById)
	//routes.HandleFunc("POST /orders", app.createOrderHandler)
	//
	//// User routes
	//routes.HandleFunc("POST /register", app.registerUserHandler)
	//routes.HandleFunc("POST /login", app.loginUserHandler)
	//routes.HandleFunc("GET /logout/{userId}", app.logoutUserHandler)
	//routes.HandleFunc("GET /users/{userId}", app.getUserDetailsHandler)
	//routes.HandleFunc("PUT /users/{userId}", app.updateUserDetailsHandler)
	//
	//// Label routes
	//routes.HandleFunc("GET /labels", app.getAllLabelsPaginated)

	// Return routes
	return app.recoverPanic(routes)
}
