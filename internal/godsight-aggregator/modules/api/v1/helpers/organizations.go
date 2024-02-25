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

func FetchListOrganization(ctx context.Context, q *url.Values) (ResponseOrganizationList, ResponseError) {

	urlPath, err := url.Parse(constants.OrganizationsPath)
	if err != nil {
		responseError := ResponseError{
			Action: "creating IAM service request url is failed",
			Error:  err,
		}
		return ResponseOrganizationList{}, responseError
	}

	response, iamErr := iam_service.RequestOrganizationList(ctx, urlPath, q)

	responseBody := ResponseOrganizationList{}

	if iamErr.Error == nil {
		responseBody.Data.Organizations = response.Data.Payload.Organizations
		responseBody.Meta = response.Data.Meta
	}

	responseError := ResponseError{
		Action: response.Action,
		Error:  iamErr.Error,
		Status: iamErr.Status,
	}

	return responseBody, responseError
}

func CreateOrganization(ctx context.Context, q *url.Values, reqBody requests.RequestOrganizationInsertionBody) (ResponseOrganizationCreate, ResponseError) {

	responseError := ResponseError{}

	urlPath, err := url.Parse(constants.OrganizationsPath)
	if err != nil {
		responseError.Error = err
		responseError.Action = "creating IAM service request url is failed"
		return ResponseOrganizationCreate{}, responseError
	}

	requestBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		responseError.Error = err
		responseError.Action = "creating IAM service request body is failed"
		return ResponseOrganizationCreate{}, responseError
	}

	response, iamErr := iam_service.RequestOrganizationCreate(ctx, urlPath, q, requestBodyJson)

	responseError.Action = response.Action
	responseError.Error = iamErr.Error
	responseError.Status = iamErr.Status

	responseBody := ResponseOrganizationCreate{}

	if iamErr.Error == nil {
		responseBody.Data.OrganizationId = response.Data.Payload.OrganizationId
		responseBody.Meta = response.Data.Meta
	}

	return responseBody, responseError

}

func UpdateOrganization(ctx context.Context, q *url.Values, reqBody requests.RequestOrganizationUpdateBody, organizationId string) ResponseError {

	responseError := ResponseError{}

	urlPath, err := url.Parse(constants.OrganizationsPath)
	if err != nil {
		responseError.Error = err
		responseError.Action = "creating IAM service request url is failed"
		return responseError
	}

	urlPath.Path = path.Join(urlPath.Path, organizationId)

	requestBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		responseError.Error = err
		responseError.Action = "creating IAM service request body is failed"
		return responseError
	}

	response, iamErr := iam_service.RequestOrganizationUpdate(ctx, urlPath, q, requestBodyJson)

	responseError.Action = response.Action
	responseError.Error = iamErr.Error
	responseError.Status = iamErr.Status

	return responseError

}

func FetchOrganizationInfo(ctx context.Context, q *url.Values, organizationId string) (ResponseOrganization, ResponseError) {

	urlPath, err := url.Parse(constants.OrganizationsPath)
	if err != nil {
		responseError := ResponseError{
			Action: "creating IAM service request url is failed",
			Error:  err,
		}
		return ResponseOrganization{}, responseError
	}

	urlPath.Path = path.Join(urlPath.Path, organizationId)

	response, iamErr := iam_service.RequestOrganizationInfo(ctx, urlPath, q)

	responseBody := ResponseOrganization{}

	if iamErr.Error == nil {
		responseBody.Data.Organization = response.Data.Payload.Organization
		responseBody.Meta = response.Data.Meta
	}

	responseError := ResponseError{
		Action: response.Action,
		Error:  iamErr.Error,
		Status: iamErr.Status,
	}

	return responseBody, responseError
}

func ActivateOrganization(ctx context.Context, q *url.Values, organizationId string) ResponseError {

	urlPath, err := url.Parse(constants.OrganizationsPath)
	if err != nil {
		responseError := ResponseError{
			Action: "creating IAM service request url is failed",
			Error:  err,
		}
		return responseError
	}

	urlPath.Path = path.Join(urlPath.Path, organizationId)

	response, iamErr := iam_service.RequestActivateOrganization(ctx, urlPath, q)

	responseError := ResponseError{
		Action: response.Action,
		Error:  iamErr.Error,
		Status: iamErr.Status,
	}

	return responseError
}

func DeleteOrganization(ctx context.Context, q *url.Values, organizationId string) ResponseError {

	urlPath, err := url.Parse(constants.OrganizationsPath)
	if err != nil {
		responseError := ResponseError{
			Action: "creating IAM service request url is failed",
			Error:  err,
		}
		return responseError
	}

	urlPath.Path = path.Join(urlPath.Path, organizationId)

	response, iamErr := iam_service.RequestDeleteOrganization(ctx, urlPath, q)

	responseError := ResponseError{
		Action: response.Action,
		Error:  iamErr.Error,
		Status: iamErr.Status,
	}

	return responseError
}
