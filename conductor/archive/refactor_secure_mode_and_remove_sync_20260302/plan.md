# Implementation Plan: Refactor Secure Mode and Remove Sync Mode

This plan details the steps to streamline the interactive TUI by removing the Sync mode and standardizing Secure mode options.

## Phase 1: Interactive TUI Refactor
- [x] **Task: Update `internal/ui/interactive.go`**
    - [x] Remove "Sync Mode" option from the main action select.
    - [x] Add "Secure Mode" confirm prompt to the "receive" case in the configuration form.
    - [x] Update the final PIN prompt condition to include the "receive" action.
- [ ] **Task: Verify TUI changes with unit tests**
    - [x] Update `internal/ui/interactive_test.go` if necessary to reflect the removed action. (No changes needed)
- [x] **Task: Conductor - User Manual Verification 'TUI Refactor' (Protocol in workflow.md)**

## Phase 2: Command Logic Cleanup
- [x] **Task: Update `cmd/root.go`**
    - [x] Remove the `case "sync"` from the interactive run logic.
- [x] **Task: End-to-end manual verification**
    - [x] Run `qshare` and verify "Sync Mode" is gone.
    - [x] Verify "Secure Mode" appears for both Send and Receive.
    - [x] Verify PIN protection works for Receive mode started via TUI.
- [x] **Task: Conductor - User Manual Verification 'Command Logic Cleanup' (Protocol in workflow.md)**

## Phase 3: Documentation Sync
- [x] **Task: Update project documents**
    - [x] Remove references to Sync mode from `conductor/product.md`. (No references found)
- [x] **Task: Conductor - User Manual Verification 'Documentation Sync' (Protocol in workflow.md)**
