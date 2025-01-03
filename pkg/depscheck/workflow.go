package depscheck

import (
	"fmt"
	"os"
	"strings"

	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v3"
)

type Workflow struct {
	Name string         `yaml:"name"`
	Jobs map[string]Job `yaml:"jobs"`
}

type Job struct {
	Name  string `yaml:"name,omitempty"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Uses string            `yaml:"uses,omitempty"`
	With map[string]string `yaml:"with,omitempty"`
}

func LoadWorkflow(filename string) (*Workflow, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var workflow Workflow
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		return nil, err
	}

	return &workflow, nil
}

func GetLatestTags(repo string) ([]string, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/tags", repo)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch tags: %s", resp.Status)
	}

	var tags []struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return nil, err
	}

	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	return tagNames, nil
}

/*
Example Output:
If the repository is "actions/checkout", calling GetLatestTags("actions/checkout") might return:

	[]string{
		"v3",
		"v2.4.0",
		"v2.3.4",
		"v2.3.3",
	}
*/
func CompareActionVersions(workflow *Workflow) map[string][]string {
	outdatedActions := make(map[string][]string)

	for jobName, job := range workflow.Jobs {
		for _, step := range job.Steps {
			if step.Uses != "" {
				parts := strings.Split(step.Uses, "@")
				if len(parts) != 2 {
					continue // Skip steps without valid versioning
				}

				repo, currentVersion := parts[0], parts[1]
				if currentVersion == "latest" {
					continue // Consider 'latest' as up-to-date
				}

				tags, err := GetLatestTags(repo)
				if err != nil {
					fmt.Printf("Error fetching tags for %s: %v\n", repo, err)
					continue
				}

				if !isVersionUpToDate(currentVersion, tags) {
					outdatedActions[jobName] = append(outdatedActions[jobName], fmt.Sprintf("%s@%s", repo, currentVersion))
				}
			}
		}
	}

	return outdatedActions
}

func isVersionUpToDate(currentVersion string, tags []string) bool {
	for _, tag := range tags {
		if strings.HasPrefix(tag, currentVersion) {
			return true
		}
	}
	return false
}
