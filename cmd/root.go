// This paclage sets up the main command for our CLI.
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "kvdb"}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v\n", err)
		os.Exit(1)
	}
}
