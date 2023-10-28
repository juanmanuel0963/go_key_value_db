// This package handles the actual database operations.
package database

import (
	"bufio"
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

// Both data and timestamps are read from and written to their respective files using buffered I/O with the bufio package,
// reducing the overhead of frequent disk I/O operations.
// Buffered writers and scanners are used to write and read data to and from files, reducing the overhead of frequent disk I/O operations.

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

	dataWriter := bufio.NewWriter(dataFile)
	encoder := json.NewEncoder(dataWriter)
	if err := encoder.Encode(data); err != nil {
		// Handle the error
		fmt.Println(err)
	}
	//The Flush method is called to ensure that the buffered data is written to the file.
	dataWriter.Flush()

	//----------Save TimeStamp file----------
	tsFile, err := os.Create(tsFile)
	if err != nil {
		// Handle the error
		fmt.Println(err)
		return
	}
	defer tsFile.Close()

	tsWriter := bufio.NewWriter(tsFile)
	tsEncoder := json.NewEncoder(tsWriter)
	if err := tsEncoder.Encode(timestamps); err != nil {
		// Handle the error
		fmt.Println(err)
	}
	//The Flush method is called to ensure that the buffered data is written to the file.
	tsWriter.Flush()
}

func loadFromDisk() {
	mu.Lock()
	defer mu.Unlock()

	//----------Load Data file----------
	dataFile, err := os.Open(dataFile)
	if err != nil {
		// Handle the error
		//fmt.Println(err)
		return
	}
	defer dataFile.Close()

	dataScanner := bufio.NewScanner(dataFile)
	for dataScanner.Scan() {
		line := dataScanner.Text()
		var entry map[string]string
		if err := json.Unmarshal([]byte(line), &entry); err == nil {
			for key, value := range entry {
				data[key] = value
			}
		}
	}

	//----------Load TimeStamp file----------
	tsFile, err := os.Open(tsFile)
	if err != nil {
		// Handle the error
		//fmt.Println(err)
		return
	}
	defer tsFile.Close()

	tsScanner := bufio.NewScanner(tsFile)
	for tsScanner.Scan() {
		line := tsScanner.Text()
		var tsEntry map[string]Timestamp
		if err := json.Unmarshal([]byte(line), &tsEntry); err == nil {
			for key, ts := range tsEntry {
				timestamps[key] = ts
			}
		}
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
