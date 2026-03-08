# Implementation Plan: Optimize Standalone Launch Experience

This plan focuses on enhancing the interactive mode into a persistent loop and adding Windows shortcut functionality.

## Phase 1: Main Menu Loop [checkpoint: eab6a3d]
- [x] Task: Refactor `RunInteractivePrompts` for exit option (55f17f8)
    - [x] Add an "Exit" option to the action selection menu in `internal/ui/interactive.go`.
    - [x] Update `InteractiveConfig` to handle the "exit" action.
- [x] Task: Refactor `rootCmd` for looping (5179163)
    - [x] Modify `cmd/root.go`'s `RunE` function.
    - [x] Wrap the `ui.RunInteractivePrompts` call in an infinite loop.
    - [x] Break the loop if the action is "exit" or if an unrecoverable error occurs.
    - [x] Log action completions and return to menu.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Main Menu Loop' (Protocol in workflow.md) (eab6a3d)

## Phase 2: Shortcut Creation
- [x] Task: Implement Windows Shortcut Logic (33a8916)
    - [x] Create `internal/workspace/shortcut_windows.go` with `//go:build windows`.
    - [x] Implement `CreateDesktopShortcut()` using a shell-based approach (e.g., `WScript.Shell` via `os/exec` or a library).
    - [x] Create stubs for other platforms in `internal/workspace/shortcut_stub.go`.
- [ ] Task: Add `--shortcut` flag to `integrate` command
    - [ ] Update `cmd/integrate.go` to include the `shortcutFlag`.
    - [ ] Integrate with `workspace.CreateDesktopShortcut()` and provide terminal feedback.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: Shortcut Creation' (Protocol in workflow.md)

## Phase 3: Final Integration and Testing
- [ ] Task: Verify standalone vs CLI behavior
    - [ ] Test `magshare send <path>` directly to ensure it still exits as expected.
    - [ ] Test `magshare` without arguments to verify the loop and exit option.
    - [ ] Verify the created shortcut launches the program correctly.
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Final Integration and Testing' (Protocol in workflow.md)
