package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tbobm/gh-depscheck/pkg/depscheck"
)

func main() {
	pflag.String("workflowfile", "workflow.yml", "Target Github Actions Workflow manifest")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	workflow := viper.GetString("workflowfile")
	localWorkflow, err := depscheck.LoadWorkflow(workflow)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	outdated := depscheck.CompareActionVersions(localWorkflow)

	fmt.Println("Jobs with outdated Actions for Workflow", localWorkflow.Name, ":", len(outdated))

	if len(outdated) > 0 {
		for job, action := range outdated {
			fmt.Println("~> Job", job, "actions:", action)
		}
	}
}
