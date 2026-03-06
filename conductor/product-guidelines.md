# Product Guidelines

## CLI Philosophy
- **Speed First:** Commands should execute instantly and provide immediate feedback.
- **Interactivity:** Prefer guided prompts (TUI) for complex operations but support direct flags for automation.
- **Clarity:** Help messages and error output must be concise and actionable.
- **Zero-Config:** Works out of the box with sensible defaults (auto-IP detection, random high-range ports).

## User Experience (UX)
- **Visual Feedback:** Use progress bars for file transfers.
- **Platform Specifics:** Render QR codes in the terminal using ANSI escape codes for broad compatibility.
- **Error Recovery:** Gracefully handle network interruptions and provide clear retry instructions.
- **Web Interface:** Keep the receiver's dropzone lightweight and mobile-responsive.

## Branding
- **Name:** `magshare` (lowercase) for the command, but `MagShare` for formal documentation.
- **Colors:** Use high-contrast terminal colors; maintain a clean, professional aesthetic.
- **Logo:** Minimalist; often represented by its name in a bold terminal font.

## Technical Standards
- **Memory Safety:** Direct streaming of data to avoid excessive memory consumption.
- **Concurrency:** Efficiently manage parallel transfers and server timeouts.
- **Extensibility:** Use standard Go conventions to ensure the codebase is easy to contribute to.
