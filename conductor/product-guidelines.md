# Product Guidelines: FileInBreach CLI

## Communication Style
- **Direct and Efficient:** All output should be concise and focused on the task at hand. Avoid verbose explanations and prioritize high-signal information for power users.
- **Immediate Feedback:** Provide instant visual confirmation for every user action and detection result.

## Branding and Visual Identity (CLI)
- **Modern Tech-Forward:** Utilize ANSI colors to clearly differentiate between information, warnings, and errors.
  - **Info:** Cyan/Blue
  - **Success:** Green
  - **Warning:** Yellow
  - **Error:** Red
- **ANSI Styling:** Use bold and underlined text sparingly to highlight critical data points like Process IDs (PIDs) or file paths.

## User Interface and Visual Language
- **Tabular Data Display:** Process lists and lock information should be presented in well-aligned tables for quick scanning.
- **Hierarchical Process Trees:** When multiple processes are involved, display them in a tree structure to show parent-child relationships.
- **Streamlined Action Logs:** Provide a clean, timestamped log of all actions performed during the session.

## User Experience (UX) Principles
- **Safety-First Design:** Mandatory confirmations are required before any process is terminated. The user must always have the final say.
- **Predictability:** The tool should follow a consistent interaction pattern: Scan -> Identify -> Confirm -> Resolve.
- **Transparency:** Clearly communicate why a process is being targeted and what the potential impact of killing it might be (especially for system processes).

## Performance Targets
- **Initial Scan:** Detection of locking processes should be near-instantaneous (under 100ms for single files).
- **Responsive Interface:** The UI should remain responsive and provide visual feedback even during recursive directory scans.
