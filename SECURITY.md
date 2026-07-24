# Security Policy

BearMount mounts and streams user Usenet content over WebDAV/FUSE and integrates with
Radarr/Sonarr/Prowlarr and other automation tools, so it handles credentials (indexer/provider
API keys, WebDAV/rclone secrets) and exposes a network-facing API. We take reports of security
issues seriously and appreciate responsible disclosure.

## Supported Versions

BearMount is a rolling-release fork (no long-term-support branches). Only the latest release
and the `main` branch are supported with security fixes.

| Version           | Supported          |
| ------------------ | ------------------ |
| Latest release      | :white_check_mark: |
| `main` (unreleased) | :white_check_mark: |
| Older releases       | :x:                 |

If you are running an older tagged release, please upgrade to the latest release before
reporting - we cannot commit to backporting fixes to unsupported versions.

## Reporting a Vulnerability

**Please do not open a public GitHub issue for security vulnerabilities.**

The preferred way to report a vulnerability is through GitHub's private vulnerability
reporting feature, which is enabled for this repository:

1. Go to the [Security tab](https://github.com/WhispersOfJ/bearmount/security).
2. Click **Report a vulnerability**.
3. Fill in as much detail as you can (see "What to include" below).

This creates a private advisory visible only to the maintainer and you, with no public
disclosure until a fix is ready.

If you are unable to use GitHub's private reporting for any reason, you may instead email
**WhispersofJ@gmail.com** with the same information. Please do not send exploit details over
any unencrypted or public channel.

### What to include

To help us triage and fix the issue quickly, please include:

- A clear description of the vulnerability and its impact (e.g. authentication bypass, path
  traversal, SSRF, credential exposure, remote code execution).
- Steps to reproduce, including the affected endpoint/route, config, or code path.
- The version/commit you tested against (`bearmount version` or the Docker image tag/digest).
- A proof-of-concept or minimal reproduction, if you have one.
- Whether the issue requires authentication, and with what privilege level.

### Scope

In scope: the `bearmount` server/CLI itself, its REST API and WebDAV/FUSE mount handling, its
authentication/session logic, its integrations with Radarr/Sonarr/Prowlarr/Stremio, and its
Docker image/build pipeline.

Generally out of scope: vulnerabilities in third-party dependencies that are already publicly
disclosed upstream (please report those to the dependency's own project first, and let us know
so we can track the update); vulnerabilities that require an already-fully-compromised host or
physical access; social engineering.

### What to expect

- **Acknowledgement**: we aim to respond within 5 business days of a report.
- **Triage**: we'll confirm whether the report is a valid vulnerability and its severity, and
  keep you updated as we investigate.
- **Fix and disclosure**: once a fix is ready, we'll coordinate a disclosure timeline with you.
  We credit reporters in the advisory/release notes unless you prefer to remain anonymous.

Since this is a small, community-maintained fork (not a funded security team), response and fix
times may vary with severity and maintainer availability - critical issues (auth bypass, RCE,
credential leakage, SSRF) are prioritized over lower-severity findings.
