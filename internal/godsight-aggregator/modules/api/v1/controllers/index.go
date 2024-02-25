package controllers

import (
	"fmt"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/modules/api/v1/helpers"
	log "github.com/Blockchain-Framework/controller/pkg/logger"
	"github.com/Blockchain-Framework/controller/pkg/serviceresponse"
	"net/http"
)

func Init(w http.ResponseWriter, r *http.Request) {

	log.Debug(r.Context()).Msg("Inside Init")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Request failed:", r)
		}
	}()

	queryParams := r.URL.Query()
	responseBody, responseError := helpers.FetchInitData(r.Context(), &queryParams)

	if responseError.Error != nil {
		log.Error(r.Context()).Msgf("failed to fetch init data: %s", responseError.Error)
		serviceError := serviceresponse.NewError(w, r)
		serviceError.Action = responseError.Action
		serviceError.AddDetail(responseError.Error)
		_ = serviceError.Write(responseError.Status, nil)
		panic(responseError.Error)
	}

	success := serviceresponse.NewSuccess(w, r)
	success.Action = "successfully fetch init data"
	success.Data = responseBody.Data
	success.Meta = responseBody.Meta

	_ = success.WriteOk()

}
