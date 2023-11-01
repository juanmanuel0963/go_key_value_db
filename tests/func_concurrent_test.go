package test

import (
	"errors"
	"kvdb/database"
	"strconv"
	"sync"
	"testing"
)

// TestConcurrentAccessToFunc tests concurrent access to the key-value database calling the functions, not the binary exec file
func TestConcurrentAccessToFunc(t *testing.T) {
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

			key := "key" + strconv.Itoa(i)
			value := "value" + strconv.Itoa(i)

			//----------Set----------

			// Set a key-value pair concurrently
			err := database.Set(key, value)
			if err != nil {
				errCh <- err
				return
			}
			// Persist data to disk
			database.SaveToDisk()

			//----------Get----------

			// Get the value and verify it
			val, err := database.Get(key)
			if err != nil {
				errCh <- err
				return
			}
			if val != value {
				errCh <- errors.New("Mismatched value")
				return
			}

			//----------Modify existing----------

			// Modify the key's value
			newValue := "new_value" + strconv.Itoa(i)
			err = database.Set(key, newValue)
			if err != nil {
				errCh <- err
				return
			}
			// Persist data to disk
			database.SaveToDisk()

			// Get the modified value and verify it
			modifiedVal, err := database.Get(key)
			if err != nil {
				errCh <- err
				return
			}
			if modifiedVal != newValue {
				errCh <- errors.New("Mismatched modified value")
				return
			}

			//----------Delete----------
			/*
				// Delete the key
				err = database.Delete(key)
				if err != nil {
					errCh <- err
					return
				}
				// Persist data to disk
				database.SaveToDisk()

				// Verify that the key is deleted
				_, err = database.Get(key)
				if err == nil {
					errCh <- errors.New("Key not deleted")
					return
				}
			*/
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
