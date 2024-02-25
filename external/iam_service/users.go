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

func RequestUserList(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values) (IAMResponseUserList, IAMResponseError) {

	//Generating Request URL
	u, err := url.Parse(config.Conf.IAMService.Url)
	if err != nil {

		return IAMResponseUserList{Action: "Url parse is failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}
	requestUrl := u.ResolveReference(iamServiceUrlPath)

	client := httpclient.NewHttpClient(ctx, config.Conf)

	resp, err := client.DoGet(requestUrl.String(), q)
	if err != nil {

		return IAMResponseUserList{Action: "IAM service request error"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return IAMResponseUserList{Action: "querying Data failed in IAM service"}, IAMResponseError{Status: http.StatusBadRequest, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponseUserList
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponseUserList{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	validate := validator.New()

	if err != nil {
		for _, user := range response.Data.Payload.Users {
			err := validate.Struct(user)
			if err != nil {
				return IAMResponseUserList{Action: "IAM service response validation failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
			}
		}
	}

	// for testing use this
	//var response IAMResponseUserList
	//response = GetUserListData()

	return response, IAMResponseError{}
}

func RequestUserCreate(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values, requestBodyJson []byte) (IAMResponseUserCreate, IAMResponseError) {

	//Generating Request URL
	u, err := url.Parse(config.Conf.IAMService.Url)
	if err != nil {

		return IAMResponseUserCreate{Action: "IAM service request url creation failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}
	requestUrl := u.ResolveReference(iamServiceUrlPath)

	client := httpclient.NewHttpClient(ctx, config.Conf)

	resp, err := client.DoPost(requestUrl.String(), q, bytes.NewBuffer(requestBodyJson))
	if err != nil {

		return IAMResponseUserCreate{Action: "IAM service request error"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	if resp.StatusCode != http.StatusCreated {
		return IAMResponseUserCreate{Action: "inserting Data failed in IAM service"}, IAMResponseError{Status: http.StatusBadRequest, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponseUserCreate
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponseUserCreate{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	validate := validator.New()

	err = validate.Struct(response.Data)
	if err != nil {
		return IAMResponseUserCreate{Action: "IAM service response validation failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	return response, IAMResponseError{}

	// for testing use this
	//response, err := GetUserCreationData(0)
	//return response, err
}

func RequestUserUpdate(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values, requestBodyJson []byte) (IAMResponse, IAMResponseError) {

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
		return IAMResponse{Action: "updating Data failed in IAM service"}, IAMResponseError{Status: http.StatusBadRequest, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponse{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	return response, IAMResponseError{}

	// for testing use this
	//response, err := GetUserUpdateData(0)
	//return response, err

}

func RequestUserInfo(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values) (IAMResponseUserInfo, IAMResponseError) {

	//Generating Request URL
	u, err := url.Parse(config.Conf.IAMService.Url)
	if err != nil {

		return IAMResponseUserInfo{Action: "Url parse is failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}
	requestUrl := u.ResolveReference(iamServiceUrlPath)

	client := httpclient.NewHttpClient(ctx, config.Conf)

	resp, err := client.DoGet(requestUrl.String(), q)
	if err != nil {

		return IAMResponseUserInfo{Action: "IAM service request error"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return IAMResponseUserInfo{Action: "querying Data failed in IAM service"}, IAMResponseError{Status: http.StatusBadRequest, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponseUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponseUserInfo{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	validate := validator.New()

	err = validate.Struct(response.Data.Payload.User)
	if err != nil {
		return IAMResponseUserInfo{Action: "IAM service response validation failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	return response, IAMResponseError{}

	// for testing use this
	//response := GetUserData()
	//return response, nil
}

func RequestActivateUser(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values) (IAMResponse, IAMResponseError) {

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
		return IAMResponse{Action: "updating Data in IAM service"}, IAMResponseError{Status: http.StatusBadRequest, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponse{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	return response, IAMResponseError{}

	// for testing use this
	//response, err := GetUserActivationData(0)
	//return response, err

}

func RequestDeleteUser(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values) (IAMResponse, IAMResponseError) {

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
		return IAMResponse{Action: "deleting Data in IAM service"}, IAMResponseError{Status: http.StatusBadRequest, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponse{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	return response, IAMResponseError{}

	// for testing use this
	//response, err := GetUserDeleteData(0)
	//return response, err

}
