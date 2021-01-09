package main

import (
	"os"

	"github.com/code7unner/rest-api-test-task/users/cmd"
	"github.com/spf13/cobra"
)

//go:generate swag i -g cmd/api.go -o docs

func main() {
	rootCmd := &cobra.Command{Use: "users"}
	rootCmd.AddCommand(cmd.NewAPICmd(), cmd.NewMigrateCmd())
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}
