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
	//Load env variables
	envList, err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("error loading env variables: %v\n", err)
	}
	//Init database
	db, err := pkg.InitDatabase(envList.PostgresUser, envList.PostgresPass, envList.PostgresDBName, envList.PostgresHost, envList.PostgresPort)
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}
	//Adds test ads into the database
	err = repository.AddTestAds(db)
	if err != nil {
		log.Fatalf("Error adding test ads into dsp database: %v", err)
	}
	//Starts DSP server
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
