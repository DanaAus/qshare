# Specification: Create Winget Package (Temos.QShare)

## Overview
This track aims to create a fully functional and compliant Winget manifest (singleton) for the QShare terminal-based utility. This will enable Windows users to install and manage QShare using the standard `winget` CLI.

## Functional Requirements
- **Manifest Creation:** Generate a valid Winget YAML manifest.
- **Package Metadata:**
  - `PackageIdentifier`: `Temos.QShare`
  - `PackageName`: `QShare`
  - `PackageVersion`: `1.0.0`
  - `Publisher`: `Temos`
  - `License`: `Apache 2.0`
  - `ShortDescription`: `Instant local network file sharing with QR codes.`
  - `Tags`: `cli`, `file-sharing`, `network`, `qr-code`, `terminal`
- **Installer Configuration:**
  - `InstallerUrl`: `https://github.com/DanaAus/qshare/releases/download/v1.0.0/qshare.exe`
  - `InstallerType`: `portable`
  - `Architecture`: `x64`
- **Checksum:** Calculate and include the SHA256 hash of the `qshare.exe` version 1.0.0 binary.

## Non-Functional Requirements
- **Schema Validation:** The manifest must conform to the `winget-cli` manifest schema (v1.2.0).
- **Correct Paths:** Ensure the installer URL points to a valid release on GitHub.

## Acceptance Criteria
- [ ] A valid `winget/qshare.yaml` file exists with the specified metadata.
- [ ] The `PackageIdentifier` is correctly set to `Temos.QShare`.
- [ ] The `InstallerSha256` is accurate for the `v1.0.0` release of `qshare.exe`.
- [ ] The manifest is successfully validated against the Winget schema.

## Out of Scope
- Submitting the package to the official `microsoft/winget-pkgs` repository.
- Creating manifests for other operating systems or package managers.