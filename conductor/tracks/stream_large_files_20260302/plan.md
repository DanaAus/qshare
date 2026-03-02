# Implementation Plan: Stream Large Files (Memory Safety)

This plan focuses on ensuring that file transfers in QShare are memory-efficient by strictly utilizing streaming for all data operations.

## Phase 1: Verification and Benchmarking [checkpoint: 258b333]
- [x] **Task: Create a memory profiling test**
    - [x] Create a test that generates a large temporary file (e.g., 500MB+).
    - [x] Implement a test server and client that performs a transfer.
    - [x] Measure peak RSS memory usage during the transfer.
- [x] **Task: Analyze current `io.Copy` and `io.MultiWriter` performance**
    - [x] Investigate if `io.MultiWriter` with `progressbar/v3` causes any unexpected buffering.
    - [x] Verify that `os.File` is read in chunks (standard `io.Copy` behavior).
- [x] **Task: Conductor - User Manual Verification 'Verification and Benchmarking' (Protocol in workflow.md)**

## Phase 2: Streaming Optimization and Robustness [checkpoint: 8755585]
- [x] **Task: Implement custom `io.Reader` wrapper for progress**
    - [x] Write a `ProgressReader` that wraps an `io.Reader` and updates the progress bar on every `Read()` call.
    - [x] Replace `io.MultiWriter` with this wrapper to minimize write-side overhead.
    - [x] Ensure the wrapper handles EOF correctly.
- [x] **Task: Enhance error handling for stream interruptions**
    - [x] Update `ServeFileWithProgress` to listen for `r.Context().Done()` to stop reading and writing immediately if the client disconnects.
    - [x] Add detailed logging for different types of stream failures (connection reset, file read error).
- [x] **Task: Conductor - User Manual Verification 'Streaming Optimization and Robustness' (Protocol in workflow.md)**

## Phase 3: Concurrency and Performance Testing [checkpoint: a127ad9]
- [x] **Task: Implement concurrent transfer tests**
    - [x] Create a test case that initiates 3+ concurrent large file transfers.
    - [x] Assert that total memory usage of the process remains below 20MB.
- [x] **Task: Verify ZIP streaming for directories**
    - [x] Ensure that `ServeDirWithProgress` also maintains a low memory footprint by streaming individual files into the zip writer.
- [x] **Task: Conductor - User Manual Verification 'Concurrency and Performance Testing' (Protocol in workflow.md)**
