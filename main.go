package main

import (
	"fmt"
	"magshare/cmd"
	"magshare/internal/logger"
	"magshare/internal/ui"
	"magshare/internal/workspace"
	"os"
)

func main() {
	// 1. Initialize Workspace
	isFirstRun, err := workspace.InitializeWorkspace()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing workspace: %v\n", err)
		os.Exit(1)
	}

	// 2. Setup Logging
	logsDir, _ := workspace.GetLogsDir()
	logPath, cleanup, err := logger.SetupLogging(logsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting up logging: %v\n", err)
	}

	l := logger.WithComponent("main")
	l.Debug("Application starting...")

	var exitCode int
	// panicOccurred := true // Not needed if we use the l instance

	// 3. Crash Recovery and Cleanup
	defer func() {
		// Recover from panic if any
		if r := recover(); r != nil {
			logger.HandlePanic(logPath, r)
			os.Exit(1)
		}

		// Cleanup redirection
		if cleanup != nil {
			cleanup()
		}

		// Delete log if exit was successful (exitCode == 0) and no panic
		if exitCode == 0 && logPath != "" {
			logger.CleanupLogs(logPath)
		}

		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}()

	// 4. Welcome Message
	if isFirstRun {
		l.Info("First run detected, displaying welcome message")
		ui.DisplayWelcomeMessage(os.Stdout)
	}

	// 5. Execute Command
	if err := cmd.Execute(); err != nil {
		l.Error(fmt.Sprintf("Command execution failed: %v", err))
		exitCode = 1
	} else {
		l.Debug("Command execution successful")
		exitCode = 0
	}
}
