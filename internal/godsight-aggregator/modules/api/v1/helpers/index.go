package helpers

import (
	"context"
	"github.com/Blockchain-Framework/controller/external/iam_service"
	"github.com/Blockchain-Framework/controller/pkg/constants"
	"net/url"
)

func FetchInitData(ctx context.Context, q *url.Values) (ResponseInitData, ResponseError) {

	urlPath, err := url.Parse(constants.InitPath)
	if err != nil {
		responseError := ResponseError{
			Action: "creating IAM service request url is failed",
			Error:  err,
		}
		return ResponseInitData{}, responseError
	}

	response, iamErr := iam_service.RequestInitData(ctx, urlPath, q)

	responseBody := ResponseInitData{}

	if iamErr.Error == nil {
		responseBody.Data.OrganizationTypes = response.Data.Payload.OrganizationTypes
		responseBody.Meta = response.Data.Meta
	}

	responseError := ResponseError{
		Action: response.Action,
		Error:  iamErr.Error,
		Status: iamErr.Status,
	}

	return responseBody, responseError
}
