# Implementation Plan: Config-Aware Application Modes

## Phase 1: Configuration Loading Core [checkpoint: dbd85ea]
- [x] Task: Implement `LoadConfig` in `internal/workspace/config.go` f8ec0c0
    - [ ] Add `LoadConfig(path string) (Config, error)` function.
    - [ ] Implement logic to read the file and unmarshal JSON.
    - [ ] Add basic validation (port range, download dir existence).
    - [ ] Write unit tests in `internal/workspace/config_test.go` for successful loads and corrupted files.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Configuration Loading Core' (Protocol in workflow.md) dbd85ea

## Phase 2: Root Command Integration [checkpoint: ed3f047]
- [x] Task: Integrate config loading in `cmd/root.go` 7890bb6
    - [ ] Add a package-level variable `appConfig workspace.Config`.
    - [ ] Create an `initConfig()` function that resolves the config path and calls `LoadConfig`.
    - [ ] Register `initConfig` with `cobra.OnInitialize()`.
    - [ ] Handle load errors by logging a warning via the structured logger.
- [x] Task: Conductor - User Manual Verification 'Phase 2: Root Command Integration' (Protocol in workflow.md) ed3f047

## Phase 3: Subcommand and Interactive Mode Update [checkpoint: b2d87bc]
- [x] Task: Update `cmd/receive.go` 0569e19
    - [ ] Modify `Run` to use `appConfig.DownloadDir` and `appConfig.SecureMode` as defaults.
    - [ ] Ensure `receiveSecure` flag correctly overrides `appConfig.SecureMode`.
- [x] Task: Update `cmd/send.go` 0569e19
    - [ ] Modify `Run` to use `appConfig.SecureMode` as the default.
    - [ ] Ensure `sendSecure` flag correctly overrides `appConfig.SecureMode`.
- [x] Task: Update `internal/ui/interactive.go` 0569e19
    - [ ] Modify `RunInteractivePrompts` to accept `appConfig` as an argument.
    - [ ] Use config values as the initial state for the TUI form.
- [x] Task: Conductor - User Manual Verification 'Phase 3: Subcommand and Interactive Mode Update' (Protocol in workflow.md) b2d87bc

## Phase 4: Final Validation [checkpoint: cd5090e]
- [x] Task: End-to-end manual verification (manual)
    - [x] Set `secure_mode: true` in config, run `send` without flags, verify PIN is required.
    - [x] Run `send --secure=false`, verify PIN is NOT required.
    - [x] Verify `receive` uses the setup download directory.
- [x] Task: Conductor - User Manual Verification 'Phase 4: Final Validation' (Protocol in workflow.md) cd5090e
