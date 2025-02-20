package cmd

import (
	"fmt"
	"log"
	"os"
	"test_project/internal/api"
	"test_project/internal/repository"
	"test_project/internal/utils"
	"test_project/pkg"

	"github.com/spf13/cobra"
)

func startDSPExec(cmd *cobra.Command, args []string) {
	fmt.Println("Starting DSP server")
	envList, err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("error loading env variables: %v\n", err)
	}
	db, err := pkg.InitDspDatabase(envList.PostgresUser, envList.PostgresPass, envList.PostgresDBName, envList.PostgresHost, envList.PostgresPort)
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}
	err = repository.AddTestAds(db)
	if err != nil {
		log.Fatalf("Error adding test ads into dsp database: %v", err)
	}
	err = api.StartDSP(envList.DSPHost, envList.DSPPort, db)
	if err != nil {
		fmt.Printf("Error starting DSP server: %v", err)
		os.Exit(1)
	}
}

var startDSPCmd = &cobra.Command{
	Use:   "startdsp",
	Short: "Start DSP server",
	Long:  "Start DSP server",
	Run:   startDSPExec,
}

func init() {
	rootCmd.AddCommand(startDSPCmd)
}
