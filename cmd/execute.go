// This paclage sets up the main command for our CLI.
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "kvdb"}

func Execute() error {

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v\n", err)
		return err
		//os.Exit(1)
	}
	return nil
}
