# Implementation Plan: Enhanced Human-Readable Logging

## Phase 1: Core Logger Implementation [checkpoint: 4720e3c]
- [x] Task: Define the new Logger structure and levels 6d7195d
    - [ ] Create `internal/logger/types.go` to define `LogLevel` (DEBUG, INFO, WARN, ERROR) and the `Logger` interface
    - [ ] Implement a `StructuredLogger` struct that holds the writer, component name, and PID
    - [ ] Write unit tests to verify level comparison and basic formatting logic
- [x] Task: Implement the formatting logic 2a1a78c
    - [ ] Add a `format(level LogLevel, msg string)` method to `StructuredLogger`
    - [ ] Ensure it correctly interpolates `[TIMESTAMP] [LEVEL] [COMPONENT] [PID]`
    - [ ] Write unit tests to verify the string output matches the spec example
- [x] Task: Conductor - User Manual Verification 'Phase 1: Core Logger Implementation' (Protocol in workflow.md) 4720e3c

## Phase 2: Hybrid Output and Filtering [checkpoint: dde475d]
- [x] Task: Implement Filtered Multi-Writer d39f0fd
    - [ ] Create a `FilteredWriter` that wraps an `io.Writer` and only writes if the `LogLevel` meets a threshold
    - [ ] Update `SetupLogging()` to use this `FilteredWriter` for `os.Stdout` (Threshold: INFO)
    - [ ] Ensure the log file `io.Writer` has a lower threshold (Threshold: DEBUG)
    - [ ] Write unit tests to verify that DEBUG messages are suppressed in one writer but captured in another
- [x] Task: Update the global singleton logger d39f0fd
    - [ ] Provide global helper functions like `logger.Info(component, msg)`, `logger.Debug(component, msg)`, etc.
    - [ ] Ensure thread safety using `sync.Mutex` or by utilizing `log.Logger` as the underlying engine
- [x] Task: Conductor - User Manual Verification 'Phase 2: Hybrid Output and Filtering' (Protocol in workflow.md) dde475d

## Phase 3: Integration and Refactoring [checkpoint: cc12de0]
- [x] Task: Replace raw `fmt.Print` calls with the new Logger 7f18d55
    - [ ] Audit `internal/server`, `internal/handlers`, and `cmd/` for terminal output
    - [ ] Migrate essential messages to `logger.Info()` and technical details to `logger.Debug()`
    - [ ] Update `main.go` to use the new logging methods for its lifecycle events
- [x] Task: Refactor Crash Recovery integration 7f18d55
    - [ ] Update `HandlePanic()` to use the new formatting for the crash report header
    - [ ] Ensure the stack trace is still captured clearly within the formatted log
- [x] Task: Conductor - User Manual Verification 'Phase 3: Integration and Refactoring' (Protocol in workflow.md) cc12de0

## Phase 4: Validation and Quality Assurance [checkpoint: acb967c]
- [x] Task: Verify >80% code coverage for the new logging system 86%
    - [ ] Run `go test -cover ./internal/logger/...`
- [x] Task: Final end-to-end manual verification (manual)
    - [ ] Run a normal session and verify `INFO` logs in terminal and `DEBUG` logs in the file
    - [ ] Trigger a failure and verify `ERROR` logs and metadata accuracy
- [x] Task: Conductor - User Manual Verification 'Phase 4: Validation and Quality Assurance' (Protocol in workflow.md) acb967c
