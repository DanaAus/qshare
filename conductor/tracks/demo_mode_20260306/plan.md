# Implementation Plan: Demo Mode

This plan outlines the steps required to implement the "Demo Mode" for magshare, faking connection information for promotional purposes while maintaining stealth.

## Phase 1: Core Logic & Flag Definition [checkpoint: 6a1ca29]
- [x] **Task: Define the Global Demo Flag** e037374
    - [ ] Add a new flag (e.g., `--demo`) to the root command in `cmd/root.go`.
    - [ ] Ensure the flag is persistent and accessible to all subcommands (`send`, `receive`).
    - [ ] Update any necessary state or context objects to carry this flag.
- [x] **Task: Create IP and URL Faking Utilities** 9fca513
    - [ ] Implement a function in `internal/network/interface.go` that returns a fake IP when demo mode is active.
    - [ ] Implement a utility that takes a URL and replaces the IP/port with fake versions for display.
- [x] **Task: Conductor - User Manual Verification 'Phase 1: Core Logic & Flag Definition' (Protocol in workflow.md)** 6a1ca29

## Phase 2: Integration into Commands [checkpoint: 5260db4]
- [x] **Task: Integrate Demo Mode into `send` Command** 14db76f
    - [ ] Modify `cmd/send.go` and `internal/handlers/send.go` to use the fake IP and URL for terminal output.
    - [ ] Ensure the QR code generator receives the faked URL.
- [x] **Task: Integrate Demo Mode into `receive` Command** c5ecc2d
    - [ ] Modify `cmd/receive.go` and `internal/handlers/receive.go` to use the fake IP and URL for terminal output.
    - [ ] Ensure the QR code generator receives the faked URL.
- [x] **Task: Integrate Demo Mode into TUI Mode** a534d36
    - [ ] Modify the `interactive.go` and related TUI logic to respect the demo flag when displaying connection details.
- [x] **Task: Conductor - User Manual Verification 'Phase 2: Integration into Commands' (Protocol in workflow.md)** 5260db4

## Phase 3: Stealth & Verification
- [ ] **Task: Ensure Stealth Attribute**
    - [ ] Review all UI components to ensure no "Demo Mode" label or status is visible when the flag is active.
- [ ] **Task: Verification and Manual Testing**
    - [ ] Verify `magshare send --demo` shows faked details.
    - [ ] Verify `magshare receive --demo` shows faked details.
    - [ ] Verify the actual server still binds to the real IP (internal functionality test).
- [ ] **Task: Conductor - User Manual Verification 'Phase 3: Stealth & Verification' (Protocol in workflow.md)**
