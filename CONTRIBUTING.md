# Contributing to Kai

Thanks for contributing to Kai. This document describes the workflow, branching model, and expectations for pull requests.

## Branching model

Kai follows the [Git Flow](https://git-flow.readthedocs.io/fr/latest/presentation.html) branching model:

- **`master`** — stable branch. Always deployable. Only receives merges from `develop` (or hotfix branches).
- **`develop`** — integration branch. All features are merged here first. Automatically published as the `dev` Docker image.
- **`feature/<name>`** — one branch per feature or fix, created from `develop`.
- **`hotfix/<name>`** — urgent fixes branched from `master`, merged back into both `master` and `develop`.
- **`release/<version>`** — optional stabilization branch cut from `develop` before tagging a release, for last-minute fixes and version bumps.

```
feature/xyz --> develop --> release/x.y.z --> master --(tag vX.Y.Z)--> production release
```

## Workflow

1. Create your branch from `develop`:
   ```bash
   git checkout develop
   git pull
   git checkout -b feature/my-feature
   ```
2. Make your changes, following the existing code style (see below).
3. Add or update tests for any behavior change.
4. Run checks locally before pushing:
   ```bash
   gofmt -l .
   go vet ./...
   go test ./...
   go build -o kai ./cmd/kai
   ```
5. Push your branch and open a pull request **against `develop`**.
6. Ensure CI (`.github/workflows/ci.yml`) passes on the PR.
7. Once approved and merged, delete the feature branch.

## Releasing

Releases are cut from `master` only:

1. Merge `develop` into `master` via pull request once it's stable.
2. Tag the release:
   ```bash
   git checkout master
   git pull
   git tag vX.Y.Z
   git push origin vX.Y.Z
   ```
3. The `docker-publish.yml` workflow builds and pushes the versioned image to GHCR and creates the GitHub Release automatically.

## Code style

- Standard Go formatting (`gofmt`); CI fails if any file is unformatted.
- Every file starts with a short English doc comment describing its purpose (`// Package x ...` or `// <FuncName> ...`).
- Keep changes minimal and focused; avoid unrelated refactors in the same PR.
- No secrets (tokens, keys) in code, comments, or committed `.env*` files. Only `.env.example` is committed.

## Commit messages

Keep commit messages short and descriptive, e.g.:

```
welcome: auto-detect channel per guild
config: drop WELCOME_CHANNEL_ID requirement
ci: run tests on develop and feature branches
```

## Reporting issues

Open a GitHub issue with a clear description, steps to reproduce (if a bug), and relevant logs. Never include your `DISCORD_TOKEN` in an issue or pull request.
