# Specification: Windows Shell Context Menu Integration

## Overview
This track integrates the previously implemented Windows registry logic into the Magshare CLI and first-run onboarding flow. It ensures users can easily enable/disable the "Share via Magshare" context menu option, which allows sharing files directly from Windows Explorer.

## Functional Requirements
1.  **Config Extension:**
    - Add a `ShellIntegration` boolean field to the `workspace.Config` struct.
    - Persist this value in `config.json`.
2.  **First-Run Onboarding Flow:**
    - Update `ui.RunFirstRunSetup` to include a prompt for Windows users: "Enable Windows Explorer integration? (Adds 'Share via Magshare' to right-click menu)".
    - If accepted, call `workspace.RegisterContextMenu()` and save the preference.
3.  **Dedicated CLI Command (`integrate`):**
    - Implement a new command `magshare integrate`.
    - **Flags:**
        - `--install`: Manually registers the context menu.
        - `--uninstall`: Manually removes the context menu.
        - `--status`: Reports whether the integration is currently active (checks registry or config).
    - Provide verbose terminal feedback (success/failure messages) using `huh` or simple print statements.
4.  **Execution Persistence:**
    - The registry keys must point to the absolute path of the current `magshare.exe`.
    - Clicking the context menu option must launch `magshare send "{path}"` even if the app was closed.

## Non-Functional Requirements
- **Platform Awareness:** Onboarding prompts and the `integrate` command should only be active/visible on Windows.
- **Robustness:** Registry failures should be logged, and manual CLI actions should provide clear error messages.

## Acceptance Criteria
- [ ] Windows users are prompted to enable Explorer integration during first-run setup.
- [ ] Running `magshare integrate --install` adds the context menu entries.
- [ ] Running `magshare integrate --uninstall` removes the context menu entries.
- [ ] `magshare integrate --status` correctly identifies the integration state.
- [ ] Right-clicking a file/folder and choosing "Share via Magshare" launches the app in `send` mode for that item.
- [ ] Integration state is saved in `config.json`.

## Out of Scope
- Integration with macOS/Linux shells in this track.
- Customizable menu labels.
