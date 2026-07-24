# AGENTS.md

Guidance for AI coding agents working in this repo. Read `CLAUDE.md` first for frontend
coding standards; this file covers repo-wide facts that aren't obvious from the code alone,
several of them hard-won during the 2026-07-24 rebrand.

## This is BearMount, a fork of javi11/altmount - rebranded 2026-07-24

Module path is `github.com/WhispersOfJ/bearmount` (was `github.com/javi11/altmount`), GitHub
repo is `WhispersOfJ/bearmount` (renamed from `WhispersOfJ/altmount` the same day), binary/CLI
is `bearmount` (was `altmount`, `cmd/altmount/` → `cmd/bearmount/`). The `upstream` git remote
still correctly points at `https://github.com/javi11/altmount.git` - **never rename or repoint
it**, `.github/workflows/sync-upstream.yml` depends on it for the daily merge, and every
`javi11/altmount` reference inside that workflow (and a couple of doc lines: `docker pull
javi11/altmount:latest`, a Docker Hub footer link, a custom-domain comment in
`docs/docusaurus.config.ts`) is a **deliberate, correct** credit/reference to the real upstream
project, not a leftover to rename. This fork has no Docker Hub namespace of its own
(`DOCKERHUB_IMAGE: laris11/altmount` in `release.yml`/`dev-image.yml` is also deliberately
unchanged for the same reason) - don't invent one.

Versioning: this fork's own tags are `fork-vX.Y.Z` (currently `fork-v0.1.0`), not plain
`vX.Y.Z` - upstream already owns that tag namespace (`v0.0.1-alpha2` through `v0.3.2` and
climbing via the daily sync), and a same-named tag would collide. `internal/version.Version`
is set by ldflags at build time from `git describe --tags` (see `Makefile`/`build-cli.yml`),
so a fresh checkout with no `fork-v*` tag reachable from `HEAD` will report a stale/wrong
version - check `git tag --list 'fork-v*'` before assuming the binary's own `--version` output
means anything.

## `vendor/` is fully committed - `go build`/`go vet`/`go test` need no network

This repo vendors all ~1780 Go dependencies (`go mod vendor`, ~85MB). Go auto-detects
`vendor/modules.txt` and uses `-mod=vendor` by default - no `GOFLAGS` or `go.mod` change
needed. **`.gitignore`'s broad patterns (`*debug*`, `go.work`/`go.work.sum`, `.env`,
`coverage.*`, etc.) will silently drop legitimate vendored files that happen to match** (e.g.
`golift.io/starr/debuglog/roundtripper.go`, several packages' own `go.work` files, a package
literally named `coverage.go`) - this broke CI once already (`cannot find module providing
package golift.io/starr/debuglog: import lookup disabled by -mod=vendor`), silently, because
`git add vendor/` just skipped the file with no error. The fix is the `!vendor/**` line at the
end of `.gitignore` - **never delete or reorder that line**, and if you ever regenerate vendor
after an ignore-pattern change, diff `find vendor -type f | wc -l` against `git ls-files vendor
| wc -l` to catch a silent drop before it reaches CI.

One vendored file is deliberately absent: `vendor/github.com/go-openapi/spec/appveyor.yml`
(dead upstream CI config carrying a hardcoded, real-looking Slack incoming-webhook URL that
trips GitHub's push protection). If a future `go mod vendor` reintroduces it, delete it again
before committing rather than trying to allow-list the secret.

## `rapidyenc` hard-limits the CLI release matrix - don't re-add these targets

`vendor/github.com/mnightingale/rapidyenc` (cgo yEnc encode/decode) ships prebuilt static
libraries for exactly four targets - confirm via its `cgo.go` LDFLAGS if in doubt:

- `linux/amd64`, `linux/arm64`, `windows/amd64`, `darwin` (universal)

**`linux/386`, `linux/arm`, `windows/386`, and `windows/arm64` cannot link and were reverted
from `build-cli.yml`** after CI caught `windows/386` failing with `undefined symbol:
_rapidyenc_decode_kernel` etc. (the other three would fail identically, untested only because
`fail-fast` cancelled them first). This is a real upstream limitation of the dependency, not a
CI config problem - don't re-add these GOOS/GOARCH combinations without first confirming
rapidyenc has shipped a library for them. `build-cli.yml`'s `strategy.fail-fast: false` is
there so a future infeasible target only fails its own job, not every other target's too.

## CI workflows

- **`bearmount-fixes-ci.yml`** (renamed from `altmount-fixes-ci.yml`): the real build/vet/lint/test
  gate for this fork's own work - runs on every push, never touches a registry. Uses
  `golangci-lint-action@v9` (not `v6` - `v6` only supports golangci-lint v1.x, which stopped
  receiving releases; this repo's own `.golangci.yml` already declares `version: '2'`, so the
  action version has to match). If lint ever fails with "the Go language version (goX.Y) used
  to build golangci-lint is lower than the targeted Go version", the action is stale again -
  check `gh release list --repo golangci/golangci-lint` for the real latest and make sure the
  action major version supports it.
- **`sync-upstream.yml`**: daily merge from `javi11/altmount` main, `-X ours` so this fork's
  fixes win any conflicting hunk. Never rework its `javi11/altmount` references (see above).
- **`build-cli.yml`** / **`dev-image.yml`** / **`release.yml`** / **`manual-build.yml`**: CLI
  binary and Docker image builds. `build-cli.yml`'s matrix is the rapidyenc-constrained one
  above.
- **`deploy-docs.yml`**: Docusaurus → GitHub Pages, path-filtered to `docs/**` changes only
  (plus `workflow_dispatch` for manual runs). **GitHub Pages had to be enabled via API**
  (`gh api -X POST repos/WhispersOfJ/bearmount/pages -f build_type=workflow`) - it wasn't
  configured at all after the repo rename (forks don't inherit Pages settings). Live at
  `https://whispersofj.github.io/bearmount/`.

## Security

Private vulnerability reporting is enabled for this repo (`gh api
repos/WhispersOfJ/bearmount/private-vulnerability-reporting` → `{"enabled": true}`) and
`SECURITY.md` points reporters at it. Don't disable it without a reason - it's the primary
private-disclosure channel documented there.
