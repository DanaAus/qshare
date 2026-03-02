# Implementation Plan: Stream Large Files (Memory Safety)

This plan focuses on ensuring that file transfers in QShare are memory-efficient by strictly utilizing streaming for all data operations.

## Phase 1: Verification and Benchmarking
- [ ] **Task: Create a memory profiling test**
    - [ ] Create a test that generates a large temporary file (e.g., 500MB+).
    - [ ] Implement a test server and client that performs a transfer.
    - [ ] Measure peak RSS memory usage during the transfer.
- [ ] **Task: Analyze current `io.Copy` and `io.MultiWriter` performance**
    - [ ] Investigate if `io.MultiWriter` with `progressbar/v3` causes any unexpected buffering.
    - [ ] Verify that `os.File` is read in chunks (standard `io.Copy` behavior).
- [ ] **Task: Conductor - User Manual Verification 'Verification and Benchmarking' (Protocol in workflow.md)**

## Phase 2: Streaming Optimization and Robustness
- [ ] **Task: Implement custom `io.Reader` wrapper for progress**
    - [ ] Write a `ProgressReader` that wraps an `io.Reader` and updates the progress bar on every `Read()` call.
    - [ ] Replace `io.MultiWriter` with this wrapper to minimize write-side overhead.
    - [ ] Ensure the wrapper handles EOF correctly.
- [ ] **Task: Enhance error handling for stream interruptions**
    - [ ] Update `ServeFileWithProgress` to listen for `r.Context().Done()` to stop reading and writing immediately if the client disconnects.
    - [ ] Add detailed logging for different types of stream failures (connection reset, file read error).
- [ ] **Task: Conductor - User Manual Verification 'Streaming Optimization and Robustness' (Protocol in workflow.md)**

## Phase 3: Concurrency and Performance Testing
- [ ] **Task: Implement concurrent transfer tests**
    - [ ] Create a test case that initiates 3+ concurrent large file transfers.
    - [ ] Assert that total memory usage of the process remains below 20MB.
- [ ] **Task: Verify ZIP streaming for directories**
    - [ ] Ensure that `ServeDirWithProgress` also maintains a low memory footprint by streaming individual files into the zip writer.
- [ ] **Task: Conductor - User Manual Verification 'Concurrency and Performance Testing' (Protocol in workflow.md)**
