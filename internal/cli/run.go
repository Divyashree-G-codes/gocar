package cli

import (
	"fmt"
	"os"
	"os/exec"

	"gocar/internal/project"
)

// RunCommand run 命令
type RunCommand struct{}

// Run 执行 run 命令
func (c *RunCommand) Run(args []string) error {
	// Get project info
	projectRoot, appName, projectMode, err := project.DetectProject()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Determine source path based on project mode
	var sourcePath string
	if projectMode == "project" {
		sourcePath = "./cmd/server"
	} else {
		sourcePath = "."
	}

	fmt.Printf("Running %s...\n\n", appName)

	runArgs := []string{"run", sourcePath}
	runArgs = append(runArgs, args...)

	cmd := exec.Command("go", runArgs...)
	cmd.Dir = projectRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		// Don't print error for normal exit
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		fmt.Printf("Run failed: %v\n", err)
		os.Exit(1)
	}

	return nil
}

// Help 返回帮助信息
func (c *RunCommand) Help() string {
	return `gocar run - Run the project

USAGE:
    gocar run [args...]

EXAMPLES:
    gocar run                Run the project
    gocar run --help         Pass --help to the application
`
}
