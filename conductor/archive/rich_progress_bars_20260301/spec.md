# Track: Rich Progress Bars

## Objective
Enhance the user experience of QShare by providing real-time visual feedback during file transfers. This includes transfer speed, estimated time of arrival (ETA), and a visual progress bar.

## Requirements
- Integrate github.com/schollz/progressbar/v3 for terminal-based progress tracking.
- Update StartSendServer in internal/handlers/send.go to display progress when a file is being sent.
- Update StartReceiveServer in internal/handlers/receive.go to display progress when a file is being received.
- Show current transfer speed in MB/s.
- Show ETA for completion.
- Ensure the progress bar works for both single files and directories (ZIP streaming).

## Technical Details
- Use progressbar.DefaultBytes or custom options to track bytes transferred.
- Wrap the io.Reader or io.Writer used in file operations with the progress bar.
- For directory transfers, calculate total size before starting the ZIP process if possible, or use a spinner/indeterminate bar.
