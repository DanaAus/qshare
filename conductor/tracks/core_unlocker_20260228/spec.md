# Specification: core_unlocker_20260228

## Overview
Build a lightning-fast Windows CLI tool in Rust that identifies locking processes and provides an interactive interface to terminate them, resolving "file in use" errors.

## Functional Requirements
- **Lock Detection:** Use Windows Win32 / Restart Manager APIs to identify processes locking a file or directory.
- **Recursive Scanning:** Support recursive scanning of directories to find locks in all nested files.
- **Process Information:** Retrieve and display Process IDs (PIDs) and process names.
- **Interactive Termination:** Allow users to interactively select which locking processes to kill.
- **System Protection:** Identify and warn before terminating critical system processes.
- **Audit Logging:** Maintain a timestamped log of all lock detections and termination actions.

## Non-Functional Requirements
- **Performance:** Lock detection should typically complete in under 50ms.
- **Portability:** Compile to a single, zero-dependency .exe binary for Windows.
- **Efficiency:** High memory efficiency, especially during deep directory scans.
- **Visuals:** Use ANSI colors for a modern, tech-forward terminal experience.

## Acceptance Criteria
- [ ] Correctly identifies processes locking a single file.
- [ ] Correctly identifies locks across a recursive directory tree.
- [ ] Successfully terminates selected processes with user confirmation.
- [ ] Prevents accidental termination of protected system processes.
- [ ] Operates as a single standalone binary on Windows 10/11.

## Out of Scope
- Network file lock resolution.
- Remote process management (local only).
