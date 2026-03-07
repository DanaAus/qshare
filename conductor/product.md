# Initial Concept
magshare is a terminal-based utility designed for instant, frictionless file sharing and receiving across a local network. It leverages ephemeral web servers and terminal-rendered QR codes to allow any device on the same Wi-Fi network to securely download or upload files without complex configuration.

# Product Guide: magshare

## Target Audience
The primary users are everyday users and cross-device users who need a quick and easy way to transfer files between their PCs and mobile devices on the same local network without relying on cloud services or physical cables.

## Core Goals
1.  **Frictionless Sharing:** Eliminate the setup burden for transferring files.
2.  **Secure Transfer:** Provide mechanisms like PIN protection to ensure transfers are authorized.
3.  **Simplicity:** Maintain a clean, intuitive interface that anyone can use.

## Key Features
*   **QR Code Discovery:** Automatically generate a scannable QR code in the terminal for instant connection.
*   **Bi-directional Sharing:** Support both sending files from the host to a client and receiving files from a client to the host via a web-based "Dropzone."
*   **TUI/Interactive Mode:** Provide a guided, prompt-based terminal user interface (using the Charmbracelet stack) for configuration.
*   **Interactive Onboarding:** Guided first-run experience to configure default download locations and security preferences.
*   **Structured & Leveled Logging:** Provides high-signal terminal output and detailed ISO-formatted diagnostic logs in the workspace, categorized by severity (DEBUG, INFO, WARN, ERROR).
*   **Secure Mode:** Optional PIN-based authentication for sensitive transfers.
*   **Flexible Configuration:** Explicitly customize ports and security PINs via terminal flags or the interactive UI.
*   **Resumable & Efficient Streaming:** Supports paused/resumed downloads (Range headers) for mobile compatibility and ultra-large files (10GB+), maintaining a low memory footprint (under 20MB).

## Technical Constraints
*   **Zero External Dependencies:** The application should be distributed as a single, standalone binary with no external runtime requirements.
*   **Low Memory Footprint:** Efficient resource usage is critical for performance and reliability across different hardware.
