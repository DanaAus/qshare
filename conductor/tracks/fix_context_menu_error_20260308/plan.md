# Implementation Plan: Fix Context Menu Launch Error

This plan focuses on investigating the execution guard that blocks context menu launches and refactoring it to allow valid argument-based execution.

## Phase 1: Investigation and Guard Refactoring [checkpoint: 2977661]
- [x] Task: Locate and Analyze the Guard Message
    - [x] Grep the codebase for "This is a command line tool..." to find the exact guard logic.
    - [x] Analyze the conditions triggering the guard (e.g., checking if `os.Stdin` is a terminal).
- [x] Task: Refactor Guard Logic (TDD) (a564e1a)
    - [x] Write a test case that simulates a non-interactive launch with valid arguments.
    - [x] Modify the guard logic to allow execution if valid subcommands or arguments are present.
    - [x] Verify that double-clicking the EXE still shows the guard/help if NO arguments are provided (if that is the intended behavior) or enters interactive mode.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Investigation and Guard Refactoring' (Protocol in workflow.md) (2977661)

## Phase 2: Context Menu Compatibility [checkpoint: bd2152c]
- [x] Task: Ensure Argument Handling for Context Menu (cda1f38)
    - [x] Verify that `send "%1"` (the registry command) is correctly parsed by Cobra. (Refactored to use direct path instead)
    - [x] Ensure the interactive loop correctly pre-fills information when a path is passed as an argument.
- [x] Task: Terminal Persistence (33114b7)
    - [x] Ensure that if an error occurs during context-menu launch, the window doesn't close immediately (using the "Press Enter to exit" logic if needed).
- [x] Task: Conductor - User Manual Verification 'Phase 2: Context Menu Compatibility' (Protocol in workflow.md) (bd2152c)

## Phase 3: Final Verification
- [ ] Task: End-to-End Test
    - [ ] Manually register the context menu (if not already done).
    - [ ] Right-click a file and verify Magshare starts the send server without the error message.
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Final Verification' (Protocol in workflow.md)
