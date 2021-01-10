package main

import (
	"os"

	"github.com/code7unner/rest-api-test-task/todos/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "todos"}
	rootCmd.AddCommand(cmd.NewAPICmd(), cmd.NewMigrateCmd())
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}
