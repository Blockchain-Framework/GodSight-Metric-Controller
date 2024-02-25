package middleware

import (
	"context"
	"net/http"

	"github.com/Blockchain-Framework/controller/pkg/constants"
	"github.com/google/uuid"
)

func TraceId(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		traceId := r.Header.Get(constants.HeaderTraceId)
		if len(traceId) <= 0 {
			traceId = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), constants.HeaderTraceId, traceId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetTraceId(ctx context.Context) string {

	if ctx != nil {
		if traceId, ok := ctx.Value(constants.HeaderTraceId).(string); ok {
			return traceId
		}
	}

	return ""
}
