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
	envList, err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("error loading env variables: %v\n", err)
	}
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

	go func() {
		repository.SspAddDeals(db, kfkReader)
	}()

	err = api.StartSSP(envList.SSPHost, envList.SSPPort, kfkWriter)
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
