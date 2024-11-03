package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)


type FileStat struct {
	exists   bool
	path     string
	modified time.Time
}

// Run a command and return output
func execute(cmdArg []string) (string, error) {
	cmd := exec.Command(cmdArg[0], cmdArg[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
    return "", fmt.Errorf("Command failed: %v\n%s", err, string(out)) //TODO: USE LOGGER
	}

	return string(out), err
}

// Create a [FileStat] from a file
func checkFile(path string) (*FileStat, error) {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &FileStat{
				path:   path,
				exists: false,
			}, nil
		}
		return nil, err
	}

	return &FileStat{
		path:     path,
		modified: stat.ModTime(),
		exists:   true,
	}, nil
}

