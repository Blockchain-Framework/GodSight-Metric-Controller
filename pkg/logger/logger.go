package logger

import (
	"context"
	"io"
	"os"
	"reflect"

	"github.com/Blockchain-Framework/controller/pkg/constants"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const RequestParams = "request_params"
const CustomParams = "custom_params"

type LogContext interface {
	GetOrgId() string
	GetTraceId() string
}

var log zerolog.Logger

func New(isDebug bool) {
	setLogLevel(isDebug)

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log = zerolog.New(os.Stderr).With().Timestamp().Logger()
}

func setLogLevel(isDebug bool) {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)
}

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	return log.Output(w)
}

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return log.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func level(level zerolog.Level) zerolog.Logger {
	return log.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return log.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return log.Hook(h)
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug(ctx context.Context) *zerolog.Event {
	return addContextDetails(ctx).Debug()
}

// Trace starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Trace(ctx context.Context) *zerolog.Event {
	return addContextDetails(ctx).Trace()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info(ctx context.Context) *zerolog.Event {
	return addContextDetails(ctx).Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn(ctx context.Context) *zerolog.Event {
	return addContextDetails(ctx).Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error(ctx context.Context) *zerolog.Event {
	return addContextDetails(ctx).Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal(ctx context.Context) *zerolog.Event {
	return addContextDetails(ctx).Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic(ctx context.Context) *zerolog.Event {
	return addContextDetails(ctx).Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return log.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log(ctx context.Context) *zerolog.Event {
	return addContextDetails(ctx).Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	log.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

func addContextDetails(ctx context.Context) *zerolog.Logger {
	slog := log

	// add mandatory params
	if traceId, ok := ctx.Value(constants.HeaderTraceId).(string); ok && len(traceId) != 0 {
		slog = log.With().Str("trace_id", traceId).Logger()
	}

	if orgId, ok := ctx.Value(constants.HeaderOrganizationId).(string); ok && len(orgId) != 0 {
		slog = slog.With().Str("org_id", orgId).Logger()
	}

	// add request params
	if exclude := ctx.Value("ExcludeRequestParams"); exclude != nil {
		return &slog
	}

	p := ctx.Value(RequestParams)
	if p == nil {
		return &slog
	}

	v := reflect.ValueOf(&p).Elem().Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		typeField := v.Type().Field(i)

		// Get the field tag value
		tag := typeField.Tag.Get("json")

		slog = slog.With().Str(tag, field.String()).Logger()
	}

	return &slog
}
