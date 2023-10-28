package database

import (
	"errors"
	"strconv"
	"sync"
	"testing"
)

// TestConcurrentAccess tests concurrent access to the key-value database.
func TestConcurrentAccess(t *testing.T) {
	// Number of concurrent goroutines
	numGoroutines := 100

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Channel to collect errors from goroutines
	errCh := make(chan error, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			key := "key" + strconv.Itoa(i)
			value := "value" + strconv.Itoa(i)

			// Set a key-value pair concurrently
			err := Set(key, value)
			if err != nil {
				errCh <- err
				return
			}
			// Persist data to disk
			SaveToDisk()

			// Get the value and verify it
			val, err := Get(key)
			if err != nil {
				errCh <- err
				return
			}
			if val != value {
				errCh <- errors.New("Mismatched value")
				return
			}

			// Modify the key's value
			newValue := "new_value" + strconv.Itoa(i)
			err = Set(key, newValue)
			if err != nil {
				errCh <- err
				return
			}
			// Persist data to disk
			SaveToDisk()

			// Get the modified value and verify it
			modifiedVal, err := Get(key)
			if err != nil {
				errCh <- err
				return
			}
			if modifiedVal != newValue {
				errCh <- errors.New("Mismatched modified value")
				return
			}

			// Delete the key
			err = Delete(key)
			if err != nil {
				errCh <- err
				return
			}
			// Persist data to disk
			SaveToDisk()

			// Verify that the key is deleted
			_, err = Get(key)
			if err == nil {
				errCh <- errors.New("Key not deleted")
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
