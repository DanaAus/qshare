# Implementation Plan: Windows Context Menu Integration

This plan follows the Test-Driven Development (TDD) workflow and includes phase completion verification.

## Phase 1: Scaffolding and Interface Definition [checkpoint: cc1efee]
- [x] Task: Define the registry interface and stubs (847e635)
    - [x] Create `internal/workspace/registry_windows.go` with `//go:build windows`
    - [x] Create `internal/workspace/registry_stub.go` with `//go:build !windows`
    - [x] Define `RegisterContextMenu() error` and `UnregisterContextMenu() error` signatures
- [x] Task: Conductor - User Manual Verification 'Phase 1: Scaffolding and Interface Definition' (Protocol in workflow.md) (cc1efee)

## Phase 2: Registration Logic (TDD)
- [ ] Task: Write failing tests for `RegisterContextMenu`
    - [ ] Create `internal/workspace/registry_windows_test.go`
    - [ ] Implement test cases that check for key existence (mocking or using a test subkey)
    - [ ] Run tests and verify failure
- [ ] Task: Implement `RegisterContextMenu`
    - [ ] Implement registry key creation for `*` and `Directory` under HKCU
    - [ ] Implement command string construction with `os.Executable()`
    - [ ] Run tests and verify success
- [ ] Task: Conductor - User Manual Verification 'Phase 2: Registration Logic (TDD)' (Protocol in workflow.md)

## Phase 3: Unregistration Logic (TDD)
- [ ] Task: Write failing tests for `UnregisterContextMenu`
    - [ ] Update `internal/workspace/registry_windows_test.go`
    - [ ] Implement test cases that verify keys are removed after being added
    - [ ] Run tests and verify failure
- [ ] Task: Implement `UnregisterContextMenu`
    - [ ] Implement logic to delete `command` subkey and then the `Magshare` shell key
    - [ ] Handle `registry.ErrNotExist` gracefully
    - [ ] Run tests and verify success
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Unregistration Logic (TDD)' (Protocol in workflow.md)

## Phase 4: Integration and Logging
- [ ] Task: Add diagnostic logging
    - [ ] Integrate `internal/logger` into registry functions
    - [ ] Ensure errors are logged as INFO or ERROR but do not return to caller if "silent" is desired (as per spec)
- [ ] Task: Final verification
    - [ ] Verify coverage > 80% for new registry code
    - [ ] Perform manual verification of context menu on a Windows machine
- [ ] Task: Conductor - User Manual Verification 'Phase 4: Integration and Logging' (Protocol in workflow.md)
