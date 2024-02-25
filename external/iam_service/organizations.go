package iam_service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/config"
	"github.com/Blockchain-Framework/controller/pkg/httpclient"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/url"
)

func RequestOrganizationList(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values) (IAMResponseOrganizationList, IAMResponseError) {

	//Generating Request URL
	u, err := url.Parse(config.Conf.IAMService.Url)
	if err != nil {

		return IAMResponseOrganizationList{Action: "Url parse is failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}
	requestUrl := u.ResolveReference(iamServiceUrlPath)

	client := httpclient.NewHttpClient(ctx, config.Conf)

	resp, err := client.DoGet(requestUrl.String(), q)
	if err != nil {

		return IAMResponseOrganizationList{Action: "IAM service request error"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return IAMResponseOrganizationList{Action: "querying Data failed in IAM service"}, IAMResponseError{Status: http.StatusInternalServerError, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponseOrganizationList
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponseOrganizationList{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	validate := validator.New()

	if err != nil {
		for _, org := range response.Data.Payload.Organizations {
			err := validate.Struct(org)
			if err != nil {
				return IAMResponseOrganizationList{Action: "IAM service response validation failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
			}
		}
	}
	// for testing use this
	//var response IAMResponseOrganizationList
	//response = GetOrganizationListData()

	return response, IAMResponseError{}
}

func RequestOrganizationCreate(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values, requestBodyJson []byte) (IAMResponseOrganizationCreate, IAMResponseError) {

	//Generating Request URL
	u, err := url.Parse(config.Conf.IAMService.Url)
	if err != nil {

		return IAMResponseOrganizationCreate{Action: "IAM service request url creation failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}
	requestUrl := u.ResolveReference(iamServiceUrlPath)

	client := httpclient.NewHttpClient(ctx, config.Conf)

	resp, err := client.DoPost(requestUrl.String(), q, bytes.NewBuffer(requestBodyJson))
	if err != nil {

		return IAMResponseOrganizationCreate{Action: "IAM service request error"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	if resp.StatusCode != http.StatusCreated {
		return IAMResponseOrganizationCreate{Action: "inserting Data failed in IAM service"}, IAMResponseError{Status: http.StatusInternalServerError, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponseOrganizationCreate
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponseOrganizationCreate{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	validate := validator.New()

	err = validate.Struct(response.Data)
	if err != nil {
		return IAMResponseOrganizationCreate{Action: "IAM service response validation failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	// for testing use this
	//var response IAMResponse
	//response = GetOrganizationCreationData(0)

	return response, IAMResponseError{}
}

func RequestOrganizationUpdate(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values, requestBodyJson []byte) (IAMResponse, IAMResponseError) {

	//Generating Request URL
	u, err := url.Parse(config.Conf.IAMService.Url)
	if err != nil {

		return IAMResponse{Action: "IAM service request url creation failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}
	requestUrl := u.ResolveReference(iamServiceUrlPath)

	client := httpclient.NewHttpClient(ctx, config.Conf)

	resp, err := client.DoPut(requestUrl.String(), q, bytes.NewBuffer(requestBodyJson))
	if err != nil {

		return IAMResponse{Action: "IAM service request error"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return IAMResponse{Action: "updating Data failed in IAM service"}, IAMResponseError{Status: http.StatusInternalServerError, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponse{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	// for testing use this
	//var response IAMResponse
	//response = GetOrganizationUpdateData(0)

	return response, IAMResponseError{}
}

func RequestOrganizationInfo(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values) (IAMResponseOrganizationInfo, IAMResponseError) {

	//Generating Request URL
	u, err := url.Parse(config.Conf.IAMService.Url)
	if err != nil {

		return IAMResponseOrganizationInfo{Action: "Url parse is failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}
	requestUrl := u.ResolveReference(iamServiceUrlPath)

	client := httpclient.NewHttpClient(ctx, config.Conf)

	resp, err := client.DoGet(requestUrl.String(), q)
	if err != nil {

		return IAMResponseOrganizationInfo{Action: "IAM service request error"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return IAMResponseOrganizationInfo{Action: "querying Data failed in IAM service"}, IAMResponseError{Status: http.StatusInternalServerError, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponseOrganizationInfo
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponseOrganizationInfo{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	validate := validator.New()

	err = validate.Struct(response.Data.Payload.Organization)
	if err != nil {
		return IAMResponseOrganizationInfo{Action: "IAM service response validation failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	// for testing use this
	//var response IAMResponseOrganizationInfo
	//response = GetOrganizationData()

	return response, IAMResponseError{}
}

func RequestActivateOrganization(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values) (IAMResponse, IAMResponseError) {

	//Generating Request URL
	u, err := url.Parse(config.Conf.IAMService.Url)
	if err != nil {

		return IAMResponse{Action: "Url parse is failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}
	requestUrl := u.ResolveReference(iamServiceUrlPath)

	client := httpclient.NewHttpClient(ctx, config.Conf)

	resp, err := client.DoPost(requestUrl.String(), q, nil)
	if err != nil {

		return IAMResponse{Action: "IAM service request error"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return IAMResponse{Action: "updating Data in IAM service"}, IAMResponseError{Status: http.StatusInternalServerError, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponse{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	// for testing use this
	//var response IAMResponse
	//response = GetOrganizationActivationData(0)

	return response, IAMResponseError{}
}

func RequestDeleteOrganization(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values) (IAMResponse, IAMResponseError) {

	//Generating Request URL
	u, err := url.Parse(config.Conf.IAMService.Url)
	if err != nil {

		return IAMResponse{Action: "Url parse is failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}
	requestUrl := u.ResolveReference(iamServiceUrlPath)

	client := httpclient.NewHttpClient(ctx, config.Conf)

	resp, err := client.DoDelete(requestUrl.String(), q)
	if err != nil {

		return IAMResponse{Action: "IAM service request error"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return IAMResponse{Action: "deleting Data in IAM service"}, IAMResponseError{Status: http.StatusInternalServerError, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponse{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	return response, IAMResponseError{}

	// for testing use this
	//response, err := GetOrganizationDeleteData(0)
	//return response, err

}
