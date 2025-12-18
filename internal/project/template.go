package project

import "fmt"

// SimpleMainTemplate 生成简单项目的 main.go 内容
func SimpleMainTemplate(appName string) string {
	return fmt.Sprintf(`package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, gocar! A golang project scaffolding tool for %s.")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
`, appName)
}

// ProjectMainTemplate 生成项目模式的 main.go 内容
func ProjectMainTemplate(appName string) string {
	return fmt.Sprintf(`package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, gocar! A golang project scaffolding tool for %s.")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
`, appName)
}

// SimpleReadmeTemplate 生成简单项目的 README.md 内容
func SimpleReadmeTemplate(appName string) string {
	return fmt.Sprintf(`# %s

A Go project created with gocar.

## Build

`+"```bash"+`
# Debug build (current platform)
gocar build

# Release build (current platform)
gocar build --release

# Cross-compile for Linux on AMD64
gocar build --target linux/amd64
`+"```"+`

## Run

`+"```bash"+`
gocar run
`+"```"+`

## Output Structure

`+"```"+`
bin/
├── debug/
│   └── <os>-<arch>/
│       └── %s
└── release/
    └── <os>-<arch>/
        └── %s
`+"```"+`

Build artifacts are organized by:
- **Build mode**: debug or release
- **Target platform**: OS and architecture (e.g., linux-amd64, darwin-arm64)

Examples:
- Debug build for current platform: `+"`./bin/debug/linux-amd64/%s`"+`
- Release build for Windows: `+"`./bin/release/windows-amd64/%s.exe`"+`
`, appName, appName, appName, appName, appName)
}

// ProjectReadmeTemplate 生成项目模式的 README.md 内容
func ProjectReadmeTemplate(appName string) string {
	return fmt.Sprintf(`# %s

A Go project created with gocar (project mode).

## Project Structure

`+"```"+`
%s/
├── cmd/
│   └── server/
│       └── main.go      # Application entry point
├── internal/            # Private application code
├── pkg/                 # Public library code
├── test/                # Integration tests
├── bin/                 # Build output
├── go.mod
└── README.md
`+"```"+`

## Build

`+"```bash"+`
# Debug build (current platform)
gocar build

# Release build (current platform)
gocar build --release

# Cross-compile for Linux
gocar build --target linux/amd64
`+"```"+`

## Run

`+"```bash"+`
gocar run
`+"```"+`

## Output Structure

`+"```"+`
bin/
├── debug/
│   └── <os>-<arch>/
│       └── %s
└── release/
    └── <os>-<arch>/
        └── %s
`+"```"+`

Build artifacts are organized by:
- **Build mode**: debug or release
- **Target platform**: OS and architecture (e.g., linux-amd64, darwin-arm64)

Examples:
- Debug build for current platform: `+"`./bin/debug/linux-amd64/%s`"+`
- Release build for Windows: `+"`./bin/release/windows-amd64/%s.exe`"+`

## Directories

- **cmd/**: Main applications for this project
- **internal/**: Private application and library code (not importable by other projects)
- **pkg/**: Library code that can be used by external applications
- **test/**: Integration tests, black-box tests
`, appName, appName, appName, appName, appName, appName)
}

// GitignoreTemplate 生成 .gitignore 内容
func GitignoreTemplate(appName string) string {
	return fmt.Sprintf(`# Binaries
%s
bin/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output of go coverage
*.out

# Dependency directories
vendor/

# IDE
.idea/
.vscode/
*.swp
*.swo

# OS files
.DS_Store
Thumbs.db
`, appName)
}
