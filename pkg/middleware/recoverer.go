package middleware

import (
	"encoding/json"
	"net/http"

	log "github.com/Blockchain-Framework/controller/pkg/logger"
)

func Recoverer(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil && err != http.ErrAbortHandler {

				if runtimeError, ok := err.(error); ok {
					log.Error(r.Context()).Stack().Err(runtimeError).Msgf("Internal error occurred. %s %s", r.Method, r.URL.String())
				} else {
					log.Error(r.Context()).Msgf("Internal error occurred. %s %s. %s", r.Method, r.URL.String(), err)
				}

				errorMessage, _ := json.Marshal(map[string]string{
					"error": "There was an code_generator server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				if _, writeError := w.Write(errorMessage); writeError != nil {
					log.Error(r.Context()).Msgf("Unable to write code_generator error response. %s %s. %s", r.Method, r.URL.String(), err)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
