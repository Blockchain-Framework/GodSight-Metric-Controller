package serviceresponse

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Blockchain-Framework/controller/pkg/constants"
	"github.com/pkg/errors"
)

const (
	DefaultPayLoadName = "payload"
)

type Success struct {
	Action      string      `json:"action,omitempty"`
	PayloadName string      `json:"-"`
	Data        interface{} `json:"data"`
	Meta        interface{} `json:"meta"`

	w http.ResponseWriter `json:"-"`
	r *http.Request       `json:"-"`
}

type Error struct {
	Action  string        `json:"action,omitempty"`
	Message string        `json:"message,omitempty"`
	Errors  []DetailError `json:"errs,omitempty"`

	w http.ResponseWriter `json:"-"`
	r *http.Request       `json:"-"`
}

type Meta struct {
	Size     int `json:"size"`
	NextPage int `json:"nextPage"`
	Count    int `json:"count"`
	Page     int `json:"page"`
}

type DetailError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

// /////////////////////////////////////////////////////////////////////////////////////////////////
// Success Response
// /////////////////////////////////////////////////////////////////////////////////////////////////
func NewSuccess(w http.ResponseWriter, r *http.Request) *Success {
	return &Success{
		Action: "action", w: w, r: r, Data: nil, Meta: nil,
	}
}

func (s *Success) DataName() string {
	if len(s.PayloadName) > 0 {
		return s.PayloadName
	}

	return DefaultPayLoadName
}

func (s *Success) WriteOk() (err error) {
	return s.Write(http.StatusOK, nil)
}

func (s *Success) Write(code int, headers map[string]string) (err error) {

	m := map[string]interface{}{}
	m["action"] = s.Action

	m["data"] = map[string]interface{}{
		s.DataName(): s.Data,
		"meta":       s.Meta,
	}

	m["traceId"] = getTraceId(s.r.Context())
	m["timestamp"] = time.Now()

	if err := write(s.w, m, code, headers); err != nil {
		return err
	}

	return nil
}

// /////////////////////////////////////////////////////////////////////////////////////////////////
// Error response
// /////////////////////////////////////////////////////////////////////////////////////////////////
func NewError(w http.ResponseWriter, r *http.Request) *Error {
	return &Error{
		Message: "", Errors: nil, w: w, r: r,
	}
}

func (e *Error) AddDetail(err error) {

	switch v := err.(type) {
	case *ServiceError:
		e.Errors = append(e.Errors, DetailError{Code: v.Code, Message: v.Message, Detail: v.Detail})
	default:
		e.Errors = append(e.Errors, DetailError{Code: 0, Message: err.Error()})
	}
}

func (e *Error) AddValidatorDetail(err error, message string) {
	e.Errors = append(e.Errors, DetailError{Code: 1234, Message: message})
}

func (e *Error) Write(code int, headers map[string]string) (err error) {
	m := map[string]interface{}{}
	m["action"] = e.Action
	m["message"] = e.Message
	m["error"] = e.Errors
	m["traceId"] = getTraceId(e.r.Context())
	m["timestamp"] = time.Now()

	if err := write(e.w, m, code, headers); err != nil {
		return err
	}

	return nil
}

// /////////////////////////////////////////////////////////////////////////////////////////////////
// Utility
// /////////////////////////////////////////////////////////////////////////////////////////////////
func write(w http.ResponseWriter, d interface{}, code int, headers map[string]string) (err error) {

	if code == 0 {
		return errors.New("invalid http status code")
	}

	for k, v := range headers {
		w.Header().Set(k, v)
	}

	response, err := json.Marshal(d)
	if err != nil {
		return errors.Wrap(err, "failed marshalling response")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if _, err := w.Write(response); err != nil {
		return errors.Wrap(err, "failed writing response")
	}

	return nil
}

func getTraceId(ctx context.Context) string {

	if ctx != nil {
		if traceId, ok := ctx.Value(constants.HeaderTraceId).(string); ok {
			return traceId
		}
	}

	return ""
}
