package config

import (
	"time"

	"github.com/Blockchain-Framework/controller/pkg/util"
	"github.com/joeshaw/envdecode"
)

var Conf *Config

type Config struct {
	Debug                 bool `env:"DEBUG,required"`
	DumpDownstreamRequest bool `env:"DUMP_DOWNSTREAM_REQUEST"`
	*util.Validator       `json:"-"`

	IAMService iamService

	Server serverConf
}

type serverConf struct {
	AppName       string        `env:"APP_NAME,required"`
	ServicePort   int           `env:"SERVICE_PORT,required"`
	ClientTimeout time.Duration `env:"CLIENT_TIMEOUT,required"`
	ContextPath   string        `env:"CONTEXT_PATH,required"`

	//These values are used for local testing
	ApiGatewayUrl string `env:"API_GATEWAY_URL"`
	ApiKey        string `env:"API_KEY"`
}

type iamService struct {
	Url string `env:"IAM_SERVICE_URL, required"`
}

func LoadFromEnvironment() error {

	var conf Config

	if err := envdecode.StrictDecode(&conf); err != nil {
		return err
	}

	Conf = &conf

	return nil
}
