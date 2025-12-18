package util

import (
	"fmt"
	"os/exec"
)

// RunCommandSilent 静默执行命令（不输出）
func RunCommandSilent(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	return cmd.Run()
}

// RunCommand 执行命令并输出结果
func RunCommand(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir

	// Capture output to display in case of error
	output, err := cmd.CombinedOutput()

	// Always display output (for progress messages, warnings, etc.)
	if len(output) > 0 {
		fmt.Print(string(output))
	}

	return err
}
