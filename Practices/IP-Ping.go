package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

func main() {
	// Ask the user for the IP address
	fmt.Print("Enter the IP address to ping: ")
	var ip string
	fmt.Scanln(&ip)

	// Ask the user for the number of pings
	fmt.Print("Enter the number of pings: ")
	var count int
	fmt.Scanln(&count)

	// Start pinging the IP address
	pingIP(ip, count)
}

func pingIP(ip string, count int) {
	// Create a wait group to wait for all pings to finish
	var wg sync.WaitGroup

	// Start pinging the IP address for the specified number of times
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Execute the ping command based on the operating system
			var cmd *exec.Cmd
			switch runtime.GOOS {
			case "windows":
				cmd = exec.Command("ping", "-n", "1", "-w", "1000", ip)
			case "darwin", "linux":
				cmd = exec.Command("ping", "-c", "1", "-W", "1", ip)
			default:
				fmt.Println("Unsupported operating system")
				return
			}

			// Run the ping command and check the output
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
				return
			}

			// Check if the ping was successful
			if strings.Contains(string(output), "time=") {
				fmt.Println("Ping successful")
			} else {
				fmt.Println("Ping failed")
			}
		}()
	}

	// Wait for all pings to finish
	wg.Wait()
}
