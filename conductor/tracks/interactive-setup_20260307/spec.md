# Specification: Interactive First-Run Setup

## Overview
Implement an interactive first-run experience for Magshare. When the application starts and detects that no configuration file exists, it will guide the user through a brief setup process to define where files are saved and set security defaults.

## Functional Requirements
1.  **First-Run Detection:**
    -   Check for the existence of `config.json` in the system's Magshare config directory (e.g., `%LOCALAPPDATA%\magshare` on Windows or `~/.config/magshare` on Linux).
2.  **Interactive Configuration Prompt:**
    -   If the config is missing, automatically trigger an interactive TUI (using `huh`).
    -   **Question 1: Download Location**
        -   Ask where received files should be saved.
        -   Default value: `Magshare Downloads` folder in the user's home directory.
        -   Validation: Must be an absolute path and writable.
    -   **Question 2: Default Security**
        -   Ask if PIN security should be enabled by default for all future transfers.
3.  **Persistence:**
    -   Save the user's choices into the `config.json` file.
    -   Ensure the Magshare workspace and logs directories are also created if they don't exist.
4.  **Resource Preparation:**
    -   Ensure the specified downloads folder is created (on demand) when the first transfer occurs, but validate its viability during setup.

## Non-Functional Requirements
-   **User-Friendly TUI:** Use the established `huh` styling consistent with existing interactive prompts.
-   **Robust Path Handling:** Correctly resolve home directory paths (`~`) and handle cross-platform differences.

## Acceptance Criteria
- [ ] Application detects missing `config.json` and starts the setup form.
- [ ] User can specify a custom path or accept the default home directory folder.
- [ ] PIN security toggle correctly saves to the config.
- [ ] Config file is written with the correct JSON structure.
- [ ] Path validation prevents invalid or unwritable directories.

## Out of Scope
-   **Re-running Setup:** A separate command to re-trigger this setup (will be handled by manual config editing for now).
-   **Cloud Sync:** Syncing the config across devices.
