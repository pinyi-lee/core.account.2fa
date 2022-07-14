package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pinyi-lee/core.account.2fa.git/internal/app/router"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/config"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/logger"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/mongo"
)

func Setup() {
	var err error

	if err = config.Setup(); err != nil {
		log.Fatal(err)
	}

	if err = logger.Setup(config.Env.LogLevel); err != nil {
		log.Fatal(err)
	}

	if err = mongo.GetInstance().Setup(mongo.Config{
		URI: config.Env.MongoURI,
	}); err != nil {
		log.Fatalf("mongo Setup, error:%v", err)
	}

	if err = router.Setup(); err != nil {
		log.Fatal(err)
	}
}

func Close() {
}

func RunServer() {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Env.Port),
		Handler:      router.Router,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("%s\n", err)
	}
}

// @title        Community Service Swagger
// @description  this service is Community Service
func main() {
	Setup()
	defer Close()
	RunServer()
}
