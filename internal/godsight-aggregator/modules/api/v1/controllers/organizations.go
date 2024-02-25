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

func ListOrganizations(w http.ResponseWriter, r *http.Request) {

	log.Debug(r.Context()).Msg("Inside ListOrganizations")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	queryParams := r.URL.Query()
	responseBody, responseError := helpers.FetchListOrganization(r.Context(), &queryParams)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed to fetch organizations: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully fetch organizations data"
	success.Data = responseBody.Data
	success.Meta = responseBody.Meta

	_ = success.WriteOk()

}

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Context()).Msg("Inside Create Organization")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	reqBody, err := requests.CreateRequestBodyForOrganizationCreation(r)

	if err != nil {
		log.Error(r.Context()).Msgf("failed to create request body for organization creation: %s", err)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = "failed to create request body for organization creation"
		serviceError.AddDetail(err)
		_ = serviceError.Write(http.StatusBadRequest, nil)
		panic(err)
	}

	queryParams := r.URL.Query()
	responseBody, responseError := helpers.CreateOrganization(r.Context(), &queryParams, reqBody)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed iam request for organization creation: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully created new organization"
	success.Data = responseBody.Data
	success.Meta = responseBody.Meta
	_ = success.Write(http.StatusCreated, nil)

}

func UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Context()).Msg("Inside Update Organization")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	organizationId := r.Context().Value(middleware.RequestParams).(organizationId).Id

	reqBody, err := requests.CreateRequestBodyForOrganizationUpdate(r)

	if err != nil {
		log.Error(r.Context()).Msgf("failed to create request body for organization update: %s", err)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = "failed to create request body for organization update"
		serviceError.AddDetail(err)
		_ = serviceError.Write(http.StatusBadRequest, nil)
		panic(err)
	}

	queryParams := r.URL.Query()
	responseError := helpers.UpdateOrganization(r.Context(), &queryParams, reqBody, organizationId)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed iam request for organization update: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully updated the organization"
	success.Data = nil
	success.Meta = nil

	_ = success.WriteOk()

}

func OrganizationInfo(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Context()).Msg("Inside Organization Info")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	organizationId := r.Context().Value(middleware.RequestParams).(organizationId).Id

	queryParams := r.URL.Query()
	responseBody, responseError := helpers.FetchOrganizationInfo(r.Context(), &queryParams, organizationId)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed to fetch organization: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully fetch organization info data"
	success.Data = responseBody.Data
	success.Meta = responseBody.Meta

	_ = success.WriteOk()
}

func ActivateOrganization(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Context()).Msg("Inside Organization Activation")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	organizationId := r.Context().Value(middleware.RequestParams).(organizationId).Id

	queryParams := r.URL.Query()
	responseError := helpers.ActivateOrganization(r.Context(), &queryParams, organizationId)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed to fetch organization: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully activate organization"
	success.Data = nil
	success.Meta = nil

	_ = success.WriteOk()
}

func DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.Context()).Msg("Inside Organization Deletion")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	organizationId := r.Context().Value(middleware.RequestParams).(organizationId).Id

	queryParams := r.URL.Query()
	responseError := helpers.DeleteOrganization(r.Context(), &queryParams, organizationId)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed to delete organization: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully delete organization"
	success.Data = nil
	success.Meta = nil

	_ = success.WriteOk()
}
