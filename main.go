package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

// getPID uses pgrep to find the PID of the process by name.
// Returns the first PID found or -1 if the process is not found
func getPID(processName string) int {
	cmd := exec.Command("pgrep", processName)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println("Error creating StdoutPipe for Cmd", err)
		return -1
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting Cmd", err)
		return -1
	}

	scanner := bufio.NewScanner(stdout)
	if scanner.Scan() {
		pidStr := scanner.Text()
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			fmt.Println("Error converting PID to int", err)
			return -1
		}
		return pid
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for Cmd", err)
		return -1
	}

	return -1
}

func main() {
	processName := "ControlCenter"
	var lastPID int = -1

	for {
		currentPID := getPID(processName)

		if (lastPID != -1 && currentPID != -1 && lastPID != currentPID) || (lastPID == -1 && currentPID != -1) {
			fmt.Printf("Process '%s' has started or restarted. New PID: %d\n", processName, currentPID)
		}

		lastPID = currentPID

		time.Sleep(2 * time.Second)
	}
}
