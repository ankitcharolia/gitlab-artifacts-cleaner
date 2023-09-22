package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var serverURL, privateToken, projectID string

	// Define command-line flags
	flag.StringVar(&serverURL, "server", "", "GitLab server URL (e.g., https://gitlab.com)")
	flag.StringVar(&privateToken, "token", "", "GitLab private token")
	flag.StringVar(&projectID, "project-id", "", "GitLab project ID")
	flag.Parse()

	// Check if required arguments are provided
	if serverURL == "" || privateToken == "" || projectID == "" {
		fmt.Println("Usage: go run main.go --server <GitLab server URL> --token <private token> --project-id <project ID>")
		os.Exit(1)
	}

	// Construct the API URL for deleting artifacts
	apiURL := fmt.Sprintf("%s/api/v4/projects/%s/artifacts", serverURL, projectID)

	// Create an HTTP client with the provided private token
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		os.Exit(1)
	}
	req.Header.Set("Private-Token", privateToken)

	// Send the DELETE request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending DELETE request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Artifacts deleted successfully.")
	} else {
		fmt.Printf("Failed to delete artifacts. Status code: %d\n", resp.StatusCode)
	}
}
