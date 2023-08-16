package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/shirou/gopsutil/process"
)

func main() {
	// Retrieve a list of all running processes
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a map to track encountered process names and their counts
	encounteredNames := make(map[string]int)

	// Iterate through the list of processes
	for _, p := range processes {
		// Get the name of the process
		name, _ := p.Name()

		// Increment the count for the encountered process name
		encounteredNames[name]++
	}

	// Create or open a file for writing process information
	outputFile, err := os.Create("processes.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	// Write the process information to the file
	fmt.Fprintln(outputFile, "List of running processes:")
	for name, count := range encounteredNames {
		fmt.Fprintf(outputFile, "Name: %s, Encountered: %d times\n", name, count)
	}

	// Print out the unique process names and their repetition counts
	fmt.Println("Unique process names and repetition counts:")
	uniqueNames := getSortedUniqueNames(encounteredNames)
	for _, name := range uniqueNames {
		fmt.Printf("Name: %s, Repeated: %d times\n", name, encounteredNames[name])
	}
}

// Helper function to get sorted unique process names from the map
func getSortedUniqueNames(namesMap map[string]int) []string {
	uniqueNames := make([]string, 0, len(namesMap))
	for name := range namesMap {
		uniqueNames = append(uniqueNames, name)
	}
	sort.Strings(uniqueNames)
	return uniqueNames
}
