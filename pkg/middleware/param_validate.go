package middleware

import (
	"context"
	"net/http"

	log "github.com/Blockchain-Framework/controller/pkg/logger"
	"github.com/Blockchain-Framework/controller/pkg/serviceresponse"
	"github.com/Blockchain-Framework/controller/pkg/util"
	validatorV10 "github.com/go-playground/validator/v10"
)

const RequestParams = "request_params"

type RulesFunc func(r *http.Request) interface{}

func ParamValidate(validator *util.Validator, rulesFunc RulesFunc) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			params := rulesFunc(r)

			if err := validator.Struct(params); err != nil {

				responseError := serviceresponse.NewError(w, r)
				responseError.Action = "validating request parameters"
				responseError.Message = "Invalid Request"

				for _, e := range err.(validatorV10.ValidationErrors) {
					responseError.AddValidatorDetail(e, e.Translate(*validator.Translator))
				}

				if err = responseError.Write(http.StatusBadRequest, nil); err != nil {
					log.Warn(r.Context()).Err(err).Msg("unable to send request param validation error")
				}

				return
			}

			ctx := context.WithValue(r.Context(), RequestParams, params)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
