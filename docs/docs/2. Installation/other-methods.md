---
title: Other Installation Methods
description: Install BearMount as a standalone binary with rclone WebDAV mounting on Linux, macOS, and Windows.
keywords: [bearmount, install, binary, rclone, webdav, linux, macos, windows, standalone]
---

# Other Installation Methods

This guide covers installing BearMount as a standalone binary and setting it up with rclone for WebDAV mounting on Linux, macOS, and Windows systems.

## Prerequisites

- Operating System: Linux (x64/ARM64), macOS (x64/ARM64), or Windows (x64)
- Available disk space: 100MB+ for the binary and basic operations
- Network access to Usenet providers and the internet
- **rclone** installed for WebDAV mounting ([installation guide](https://rclone.org/install/))

## Download Methods

### Method 1: GitHub Releases (Recommended)

Download the latest pre-built binary from the GitHub releases page:

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

<Tabs>
  <TabItem value="linux" label="Linux x64" default>

```bash
# Download latest release
wget https://github.com/WhispersOfJ/bearmount/releases/latest/download/bearmount-linux-amd64

# Make executable
chmod +x bearmount-linux-amd64

# Move to system path (optional)
sudo mv bearmount-linux-amd64 /usr/local/bin/bearmount
```

  </TabItem>
  <TabItem value="linux-arm" label="Linux ARM64">

```bash
# Download ARM64 version
wget https://github.com/WhispersOfJ/bearmount/releases/latest/download/bearmount-linux-arm64

# Make executable
chmod +x bearmount-linux-arm64

# Move to system path (optional)
sudo mv bearmount-linux-arm64 /usr/local/bin/bearmount
```

  </TabItem>
  <TabItem value="macos" label="macOS">

```bash
# Download for macOS
wget https://github.com/WhispersOfJ/bearmount/releases/latest/download/bearmount-darwin-amd64

# Make executable
chmod +x bearmount-darwin-amd64

# Move to system path (optional)
sudo mv bearmount-darwin-amd64 /usr/local/bin/bearmount
```

  </TabItem>
  <TabItem value="windows" label="Windows">

```powershell
# Download using PowerShell
Invoke-WebRequest -Uri "https://github.com/WhispersOfJ/bearmount/releases/latest/download/bearmount-windows-amd64.exe" -OutFile "bearmount.exe"

# Or download using curl (if available)
curl -L -o bearmount.exe https://github.com/WhispersOfJ/bearmount/releases/latest/download/bearmount-windows-amd64.exe
```

  </TabItem>
</Tabs>

_[Screenshot placeholder: GitHub releases page showing available downloads for different platforms]_

### Method 2: Build from Source

If you prefer to build from source or need a custom build:

```bash
# Prerequisites: Go 1.24.5+ and Bun
git clone https://github.com/WhispersOfJ/bearmount.git
cd bearmount

# Build everything (frontend + backend)
make

# The binary is now available as ./bearmount
```

## Configuration

### Initial Setup

1. **Create configuration directory**:

   ```bash
   mkdir -p ~/.config/bearmount
   cd ~/.config/bearmount
   ```

2. **Download sample configuration**:

   ```bash
   wget https://raw.githubusercontent.com/WhispersOfJ/bearmount/main/config.sample.yaml -O config.yaml
   ```

3. **Create required directories**:
   ```bash
   mkdir -p ./metadata ./logs
   ```

_[Screenshot placeholder: Directory structure showing config.yaml, metadata/, and logs/ folders]_

### Basic Configuration

Edit the `config.yaml` file with your settings. At minimum, you need to configure:

1. **NNTP Providers** (at least one is required):

   ```yaml
   providers:
     - host: "ssl-news.provider.com"
       port: 563
       username: "your_username"
       password: "your_password"
       max_connections: 20
       tls: true
   ```

2. **WebDAV Settings** (optional, uses defaults if not specified):

   ```yaml
   webdav:
     port: 8080
     user: "usenet"
     password: "usenet"
   ```

3. **Metadata Path**:
   ```yaml
   metadata:
     root_path: "./metadata"
   ```

## Running BearMount

### Basic Usage

```bash
# Run with default config (./config.yaml)
bearmount serve

# Run with specific config file
bearmount serve --config=/path/to/config.yaml

# Get help
bearmount --help
bearmount serve --help
```

_[Screenshot placeholder: Terminal showing successful BearMount startup with configuration summary and listening ports]_

### rclone WebDAV Mount Setup

Once BearMount is running, set up rclone to mount the WebDAV interface:

1. **Configure rclone remote**:

   ```bash
   rclone config create bearmount webdav \
     url=http://localhost:8080 \
     vendor=other \
     user=usenet \
     pass=$(rclone obscure "usenet")
   ```

2. **Create mount point**:

   ```bash
   # Linux/macOS
   sudo mkdir -p /mnt/remotes/bearmount
   sudo chown $USER:$USER /mnt/remotes/bearmount

   # Windows (PowerShell as Administrator)
   # Creates a network drive mapping
   ```

3. **Mount WebDAV**:

   ```bash
   # Linux/macOS
   rclone mount bearmount: /mnt/remotes/bearmount \
     --vfs-cache-mode writes \
     --vfs-read-chunk-size 32M \
     --buffer-size 64M \
     --allow-other \
     --daemon

   # Windows
   rclone mount bearmount: Z: \
     --vfs-cache-mode writes \
     --vfs-read-chunk-size 32M \
     --buffer-size 64M
   ```

### Verification

Verify the installation and rclone mount:

1. **Check BearMount web interface**: Open http://localhost:8080 in your browser
2. **Verify rclone mount**: `ls -la /mnt/remotes/bearmount` (Linux/macOS) or check `Z:` drive (Windows)
3. **Test health endpoint**: `curl http://localhost:8080/live` should return `OK`

## Next Steps

- [Configure NNTP Providers](../3. Configuration/providers.md)
