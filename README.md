<p align="center">
  <img src="frontend/public/brand-mark.png" alt="Beacon" width="96" height="112" />
</p>

<h1 align="center">Beacon</h1>

<p align="center">
  <strong>Self-hosted container observability</strong> — live logs, RBAC, audit, health metrics,<br/>
  alerts, notifications, and optional Kubernetes visibility.
</p>

<p align="center">
  <a href="#-getting-started"><img src="https://img.shields.io/badge/setup-%3C%202%20min-4B56D4?style=flat-square" alt="Setup under 2 minutes" /></a>
  <a href="#-license"><img src="https://img.shields.io/badge/license-MIT-4B56D4?style=flat-square" alt="MIT License" /></a>
  <a href="#-tech-stack"><img src="https://img.shields.io/badge/stack-Go%20%2B%20Vue%203-4B56D4?style=flat-square" alt="Go + Vue 3" /></a>
  <a href="#-features"><img src="https://img.shields.io/badge/runtime-Docker%20%7C%20Kubernetes-4B56D4?style=flat-square" alt="Docker | Kubernetes" /></a>
</p>

---

Beacon is a lightweight, team-safe observability console for Docker (and optionally Kubernetes).  
Deploy one container, stream logs in real time, manage who can see and control which workloads, and keep a full audit trail — without standing up a heavyweight observability stack.

> Average setup time: **under 2 minutes**.

---

## Table of contents

- [Why Beacon?](#-why-beacon)
- [Features](#-features)
- [Tech stack](#-tech-stack)
- [Getting started](#-getting-started)
- [Configuration](#️-configuration)
- [User roles](#-user-roles)
- [CLI](#-cli)
- [Local development](#-local-development)
- [Docker socket security](#-docker-socket-security)
- [Documentation](#-documentation)
- [Performance](#-performance)
- [Roadmap](#️-roadmap)
- [Contributing](#-contributing)
- [License](#-license)

---

## Why Beacon?

Most Docker log viewers are built for a single administrator. Beacon is built for the entire team.

| Principle | What it means in practice |
| --- | --- |
| **Team-first security** | Wildcard and regex RBAC so developers only see the containers they need |
| **Audit everything** | Full trail of who started, stopped, restarted, or deleted workloads |
| **Zero-config deploy** | No external database — single container, embedded Vue UI, SQLite on disk |
| **Small footprint** | Go backend with a typical ~30–50 MB RAM usage |
| **Modern UI** | Sharp LaunchOps-family chrome, system / light / dark themes |

---

## Features

### Real-time log streaming

- WebSocket live streaming with JWT subprotocol auth (no tokens in URLs)
- Infinite scroll and manual history loading
- Log search, highlighting, and safe HTML rendering
- Auto-scroll with reconnect handling
- RFC3339Nano timestamp filtering

### Advanced RBAC

- Wildcard permissions (`backend-*`) and full regex (`^prod-.*$`)
- Per-user start / stop / restart / delete rights
- Staff management and container-level isolation

### Audit & security

- JWT authentication with session invalidation on password change
- First-login forced password reset
- Login rate limiting
- Client access control for the web UI (origin validation)
- Native API clients via standard JWT
- Full audit trail for sensitive actions

See [Security & RBAC](docs/SECURITY.md) for details.

### Kubernetes (optional)

When `RUNTIME_MODE` is `kubernetes` or `both`:

- Namespace-scoped pod list, detail views, live logs, and shell access
- Workload hub for deployments, HPAs, services, and cluster events
- Namespace allowlist via `K8S_NAMESPACES`
- Graceful startup when the cluster is unreachable (UI shows connection status)

### Alerts & notifications

- Alert rules for CPU, memory, restarts, offline detection, and log patterns
- Delivery to Slack, Microsoft Teams, Discord, or custom HTTPS webhooks
- Per-channel event filters and a master enable/disable switch
- Version-update notifications when a newer Beacon release is available

### Modern UI

- Beacon tower brand mark with primary accent `#4B56D4`
- True-black application chrome; terminal/log palette kept distinct
- **Auto / Light / Dark** theme (follows system by default)
- Mobile-friendly layout with responsive sidebar
- Lazy-loaded routes for faster initial load

### System monitoring

- Host CPU and memory with live WebSocket stats
- Per-container metrics and historical charts
- Health dashboard with usage trends

### High performance

- ~30–50 MB RAM typical
- 10k+ log lines/sec throughput
- Embedded frontend in a single Go binary
- Optimized log rendering (capped DOM buffer)

---

## Tech stack

| Layer | Technology |
| --- | --- |
| Backend | Go 1.26, Echo, JWT, Gorilla WebSocket |
| Frontend | Vue 3, Vue Router, Vite, Chart.js, xterm.js |
| Storage | SQLite (`modernc.org/sqlite`, no CGO required) |
| Runtimes | Docker (Moby SDK), optional Kubernetes client-go |
| Packaging | Multi-stage Dockerfile → single `beacon` binary |

---

## Getting started

### Default login

| Username | Password |
| --- | --- |
| `admin` | `admin123` |

> [!WARNING]
> Change the default administrator password immediately after first login.

### Docker Compose (recommended)

```yaml
services:
  beacon:
    image: beacon:latest
    container_name: beacon
    ports:
      - "8888:8000"
    environment:
      - SECRET_KEY=your-secure-key-here
      - DB_PATH=/app/data/app.db
      - CLIENT_ACCESS=strict
      - ENV=production
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - /proc:/host/proc:ro
      - ./data:/app/data
    restart: unless-stopped
```

Or build and run from this repository:

```bash
docker compose up --build -d
```

Open **[http://localhost:8888](http://localhost:8888)**.

#### No-auth mode (development only)

```yaml
environment:
  - DISABLE_AUTH=true
```

### Direct Docker run

```bash
docker run -d \
  --name beacon \
  -p 8888:8000 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /proc:/host/proc:ro \
  -v "$(pwd)/data:/app/data" \
  -e SECRET_KEY=your-secure-key-here \
  -e DB_PATH=/app/data/app.db \
  -e CLIENT_ACCESS=strict \
  -e HOST_PROC=/host/proc \
  --restart unless-stopped \
  beacon:latest
```

Build the image locally:

```bash
make docker-build
# → beacon:latest
```

---

## Configuration

### Environment variables

| Variable | Description | Default |
| --- | --- | --- |
| `SECRET_KEY` | JWT signing secret. **Required in production.** | `secret-key-change-this` |
| `DB_PATH` | SQLite database path | `app.db` |
| `PORT` | HTTP listen port | `8000` |
| `DOCKER_HOST` | Docker daemon socket | `unix:///var/run/docker.sock` |
| `RUNTIME_MODE` | Runtime backend (`docker`, `kubernetes`, `both`) | `docker` |
| `K8S_NAMESPACES` | Comma-separated namespace allowlist. Empty = all namespaces allowed by cluster RBAC | _(empty)_ |
| `K8S_CONTEXT` | Kubernetes context override | _(empty)_ |
| `KUBECONFIG` | Path to kubeconfig | _(empty)_ |
| `K8S_API_SERVER` | Override Kubernetes API server URL from kubeconfig | _(empty)_ |
| `K8S_REWRITE_LOCALHOST` | Rewrite `127.0.0.1` / `localhost` in kubeconfig to `host.docker.internal` when running in Docker | auto (`true` in Docker) |
| `K8S_INSECURE_SKIP_TLS_VERIFY` | Skip TLS verification for the Kubernetes API (local/dev only) | `false` |
| `DISABLE_AUTH` | Disable auth (in-memory DB, no login). `NO_AUTH=true` is an alias | `false` |
| `CLIENT_ACCESS` | Restrict `/api` and `/ws` to web UI + native clients (`strict` or `off`) | `strict` |
| `ALLOWED_ORIGINS` | Extra browser origins for the Vue UI (comma-separated URLs) | _(empty)_ |
| `TRUST_PROXY` | Honor `X-Forwarded-*` headers when behind a reverse proxy | `false` |
| `DEBUG_MODE` | Enable verbose internal debug logs | `false` |
| `ENV` | Set to `production` to disable localhost origin bypass | _(empty)_ |
| `ALLOW_START` | Allow start action (also used as defaults in no-auth mode) | `false` |
| `ALLOW_STOP` | Allow stop action | `false` |
| `ALLOW_RESTART` | Allow restart action | `false` |
| `ALLOW_DELETE` | Allow delete action | `false` |
| `ALLOW_SHELL` | Allow interactive shell over WebSocket (`ALLOW_BASH` is an alias) | `false` |
| `EXCLUDE_CONTAINERS` | Comma-separated container names to hide from the dashboard | _(empty)_ |
| `HOST_PROC` | Host `/proc` mount path for real host CPU/RAM stats inside Docker (e.g. `/host/proc`) | auto (`/host/proc` if mounted) |

The Beacon container itself is **always hidden** (matched by its own service name or image).

### Notifications (Slack, Teams, Discord, custom webhooks)

Configure notifications from **Admin → Notifications** in the UI. Webhook URLs are stored in SQLite; there are no environment variables for notification delivery.

- **Master switch** turns delivery on or off for all channels.
- **Per channel** (Slack, Microsoft Teams, Discord, **Custom webhook**): webhook URL, enable/disable, and event filters.
- **Custom webhook** — POST JSON to any HTTPS endpoint (PagerDuty, n8n, Zapier, your own service).
- **Event types** (chosen per channel): container actions, security events, admin actions, intelligent alerts, and Beacon version updates.
- **Email (SMTP)** is planned; the UI lists it as coming soon.

**Custom webhook JSON payloads** (`Content-Type: application/json`):

- **Audit / admin events** — `type: "audit"` with `title`, `user`, `action`, `resource`, `status`, `message`, `timestamp`, `source`
- **Alert rules** — `type: "alert"` with `title`, `ruleId`, `severity`, `container`, `host`, `source`, `message`, `recovery`, `timestamp`

Legacy env-based webhook URLs are migrated into `notification_channels` on first startup if present.

### Kubernetes (Docker Compose)

When `RUNTIME_MODE` is `kubernetes` or `both`, Beacon needs API access to a cluster. Without a kubeconfig, Beacon **starts anyway** and shows a warning; the Pods page reports the connection error.

**Option A — mount your local kubeconfig:**

```yaml
environment:
  - RUNTIME_MODE=both
  - KUBECONFIG=/app/kube/config
  - K8S_NAMESPACES=default,prod
volumes:
  - ${HOME}/.kube:/app/kube:ro
```

**Option B — run Beacon inside the cluster** with a ServiceAccount that can list pods/namespaces (in-cluster config is detected automatically).

**Docker Desktop / local clusters:** kubeconfig often points at `https://127.0.0.1:PORT`. Inside a container that address is unreachable, so Beacon rewrites it to `host.docker.internal` and sets TLS `ServerName=localhost` to match the Docker Desktop certificate. If you still see TLS errors after `docker compose up -d --build`, set `K8S_INSECURE_SKIP_TLS_VERIFY=true`.

### Production checklist

1. Generate and set a strong `SECRET_KEY`:

```bash
openssl rand -base64 32
```

2. Set `ENV=production` (or `GO_ENV=production`).
3. Keep `CLIENT_ACCESS=strict`.
4. Run behind Nginx, Traefik, or Caddy with HTTPS and set `TRUST_PROXY=true`.
5. Change the default `admin` / `admin123` password on first login.
6. Restrict network access to trusted users only (Docker socket access is high privilege).

### Client access

| Client | How it connects |
| --- | --- |
| **Vue web UI** | Served by Beacon; sends `X-Beacon-Client: web` and passes origin checks |
| **Native mobile app** | Standard JWT after `POST /api/token` |
| **curl / random sites** | Blocked when `CLIENT_ACCESS=strict` |

Set `CLIENT_ACCESS=off` only for local debugging.

### Beacon Mobile (companion app)

The official Flutter client supports **Advanced Settings** when adding a server:

- **Resolve Host to IP** — `/etc/hosts`-style override when the URL hostname is not in public DNS
- **Skip TLS verification** — self-signed certificates
- **Custom HTTP headers** — Cloudflare Access and similar proxies

---

## User roles

### Administrator

Full container visibility, user management, audit log access, and container control.

### Staff member

Container / pod visibility is controlled with patterns such as:

```text
redis
backend-*
prod-*, *-app
^prod-.*$
```

Users only see containers matching their assigned rules.

---

## CLI

| Command | Description |
| --- | --- |
| `beacon` | Run full dashboard (default) |
| `beacon server` | Full API + WebSockets + embedded Vue UI |
| `beacon agent` | Fleet agent mode (local UI + API; optional `CONTROL_PLANE_URL`) |
| `beacon agent-only` | Headless agent (API/WebSockets only, no bundled UI) |
| `beacon reset-password <user> <pass>` | Reset a user password in SQLite |
| `beacon config` | Print non-secret configuration |
| `beacon version` | Print version |
| `beacon help [command]` | Show help |

```bash
beacon help agent-only
docker exec beacon beacon reset-password admin 'NewSecurePass1'
```

---

## Local development

Prerequisites: **Go 1.26+**, **pnpm**, and a local Docker daemon (for container features).

```bash
# Build frontend + backend binary
make build

# Install to PATH (optional)
make install          # → /usr/local/bin/beacon

# Run server (serves frontend/dist on :8000)
beacon
# same as: beacon server

# Tests
make test
```

Frontend dev server (separate terminal):

```bash
cd frontend && pnpm install && pnpm dev
```

Point the Vite proxy / API calls at `http://localhost:8000`. For unrestricted local API testing, set `CLIENT_ACCESS=off`.

Useful Make targets:

| Target | Action |
| --- | --- |
| `make build` | Build Vue UI + `beacon` binary |
| `make install` | Install binary to `/usr/local/bin` |
| `make docker-build` | Build `beacon:latest` image |
| `make up` / `make down` | Compose up/down |
| `make test` | Run Go tests |
| `make clean` | Remove binary and `frontend/dist` |

---

## Docker socket security

Beacon requires access to `/var/run/docker.sock`, which grants Docker API access to the application.

Recommended hardening:

- Reverse proxy with TLS
- Firewall or VPN for admin access
- Strong passwords and a rotated `SECRET_KEY`
- Keep `ALLOW_SHELL=false` unless explicitly needed
- Prefer read-only `/proc` mounts (`HOST_PROC`) rather than privileged mode

---

## Documentation

- [Architecture Overview](docs/ARCHITECTURE.md)
- [Security & RBAC](docs/SECURITY.md)

---

## Performance

| Metric | Value |
| --- | --- |
| RAM usage | ~30–50 MB |
| Log throughput | 10k+ lines/sec |
| Deployment | Single container |

---

## Roadmap

- Log retention controls
- Multi-host / fleet control plane
- External authentication providers (OIDC / SSO)
- Email (SMTP) notification channel

---

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/my-change`)
3. Make your changes with tests where practical (`make test`)
4. Open a pull request with a clear summary and test notes

Please do not commit secrets, local databases, or `frontend/node_modules`. See [`.gitignore`](.gitignore).

---

## License

MIT License. See [LICENSE](LICENSE).

---

<p align="center">
  Built for developers who want real-time Docker visibility<br/>
  without deploying an observability cathedral.
</p>
