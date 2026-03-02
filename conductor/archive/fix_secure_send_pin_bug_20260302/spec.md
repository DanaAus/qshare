# Specification: Fix Secure Send PIN Bug

## Overview
When using `qshare send --secure`, the generated URL currently requires a `?pin=XXXX` query parameter to function. If a user opens the link directly in a browser without this parameter, the server immediately returns a raw text "Invalid PIN" error. This is confusing for users who expect a way to enter the PIN.

## Problem Description
- **Immediate Error**: The `downloadPath` handler in `internal/handlers/send.go` checks the PIN immediately upon request.
- **No UI**: There is no HTML page for the user to enter the PIN for downloads.
- **Inconsistent UX**: The `receive` command has a UI for PIN entry, but `send` does not.

## Proposed Solution
1. **Create PIN Entry UI**: Create a simple, embedded HTML template for PIN entry (`ui/pin.html`).
2. **Update Send Handler**: Modify `ServeFileWithProgress` or the `StartSendServer` handler to check for the PIN. If the PIN is missing or incorrect, instead of returning an error, it should serve the PIN entry HTML page.
3. **Handle PIN Submission**: The PIN entry page should submit the PIN (e.g., via a query parameter or a POST form) to the same URL.

## Acceptance Criteria
- [ ] Running `qshare send file.txt --secure` provides a URL.
- [ ] Opening the URL in a browser displays a PIN entry form.
- [ ] Entering the correct PIN allows the file download to begin.
- [ ] Entering an incorrect PIN shows an error message on the same PIN entry page.
- [ ] Direct access with `?pin=XXXX` still works (for CLI tools like `curl`).

## Out of Scope
- Implementing full session management or cookies (query params are enough for this ephemeral use case).
