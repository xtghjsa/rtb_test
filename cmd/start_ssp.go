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

func startSSPExec(cmd *cobra.Command, args []string) {
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

	//Init kafka writer
	kfkWriter := repository.NewWriterKafka()
	defer kfkWriter.Close()
	//Init kafka reader
	kfkReader := repository.NewReaderKafka()
	defer kfkReader.Close()

	//Load auction config
	cfg, err := utils.LoadSSPConfig()
	if err != nil {
		fmt.Printf("Error loading ssp config: %v\n", err)
		os.Exit(1)
	}
	//Checks for messages in kafka, loads them into postgres database
	go func() {
		repository.SspAddDeals(db, kfkReader)
	}()
	//Starts SSP server
	err = api.StartSSP(envList.SSPHost, envList.SSPPort, kfkWriter, cfg)
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
