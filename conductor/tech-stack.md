# Technology Stack: FileInBreach CLI

## Core Languages
- **Rust:** The primary language for the tool, chosen for its safety, performance, and excellent Windows API interoperability.

## Windows API & System Libraries
- **windows-rs (Official Windows Crate):** Used for direct interaction with the Windows Win32 and Restart Manager APIs for detecting file locks.
- **sysinfo:** For retrieving process information, names, and memory usage.
- **winapi / windows-sys:** As fallback for low-level system calls if required by specific lock detection scenarios.

## CLI & User Interface
- **clap (Command Line Argument Parser):** The robust framework for handling CLI flags, subcommands, and generating help documentation.
- **inquire:** For creating the interactive prompts and selection menus requested in the UX guidelines.
- **indicatif:** To provide progress bars and spinners during long-running recursive directory scans.
- **colored or console:** For ANSI color support to meet the tech-forward branding requirements.

## Build & Distribution
- **Cargo:** Rust's built-in package manager and build system.
- **Static Linking:** All dependencies will be statically linked into a single, portable .exe binary for easy distribution on Windows systems.

## Testing & Quality Assurance
- **	okio (Optional):** For asynchronous operations if multiple concurrent scans are required.
- **Standard Rust Test Suite:** For unit and integration testing.
