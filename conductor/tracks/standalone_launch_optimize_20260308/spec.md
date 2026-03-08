# Specification: Optimize Standalone Launch Experience

## Overview
This track enhances the experience of using Magshare as a standalone application, specifically when launched by double-clicking the executable. It moves the interactive TUI from a "one-and-done" script to a persistent application loop and adds shortcut creation capabilities.

## Functional Requirements
1.  **Main Menu Loop:**
    - Refactor the `rootCmd` (interactive mode) to run in a loop.
    - After a "Send" or "Receive" operation completes (or is cancelled by the user), the application should return to the main action selection menu.
    - Provide a clear "Exit" option in the main menu to gracefully close the application.
2.  **Shortcut Creation:**
    - Extend the `magshare integrate` command with a new flag: `--shortcut`.
    - When run with `--shortcut` on Windows, create a Desktop shortcut pointing to the current `magshare.exe`.
    - The shortcut should have a descriptive name like "Magshare".
3.  **TUI Refinement:**
    - Ensure that cancellations within sub-menus (e.g., cancelling a path input) return the user to the main menu instead of exiting the entire program.

## Non-Functional Requirements
- **Platform Awareness:** Shortcut creation logic should be Windows-specific.
- **Maintainability:** Ensure the loop logic doesn't complicate the single-command CLI usage (e.g., `magshare send file.txt` should still exit after completion).

## Acceptance Criteria
- [ ] Launching `magshare.exe` without arguments enters a loop.
- [ ] Finishing a "Send" session returns the user to the main menu.
- [ ] Choosing "Exit" from the main menu closes the program.
- [ ] Running `magshare integrate --shortcut` creates a working shortcut on the Windows Desktop.
- [ ] Standard CLI commands (e.g., `magshare receive`) still exit upon completion/cancellation.

## Out of Scope
- Creating an installer (MSI/EXE).
- Start Menu integration (deferred to a future track).
- Automatic updates.
