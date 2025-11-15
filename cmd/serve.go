package cmd

import (
	"fmt"
	"hesdastore/api-ppob/config"
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/middlewares"
	"hesdastore/api-ppob/repositories"
	"hesdastore/api-ppob/services"

	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		db, err := config.InitDatabase(cfg.Database)
		if err != nil {
			panic(err)
		}

		repository := repositories.NewRepositoryRegistry(db)
		service := services.NewServiceRegistry(repository)
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
