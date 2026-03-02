# Specification: Interactive CLI Prompts (charmbracelet/huh)

## Overview
This track implements an interactive TUI for QShare using the `charmbracelet/huh` library. Instead of forcing users to remember and type complex flags, running `qshare` with no arguments will launch an interactive session to guide the user through the most common tasks.

## Functional Requirements
- **Interactive Trigger**: The interactive TUI must launch ONLY when the `qshare` root command is executed without any subcommands or flags.
- **Library Choice**: Use `github.com/charmbracelet/huh` for building the interactive prompts and forms.
- **Inline Style**: The prompts should appear as inline terminal elements (standard `huh` behavior), not a full-screen application.
- **Primary Action Prompt**:
    - Ask the user to choose between: `Send`, `Receive`, and `Sync`.
- **Feature-Specific Configuration**:
    - If `Send` or `Sync` is chosen:
        - Prompt for the `Port` (with a default value).
        - Prompt for a `PIN` (optional, default to none).
        - Prompt for `Secure` mode (Yes/No toggle).
        - Prompt for the `File/Directory Path`.
    - If `Receive` is chosen:
        - Prompt for the `Port` (with a default value).
        - Prompt for the `Destination Directory`.
- **CLI Flag Parity**: The interactive inputs must be mapped directly to the existing CLI flags and logic.

## Non-Functional Requirements
- **Frictionless Entry**: The interaction should be smooth and intuitive, following the product vision of "frictionless sharing."
- **Terminal Aesthetics**: The prompts should be visually appealing and clearly indicate the current step.

## Acceptance Criteria
- [ ] Running `qshare` (with no args) launches the `huh` interactive UI.
- [ ] Selecting an action (e.g., Send) and providing the required inputs (e.g., file path) correctly triggers the underlying command logic.
- [ ] The user can successfully complete a file transfer using only the interactive prompts.
- [ ] Traditional CLI usage (e.g., `qshare send file.txt`) remains unaffected.

## Out of Scope
- A full-screen TUI dashboard.
- Interactive prompts for advanced or rarely used flags.
- Interactive file browser (for now, simple path input is enough).
