# Specification: Workspace Setup & Crash Logging

## Overview
Implement a robust Go `main` function for `magshare` that handles first-run setup, redirects output to a log file, and provides crash recovery. The goal is to ensure a seamless setup in the user's system configuration directory and capture vital diagnostic information if an error occurs.

## Functional Requirements
1.  **System Workspace Creation:**
    -   Automatically and silently create a `magshare` directory in the user's local application data folder (e.g., `%LOCALAPPDATA%\magshare` on Windows).
    -   Create a subfolder named `logs` within the workspace.
2.  **First Run Setup:**
    -   If the workspace directory does not exist, perform "first-run" tasks:
        -   Create the `magshare` and `logs` folder structure.
        -   Generate a default `config.json` file in the workspace.
        -   Display a one-time "Welcome to magshare" message in the terminal.
3.  **Output Redirection & Logging:**
    -   At startup, create a temporary log file in the `logs` directory.
    -   Use a "Timestamp Based" naming convention (e.g., `magshare-20231027-143005.log`).
    -   Redirect both `stdout` and `stderr` to both the terminal (for real-time feedback) and the temporary log file.
4.  **Crash Recovery (Panic Handling):**
    -   Wrap the main application logic in a `defer` block with `recover`.
    -   If a panic occurs:
        -   Capture the panic message and stack trace into the log file.
        -   Print the full path of the log file to the terminal.
        -   Keep the terminal window open for exactly 5 seconds before exiting.
5.  **Clean Exit:**
    -   If the application exits normally (without a panic), delete the temporary log file.

## Non-Functional Requirements
-   **Low Overhead:** The logging and redirection should not significantly impact performance or memory usage.
-   **Robustness:** The setup logic should handle file system permission issues gracefully (e.g., if the user doesn't have write access to AppData).
-   **Platform Consistency:** While focused on Windows (LocalAppData), the implementation should be adaptable to other OS config patterns (e.g., XDG_CONFIG_HOME on Linux).

## Acceptance Criteria
- [ ] On the first run, `magshare` creates `%LOCALAPPDATA%\magshare` with a `logs` folder and a default `config.json`.
- [ ] A "Welcome" message is displayed only during the initial setup.
- [ ] Normal output (e.g., `magshare send --help`) is printed to both terminal and log.
- [ ] On normal exit, no temporary log file remains.
- [ ] On panic, the application prints the log path and waits for 5 seconds before the window closes.
- [ ] The crash log contains the panic details and stack trace.

## Out of Scope
-   **Log Rotation:** Automated deletion of old crash logs (beyond the temporary one).
-   **Remote Error Reporting:** Sending crash logs to an external server.
-   **GUI Setup:** Graphical installer or setup wizard.
