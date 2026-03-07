# Implementation Plan: Workspace Setup & Crash Logging

## Phase 1: Foundation and Utilities [checkpoint: c3c7002]
- [x] Task: Set up core directory and path utility functions 1b5f374
    - [ ] Create `internal/workspace/path.go` to handle system config directory discovery (e.g., `os.UserConfigDir`)
    - [ ] Implement `GetWorkspaceRoot()` and `GetLogsDir()` functions
    - [ ] Write unit tests in `internal/workspace/path_test.go` to verify correct path resolution for Windows/macOS/Linux
- [x] Task: Implement basic folder and file existence checks fdbefda
    - [ ] Add `EnsureDirectoryExists(path string)` utility
    - [ ] Add `FileExists(path string)` utility
    - [ ] Write tests to verify directory creation and file detection logic
- [x] Task: Conductor - User Manual Verification 'Phase 1: Foundation and Utilities' (Protocol in workflow.md) c3c7002

## Phase 2: First-Run Logic and Config Creation [checkpoint: f571656]
- [x] Task: Implement first-run detection and initialization 65631fa
    - [ ] Add `InitializeWorkspace()` function in `internal/workspace/init.go`
    - [ ] Logic: Check for workspace root; if missing, create structure and return `isFirstRun = true`
    - [ ] Write unit tests to verify directory structure creation and first-run flag detection
- [x] Task: Implement default config file generation 1016094
    - [ ] Define basic `Config` struct (e.g., `Port`, `SecureMode`)
    - [ ] Add `CreateDefaultConfig(path string)` function using `encoding/json`
    - [ ] Write unit tests to verify JSON content and file permissions
- [x] Task: Integrate "Welcome" message in terminal 1846201
    - [ ] Add `DisplayWelcomeMessage()` in `internal/ui/welcome.go` using `lipgloss` for styling
    - [ ] Write a test/verification to ensure it only prints when `isFirstRun` is true
- [x] Task: Conductor - User Manual Verification 'Phase 2: First-Run Logic and Config Creation' (Protocol in workflow.md) f571656

## Phase 3: Output Redirection and Logging [checkpoint: 4e02879]
- [x] Task: Implement Multi-Writer for terminal and log file 935237a
    - [ ] Create `internal/logger/logger.go` with `SetupLogging()` function
    - [ ] Implement `io.MultiWriter` to pipe `os.Stdout` and `os.Stderr` to both terminal and a temporary log file
    - [ ] Implement timestamp-based naming (e.g., `magshare-20231027-143005.log`)
    - [ ] Write unit tests to verify output is written to both destinations
- [x] Task: Implement session cleanup on success 660c2d9
    - [ ] Add `CleanupLogs(tempLogPath string)` function to delete the temporary file on normal exit
    - [ ] Write unit tests to verify file deletion after a successful run simulation
- [x] Task: Conductor - User Manual Verification 'Phase 3: Output Redirection and Logging' (Protocol in workflow.md) 4e02879

## Phase 4: Crash Recovery and Exit Logic [checkpoint: 812eddb]
- [x] Task: Implement Panic handling with `defer` and `recover` fad451f
    - [ ] Add `HandlePanic(logPath string)` function
    - [ ] Logic: Recover from panic, write stack trace to log, print log path to terminal
    - [ ] Implement "Wait 5 Seconds" logic using `time.Sleep` and a countdown timer if possible
    - [ ] Write unit tests to simulate a panic and verify log content and path printing
- [x] Task: Refactor `main.go` for integration 2dac42d
    - [ ] Update `main()` to use the new `internal/workspace`, `internal/logger`, and `HandlePanic` logic
    - [ ] Ensure proper ordering of initialization, logging setup, and panic deferment
- [x] Task: Conductor - User Manual Verification 'Phase 4: Crash Recovery and Exit Logic' (Protocol in workflow.md) 812eddb

## Phase 5: Quality Assurance and Final Checks [checkpoint: c0f2fff]
- [x] Task: Final end-to-end manual verification (manual)
    - [ ] Verify first-run setup in a clean environment (delete existing workspace)
    - [ ] Verify normal execution (no log left behind)
    - [ ] Verify crash execution (forced panic, log saved, 5-second wait)
- [x] Task: Verify >80% code coverage for new modules d0c84d2
    - [ ] Run `go test -cover ./internal/workspace/... ./internal/logger/...`
- [x] Task: Conductor - User Manual Verification 'Phase 5: Quality Assurance and Final Checks' (Protocol in workflow.md) c0f2fff
