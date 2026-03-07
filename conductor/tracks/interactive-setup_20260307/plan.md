# Implementation Plan: Interactive First-Run Setup

## Phase 1: Config Struct Update [checkpoint: 5ec3b5c]
- [x] Task: Update `internal/workspace/config.go` 10dc16f
    - [ ] Add `DownloadDir` string field to the `Config` struct.
    - [ ] Update `CreateDefaultConfig` to accept a `Config` object instead of creating a fixed one.
    - [ ] Write unit tests to verify the new struct and JSON serialization.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Config Struct Update' (Protocol in workflow.md) 5ec3b5c

## Phase 2: Interactive Setup Implementation [checkpoint: 37166d7]
- [x] Task: Create `internal/ui/setup.go` e1d1e03
    - [ ] Implement `RunFirstRunSetup()` using `huh`.
    - [ ] Logic:
        - Resolve user home directory.
        - Set default path to `filepath.Join(home, "Magshare Downloads")`.
        - Form with Input (Download Path) and Confirm (PIN Security).
        - Validation logic for the path (absolute, writable).
- [x] Task: Write tests/mocks for the setup logic 6bd27d7
    - [ ] Since TUIs are hard to unit test, focus on validating the path-checking logic in a separate utility.
- [x] Task: Conductor - User Manual Verification 'Phase 2: Interactive Setup Implementation' (Protocol in workflow.md) 37166d7

## Phase 3: Integration into Initialization [checkpoint: 71f25bd]
- [x] Task: Update `internal/workspace/init.go` 303ad69
    - [ ] Modify `InitializeWorkspaceAtPath` to detect missing `config.json`.
    - [ ] If missing, call `ui.RunFirstRunSetup()`.
    - [ ] Save the returned configuration using the updated `CreateDefaultConfig`.
- [x] Task: Update `main.go` a512306
    - [ ] Ensure the welcome message and setup flow happen in the correct order.
- [x] Task: Conductor - User Manual Verification 'Phase 3: Integration into Initialization' (Protocol in workflow.md) 71f25bd

## Phase 4: Final Validation [checkpoint: 7e9734a]
- [x] Task: End-to-end manual verification (Verified in Phase 3)
- [x] Task: Conductor - User Manual Verification 'Phase 4: Final Validation' (Protocol in workflow.md) 7e9734a
