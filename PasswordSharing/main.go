package main

import (
	"github.com/porwalameet/go-projects/PasswordSharing/config"
	"github.com/porwalameet/go-projects/PasswordSharing/controller"
	"github.com/porwalameet/go-projects/PasswordSharing/database"
	"github.com/porwalameet/go-projects/PasswordSharing/health"
	"github.com/porwalameet/go-projects/PasswordSharing/helper"
	"github.com/porwalameet/go-projects/PasswordSharing/logger"
	"github.com/porwalameet/go-projects/PasswordSharing/server"
	"github.com/porwalameet/go-projects/PasswordSharing/service"
)

func main() {
	appConfiguration, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	appLogger := logger.NewLoggerFactory(appConfiguration)

	encoder := helper.NewEncoder(appConfiguration)
	databaseFactory := database.NewFactory(appConfiguration, appLogger)
	randomFactory := helper.NewRandomFactory()
	service := service.NewPasswordService(databaseFactory, appConfiguration, randomFactory, appLogger, encoder)

	pgHealthCheck := health.NewPgHealthCheck(databaseFactory, appLogger)

	server := server.NewServer(
		appLogger,
		appConfiguration,
		controller.NewCreateLinkController(service, appConfiguration),
		controller.NewGetLinkController(service),
		controller.NewHealthController(pgHealthCheck),
	)

	if err = server.Run(); err != nil {
		panic(err)
	}
}
