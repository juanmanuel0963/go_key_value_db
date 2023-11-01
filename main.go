package main

import (
	"fmt"
	"kvdb/cmd"
	"os"
)

func main() {

	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
