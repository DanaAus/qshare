# Project Tracks

This file tracks all major tracks for the project. Each track has its own detailed plan in its respective folder.

---

- [ ] **Track: Write a Go main function for first run setup and crash logging. Silently create a Magshare workspace in the system user config directory (e.g: The local appdata folder on Windows). Redirect standard output to both the terminal and a temporary log file. Use defer and recover so if the app crashes, it moves the temp log to a logs folder in the workspace, prints the path, and keeps the window open for five seconds. On success, simply delete the temporary file.**
*Link: [./tracks/workspace-setup_20260307/](./tracks/workspace-setup_20260307/)*
