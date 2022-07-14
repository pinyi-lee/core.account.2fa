package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/pinyi-lee/core.account.2fa.git/internal/app/router"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/config"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/logger"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/mongo"
	"github.com/pinyi-lee/core.account.2fa.git/test/container"
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	mongoContainer, setupMongoErr := container.SetupMongo(ctx)

	if setupMongoErr != nil {
		log.Fatalf("SetupMongo Fail, %s\n", setupMongoErr)
	}
	defer mongoContainer.Terminate(ctx)

	Setup()
	defer Close()

	r := m.Run()
	os.Exit(r)
}

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
