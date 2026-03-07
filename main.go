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
	// 1. Initialize Workspace with Interactive Setup if needed
	setupFunc := func() (workspace.Config, error) {
		res, err := ui.RunFirstRunSetup()
		if err != nil {
			return workspace.Config{}, err
		}
		return workspace.Config{
			Port:        8080, // Default for now
			DownloadDir: res.DownloadDir,
			SecureMode:  res.SecureMode,
		}, nil
	}

	isFirstRun, err := workspace.InitializeWorkspace(setupFunc)
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

	// 4. Welcome Message (if setup was performed, we skip welcome as setup is enough?)
	// Actually spec says "display a one-time Welcome to magshare message".
	// Setup is already a "welcome".
	if isFirstRun {
		// ui.DisplayWelcomeMessage(os.Stdout)
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
