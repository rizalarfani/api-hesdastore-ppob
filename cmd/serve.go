package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	clientsCfg "hesdastore/api-ppob/clients/config"
	digiclient "hesdastore/api-ppob/clients/digiflazz"
	"hesdastore/api-ppob/config"
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/middlewares"
	"hesdastore/api-ppob/pkg/rabbitmq"
	"hesdastore/api-ppob/repositories"
	"hesdastore/api-ppob/services"
	"log"
	"os"
	"os/signal"

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

		rabbitMqConfig := rabbitmq.Config{
			Host:     cfg.RabbitMQ.Host,
			Port:     cfg.RabbitMQ.Port,
			Username: cfg.RabbitMQ.Username,
			Password: cfg.RabbitMQ.Password,
			VHost:    cfg.RabbitMQ.VHost,
		}
		managerRabbitMQ, err := rabbitmq.NewManager(rabbitMqConfig)
		if err != nil {
			log.Fatalf("failed to create rabbitmq manager: %v", err)
		}
		defer managerRabbitMQ.Close()
		managerRabbitMQ.SetupWebhook(ctx)

		service := services.NewServiceRegistry(repository, client, clientConfig, managerRabbitMQ)
		controller := controllers.NewControllerRegistry(service)
		middleware := middlewares.NewAuthMiddleware(service)

		router := config.NewRoutes(controller, middleware)

		// start serve https
		go func() {
			port := fmt.Sprintf(":%d", cfg.Port)
			router.Run(port)
		}()

		// start consumer webhook
		go func() {
			consumer := managerRabbitMQ.GetConsumer()
			handler := func(ctx context.Context, body []byte) error {
				var event dto.TransactionUpdateEventWebhook
				if err := json.Unmarshal(body, &event); err != nil {
					return err
				}

				err := service.Webhook().SendWebhook(ctx, &event)
				if err != nil {
					return err
				}

				return nil
			}

			consumer.Consume(ctx, "transaction.update.webhook", handler)
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)

		<-quit
		log.Println("Shutting down gracefully...")

	},
}

func Run() {
	err := command.Execute()
	if err != nil {
		panic(err)
	}
}
