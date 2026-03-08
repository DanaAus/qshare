# Implementation Plan: Windows Shell Context Menu Integration

This plan follows the TDD workflow and includes platform-specific conditional logic.

## Phase 1: Config Extension and Onboarding
- [x] Task: Update `workspace.Config` struct (ed9ffdd)
    - [x] Add `ShellIntegration bool` to `Config` in `internal/workspace/config.go`.
    - [x] Update `config_test.go` to ensure JSON serialization works.
- [ ] Task: Integrate into First-Run Setup
    - [ ] Update `ui.SetupResult` and `ui.RunFirstRunSetup` in `internal/ui/setup.go` to include the integration prompt.
    - [ ] Add platform check (`runtime.GOOS == "windows"`) to show/hide the prompt.
    - [ ] Update `main.go` to call `workspace.RegisterContextMenu()` if the user opts-in during setup.
- [ ] Task: Conductor - User Manual Verification 'Phase 1: Config Extension and Onboarding' (Protocol in workflow.md)

## Phase 2: CLI Command 'integrate'
- [ ] Task: Scaffolding for `integrate` command
    - [ ] Create `cmd/integrate.go`.
    - [ ] Define flags: `--install`, `--uninstall`, `--status`.
    - [ ] Implement platform-specific error (exit if not on Windows).
- [ ] Task: Implement `integrate` logic
    - [ ] Implement `--install` (calls `workspace.RegisterContextMenu`).
    - [ ] Implement `--uninstall` (calls `workspace.UnregisterContextMenu`).
    - [ ] Implement `--status` (verifies registry entries).
    - [ ] Use `huh` or fmt to provide verbose feedback to the terminal.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: CLI Command integrate' (Protocol in workflow.md)

## Phase 3: Robustness and Finalization
- [ ] Task: Verify end-to-end flow
    - [ ] Run first-run setup on Windows and verify menu appears.
    - [ ] Use CLI to uninstall and verify menu disappears.
    - [ ] Verify clicking the menu option launches `magshare send`.
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Robustness and Finalization' (Protocol in workflow.md)
