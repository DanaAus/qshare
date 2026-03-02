# Implementation Plan: Fix Secure Send PIN Bug

This plan focuses on implementing a user-friendly PIN entry page for secure file sharing.

## Phase 1: UI Implementation
- [ ] **Task: Create `ui/pin.html` template**
    - [ ] Design a simple, dark-themed PIN entry page consistent with `dropzone.html`.
    - [ ] Implement a form that submits the PIN via a GET request (adding `?pin=` to the URL).
- [ ] **Task: Register template in `ui/embed.go`**
    - [ ] Ensure the new template is included in the embedded FS.
- [ ] **Task: Conductor - User Manual Verification 'UI Implementation' (Protocol in workflow.md)**

## Phase 2: Server Logic Update
- [ ] **Task: Update `internal/handlers/send.go` to serve PIN page**
    - [ ] Modify the handler logic to serve `pin.html` if `opts.Secure` is true and the PIN is incorrect or missing.
    - [ ] Ensure the file download only triggers when the PIN is correct.
- [ ] **Task: Verify functionality with automated tests**
    - [ ] Create a test case in `internal/handlers/send_test.go` that verifies the PIN page is served.
- [ ] **Task: Conductor - User Manual Verification 'Server Logic Update' (Protocol in workflow.md)**

## Phase 3: Final Verification
- [ ] **Task: End-to-end manual verification**
    - [ ] Run `qshare send` with secure mode.
    - [ ] Verify the full flow: Open URL -> Enter PIN -> Download.
- [ ] **Task: Conductor - User Manual Verification 'Final Verification' (Protocol in workflow.md)**
