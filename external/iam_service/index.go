package iam_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/config"
	"github.com/Blockchain-Framework/controller/pkg/httpclient"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/url"
)

func RequestInitData(ctx context.Context, iamServiceUrlPath *url.URL, q *url.Values) (IAMResponseInit, IAMResponseError) {

	//Generating Request URL
	u, err := url.Parse(config.Conf.IAMService.Url)
	if err != nil {

		return IAMResponseInit{Action: "Url parse is failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}
	requestUrl := u.ResolveReference(iamServiceUrlPath)

	client := httpclient.NewHttpClient(ctx, config.Conf)

	resp, err := client.DoGet(requestUrl.String(), q)
	if err != nil {

		return IAMResponseInit{Action: "IAM service request error"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	if resp.StatusCode != http.StatusOK {
		return IAMResponseInit{Action: "querying Data failed in IAM service"}, IAMResponseError{Status: http.StatusInternalServerError, Error: fmt.Errorf("iam service request failed")}
	}

	var response IAMResponseInit
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return IAMResponseInit{Action: "IAM service response decode failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
	}

	validate := validator.New()

	if err != nil {
		for _, type_ := range response.Data.Payload.OrganizationTypes {
			err := validate.Struct(type_)
			if err != nil {
				return IAMResponseInit{Action: "IAM service response validation failed"}, IAMResponseError{Status: http.StatusBadRequest, Error: err}
			}
		}
	}

	return response, IAMResponseError{}
}
