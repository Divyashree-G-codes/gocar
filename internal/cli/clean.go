package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"gocar/internal/project"
)

// CleanCommand clean 命令
type CleanCommand struct{}

// Run 执行 clean 命令
func (c *CleanCommand) Run(args []string) error {
	projectRoot, appName, _, err := project.DetectProject()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	binDir := filepath.Join(projectRoot, "bin")

	// Remove bin directory contents
	entries, err := os.ReadDir(binDir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Nothing to clean.")
			return nil
		}
		fmt.Printf("Error reading bin directory: %v\n", err)
		os.Exit(1)
	}

	if len(entries) == 0 {
		fmt.Println("Nothing to clean.")
		return nil
	}

	for _, entry := range entries {
		path := filepath.Join(binDir, entry.Name())
		if err := os.RemoveAll(path); err != nil {
			fmt.Printf("Error removing %s: %v\n", path, err)
		}
	}

	fmt.Printf("Cleaned build artifacts for '%s'\n", appName)
	return nil
}

// Help 返回帮助信息
func (c *CleanCommand) Help() string {
	return `gocar clean - Clean build artifacts

USAGE:
    gocar clean

DESCRIPTION:
    Remove all build artifacts from the bin/ directory.
`
}
