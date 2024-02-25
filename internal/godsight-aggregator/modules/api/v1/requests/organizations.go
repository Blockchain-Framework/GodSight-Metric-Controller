package requests

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func CreateRequestBodyForOrganizationCreation(r *http.Request) (RequestOrganizationInsertionBody, error) {

	var reqBody RequestOrganizationInsertionBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return RequestOrganizationInsertionBody{}, err
	}

	validate := validator.New()

	err := validate.Struct(reqBody)
	if err != nil {
		return RequestOrganizationInsertionBody{}, err
	}

	return reqBody, nil

}

func CreateRequestBodyForOrganizationUpdate(r *http.Request) (RequestOrganizationUpdateBody, error) {

	var reqBody RequestOrganizationUpdateBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return RequestOrganizationUpdateBody{}, err
	}

	validate := validator.New()

	err := validate.Struct(reqBody)
	if err != nil {
		return RequestOrganizationUpdateBody{}, err
	}

	return reqBody, nil

}
