// This package handles the actual database operations.
package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	data       = make(map[string]string)
	timestamps = make(map[string]Timestamp)
	mu         sync.RWMutex
	dataFolder = "data"                          // Folder to store data files
	dataFile   = dataFolder + "/data.json"       // File to store data
	tsFile     = dataFolder + "/timestamps.json" // File to store timestamps
)

type Timestamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	// Ensure that the "data" folder exists
	if _, err := os.Stat(dataFolder); os.IsNotExist(err) {
		os.Mkdir(dataFolder, os.ModePerm)
	}
	loadFromDisk() // Load data when the application starts
}

func SaveToDisk() {
	mu.RLock()
	defer mu.RUnlock()

	//----------Save Data file----------
	dataFile, err := os.Create(dataFile)
	if err != nil {
		// Handle the error
		fmt.Println(err)
		return
	}
	defer dataFile.Close()

	encoder := json.NewEncoder(dataFile)
	if err := encoder.Encode(data); err != nil {
		// Handle the error
		fmt.Println(err)
	}

	//----------Save TimeStamp file----------
	tsFile, err := os.Create(tsFile)
	if err != nil {
		// Handle the error
		fmt.Println(err)
		return
	}
	defer tsFile.Close()

	tsEncoder := json.NewEncoder(tsFile)
	if err := tsEncoder.Encode(timestamps); err != nil {
		// Handle the error
		fmt.Println(err)
	}
}

func loadFromDisk() {
	mu.Lock()
	defer mu.Unlock()

	//----------Load Data file----------
	dataFile, err := os.Open(dataFile)
	if err != nil {
		// Handle the error
		fmt.Println(err)
		return
	}
	defer dataFile.Close()

	dataDecoder := json.NewDecoder(dataFile)
	if err := dataDecoder.Decode(&data); err != nil {
		// Handle the error
		fmt.Println(err)
	}

	//----------Load TimeStamp file----------
	tsFile, err := os.Open(tsFile)
	if err != nil {
		// Handle the error
		fmt.Println(err)
		return
	}
	defer tsFile.Close()

	tsDecoder := json.NewDecoder(tsFile)
	if err := tsDecoder.Decode(&timestamps); err != nil {
		// Handle the error
		fmt.Println(err)
	}
}

func Set(key, value string) error {
	mu.Lock()
	defer mu.Unlock()

	// Check if the key exists to determine if it's an update
	_, exists := data[key]

	// Update or create the key-value pair
	data[key] = value

	if exists {
		// If the key existed before, update the "UpdatedAt" timestamp
		timestamp := timestamps[key]
		timestamp.UpdatedAt = time.Now()
		timestamps[key] = timestamp
	} else {
		// If it's a new key, set the "CreatedAt" and "UpdatedAt" timestamps
		timestamps[key] = Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	return nil
}

func Get(key string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	if value, ok := data[key]; ok {
		return value, nil
	}

	return "", errors.New("key not found")
}

func Delete(key string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := data[key]; !ok {
		return errors.New("key not found")
	}

	delete(data, key)
	delete(timestamps, key)

	return nil
}

func Timestamps(key string) (time.Time, time.Time, error) {
	mu.RLock()
	defer mu.RUnlock()

	timestamps, ok := timestamps[key]
	if !ok {
		return time.Time{}, time.Time{}, errors.New("key not found")
	}

	return timestamps.CreatedAt, timestamps.UpdatedAt, nil
}
