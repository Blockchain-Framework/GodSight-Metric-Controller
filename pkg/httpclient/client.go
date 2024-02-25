package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/config"
	"github.com/Blockchain-Framework/controller/pkg/middleware"

	log "github.com/Blockchain-Framework/controller/pkg/logger"
)

const (
	XApiKey = "x-api-key"
)

type Client struct {
	context.Context
	*config.Config
}

func NewHttpClient(ctx context.Context, config *config.Config) *Client {

	return &Client{
		ctx,
		config,
	}
}

func (c *Client) DoGet(requestUrl string, q *url.Values) (*http.Response, error) {

	req, err := http.NewRequestWithContext(c.Context, http.MethodGet, requestUrl, nil)

	if err != nil {
		log.Error(c.Context).Err(err).Msgf("error creating %s %s", http.MethodGet, requestUrl)
		return nil, err
	}

	if q != nil {
		req.URL.RawQuery = q.Encode()
	}

	if len(c.Config.Server.ApiKey) != 0 {
		req.Header.Set(XApiKey, c.Config.Server.ApiKey)
	}

	middleware.SetMandatoryHeaders(req, c.Context)

	httpClient := &http.Client{}

	log.Info(c.Context).Msgf("invoking %s %s", http.MethodGet, requestUrl)

	dumpRequest(c, req)

	return httpClient.Do(req)
}

func (c *Client) DoPost(requestUrl string, q *url.Values, reqBody io.Reader) (*http.Response, error) {

	req, err := http.NewRequestWithContext(c.Context, http.MethodPost, requestUrl, reqBody)

	if err != nil {
		log.Error(c.Context).Err(err).Msgf("error creating %s %s", http.MethodPost, requestUrl)
		return nil, err
	}

	if q != nil {
		req.URL.RawQuery = q.Encode()
	}

	if len(c.Config.Server.ApiKey) != 0 {
		req.Header.Set(XApiKey, c.Config.Server.ApiKey)
	}

	middleware.SetMandatoryHeaders(req, c.Context)

	httpClient := &http.Client{}

	log.Info(c.Context).Msgf("invoking %s %s", http.MethodPost, requestUrl)

	dumpRequest(c, req)

	return httpClient.Do(req)
}

func (c *Client) DoPut(requestUrl string, q *url.Values, reqBody io.Reader) (*http.Response, error) {

	req, err := http.NewRequestWithContext(c.Context, http.MethodPut, requestUrl, reqBody)

	if err != nil {
		log.Error(c.Context).Err(err).Msgf("error creating %s %s", http.MethodPut, requestUrl)
		return nil, err
	}

	if q != nil {
		req.URL.RawQuery = q.Encode()
	}

	if len(c.Config.Server.ApiKey) != 0 {
		req.Header.Set(XApiKey, c.Config.Server.ApiKey)
	}

	middleware.SetMandatoryHeaders(req, c.Context)

	httpClient := &http.Client{}

	log.Info(c.Context).Msgf("invoking %s %s", http.MethodPut, requestUrl)

	dumpRequest(c, req)

	return httpClient.Do(req)
}

func (c *Client) DoDelete(requestUrl string, q *url.Values) (*http.Response, error) {

	req, err := http.NewRequestWithContext(c.Context, http.MethodDelete, requestUrl, nil)

	if err != nil {
		log.Error(c.Context).Err(err).Msgf("error creating %s %s", http.MethodDelete, requestUrl)
		return nil, err
	}

	if q != nil {
		req.URL.RawQuery = q.Encode()
	}

	if len(c.Config.Server.ApiKey) != 0 {
		req.Header.Set(XApiKey, c.Config.Server.ApiKey)
	}

	middleware.SetMandatoryHeaders(req, c.Context)

	httpClient := &http.Client{}

	log.Info(c.Context).Msgf("invoking %s %s", http.MethodDelete, requestUrl)

	dumpRequest(c, req)

	return httpClient.Do(req)
}

func dumpRequest(c *Client, req *http.Request) {

	if c.Config.DumpDownstreamRequest {
		reqDump, err := httputil.DumpRequest(req, true)
		if err == nil {
			fmt.Println(string(reqDump))
		}
	}
}
