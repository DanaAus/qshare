# Specification: Demo Mode (Stealth Mode)

## Overview
Implement a "Demo Mode" that fakes all connection information (IP address, QR Code, etc.) in both "receive" and "send" modes. This mode is designed for recording promotional content without exposing real network details. It must be "stealthy," meaning no indication of its activation should be visible in the normal UI.

## Requirements
- **Activation:** The mode must be activated via a specific, non-obvious command or flag.
- **Data Faking:**
    - Fake IP Address: Replace real local IP (e.g., `192.168.1.5`) with a standard fake one (e.g., `192.168.100.100` or `demo.magshare.local`).
    - Fake QR Code: The QR code rendered in the terminal must point to the fake URL.
    - Fake Port: Optionally fake the port number if needed.
- **Stealth:**
    - No "Demo Mode Active" banner or message should be displayed.
    - All other UI elements should remain identical to the standard mode.
- **Scope:**
    - Applies to `magshare send`.
    - Applies to `magshare receive`.
    - Applies to the interactive TUI mode.

## Technical Considerations
- The logic for IP detection and QR generation needs to be intercepted when Demo Mode is active.
- A global or context-aware flag should be used to toggle the faking logic.
- Ensure that the faking logic is only applied to the *display* and *terminal output*, not the actual server binding (the server must still work for the demo).

## Success Criteria
- Running `magshare send --demo` (or similar) displays a fake IP and QR code.
- Running `magshare receive --demo` displays a fake IP and QR code.
- No visible status message indicates the mode is active.
- The program still functions correctly (files can be sent/received if the *real* IP is known).
