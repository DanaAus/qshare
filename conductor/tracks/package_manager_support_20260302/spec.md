# Specification: Package Manager Compatibility (Scoop, WinGet, Bun)

## Overview
This track aims to make QShare easily installable through popular package managers on Windows (Scoop, WinGet) and the JavaScript/TypeScript ecosystem (Bun/NPM).

## Requirements

### 1. Scoop Compatibility
- Create a Scoop manifest template (`qshare.json`).
- Ensure it points to a future release structure (e.g., GitHub Releases).
- Define architecture-specific binary paths (Windows x64).

### 2. WinGet Compatibility
- Create a WinGet manifest template (YAML).
- Include required fields: PackageIdentifier, PackageVersion, Publisher, License, etc.
- Define the installer type as `portable`.

### 3. Bun (NPM) Compatibility
- Create a `package.json` file.
- Implement a mechanism to distribute the Go binary through NPM/Bun.
- This typically involves a `bin` entry pointing to a wrapper script that detects the platform and executes the correct binary.
- Allow users to run `bun x qshare` or `bun install -g qshare`.

## Acceptance Criteria
- [ ] A valid Scoop manifest file is present in the repository.
- [ ] A valid WinGet manifest template is present.
- [ ] A `package.json` is configured such that `bun x .` (locally) or a published version would execute the QShare binary.
- [ ] Installation instructions for all three managers are added to the documentation.

## Out of Scope
- Actually publishing to the official Scoop/WinGet/NPM repositories (this requires manual steps and a public release).
- Automated CI/CD for publishing (will be a separate track).
