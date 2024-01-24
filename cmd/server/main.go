// just a simple example buildable
package main

import (
	"github.com/SOAT1StackGoLang/msvc-production/internal/service"
	"github.com/SOAT1StackGoLang/msvc-production/internal/transport"
	"os"

	logger "github.com/SOAT1StackGoLang/msvc-payments/pkg/middleware"
)

func main() {
	redisStore, err := initializeApp()
	if err != nil {
		os.Exit(1)
	}

	cacheSvc := service.NewCacheService(*redisStore, logger.InfoLogger)
	publisher := service.NewPublisher(*redisStore)

	svc := service.NewProductionService(cacheSvc, publisher)

	httpHandler := transport.NewHttpHandler(svc, logger.InfoLogger)

	logger.Info("Starting http server...")
	transport.NewHTTPServer(":8080", httpHandler)
}
