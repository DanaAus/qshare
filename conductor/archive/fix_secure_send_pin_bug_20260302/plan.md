# Implementation Plan: Fix Secure Send PIN Bug

This plan focuses on implementing a user-friendly PIN entry page for secure file sharing.

## Phase 1: UI Implementation
- [x] **Task: Create `ui/pin.html` template**
    - [x] Design a simple, dark-themed PIN entry page consistent with `dropzone.html`.
    - [x] Implement a form that submits the PIN via a GET request (adding `?pin=` to the URL).
- [x] **Task: Register template in `ui/embed.go`**
    - [x] Ensure the new template is included in the embedded FS.
- [x] **Task: Conductor - User Manual Verification 'UI Implementation' (Protocol in workflow.md)**

## Phase 2: Server Logic Update
- [x] **Task: Update `internal/handlers/send.go` to serve PIN page**
    - [x] Modify the handler logic to serve `pin.html` if `opts.Secure` is true and the PIN is incorrect or missing.
    - [x] Ensure the file download only triggers when the PIN is correct.
- [x] **Task: Verify functionality with automated tests**
    - [x] Create a test case in `internal/handlers/send_test.go` that verifies the PIN page is served.
- [x] **Task: Conductor - User Manual Verification 'Server Logic Update' (Protocol in workflow.md)**

## Phase 3: Final Verification
- [x] **Task: End-to-end manual verification**
    - [x] Run `qshare send` with secure mode.
    - [x] Verify the full flow: Open URL -> Enter PIN -> Download.
- [x] **Task: Conductor - User Manual Verification 'Final Verification' (Protocol in workflow.md)**
