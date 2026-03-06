# Technology Stack

## Core Technologies
- **Language:** Go (1.25.0) - Provides static compilation, cross-platform support, and strong networking primitives.
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra) - Industry standard for Go-based CLI tools with robust flag handling and help generation.
- **TUI/Interactivity:** [Charmbracelet huh](https://github.com/charmbracelet/huh), [Bubbletea](https://github.com/charmbracelet/bubbletea), and [Lipgloss](https://github.com/charmbracelet/lipgloss) - Modern, aesthetic terminal interfaces and layout.

## Networking & Security
- **HTTP Server:** Go's standard `net/http` package for ephemeral, high-performance file serving.
- **Real-time Communication:** [Gorilla WebSocket](https://github.com/gorilla/websocket) - Used for real-time progress updates and bidirectional signaling.
- **IP Discovery:** Custom logic using Go's `net` package for reliable local network interface discovery.

## Assets & UI
- **QR Rendering:** [qrterminal](https://github.com/mdp/qrterminal) - Renders high-quality QR codes directly into ANSI-compliant terminals.
- **Frontend Assets:** Embedded HTML/CSS using Go's `embed` package for a zero-dependency binary.
- **Progress Tracking:** [progressbar/v3](https://github.com/schollz/progressbar) - Provides visual streaming feedback for large file transfers.

## Distribution & Packaging
- **Node.js/Bun Wrapper:** Lightweight JavaScript entry point (`magshare.js`) to facilitate distribution via `npm` and `bun`.
- **Windows Packaging:** Native support for [Scoop](https://scoop.sh/) and [WinGet](https://github.com/microsoft/winget-cli) for Windows power users.
