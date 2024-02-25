package helpers

import (
	"context"
	"encoding/json"
	"github.com/Blockchain-Framework/controller/external/iam_service"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/modules/api/v1/requests"
	"github.com/Blockchain-Framework/controller/pkg/constants"
	"net/url"
	"path"
)

func FetchListUsers(ctx context.Context, q *url.Values) (ResponseUserList, ResponseError) {

	urlPath, err := url.Parse(constants.UsersPath)
	if err != nil {
		responseError := ResponseError{
			Action: "creating IAM service request url is failed",
			Error:  err,
		}
		return ResponseUserList{}, responseError
	}

	response, iamErr := iam_service.RequestUserList(ctx, urlPath, q)

	responseBody := ResponseUserList{}

	if iamErr.Error == nil {
		responseBody.Data.Organizations = response.Data.Payload.Users
		responseBody.Meta = response.Data.Meta
	}

	responseError := ResponseError{
		Action: response.Action,
		Error:  iamErr.Error,
		Status: iamErr.Status,
	}

	return responseBody, responseError
}

func CreateUser(ctx context.Context, q *url.Values, reqBody requests.RequestUserBody) (ResponseUserCreate, ResponseError) {

	responseError := ResponseError{}
	responseBody := ResponseUserCreate{}

	urlPath, err := url.Parse(constants.UsersPath)
	if err != nil {
		responseError.Error = err
		responseError.Action = "creating IAM service request url is failed"
		return responseBody, responseError
	}

	requestBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		responseError.Error = err
		responseError.Action = "creating IAM service request body is failed"
		return responseBody, responseError
	}

	response, iamErr := iam_service.RequestUserCreate(ctx, urlPath, q, requestBodyJson)

	if iamErr.Error == nil {
		responseBody.Data.UserId = response.Data.Payload.UserId
		responseBody.Meta = response.Data.Meta
	}

	responseError.Action = response.Action
	responseError.Error = iamErr.Error
	responseError.Status = iamErr.Status

	return responseBody, responseError

}

func UpdateUser(ctx context.Context, q *url.Values, reqBody requests.RequestUserBody, userId string) ResponseError {

	responseError := ResponseError{}

	urlPath, err := url.Parse(constants.UsersPath)
	if err != nil {
		responseError.Error = err
		responseError.Action = "creating IAM service request url is failed"
		return responseError
	}

	urlPath.Path = path.Join(urlPath.Path, userId)

	requestBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		responseError.Error = err
		responseError.Action = "creating IAM service request body is failed"
		return responseError
	}

	response, iamErr := iam_service.RequestUserUpdate(ctx, urlPath, q, requestBodyJson)

	responseError.Action = response.Action
	responseError.Error = iamErr.Error
	responseError.Status = iamErr.Status

	return responseError

}

func FetchUserInfo(ctx context.Context, q *url.Values, userId string) (ResponseUser, ResponseError) {

	urlPath, err := url.Parse(constants.UsersPath)
	if err != nil {
		responseError := ResponseError{
			Action: "creating IAM service request url is failed",
			Error:  err,
		}
		return ResponseUser{}, responseError
	}

	urlPath.Path = path.Join(urlPath.Path, userId)

	response, iamErr := iam_service.RequestUserInfo(ctx, urlPath, q)

	responseBody := ResponseUser{}

	if iamErr.Error == nil {
		responseBody.Data.Organization = response.Data.Payload.User
		responseBody.Meta = response.Data.Meta
	}

	responseError := ResponseError{
		Action: response.Action,
		Error:  iamErr.Error,
		Status: iamErr.Status,
	}

	return responseBody, responseError
}

func ActivateUser(ctx context.Context, q *url.Values, userId string) ResponseError {

	urlPath, err := url.Parse(constants.UsersPath)
	if err != nil {
		responseError := ResponseError{
			Action: "creating IAM service request url is failed",
			Error:  err,
		}
		return responseError
	}

	urlPath.Path = path.Join(urlPath.Path, userId)

	response, iamErr := iam_service.RequestActivateUser(ctx, urlPath, q)

	responseError := ResponseError{
		Action: response.Action,
		Error:  iamErr.Error,
		Status: iamErr.Status,
	}

	return responseError
}

func DeleteUser(ctx context.Context, q *url.Values, userId string) ResponseError {

	urlPath, err := url.Parse(constants.UsersPath)
	if err != nil {
		responseError := ResponseError{
			Action: "creating IAM service request url is failed",
			Error:  err,
		}
		return responseError
	}

	urlPath.Path = path.Join(urlPath.Path, userId)

	response, iamErr := iam_service.RequestDeleteUser(ctx, urlPath, q)

	responseError := ResponseError{
		Action: response.Action,
		Error:  iamErr.Error,
		Status: iamErr.Status,
	}

	return responseError
}
