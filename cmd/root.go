package cmd

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func Execute() {
	var serverURL, privateToken, projectID string

	// Define command-line flags
	flag.StringVar(&serverURL, "server", "", "GitLab server URL (e.g., https://gitlab.com)")
	flag.StringVar(&privateToken, "token", "", "GitLab private token")
	flag.StringVar(&projectID, "project-id", "", "GitLab project ID")
	help := flag.Bool("help", false, "gitlab-artifacts-cleaner help command")
	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage: %s [flags] [<options>]\n", "gitlab-artifacts-cleaner")
		fmt.Println("Flags:")
		flag.VisitAll(func(f *flag.Flag) {
			padding := strings.Repeat(" ", 14-len(f.Name))
			fmt.Printf("  --%s%s%s\n", f.Name, padding, f.Usage)
		})
	}

	if serverURL == "" || privateToken == "" || projectID == "" {
		fmt.Println("Error: Missing required arguments.")
		flag.Usage()
		os.Exit(1)
	}

	// Show usage instructions if help flag is set
	if *help {
		flag.Usage()
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
		log.Printf("Failed to delete artifacts. Status code: %d\n", resp.StatusCode)
	}
}
