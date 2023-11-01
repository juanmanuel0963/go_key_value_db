package test

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"testing"
)

// TestConcurrentAccessToExec tests concurrent access to the key-value database calling the binary exec kvdb file
func TestConcurrentAccessToExec(t *testing.T) {
	// Number of concurrent goroutines
	numGoroutines := 50

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Channel to collect errors from goroutines
	errCh := make(chan error, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Simulate concurrent calls to the kvdb application
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)

			setCmd := exec.Command("../kvdb", "set", key, value)
			getCmd := exec.Command("../kvdb", "get", key)
			//delCmd := exec.Command("kvdb", "del", key)
			tsCmd := exec.Command("../kvdb", "ts", key)

			setCmd.Stdout = os.Stdout
			getCmd.Stdout = os.Stdout
			//delCmd.Stdout = os.Stdout
			tsCmd.Stdout = os.Stdout

			if err := setCmd.Run(); err != nil {
				fmt.Printf("Error calling 'set': %v\n", err)
				errCh <- err
				return
			}

			if err := getCmd.Run(); err != nil {
				fmt.Printf("Error calling 'get': %v\n", err)
				errCh <- err
				return
			}
			/*
				if err := delCmd.Run(); err != nil {
					fmt.Printf("Error calling 'del': %v\n", err)
					errCh <- err
					return
				}
			*/
			if err := tsCmd.Run(); err != nil {
				fmt.Printf("Error calling 'ts': %v\n", err)
				errCh <- err
				return
			}

		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errCh)

	// Check for any errors reported by goroutines
	for err := range errCh {
		t.Errorf("Concurrent test error: %v", err)
	}
}
