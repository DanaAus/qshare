package logger

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

// HandlePanic processes a recovered panic by logging it and waiting before exit.
// r is the value returned by recover().
func HandlePanic(logPath string, r interface{}) {
	if r == nil {
		return
	}

	// 1. Log to whatever global logger is configured (usually terminal + file)
	l := WithComponent("panic")
	l.Error(fmt.Sprintf("magshare has crashed: %v", r))

	// 2. Ensure it's in the file even if global logger failed or was misconfigured
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		f.WriteString(fmt.Sprintf("[%s] [ERROR] [panic] [%d] magshare has crashed: %v\n", timestamp, os.Getpid(), r))
		
		// Capture stack trace
		stack := debug.Stack()
		f.WriteString("\n========== STACK TRACE ==========\n")
		f.WriteString(string(stack))
		f.WriteString("==================================\n")
		f.Close()
	}

	// Write status to terminal via Stderr (unfiltered)
	fmt.Fprintf(os.Stderr, "\nCrash log saved to: %s\n", logPath)

	// Countdown timer for 5 seconds
	fmt.Fprintf(os.Stderr, "\nClosing in ")
	for i := 5; i > 0; i-- {
		fmt.Fprintf(os.Stderr, "%d... ", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintln(os.Stderr)
}
