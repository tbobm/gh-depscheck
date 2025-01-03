package main

import (
	"fmt"

	"github.com/tbobm/gh-depscheck/pkg/depscheck"
)

func main() {
	tags, err := depscheck.GetLatestTags("actions/checkout")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Tags:", tags)
	localWorkflow, err := depscheck.LoadWorkflow("workflow.yml")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Loaded Workflow: %+v\n", localWorkflow)
	outdated := depscheck.CompareActionVersions(localWorkflow)
	fmt.Println("Outdated for archiver Actions:", outdated)
}
