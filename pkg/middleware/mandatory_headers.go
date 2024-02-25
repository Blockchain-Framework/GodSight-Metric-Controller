package middleware

import (
	"context"
	"net/http"

	"github.com/Blockchain-Framework/controller/pkg/constants"
)

type partialContextFn func(context.Context) context.Context

func MandatoryHeaders(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		primedCtx := r.WithContext(buildContext(
			r.Context(),
			traceId(r),
			requestedUser(r),
			requestedUserEmail(r),
		))

		next.ServeHTTP(w, primedCtx)
	})
}

func SetMandatoryHeaders(req *http.Request, ctx context.Context) {
	req.Header.Set(constants.HeaderTraceId, getContextValue(ctx, constants.HeaderTraceId))
	req.Header.Set(constants.HeaderOrganizationId, getContextValue(ctx, constants.HeaderOrganizationId))
	req.Header.Set(constants.HeaderRequestedUser, getContextValue(ctx, constants.HeaderRequestedUser))
	req.Header.Set(constants.HeaderRequestedUserEmail, getContextValue(ctx, constants.HeaderRequestedUserEmail))
}

func getContextValue(ctx context.Context, key string) string {
	if ctx != nil {
		if value, ok := ctx.Value(key).(string); ok {
			return value
		}
	}
	return ""
}

func traceId(r *http.Request) partialContextFn {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, constants.HeaderTraceId, GetTraceId(r.Context()))
	}
}

func requestedUser(r *http.Request) partialContextFn {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, constants.HeaderRequestedUser, r.Header.Get(constants.HeaderRequestedUser))
	}
}

func requestedUserEmail(r *http.Request) partialContextFn {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, constants.HeaderRequestedUserEmail, r.Header.Get(constants.HeaderRequestedUserEmail))
	}
}

func buildContext(ctx context.Context, ctxFns ...partialContextFn) context.Context {
	for _, f := range ctxFns {
		ctx = f(ctx)
	}

	return ctx
}
