# Implementation Plan: Interactive CLI Prompts (charmbracelet/huh)

This plan outlines the steps for implementing an interactive TUI for QShare using the `charmbracelet/huh` library.

## Phase 1: Setup and Prototyping [checkpoint: b572995]
- [x] **Task: Update Tech Stack and Install Dependencies** 77a6792
    - [x] Add `charmbracelet/huh` to `tech-stack.md`.
    - [x] Run `go get github.com/charmbracelet/huh` to add the library to the project.
- [x] **Task: Create a separate 'interactive' package for UI logic** 364c908
    - [x] Create `internal/ui/interactive.go` (or a similar location) to isolate the `huh` forms.
    - [x] Write a test in `internal/ui/interactive_test.go` to verify the form structure (mocking input if possible).
- [x] **Task: Prototype the 'Primary Action' prompt** 364c908
    - [x] Implement a `huh` Select field for Send, Receive, and Sync.
- [x] **Task: Conductor - User Manual Verification 'Setup and Prototyping' (Protocol in workflow.md)**

## Phase 2: Core Feature Implementation
- [ ] **Task: Implement feature-specific configuration forms**
    - [ ] Create sub-forms for `Send`, `Receive`, and `Sync` with appropriate fields (Path, Port, PIN, Secure).
    - [ ] Add validation logic for the input fields (e.g., non-empty paths, valid port numbers).
- [ ] **Task: Integrate into Root Command**
    - [ ] Modify `cmd/root.go` to detect if no arguments or flags are provided.
    - [ ] Call the interactive prompt logic when the trigger conditions are met.
- [ ] **Task: Map Interactive Inputs to Command Logic**
    - [ ] Ensure that once the interactive session is complete, it calls the corresponding internal handlers (e.g., `handlers.StartSendServer`).
- [ ] **Task: Conductor - User Manual Verification 'Core Feature Implementation' (Protocol in workflow.md)**

## Phase 3: Refinement and Polishing
- [ ] **Task: Refine UI and Error Handling**
    - [ ] Improve the visual presentation of the prompts.
    - [ ] Ensure that cancelling the interactive session (e.g., Ctrl+C) exits gracefully.
- [ ] **Task: Final verification of CLI flag parity**
    - [ ] Ensure that running the CLI with traditional flags still works as expected.
- [ ] **Task: Conductor - User Manual Verification 'Refinement and Polishing' (Protocol in workflow.md)**
