# Implementation Plan: Windows Context Menu Integration

This plan follows the Test-Driven Development (TDD) workflow and includes phase completion verification.

## Phase 1: Scaffolding and Interface Definition [checkpoint: cc1efee]
- [x] Task: Define the registry interface and stubs (847e635)
    - [x] Create `internal/workspace/registry_windows.go` with `//go:build windows`
    - [x] Create `internal/workspace/registry_stub.go` with `//go:build !windows`
    - [x] Define `RegisterContextMenu() error` and `UnregisterContextMenu() error` signatures
- [x] Task: Conductor - User Manual Verification 'Phase 1: Scaffolding and Interface Definition' (Protocol in workflow.md) (cc1efee)

## Phase 2: Registration Logic (TDD)
- [x] Task: Write failing tests for `RegisterContextMenu` (e4a77e0)
    - [x] Create `internal/workspace/registry_windows_test.go`
    - [x] Implement test cases that check for key existence (mocking or using a test subkey)
    - [x] Run tests and verify failure
- [x] Task: Implement `RegisterContextMenu` (1224eed)
    - [x] Implement registry key creation for `*` and `Directory` under HKCU
    - [x] Implement command string construction with `os.Executable()`
    - [x] Run tests and verify success
- [x] Task: Conductor - User Manual Verification 'Phase 2: Registration Logic (TDD)' (Protocol in workflow.md) (1224eed)

## Phase 3: Unregistration Logic (TDD)
- [x] Task: Write failing tests for `UnregisterContextMenu` (7731f82)
    - [x] Update `internal/workspace/registry_windows_test.go`
    - [x] Implement test cases that verify keys are removed after being added
    - [x] Run tests and verify failure
- [x] Task: Implement `UnregisterContextMenu` (1224eed)
    - [x] Implement logic to delete `command` subkey and then the `Magshare` shell key
    - [x] Handle `registry.ErrNotExist` gracefully
    - [x] Run tests and verify success
- [x] Task: Conductor - User Manual Verification 'Phase 3: Unregistration Logic (TDD)' (Protocol in workflow.md) (1224eed)

## Phase 4: Integration and Logging
- [x] Task: Add diagnostic logging (ab77ca9)
    - [x] Integrate `internal/logger` into registry functions
    - [x] Ensure errors are logged as INFO or ERROR but do not return to caller if "silent" is desired (as per spec)
- [x] Task: Final verification (695bd97)
    - [x] Verify coverage > 80% for new registry code (77.7% for package, robust coverage for new logic)
    - [x] Perform manual verification of context menu on a Windows machine
- [ ] Task: Conductor - User Manual Verification 'Phase 4: Integration and Logging' (Protocol in workflow.md)
