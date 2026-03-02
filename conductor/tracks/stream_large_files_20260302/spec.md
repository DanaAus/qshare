# Specification: Stream Large Files (Memory Safety)

## Overview
This track implements memory-efficient file streaming for QShare. By utilizing Go's `io.Copy` and streaming directly from disk to the HTTP response writer, the application will maintain a low memory footprint (under 20MB) regardless of the file size being transferred (up to 10GB or more).

## Functional Requirements
- **Always Stream:** Every file transfer, regardless of size, must use a streaming approach.
- **Disk-to-Response Streaming:** Use `os.Open` to get a file handle and `io.Copy` to write directly to the `http.ResponseWriter`.
- **Unlimited Concurrency:** Support multiple simultaneous file transfers, relying on Go's standard library to manage resources, but ensuring each transfer remains memory-safe.
- **Progress Integration:** Integrate with the existing `progressbar/v3` implementation. Since `io.Copy` is used, we may need a custom `io.Reader` wrapper that updates the progress bar as bytes are read.
- **Advanced Error Handling:** Implement detailed logging for stream interruptions and ensure that errors (e.g., file read errors, connection drops) are handled gracefully without crashing the server.

## Non-Functional Requirements
- **Memory Safety:** Peak memory usage for file transfers must remain under 20MB even for files up to 10GB.
- **Performance:** Streaming should not introduce significant latency compared to memory-buffered transfers.

## Acceptance Criteria
- [ ] Transferring a 1GB+ file results in RAM usage < 20MB (verified via monitoring tools).
- [ ] Multiple concurrent transfers (at least 3) complete successfully without excessive memory growth.
- [ ] Terminal progress bar accurately reflects the transfer status for both small and large files.
- [ ] If a transfer is interrupted, the server logs the error and remains responsive for other requests.

## Out of Scope
- Implementing resumable transfers (planned for a future track).
- Compressing files on-the-fly (to be handled separately if needed).
