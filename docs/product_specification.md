# Product Specification Document: magshare (Instant Local Network Share & Drop)

**Document Version:** 1.0
**Target Platform:** Windows (Primary)
**App Category:** CLI / Networking Utility

---

## 1. Product Overview
### 1.1 The Problem
Transferring files between a PC and a mobile device, or two PCs on the same local network, is unnecessarily frictionless. Users resort to sending files to themselves via email, using messaging apps (which compress media), uploading to cloud storage (Google Drive/Dropbox), or dealing with complex SMB network sharing configurations.

### 1.2 The Solution
**magshare** is a terminal-based utility that eliminates this friction. By executing a single command, the tool spins up an ephemeral (temporary) local web server, binds to the machine's local IP address, and generates a QR code directly in the terminal. Any device on the same Wi-Fi network can scan the QR code to securely download the file or upload files directly to the host PC. 

---

## 2. Core Features & Capabilities

### 2.1 "Send" Mode (Host -> Client)
*   **Single File Sharing:** Serve a specific file. The client's browser prompts a direct download.
*   **Directory Sharing (Zip-on-the-fly):** If the user targets a folder (`magshare send ./documents`), the tool dynamically compresses the folder into a `.zip` stream as the client downloads it, saving local disk space.
*   **QR Code Terminal Rendering:** Generates an ASCII/ANSI-based QR code in the console so mobile devices can connect instantly without typing an IP address.

### 2.2 "Receive" Mode (Client -> Host)
*   **Dropzone UI:** Running `magshare receive` serves a minimalist, responsive HTML5 web page to the client. It features a drag-and-drop zone and a file-picker button.
*   **Direct Saving:** Files uploaded from the client are saved directly into the directory where the terminal command was executed.

### 2.3 Network Interface Auto-Discovery
*   Intelligently scans the host machine's network interfaces.
*   Ignores virtual adapters (e.g., Docker, WSL, VirtualBox) and selects the active Wi-Fi or Ethernet IPv4 address to construct the sharing URL.
*   Allows manual interface selection if multiple active networks are found.

### 2.4 Ephemeral Lifecycle (Auto-Close)
*   The server is not meant to run permanently. 
*   **Download limit:** Automatically shuts down the server once the file is successfully downloaded once (default behavior).
*   **Timeout limit:** Automatically closes the server if no connection is made within 5 minutes.

---

## 3. Security & Privacy Methods

Opening a port on a local network can expose data, especially on public Wi-Fi (e.g., coffee shops). magshare implements several security layers:

### 3.1 Network Isolation
*   **Randomized Ports:** Never uses standard ports (like 80 or 8080). Randomly assigns a high, unused port (e.g., `49152-65535`) for each session.
*   **Unpredictable URLs:** Instead of serving the file at the root `http://192.168.1.5:54321/`, the file path is protected by an automatically generated cryptographic hash (e.g., `http://192.168.1.5:54321/download/a7b8c9d0`). 

### 3.2 Access Control & Authentication
*   **One-Time PIN (Optional):** If the `--secure` flag is used, the CLI generates a 4-digit PIN. The client must enter this PIN on a web page before the file download begins.
*   **Strict CORS Policy:** Prevents Cross-Origin Resource Sharing. Only the generated specific domain/IP can interact with the server.

### 3.3 Execution Safeguards
*   **Path Traversal Prevention:** Hardcoded security checks ensure that clients cannot modify the URL to access files outside the specifically shared directory (e.g., blocking requests containing `../`).
*   **Upload Validation:** In `receive` mode, prevents execution of malicious uploads by sanitizing file names and stripping potentially dangerous extensions if configured.

---

## 4. User Workflows (CLI Commands)

**Scenario 1: Sending a file**
```bash
> magshare send presentation.pdf

[Network] Using active interface: Wi-Fi (192.168.1.15)
[Server]  Started on port 51234
[URL]     http://192.168.1.15:51234/d/8f4c2e
[QR]      (ASCII QR CODE RENDERED HERE)

Status: Waiting for connection... (Server will close after 1 download)
```

**Scenario 2: Receiving files securely**
```bash
> magshare receive --secure

[Network] Using active interface: Wi-Fi (192.168.1.15)
[Server]  Dropzone started on port 60112[Auth]    PIN REQUIRED: 4892
[QR]      (ASCII QR CODE RENDERED HERE)

Status: Ready to receive files... Press Ctrl+C to stop.
```

---

## 5. Technical Architecture & Tech Stack

To ensure the tool is incredibly easy to install on Windows (no dependencies, no Python environments, no Node modules), it will be built as a **single statically compiled executable**.

### 5.1 Programming Language
*   **Go (Golang):** The undisputed best language for CLI networking utilities. It compiles to a tiny, fast `.exe` with zero runtime dependencies. Its standard library handles HTTP servers and network interfaces natively and securely.

### 5.2 Key Libraries & Packages (Go Ecosystem)
*   **CLI Framework:** `github.com/spf13/cobra` (Industry standard for building clean CLI commands, flags, and help menus).
*   **QR Code Generator:** `github.com/mdp/qrterminal` (Renders QR codes directly into the terminal using ANSI escape codes and block characters).
*   **Web Server:** Go's native `net/http` (Extremely robust, handles concurrent uploads/downloads flawlessly).
*   **In-Memory HTML:** Go's `embed` and `html/template` standard packages. The HTML/CSS/JS for the client "Receive Dropzone" UI will be compiled directly *into* the binary. No external asset folders required.

### 5.3 Client-Side Tech (The UI served to mobile/other PCs)
*   **HTML5 / CSS3:** Minimal, dependency-free CSS (no React, no heavy frameworks). 
*   **Vanilla JS:** To handle drag-and-drop events, file chunking, and progress bars during uploads.
