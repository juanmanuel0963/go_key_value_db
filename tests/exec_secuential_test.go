package test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// TestSecuentialAccessToExec tests secuential access to the key-value database calling the binary exec kvdb file
func TestSecuentialAccessToExec(t *testing.T) {

	// Number of requests
	max := 50

	for i := 0; i < max; i++ {

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
			t.Errorf("Error calling 'set': %v\n", err)
		}

		if err := getCmd.Run(); err != nil {
			t.Errorf("Error calling 'get': %v\n", err)
		}
		/*
			if err := delCmd.Run(); err != nil {
				t.Errorf("Error calling 'del': %v\n", err)
			}
		*/
		if err := tsCmd.Run(); err != nil {
			t.Errorf("Error calling 'ts': %v\n", err)
		}

	}
}
