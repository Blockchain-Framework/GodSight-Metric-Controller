package middleware

import (
	"github.com/rs/zerolog"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type logEntry struct {
	ReceivedTime  time.Time
	StatusCode    int
	RequestMethod string
	RequestURL    string
	UserAgent     string
	Referer       string
	Proto         string
	RemoteIP      string
	ServerIP      string
	Latency       time.Duration
	TraceId       string
}

type LogRecord struct {
	http.ResponseWriter
	*logEntry
}

func (r *LogRecord) Write(p []byte) (int, error) {
	return r.ResponseWriter.Write(p)
}

func (r *LogRecord) WriteHeader(status int) {
	r.logEntry.StatusCode = status
	r.ResponseWriter.WriteHeader(status)
}

func RequestLogger(next http.Handler) http.Handler {

	whitelistUrls := make(map[string]bool)

	urls := os.Getenv("RL_WHITELIST_URLS")
	if len(urls) != 0 {
		entries := strings.Split(urls, ",")
		for _, e := range entries {
			whitelistUrls[e] = true
		}
	}

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if _, ok := whitelistUrls[r.URL.String()]; ok {
			next.ServeHTTP(w, r)
			return
		}

		start := time.Now()

		le := &logEntry{
			ReceivedTime:  start,
			RequestMethod: r.Method,
			RequestURL:    r.URL.String(),
			UserAgent:     r.UserAgent(),
			Referer:       r.Referer(),
			Proto:         r.Proto,
			RemoteIP:      ipFromHostPort(r.RemoteAddr),
			TraceId:       GetTraceId(r.Context()),
		}

		if addr, ok := r.Context().Value(http.LocalAddrContextKey).(net.Addr); ok {
			le.ServerIP = ipFromHostPort(addr.String())
		}

		record := &LogRecord{
			ResponseWriter: w,
			logEntry:       le,
		}

		next.ServeHTTP(record, r)

		le.Latency = time.Since(start)

		log.Info().
			Str("trace_id", le.TraceId).
			Time("received_time", le.ReceivedTime).
			Int("status_code", le.StatusCode).
			Str("method", le.RequestMethod).
			Str("url", le.RequestURL).
			Str("agent", le.UserAgent).
			Str("referer", le.Referer).
			Str("proto", le.Proto).
			Str("remote_ip", le.RemoteIP).
			Str("server_ip", le.ServerIP).
			Dur("latency (ns)", le.Latency).
			Msg("")
	})
}

func ipFromHostPort(hp string) string {

	h, _, err := net.SplitHostPort(hp)
	if err != nil {
		return ""
	}

	if len(h) > 0 && h[0] == '[' {
		return h[1 : len(h)-1]
	}

	return h
}
