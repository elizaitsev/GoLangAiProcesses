package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strings"
	"github.com/shirou/gopsutil/process"
)

// ServiceResponse represents the response from the API.
type ServiceResponse struct {
	ProcessName string `json:"process_name"`
	Response    string `json:"response"`
}

func main() {
	// Retrieve a list of all running processes
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Map to track encountered service names
	encounteredProcess := make(map[string]bool)

	// Loop through each service and make API calls
	for _, p := range processes {
		// Get the name of the process
		name, _ := p.Name()

		// Increment the count for the encountered process name
		// Check if the service name has already been encountered
		if !encounteredProcess[name] {
			encounteredProcess[name] = true

			// Send API call for the service
			apiURL := "http://www.myapilink.com/API"
			query := fmt.Sprintf("%s=%s", "process_name", name)
			fullURL := apiURL + "?" + query

			response, err := http.Get(fullURL)
			if err != nil {
				fmt.Printf("Error sending API call for %s: %v\n", name, err)
				continue
			}
			defer response.Body.Close()

			// Parse the API response
			var apiResponse ServiceResponse
			err = json.NewDecoder(response.Body).Decode(&apiResponse)
			if err != nil {
				fmt.Printf("Error decoding API response for %s: %v\n", name, err)
				continue
			}

			// Print the API response
			fmt.Printf("Service: %s, API Response: %s\n", apiResponse.ProcessName, apiResponse.Response)
		}
	}
}
