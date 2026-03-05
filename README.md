<div align="center">

# 🚀 magshare

[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)
<<<<<<< HEAD
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/DanaAus/magshare)](https://goreportcard.com/report/github.com/DanaAus/magshare)
=======
[![Go Report Card](https://goreportcard.com/badge/github.com/DanaAus/qshare)](https://goreportcard.com/report/github.com/DanaAus/qshare)
>>>>>>> 5c68f059b69091d2b7c148c46fbf3e3361db5158

**Instant, frictionless file sharing and receiving across your local network via CLI.**

</div>

---

**magshare** is a blazing-fast, terminal-based utility designed to eliminate the friction of transferring files between devices on the same local network. By spinning up an ephemeral local web server and rendering a QR code directly in your terminal, it allows any mobile device or PC to securely download or upload files in seconds—no cables, no cloud drives, and no client-side app installations required.

## ✨ Core Features

*   📱 **Instant QR Generation:** Automatically detects your local IP and renders an access URL as a scannable QR code directly within your terminal window.
*   ⚡ **Interactive TUI Mode:** Don't want to remember command flags? Launch the guided, prompt-based UI to effortlessly configure your sharing session.
*   🪶 **Memory-Efficient Streaming:** Engineered to handle ultra-large files (10GB+) by streaming directly from disk, keeping RAM consumption safely under 20MB.
*   🛡️ **Secure Mode:** Protect sensitive network transfers by requiring a dynamically generated 4-digit PIN before a download or upload can begin.
*   🌐 **Web Dropzone:** Running in "receive" mode serves a lightweight, responsive HTML5 dropzone to the client device for seamless drag-and-drop uploads.

---

## 📦 Installation

magshare is distributed as a standalone executable. Choose your preferred package manager below:

### Windows Native
**Using Scoop:**
```powershell
scoop bucket add magshare https://github.com/DanaAus/magshare
scoop install magshare
```

**Using WinGet:**
```powershell
winget install magshare
```

### Cross-Platform (Node / JavaScript Ecosystem)
If you already have a Node or Bun environment, you can install or run magshare globally:
```bash
bun x magshare
# OR
npm install -g magshare
```

### From Source (Go)
For developers who want to compile the latest version directly:
```bash
go install github.com/DanaAus/magshare@latest
```

---

## 🚀 Usage

magshare is designed to be completely intuitive. You can use the guided TUI, or pass commands directly.

### 1. Interactive TUI Mode
Simply run the command with no arguments to launch the interactive terminal interface:
```bash
magshare
```

### 2. Sending a File (Host ➔ Client)
Share a specific file. The client's browser will automatically prompt a direct download upon scanning the QR code.
```bash
magshare send ./my-file.txt
```
*Tip: Add the `--secure` flag to generate a one-time PIN for the transfer.*

### 3. Receiving Files (Client ➔ Host)
Spin up a temporary web server that allows anyone on the network to upload files directly to your current terminal directory.
```bash
magshare receive
```

---

## 🤝 Contributing
magshare is open-source software, and contributions are always welcome! If you'd like to improve the codebase, add a feature, or report a bug:
1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

## 📄 License
This project is distributed under the [Apache 2.0 License](LICENSE). Feel free to use, modify, and distribute it as you see fit.
