# Implementation Plan: Windows Shell Context Menu Integration

This plan follows the TDD workflow and includes platform-specific conditional logic.

## Phase 1: Config Extension and Onboarding [checkpoint: 7e6c027]
- [x] Task: Update `workspace.Config` struct (ed9ffdd)
    - [x] Add `ShellIntegration bool` to `Config` in `internal/workspace/config.go`.
    - [x] Update `config_test.go` to ensure JSON serialization works.
- [x] Task: Integrate into First-Run Setup (085f92a)
    - [x] Update `ui.SetupResult` and `ui.RunFirstRunSetup` in `internal/ui/setup.go` to include the integration prompt.
    - [x] Add platform check (`runtime.GOOS == "windows"`) to show/hide the prompt.
    - [x] Update `main.go` to call `workspace.RegisterContextMenu()` if the user opts-in during setup.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Config Extension and Onboarding' (Protocol in workflow.md) (7e6c027)

## Phase 2: CLI Command 'integrate' [checkpoint: 7aea748]
- [x] Task: Scaffolding for `integrate` command (cfb1c3d)
    - [x] Create `cmd/integrate.go`.
    - [x] Define flags: `--install`, `--uninstall`, `--status`.
    - [x] Implement platform-specific error (exit if not on Windows).
- [x] Task: Implement `integrate` logic (7aea748)
    - [x] Implement `--install` (calls `workspace.RegisterContextMenu`).
    - [x] Implement `--uninstall` (calls `workspace.UnregisterContextMenu`).
    - [x] Implement `--status` (verifies registry entries).
    - [x] Use `huh` or fmt to provide verbose feedback to the terminal.
- [x] Task: Conductor - User Manual Verification 'Phase 2: CLI Command integrate' (Protocol in workflow.md) (7aea748)

## Phase 3: Robustness and Finalization
- [x] Task: Verify end-to-end flow (7aea748)
    - [x] Run first-run setup on Windows and verify menu appears.
    - [x] Use CLI to uninstall and verify menu disappears.
    - [x] Verify clicking the menu option launches `magshare send`.
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Robustness and Finalization' (Protocol in workflow.md)
