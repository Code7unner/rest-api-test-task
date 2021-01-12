package cmd

import (
	"context"
	"fmt"
	"github.com/code7unner/rest-api-test-task/users/internal/models"
	"github.com/code7unner/rest-api-test-task/users/internal/server"
	"os"
	"os/signal"

	"github.com/code7unner/rest-api-test-task/users/config"
	"github.com/code7unner/rest-api-test-task/users/internal/db"
	"github.com/code7unner/rest-api-test-task/users/internal/service"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

// @title Users API
// @version 1.0
// @description This is a users microservice.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func API(_ *cobra.Command, _ []string) {
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	if cfg.DebugMode {
		log.SetLevel(log.DEBUG)
	}

	d, err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	uModel := models.NewUsersModel(d)

	e := server.New(
		service.New(uModel, cfg.JWTSecret, cfg.ExpiresInMinutes),
		[]byte(cfg.JWTSecret),
	)
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.ServerAddress)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)
	<-terminate
	signal.Stop(terminate)
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	cancel()
}

// NewAPICmd return api command
func NewAPICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "run api",
		Long:  "starts users API server",
		Run:   API,
	}
}
