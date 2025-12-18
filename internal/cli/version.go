package cli

import "fmt"

// VersionCommand version 命令
type VersionCommand struct{}

// Run 执行 version 命令
func (c *VersionCommand) Run(args []string) error {
	fmt.Printf("gocar %s\n", Version)
	return nil
}

// Help 返回帮助信息
func (c *VersionCommand) Help() string {
	return `gocar version - Print version info

USAGE:
    gocar version
`
}
