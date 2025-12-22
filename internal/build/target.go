package build

import (
	"fmt"
	"strings"
)

// 常见的目标平台
var CommonTargets = []string{
	"linux/amd64",
	"linux/arm64",
	"darwin/amd64",
	"darwin/arm64",
	"windows/amd64",
}

// ParseTarget 解析目标平台字符串
func ParseTarget(target string) (os, arch string, err error) {
	parts := strings.Split(target, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid target format '%s', expected format: <os>/<arch>", target)
	}
	return parts[0], parts[1], nil
}

// ValidateTarget 验证目标平台
func ValidateTarget(os, arch string) error {
	// 这里可以添加更严格的验证逻辑
	if os == "" || arch == "" {
		return fmt.Errorf("OS and architecture cannot be empty")
	}
	return nil
}
