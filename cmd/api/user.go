package main

//
//import (
//	"TechStore/cache"
//	db "TechStore/db/sqlc"
//	"TechStore/internal/dto/payload"
//	"TechStore/utils"
//	"crypto/rand"
//	"encoding/base32"
//	"errors"
//	"github.com/alexedwards/argon2id"
//	"github.com/google/uuid"
//	"net/http"
//	"time"
//)
//
//func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
//
//	var request payload.RegisterUserRequest
//	err := app.readJson(w, r, &request)
//	if err != nil {
//		app.badRequestErrorResponse(w, err.Error())
//		return
//	}
//
//	checkEmail := utils.IsValidEmail(request.Data.Email)
//	if !checkEmail {
//		app.badRequestErrorResponse(w, "Email is invalid !")
//		return
//	}
//
//	emailExistence, err := app.queries.CheckEmailExistence(r.Context(), request.Data.Email)
//	if err != nil {
//		app.serverErrorResponse(w, err)
//		return
//	}
//	if emailExistence > 0 {
//		app.badRequestErrorResponse(w, "This email is already registered !")
//		return
//	}
//
//	usernameExistence, err := app.queries.CheckUsernameExistence(r.Context(), request.Data.Username)
//	if err != nil {
//		app.serverErrorResponse(w, err)
//		return
//	}
//	if usernameExistence > 0 {
//		app.badRequestErrorResponse(w, "This username is already registered !")
//		return
//	}
//
//	hashPassword, err := argon2id.CreateHash(request.Data.Password, argon2id.DefaultParams)
//	if err != nil {
//		app.serverErrorResponse(w, err)
//		return
//	}
//
//	err = app.queries.RegisterUser(r.Context(), db.RegisterUserParams{
//		ID:       uuid.NewString(),
//		Username: request.Data.Username,
//		Password: hashPassword,
//		Email:    request.Data.Email,
//	})
//	if err != nil {
//		app.badRequestErrorResponse(w, err.Error())
//		return
//	}
//
//	err = app.writeJson(w, http.StatusOK, payload.BaseResponse{
//		ResultCode:    SuccessCode,
//		ResultMessage: SuccessMessage,
//		Data:          nil,
//	}, nil)
//
//	if err != nil {
//		app.serverErrorResponse(w, err)
//	}
//}
//
//func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
//
//	var request payload.LoginUserRequest
//	err := app.readJson(w, r, &request)
//	if err != nil {
//		app.badRequestErrorResponse(w, err.Error())
//		return
//	}
//
//	user, err := app.queries.LoginUser(r.Context(), request.Data.Username)
//	if err != nil {
//		app.badRequestErrorResponse(w, "Wrong username or password !")
//		return
//	}
//
//	match, err := argon2id.ComparePasswordAndHash(request.Data.Password, user.Password)
//	if err != nil {
//		app.serverErrorResponse(w, err)
//		return
//	}
//	if !match {
//		app.badRequestErrorResponse(w, "Wrong username or password !")
//		return
//	}
//
//	userToken := RandomStringGenerator()
//	if userToken == "" {
//		app.serverErrorResponse(w, errors.New("System Error ! Please try again !"))
//		return
//	}
//
//	cache.Store.Set(userToken, user.ID, 7*24*time.Hour)
//
//	err = app.writeJson(w, http.StatusOK, payload.BaseResponse{
//		ResultCode:    SuccessCode,
//		ResultMessage: SuccessMessage,
//		Data: payload.LoginUserResponse{
//			UserId:     user.ID,
//			Token:      userToken,
//			ExpiryDate: time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339),
//		},
//	}, nil)
//
//	if err != nil {
//		app.serverErrorResponse(w, err)
//	}
//}
//
//func RandomStringGenerator() string {
//	randomBytes := make([]byte, 32)
//	_, err := rand.Read(randomBytes)
//	if err != nil {
//		return ""
//	}
//	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
//}
//
//func (app *application) logoutUserHandler(w http.ResponseWriter, r *http.Request) {
//
//	userId, err := app.readIDParam(r, "userId")
//	if err != nil {
//		app.badRequestErrorResponse(w, "Can not get userId from params !")
//		return
//	}
//
//	token := r.Header.Get("Authorization")
//	if len(token) == 0 {
//		app.badRequestErrorResponse(w, "Can not get token from header !")
//		return
//	}
//
//	inMemoryUserId, ok := cache.Store.Get(token)
//	if !ok {
//		app.badRequestErrorResponse(w, "Invalid token !")
//		return
//	}
//
//	if inMemoryUserId != userId.String() {
//		app.badRequestErrorResponse(w, "Invalid token !")
//		return
//	}
//
//	cache.Store.Delete(token)
//
//	err = app.writeJson(w, http.StatusOK, payload.BaseResponse{
//		ResultCode:    SuccessCode,
//		ResultMessage: SuccessMessage,
//		Data:          nil,
//	}, nil)
//
//	if err != nil {
//		app.serverErrorResponse(w, err)
//	}
//}
//
//func (app *application) getUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
//
//	userId, err := app.readIDParam(r, "userId")
//	if err != nil {
//		app.badRequestErrorResponse(w, "Can not get userId from params !")
//		return
//	}
//
//	token := r.Header.Get("Authorization")
//	if len(token) == 0 {
//		app.badRequestErrorResponse(w, "Can not get token from header !")
//		return
//	}
//
//	inMemoryUserId, ok := cache.Store.Get(token)
//	if !ok {
//		app.badRequestErrorResponse(w, "Invalid token !")
//		return
//	}
//
//	if inMemoryUserId != userId.String() {
//		app.badRequestErrorResponse(w, "Invalid token !")
//		return
//	}
//
//	user, err := app.queries.GetUserDetails(r.Context(), userId.String())
//	if err != nil {
//		if user.ID == "" || user.Username == "" || user.Email == "" {
//			app.notFoundErrorResponse(w, "User does not exist !")
//			return
//		}
//	}
//
//	err = app.writeJson(w, http.StatusOK, payload.BaseResponse{
//		ResultCode:    SuccessCode,
//		ResultMessage: SuccessMessage,
//		Data: payload.GetUserDetailsResponse{
//			Id:       user.ID,
//			Username: user.Username,
//			Email:    user.Email,
//			Phone:    user.Phone,
//			Address:  user.Address,
//		},
//	}, nil)
//
//	if err != nil {
//		app.serverErrorResponse(w, err)
//	}
//}
//
//func (app *application) updateUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
//
//	userId, err := app.readIDParam(r, "userId")
//	if err != nil {
//		app.badRequestErrorResponse(w, "Can not get userId from params !")
//		return
//	}
//
//	token := r.Header.Get("Authorization")
//	if len(token) == 0 {
//		app.badRequestErrorResponse(w, "Can not get token from header !")
//		return
//	}
//
//	inMemoryUserId, ok := cache.Store.Get(token)
//	if !ok {
//		app.badRequestErrorResponse(w, "Invalid token !")
//		return
//	}
//
//	if inMemoryUserId != userId.String() {
//		app.badRequestErrorResponse(w, "Invalid token !")
//		return
//	}
//
//	var request payload.UpdateUserDetailsRequest
//	err = app.readJson(w, r, &request)
//	if err != nil {
//		app.badRequestErrorResponse(w, err.Error())
//		return
//	}
//
//	user, err := app.queries.GetUserDetails(r.Context(), userId.String())
//	if err != nil {
//		if user.ID == "" || user.Username == "" || user.Email == "" {
//			app.notFoundErrorResponse(w, "User does not exist !")
//			return
//		}
//	}
//
//	var model db.UpdateUserDetailsParams
//	model.ID = userId.String()
//
//	if request.Data.Email != "" {
//		checkEmail := utils.IsValidEmail(request.Data.Email)
//		if !checkEmail {
//			app.badRequestErrorResponse(w, "Email is invalid !")
//			return
//		}
//		emailExistence, err := app.queries.CheckEmailExistence(r.Context(), request.Data.Email)
//		if err != nil {
//			app.serverErrorResponse(w, err)
//			return
//		}
//		if emailExistence > 0 {
//			app.badRequestErrorResponse(w, "This email is already registered !")
//			return
//		}
//		model.Email = request.Data.Email
//	} else {
//		model.Email = user.Email
//	}
//
//	if request.Data.Password != "" {
//		hashPassword, err := argon2id.CreateHash(request.Data.Password, argon2id.DefaultParams)
//		if err != nil {
//			app.serverErrorResponse(w, err)
//			return
//		}
//		model.Password = hashPassword
//	} else {
//		model.Password = user.Password
//	}
//
//	if request.Data.Phone != "" {
//		checkPhone := utils.IsValidPhoneNumber(request.Data.Phone)
//		if !checkPhone {
//			app.badRequestErrorResponse(w, "Phone number is invalid !")
//			return
//		}
//		model.Phone = request.Data.Phone
//	} else {
//		model.Phone = user.Phone
//	}
//
//	if request.Data.Address != "" {
//		model.Address = request.Data.Address
//	} else {
//		model.Address = user.Address
//	}
//
//	err = app.queries.UpdateUserDetails(r.Context(), model)
//
//	err = app.writeJson(w, http.StatusOK, payload.BaseResponse{
//		ResultCode:    SuccessCode,
//		ResultMessage: SuccessMessage,
//		Data:          nil,
//	}, nil)
//
//	if err != nil {
//		app.serverErrorResponse(w, err)
//	}
//}
