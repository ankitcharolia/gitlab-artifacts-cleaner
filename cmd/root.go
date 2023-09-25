package cmd

import (
	"flag"
	"fmt"
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

	// Show usage instructions if help flag is set
	if *help {
		flag.Usage()
	}
}
