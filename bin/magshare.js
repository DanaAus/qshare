#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');
const os = require('os');

// Determine binary name based on platform
const platform = os.platform();
const arch = os.arch();

let binaryName = 'magshare';
if (platform === 'win32') {
    binaryName += '.exe';
}

// In a real published package, the binaries would be in a specific folder
// for the architecture. For local development/testing, we use the root.
const binaryPath = path.join(__dirname, '..', binaryName);

const args = process.argv.slice(2);

const child = spawn(binaryPath, args, {
    stdio: 'inherit',
    shell: false
});

child.on('error', (err) => {
    if (err.code === 'ENOENT') {
        console.error(`Error: magshare binary not found at ${binaryPath}`);
        console.error('Please ensure the project is built using "go build"');
    } else {
        console.error(`Error: ${err.message}`);
    }
    process.exit(1);
});

child.on('exit', (code) => {
    process.exit(code || 0);
});
