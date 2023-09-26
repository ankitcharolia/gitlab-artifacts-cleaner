package cmd

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Execute() error {
	var serverURL, privateToken, projectID string

	// Define command-line flags
	flag.StringVar(&serverURL, "server", "", "GitLab Server URL (e.g., https://gitlab.com)")
	flag.StringVar(&privateToken, "token", "", "GitLab Personal Access Token")
	flag.StringVar(&projectID, "project-id", "", "GitLab Project ID")
	help := flag.Bool("help", false, "gitlab-artifacts-cleaner Help Command")
	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage: %s [flags] [<options>]\n", "gitlab-artifacts-cleaner")
		fmt.Println("Flags:")
		flag.VisitAll(func(f *flag.Flag) {
			padding := strings.Repeat(" ", 14-len(f.Name))
			fmt.Printf("  --%s%s%s\n", f.Name, padding, f.Usage)
		})
	}

	// Show usage instructions if no arguments are provided
	if flag.NFlag() == 0 {
		flag.Usage()
		return nil
	} else if serverURL == "" || privateToken == "" || projectID == "" {
		flag.Usage()
		return nil
	} else if *help {
		flag.Usage()
		return nil
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
	if resp.StatusCode == http.StatusAccepted {
		fmt.Println("Artifacts deleted successfully.")
	} else {
		fmt.Printf("Failed to delete artifacts. Status code: %d\n", resp.StatusCode)
	}
	return nil
}
