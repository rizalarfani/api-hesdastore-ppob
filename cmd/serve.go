package cmd

import (
	"context"
	"fmt"
	clientsCfg "hesdastore/api-ppob/clients/config"
	digiclient "hesdastore/api-ppob/clients/digiflazz"
	"hesdastore/api-ppob/config"
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/middlewares"
	"hesdastore/api-ppob/repositories"
	"hesdastore/api-ppob/services"
	"log"

	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cfg := config.Load()
		db, err := config.InitDatabase(cfg.Database)
		if err != nil {
			panic(err)
		}

		repository := repositories.NewRepositoryRegistry(db)
		configRepo, err := repository.Config().GetConfigDigiflazz(ctx)

		if err != nil {
			log.Fatalf("failed load digiflazz config: %v", err)
		}

		cfg.WithDigiflazz(config.Digiflazz{
			Host:     "https://digiflazz.hesda-store.com/v1",
			Username: configRepo.Username,
			ApiKey:   configRepo.ApiKey,
		})

		clientConfig := clientsCfg.NewClientConfig(
			clientsCfg.WithBaseURL(cfg.Digiflazz.Host),
			clientsCfg.WithSignatureKey(cfg.Digiflazz.ApiKey),
			clientsCfg.WithUsername(cfg.Digiflazz.Username),
		)
		client := digiclient.NewDigiflazzClient(clientConfig)

		service := services.NewServiceRegistry(repository, client, clientConfig)
		controller := controllers.NewControllerRegistry(service)
		middleware := middlewares.NewAuthMiddleware(service)

		router := config.NewRoutes(controller, middleware)

		port := fmt.Sprintf(":%d", cfg.Port)
		router.Run(port)
	},
}

func Run() {
	err := command.Execute()
	if err != nil {
		panic(err)
	}
}
