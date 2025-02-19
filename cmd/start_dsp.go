package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"test_project/internal/api"

	"github.com/spf13/cobra"
)

var DBConn *sql.DB

func SetDBConn(db *sql.DB) {
	DBConn = db
}

var startDSPCmd = &cobra.Command{
	Use:   "startdsp",
	Short: "Start DSP server",
	Long:  "Start DSP server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting DSP server")
		host := os.Getenv("DSP_HOST")
		port := os.Getenv("DSP_PORT")

		err := api.StartDSP(host, port, DBConn)
		if err != nil {
			fmt.Printf("Error starting DSP server: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(startDSPCmd)
}
