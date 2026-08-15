<div align="center">

# kai-bot

Kai is a modular Discord bot for community onboarding, moderation, task management, scheduling, and workflow automation.

[![Go](https://img.shields.io/badge/Go-1.26.6-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-ghcr.io%2Fsony--level%2Fkai--bot-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://ghcr.io/sony-level/kai-bot)

</div>

## Add Kai to your server

<p align="center">
  <a href="https://discord.com/api/oauth2/authorize?client_id=1537971825605349436&permissions=3072&scope=bot">
    <img src="https://img.shields.io/badge/Invite-Kai-5865F2?style=for-the-badge&logo=discord&logoColor=white" alt="Invite Kai to your server" />
  </a>
</p>

---

## Requirements

- Go 1.26.6+
- A Discord bot application and token

## Configuration

Copy `.env.example` to `.env.dev` (development) or `.env.prod` (production) and fill in the real value:

```env
DISCORD_TOKEN=your_token
```

Welcome messages are built into the bot, picked at random when a member joins.

`.env.dev` is loaded by default. Set `ENV=production` to load `.env.prod` instead.

## Run

Development:

```bash
go run ./cmd/kai
```

Production:

```bash
ENV=production go build -o kai ./cmd/kai
./kai
```

## Run with Docker

```bash
docker compose up -d
```

Uses the image published on GHCR (`ghcr.io/sony-level/kai-bot`) and reads `.env.prod`.

## Tests

```bash
go test ./...
```

## Permissions

The bot needs the `GuildMembers` privileged intent (enable it in the Discord Developer Portal under **Bot → Privileged Gateway Intents**) and permission to view channels / send messages.

## Branching & CI/CD

- `feature/*` — feature branches, merged into `develop` via pull request.
- `develop` — integration branch. Every push runs CI (`.github/workflows/ci.yml`) and publishes a `dev` Docker image (`.github/workflows/docker-publish.yml`).
- `master` — stable branch. Every push runs CI; no image is published automatically.
- `vX.Y.Z` tag on `master` — triggers the production release: versioned Docker images (`X.Y.Z`, `X.Y`, `latest`) are published and a GitHub Release is created.

See `CONTRIBUTING.md` for the full contribution workflow.

## Contributors

<p align="center">
  <a href="https://github.com/sony-level">
    <img src="https://github.com/sony-level.png" width="50" height="50" alt="sony-level" />
  </a>
  <br />
  <a href="https://github.com/sony-level">sony-level</a>
</p>

## Releases

### Version 0.1.0 — Welcome message

The first version of Kai:

- Announces itself (`✅ Kai is now online!`) in every server it is installed on.
- Sends a random welcome message to new members, mentioning them.

The target channel is auto-detected per server (the system channel, or the first available text channel) — no channel ID configuration needed.
