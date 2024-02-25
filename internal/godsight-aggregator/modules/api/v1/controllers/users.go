package controllers

import (
	"fmt"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/modules/api/v1/helpers"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/modules/api/v1/requests"
	log "github.com/Blockchain-Framework/controller/pkg/logger"
	"github.com/Blockchain-Framework/controller/pkg/middleware"
	"github.com/Blockchain-Framework/controller/pkg/serviceresponse"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {

	log.Debug(r.Context()).Msg("Inside ListUsers")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	queryParams := r.URL.Query()
	responseBody, responseError := helpers.FetchListUsers(r.Context(), &queryParams)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed to fetch users: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully fetch users data"
	success.Data = responseBody.Data
	success.Meta = responseBody.Meta

	_ = success.WriteOk()

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Context()).Msg("Inside Create User")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	reqBody, err := requests.CreateRequestBodyForUser(r)

	if err != nil {
		log.Error(r.Context()).Msgf("failed to create request body for user creation: %s", err)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = "failed to create request body for user creation"
		serviceError.AddDetail(err)
		_ = serviceError.Write(http.StatusBadRequest, nil)
		panic(err)
	}

	queryParams := r.URL.Query()
	responseBody, responseError := helpers.CreateUser(r.Context(), &queryParams, reqBody)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed iam request for user creation: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully created new user"
	success.Data = responseBody.Data
	success.Meta = responseBody.Meta
	_ = success.Write(http.StatusCreated, nil)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Context()).Msg("Inside Update User")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	userId := r.Context().Value(middleware.RequestParams).(userId).Id

	reqBody, err := requests.CreateRequestBodyForUser(r)

	if err != nil {
		log.Error(r.Context()).Msgf("failed to create request body for user update: %s", err)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = "failed to create request body for user update"
		serviceError.AddDetail(err)
		_ = serviceError.Write(http.StatusBadRequest, nil)
		panic(err)
	}

	queryParams := r.URL.Query()
	responseError := helpers.UpdateUser(r.Context(), &queryParams, reqBody, userId)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed iam request for user update: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully updated the user"
	success.Data = nil
	success.Meta = nil

	_ = success.WriteOk()

}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Context()).Msg("Inside User Info")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	userId := r.Context().Value(middleware.RequestParams).(userId).Id

	queryParams := r.URL.Query()
	responseBody, responseError := helpers.FetchUserInfo(r.Context(), &queryParams, userId)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed to fetch user: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully fetch user info data"
	success.Data = responseBody.Data
	success.Meta = responseBody.Meta

	_ = success.WriteOk()
}

func ActivateUser(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Context()).Msg("Inside User Activation")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	userId := r.Context().Value(middleware.RequestParams).(userId).Id

	queryParams := r.URL.Query()
	responseError := helpers.ActivateUser(r.Context(), &queryParams, userId)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed to activate user: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	// NEED TO IMPLEMENT SENDING NOTIFICATION ALTER TO USER

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully activate user"
	success.Data = nil
	success.Meta = nil

	_ = success.WriteOk()
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Context()).Msg("Inside User Deletion")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	userId := r.Context().Value(middleware.RequestParams).(userId).Id

	queryParams := r.URL.Query()
	responseError := helpers.DeleteUser(r.Context(), &queryParams, userId)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed to delete user: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully delete user"
	success.Data = nil
	success.Meta = nil

	_ = success.WriteOk()
}
