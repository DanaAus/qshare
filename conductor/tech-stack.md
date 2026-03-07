# Technology Stack: magshare

## Core Language & Runtime
*   **Go (Golang):** Chosen for its excellent networking support, static typing, and ability to compile to a single, zero-dependency binary. It's ideal for a cross-platform CLI tool.

## Frameworks & Libraries
*   **CLI Framework:** **spf13/cobra** - For robust command-line argument parsing and command structure.
*   **Terminal UI (TUI):** **Charmbracelet ecosystem (Bubbletea, Huh, Lipgloss)** - For building modern, interactive, and visually appealing terminal interfaces.
*   **Networking:** **Standard net/http** - For a lightweight and performant local web server.
*   **Logging:** **Custom Structured Logger** - Thread-safe, leveled logging with hybrid terminal/file output and rich metadata.
*   **Real-time Communication:** **Gorilla WebSocket** - To provide live progress updates and a responsive user experience in the web UI.
*   **QR Rendering:** **mdp/qrterminal** - To display scannable QR codes directly within the terminal.

## Frontend (Dropzone UI)
*   **HTML5 / Vanilla JS:** To keep the web interface lightweight and fast, minimizing dependencies on external frameworks for the client browser.
*   **Embedded Assets:** **Go embed** - To bundle all web assets (HTML, CSS, JS) directly into the magshare binary.

## Infrastructure & Tooling
*   **Build System:** **Standard Go Toolchain** (e.g., `go build`, `go test`).
*   **Version Control:** **Git**.
