# Implementation Plan: Resumable & Secure Large File Streaming

## Phase 1: Security and Path Utilities [checkpoint: 15dbb1f]
- [x] Task: Implement Path Security Utilities b4b0079
    - [ ] Create `internal/handlers/security.go`
    - [ ] Add `SanitizePath(base, target string) (string, error)` to resolve the absolute path and ensure the target is within or equal to the intended base.
    - [ ] Write unit tests to verify traversal prevention.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Security and Path Utilities' (Protocol in workflow.md) 15dbb1f

## Phase 2: Progress ReadSeeker
- [ ] Task: Implement `ProgressReadSeeker`
    - [ ] Update `internal/handlers/progress_reader.go`
    - [ ] Create `type ProgressReadSeeker struct` that implements `io.ReadSeeker`.
    - [ ] Logic:
        - `Read`: Call underlying reader and update progress bar.
        - `Seek`: Call underlying seeker.
    - [ ] Ensure it correctly interacts with the existing `progressbar` library.
- [ ] Task: Write unit tests for `ProgressReadSeeker`
    - [ ] Verify that seeking doesn't break the reader and that progress is tracked correctly after seeks.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: Progress ReadSeeker' (Protocol in workflow.md)

## Phase 3: Refactor File Serving
- [ ] Task: Refactor `ServeFileWithProgress` in `internal/handlers/send.go`
    - [ ] Replace `io.Copy` logic with `http.ServeContent`.
    - [ ] Wrap the file in `ProgressReadSeeker` before passing to `ServeContent`.
    - [ ] Ensure correct `modtime` and `Content-Disposition` headers are handled by `ServeContent`.
- [ ] Task: Implement memory usage verification
    - [ ] Create a specialized integration test `internal/handlers/send_memory_test.go` that serves a large file (using a sparse file or mock) and monitors `runtime.MemStats`.
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Refactor File Serving' (Protocol in workflow.md)

## Phase 4: Integration and Range Verification
- [ ] Task: Verify Range Support
    - [ ] Create an integration test `internal/handlers/range_test.go` that uses `httptest` to request specific byte ranges and verifies the status code `206 Partial Content`.
- [ ] Task: Final end-to-end manual verification
    - [ ] Test with a large file using `curl --range` and verify the terminal progress bar and file integrity.
- [ ] Task: Conductor - User Manual Verification 'Phase 4: Integration and Range Verification' (Protocol in workflow.md)
