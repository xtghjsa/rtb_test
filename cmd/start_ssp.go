package cmd

import (
	"fmt"
	"os"
	"test_project/internal/api"

	"github.com/spf13/cobra"
)

var startSSPCmd = &cobra.Command{
	Use:   "startssp",
	Short: "Start SSP server",
	Long:  "Start SSP server",
	Run: func(cmd *cobra.Command, args []string) {
		// Read host and port from environment variables, with defaults if not set
		host := os.Getenv("SSP_HOST")
		port := os.Getenv("SSP_PORT")

		fmt.Printf("Starting SSP server at %s:%s\n", host, port)
		err := api.StartSSP(host, port)
		if err != nil {
			fmt.Printf("Error starting SSP server: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(startSSPCmd)
}
