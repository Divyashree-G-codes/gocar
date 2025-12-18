package util

import (
	"fmt"
	"os"
)

// WriteFile 写入文件内容
func WriteFile(path, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}
	return nil
}

// FileExists 检查文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DirExists 检查目录是否存在
func DirExists(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.IsDir()
}

// EnsureDir 确保目录存在
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}
