// This file defines the subcommands (set, get, del, ts) for our CLI.
package cmd

import (
	"fmt"
	"kvdb/database"
	"sync"

	"github.com/spf13/cobra"
)

// Define a WaitGroup to wait for all tasks to complete
var wg sync.WaitGroup

func init() {
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(delCmd)
	rootCmd.AddCommand(tsCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a key-value pair",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		err := database.Set(key, value)
		if err != nil {
			// Handle the error
			fmt.Println(err)
		}
		database.SaveToDisk() // Save data after setting a key
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value for a key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value, err := database.Get(key)
		if err != nil {
			// Handle the error
			fmt.Println(err)
		}
		// Print or use the value
		fmt.Println(value)
	},
}

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		err := database.Delete(key)
		if err != nil {
			// Handle the error
			fmt.Println(err)
		}
		database.SaveToDisk() // Save data after deleting a key
	},
}

var tsCmd = &cobra.Command{
	Use:   "ts",
	Short: "Get timestamps for a key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		createdAt, updatedAt, err := database.Timestamps(key)
		if err != nil {
			// Handle the error
			fmt.Println(err)
		} else {
			// Print or use the timestamps
			fmt.Println(createdAt)
			fmt.Println(updatedAt)
		}
	},
}
