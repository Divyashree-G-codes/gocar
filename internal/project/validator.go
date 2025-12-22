package project

import (
	"fmt"
	"regexp"
	"strings"
)

// 保留的 Go 关键字
var reservedNames = []string{"test", "main", "init", "internal", "vendor"}

// ValidateProjectName 验证项目名称
func ValidateProjectName(name string) error {
	// Check if empty
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	// Check if starts with a dash or dot
	if strings.HasPrefix(name, "-") || strings.HasPrefix(name, ".") {
		return fmt.Errorf("project name cannot start with '-' or '.'")
	}

	// Check for valid characters (alphanumeric, dash, underscore)
	validName := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)
	if !validName.MatchString(name) {
		return fmt.Errorf("project name must start with a letter and contain only letters, numbers, dashes, or underscores")
	}

	// Check for reserved names
	for _, r := range reservedNames {
		if strings.ToLower(name) == r {
			return fmt.Errorf("'%s' is a reserved name in Go", name)
		}
	}

	return nil
}
