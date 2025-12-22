package util

// InitGit 初始化 Git 仓库
func InitGit(appName string) error {
	// git init with main as default branch
	if err := RunCommandSilent(appName, "git", "init", "-b", "main"); err != nil {
		return err
	}

	// git add .
	if err := RunCommandSilent(appName, "git", "add", "."); err != nil {
		return err
	}

	return nil
}
