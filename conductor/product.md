# Initial Concept
Instant, frictionless file sharing and receiving across a local network via CLI.

# Product Definition

## Vision
To provide the fastest and most seamless way to transfer files between devices on a local network without the need for cloud services, cables, or specialized software installations on client devices.

## Target Users
- Developers and power users who prefer the command line.
- Users who need to quickly move files from their PC to a mobile device or vice versa.
- Teams working in a local environment where cloud sharing is impractical or restricted.

## Core Features
- **Instant QR Generation:** Automatic detection of local IP and rendering of a scannable QR code.
- **Interactive TUI Mode:** A guided, prompt-based interface for easy configuration.
- **Memory-Efficient Streaming:** Direct streaming from disk to handle large files (10GB+) with low RAM usage.
- **Secure Mode:** PIN-protected transfers for sensitive files.
- **Web Dropzone:** A responsive HTML5 interface for receiving files from client devices.

## Success Criteria
- Files can be transferred between devices within seconds of starting the command.
- Zero-configuration requirement for the client device.
- Robust handling of large files without system performance degradation.
- High accessibility via standard package managers (Scoop, WinGet, npm/bun).
