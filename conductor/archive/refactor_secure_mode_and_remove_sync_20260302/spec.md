# Specification: Refactor Secure Mode and Remove Sync Mode

## Overview
This track aims to simplify the QShare interactive TUI by removing the incomplete and buggy "Sync" mode and ensuring that the "Secure Mode" (PIN protection) is consistently offered as an option for both sending and receiving files.

## Problem Description
- **Buggy Sync Mode**: The "Sync" mode is currently a prototype and considered buggy by the user.
- **Inconsistent Secure Mode**: The interactive TUI only prompts for "Secure Mode" when sending files, but the underlying handler supports it for receiving files as well.

## Proposed Changes
1. **Remove Sync Mode**:
    - Remove "Sync Mode" from the primary action selection in the interactive TUI.
    - Remove "sync" case from the root command handler.
    - Remove "sync" related logic from `internal/ui/interactive.go`.
2. **Standardize Secure Mode**:
    - Add the "Secure Mode" confirmation prompt to the "Receive" flow in the interactive TUI.
    - Ensure both "Send" and "Receive" flows allow for custom PIN entry if "Secure Mode" is enabled.
    - Update `cmd/root.go` to correctly pass these options to the handlers.

## Acceptance Criteria
- [ ] Running `qshare` (interactive mode) no longer shows "Sync Mode".
- [ ] Selecting "Receive File" in interactive mode now prompts for "Secure Mode".
- [ ] If "Secure Mode" is selected for "Receive", the user is prompted to set a PIN (consistent with the "Send" flow).
- [ ] Entering a PIN in interactive "Receive" mode correctly sets it in the dropzone server.
- [ ] Traditional CLI commands for send and receive still work.

## Out of Scope
- Fixing the Sync mode (it is being removed entirely).
- Implementing new security features beyond the existing PIN protection.
