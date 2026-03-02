# Implementation Plan: Refactor Secure Mode and Remove Sync Mode

This plan details the steps to streamline the interactive TUI by removing the Sync mode and standardizing Secure mode options.

## Phase 1: Interactive TUI Refactor
- [ ] **Task: Update `internal/ui/interactive.go`**
    - [ ] Remove "Sync Mode" option from the main action select.
    - [ ] Add "Secure Mode" confirm prompt to the "receive" case in the configuration form.
    - [ ] Update the final PIN prompt condition to include the "receive" action.
- [ ] **Task: Verify TUI changes with unit tests**
    - [ ] Update `internal/ui/interactive_test.go` if necessary to reflect the removed action.
- [ ] **Task: Conductor - User Manual Verification 'TUI Refactor' (Protocol in workflow.md)**

## Phase 2: Command Logic Cleanup
- [ ] **Task: Update `cmd/root.go`**
    - [ ] Remove the `case "sync"` from the interactive run logic.
- [ ] **Task: End-to-end manual verification**
    - [ ] Run `qshare` and verify "Sync Mode" is gone.
    - [ ] Verify "Secure Mode" appears for both Send and Receive.
    - [ ] Verify PIN protection works for Receive mode started via TUI.
- [ ] **Task: Conductor - User Manual Verification 'Command Logic Cleanup' (Protocol in workflow.md)**

## Phase 3: Documentation Sync
- [ ] **Task: Update project documents**
    - [ ] Remove references to Sync mode from `conductor/product.md`.
- [ ] **Task: Conductor - User Manual Verification 'Documentation Sync' (Protocol in workflow.md)**
