package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	// Get current time
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)

	resp, err := http.Get("https://worldtimeapi.org/api/ip")
	if err != nil {
		log.Fatal("Error fetching time from API:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status: %d", resp.StatusCode)
		return
	}

	type timeResponse struct {
		Datetime string `json:"datetime"`
	}

	var ct timeResponse

	// Decode JSON response and populate the struct instance
	err = json.NewDecoder(resp.Body).Decode(&ct)
	if err != nil {
		log.Fatal("Error decoding JSON response:", err)
		return
	}

	parsedTime, err := time.Parse(time.RFC3339, ct.Datetime)
	if err != nil {
		log.Fatal("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed API time:", parsedTime)

	if !currentTime.Equal(parsedTime) {
		cmd := exec.Command("sudo", "date", "--set", parsedTime.Format("2006-01-02 15:04:05"))

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error setting date and time:", string(output), err)
			log.Fatal(err)
		} else {
			fmt.Println("Date and time corrected successfully!")
		}
	} else {
		fmt.Println("Date and time already correct")
	}
}
