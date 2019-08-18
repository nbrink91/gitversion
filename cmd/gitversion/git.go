package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func validateGitInstalled() error {
	_, err := exec.LookPath("git")

	return err
}

func validateGitRepo(basePath string) error {
	path := filepath.Join(basePath, ".git")

	stats, err := os.Stat(path)

	if err != nil || !stats.IsDir() {
		return fmt.Errorf("path %s is not a GIT repository", basePath)
	}

	return nil
}
