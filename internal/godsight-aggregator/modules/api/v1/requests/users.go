package requests

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func CreateRequestBodyForUser(r *http.Request) (RequestUserBody, error) {

	var reqBody RequestUserBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return RequestUserBody{}, err
	}

	validate := validator.New()

	err := validate.Struct(reqBody)
	if err != nil {
		return RequestUserBody{}, err
	}

	return reqBody, nil

}
