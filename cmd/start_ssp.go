package cmd

import (
	"fmt"
	"log"
	"os"
	"test_project/internal/api"
	"test_project/internal/utils"

	"github.com/spf13/cobra"
)

func startSSPExec(cmd *cobra.Command, args []string) {
	envList, err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("error loading env variables: %v\n", err)
	}

	err = api.StartSSP(envList.SSPHost, envList.SSPPort)
	if err != nil {
		fmt.Printf("Error starting SSP server: %v\n", err)
		os.Exit(1)
	}
}

var startSSPCmd = &cobra.Command{
	Use:   "startssp",
	Short: "Start SSP server",
	Long:  "Start SSP server",
	Run:   startSSPExec,
}

func init() {
	rootCmd.AddCommand(startSSPCmd)
}
