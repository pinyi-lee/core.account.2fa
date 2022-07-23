package config

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/caarlos0/env/v6"
)

func Setup() error {
	err := env.Parse(&Env)
	if err != nil {
		fmt.Printf("config parse fail, %+v\n", err)
		return err
	}

	err = Env.Validate()
	if err != nil {
		fmt.Printf("config validate fail, %+v\n", err)
		return err
	}

	return nil
}

const (
	logLevelError   = "error"
	logLevelDebug   = "debug"
	logLevelWarning = "warn"
	logLevelInfo    = "info"

	deployEnvDevelop    = "develop"
	deployEnvStage      = "stage"
	deployEnvProduction = "production"
)

var Env EnvVariable

type EnvVariable struct {
	Version           string `env:"VERSION" envDefault:"0.1.0"`
	Port              string `env:"GO_HTTP_PORT,required"`
	LogLevel          string `env:"LOG_LEVEL" envDefault:"INFO"`
	DeployEnvironment string `env:"DEPLOY_ENVIRONMENT" envDefault:"DEVELOP"`
	MongoURI          string `env:"MONGO_URI,required"`
	ImageWidth        int    `env:"IMAGE_WIDTH" envDefault:"256"`
	ImageHeight       int    `env:"IMAGE_HEIGHT" envDefault:"256"`
}

func (env EnvVariable) Validate() (err error) {
	port, err := strconv.ParseUint(env.Port, 10, 16)
	if err != nil || port <= 0 || port > uint64(65535) {
		err = errors.New("required environment variable \"GO_HTTP_PORT\" should be 0~65535")
		return
	}
	if d := strings.ToLower(env.LogLevel); d != logLevelError && d != logLevelDebug && d != logLevelWarning && d != logLevelInfo {
		err = errors.New("required environment variable \"LOG_LEVEL\" should be \"ERROR|DEBUG|WARN|INFO\"")
		return
	}
	if d := strings.ToLower(env.DeployEnvironment); d != deployEnvDevelop && d != deployEnvStage && d != deployEnvProduction {
		err = errors.New("required environment variable \"DEPLOY_ENVIRONMENT\" should be \"DEVELOP|STAGE|PRODUCTION\"")
		return
	}

	return
}

func IsProduction() bool {
	return strings.ToLower(Env.DeployEnvironment) == deployEnvProduction
}
