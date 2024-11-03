package main

import (
	"fmt"
	"path"
	"strings"
	"time"
)

const (
	// Color codes
	colorReset  = "\033[0m"
	colorBlue   = "\033[1;34m"
	colorGreen  = "\033[1;32m"
	colorYellow = "\033[1;33m"
	colorRed    = "\033[1;31m"
	colorGray   = "\033[1;30m"
	colorCyan   = "\033[1;36m"

	// Symbols
	symbolFile    = "📁"
	symbolSuccess = "✓"
	symbolError   = "✗"
	symbolWatch   = "🐶"

	// Templates
	startupTemplate = `
%s%s Watcher Started%s
  └── Time: %s
  └── Command: %s
  └── Refresh: %s
  └── Watching %d file(s):
%s`

	fileListTemplate   = "      └──  %s%s%s"
	fileChangeTemplate = `%s%s [%s] %s
  └── Modified at: %s
  └── Reloading...%s
`
	commandStartTemplate = `%s  └── Running: %s%s
`
	commandOutputTemplate = `%s      │ %s%s
`
	commandEndTemplate = `%s      └── %s (took %s)%s
`
)

type Logger struct {
	silent bool
}

func NewLogger(silent bool) *Logger {
	return &Logger{silent: silent}
}

func (l *Logger) LogFileChange(file string, modTime time.Time) {
	if l.silent {
		return
	}

	fmt.Printf(fileChangeTemplate,
		colorBlue,
		symbolFile,
		time.Now().Format("15:04:05"),
		path.Base(file),
		modTime.Format("15:04:05.000"),
		colorReset,
	)
}

func (l *Logger) LogStartup(files []string, command string, args []string, interval time.Duration) {
	if l.silent {
		return
	}

	// Format file list
	var fileList string
	for _, file := range files {
		fileList += fmt.Sprintf(fileListTemplate,
			colorCyan,
			file,
			colorReset) + "\n"
	}

	// Format full command
	fullCommand := fmt.Sprintf("%s %s", command, strings.Join(args, " "))

	fmt.Printf(startupTemplate,
		colorGreen,
		symbolWatch,
		colorReset,
		time.Now().Format("2006-01-02 15:04:05"),
		fullCommand,
		interval.String(),
		len(files),
		fileList,
	)
}

func (l *Logger) LogCommandStart(command string, args []string) {
	if l.silent {
		return
	}

	fullCommand := fmt.Sprintf("%s %s", command, strings.Join(args, " "))
	fmt.Printf(commandStartTemplate,
		colorBlue,
		fullCommand,
		colorReset,
	)
}

func (l *Logger) LogCommandOutput(success bool, output string) {
	if l.silent {
		return
	}

	var color = colorGreen

	if !success {
		color = colorRed
	}

	// Split output into lines and format each line
	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		if line != "" {
			fmt.Printf(commandOutputTemplate,
				color,
				line,
				colorReset,
			)
		}
	}
}

func (l *Logger) LogCommandEnd(success bool, duration time.Duration) {
	if l.silent {
		return
	}

	status := fmt.Sprintf("%s Completed", symbolSuccess)
	color := colorGreen

	if !success {
		status = fmt.Sprintf("%s Failed", symbolError)
		color = colorRed
	}

	fmt.Printf(commandEndTemplate,
		color,
		status,
		formatDuration(duration),
		colorReset,
	)
}

// formatDuration formats duration in a human-readable way
func formatDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%dms", d.Milliseconds())
	}
	return fmt.Sprintf("%.1fs", d.Seconds())
}
