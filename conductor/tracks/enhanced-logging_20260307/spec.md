# Specification: Enhanced Human-Readable Logging

## Overview
Transform the current raw output redirection into a structured, human-readable logging system. The goal is to produce log files that are easy to scan, categorized by severity, and rich with diagnostic metadata, while keeping the terminal output focused and clean.

## Functional Requirements
1.  **Standardized Log Format:**
    -   Adopt a plain text format with clear prefixes: `[TIMESTAMP] [LEVEL] [COMPONENT] [PID] MESSAGE`.
    -   Example: `[2023-10-27 14:30:05] [INFO] [server] [1234] Started listening on :8080`
2.  **Rich Metadata:**
    -   **Timestamp:** ISO-8601 or similar readable format on every line.
    -   **Component Name:** Identify which part of the app generated the log (e.g., `main`, `server`, `handler`, `workspace`).
    -   **Process Info:** Include the current Process ID (PID) for correlation.
3.  **Categorized Log Levels:**
    -   Implement and support `DEBUG`, `WARN`, and `ERROR` levels.
    -   `INFO` remains the default for standard events.
4.  **Hybrid Output Strategy:**
    -   **Terminal:** Keep it high-signal. Only show essential progress and errors.
    -   **Log File:** Verbose. Capture all levels (including `DEBUG`) to provide a full audit trail for troubleshooting.
5.  **Backward Compatibility:**
    -   Retain the existing crash recovery logic, ensuring `recover()` still writes the panic and stack trace to this new formatted log.

## Non-Functional Requirements
-   **Performance:** Formatting should be efficient to avoid slowing down high-throughput operations (like file streaming).
-   **Thread Safety:** Ensure logs from different goroutines don't interleave or corrupt each other.

## Acceptance Criteria
- [ ] Log files use the `[TIMESTAMP] [LEVEL] [COMPONENT] [PID]` format consistently.
- [ ] `DEBUG` messages appear in the log file but NOT in the standard terminal output.
- [ ] `WARN` and `ERROR` messages are visually distinct in the log.
- [ ] Log lines include the correct component name.
- [ ] The system handles concurrent logging from multiple goroutines safely.

## Out of Scope
-   **Log Rotation/Pruning:** Managing the lifecycle of multiple historical logs.
-   **Remote Centralized Logging:** Sending logs to a cloud provider.
-   **Colors in Log Files:** ANSI escape codes should be excluded from the file to ensure compatibility with all text editors.
