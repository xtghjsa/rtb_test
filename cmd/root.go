package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rtb_test",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting RTB Test Application, type exit to stop the application")
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "exit" {
				break
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Error executing root command: %v", err)
		os.Exit(1)
	}
}
