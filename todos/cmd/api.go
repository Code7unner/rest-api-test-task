package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/code7unner/rest-api-test-task/todos/config"
	"github.com/code7unner/rest-api-test-task/todos/internal/db"
	"github.com/code7unner/rest-api-test-task/todos/internal/models"
	"github.com/code7unner/rest-api-test-task/todos/internal/server"
	"github.com/code7unner/rest-api-test-task/todos/internal/service"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

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

	tModel := models.NewTodosModel(d)

	e := server.New(
		service.New(tModel, cfg.JWTSecret, cfg.ExpiresInMinutes),
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
		Long:  "starts todos API server",
		Run:   API,
	}
}
