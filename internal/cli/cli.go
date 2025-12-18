package cli

import (
	"fmt"
)

// Version 版本号
const Version = "0.1.3"

// App CLI 应用
type App struct {
	commands map[string]Command
}

// Command 命令接口
type Command interface {
	Run(args []string) error
	Help() string
}

// NewApp 创建 CLI 应用
func NewApp() *App {
	app := &App{
		commands: make(map[string]Command),
	}

	// 注册命令
	app.commands["new"] = &NewCommand{}
	app.commands["build"] = &BuildCommand{}
	app.commands["run"] = &RunCommand{}
	app.commands["clean"] = &CleanCommand{}
	app.commands["add"] = &AddCommand{}
	app.commands["update"] = &UpdateCommand{}
	app.commands["tidy"] = &TidyCommand{}

	return app
}

// Run 运行应用
func (a *App) Run(args []string) error {
	if len(args) < 2 {
		printHelp()
		return nil
	}

	cmdName := args[1]

	// 处理特殊命令
	switch cmdName {
	case "help", "-h", "--help":
		printHelp()
		return nil
	case "version", "-v", "--version":
		fmt.Printf("gocar %s\n", Version)
		return nil
	}

	// 执行命令
	cmd, ok := a.commands[cmdName]
	if !ok {
		fmt.Printf("Unknown command: %s\n", cmdName)
		printHelp()
		return fmt.Errorf("unknown command: %s", cmdName)
	}

	return cmd.Run(args[2:])
}

// printHelp 打印帮助信息
func printHelp() {
	help := `gocar - A cargo-like tool for Go projects

USAGE:
    gocar <COMMAND> [OPTIONS]

COMMANDS:
    new <name> [--mode simple|project]     Create a new Go project
    build [--release]                      Build the project
    run [args...]                          Run the project
    clean                                  Clean build artifacts
    add <package>...                       Add dependencies to go.mod
    update [package]...                    Update dependencies
    tidy                                   Tidy up go.mod and go.sum
    help                                   Print this help message
    version                                Print version info

EXAMPLES:
    gocar new myapp                        Create a simple project
    gocar new myapp --mode project         Create a project-mode project
    gocar build                            Build in debug mode
    gocar build --release                  Build in release mode
    gocar run                              Build and run the project
    gocar add github.com/gin-gonic/gin     Add a dependency
    gocar update                           Update all dependencies
    gocar tidy                             Clean up go.mod
`
	fmt.Print(help)
}
