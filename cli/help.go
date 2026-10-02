package cli

import "fmt"

// PrintHelp prints CLI usage. Optional topic shows help for a single command.
func PrintHelp(topic []string) {
	if len(topic) > 0 {
		switch topic[0] {
		case "server":
			fmt.Println(`beacon server: run the full Beacon dashboard (API, WebSockets, embedded Vue UI).

Environment variables configure auth, permissions, and ports. See README.md.`)
			return
		case "agent":
			fmt.Println(`beacon agent: run Beacon as a fleet agent with the local UI enabled.

Sets BEACON_MODE=agent. Optional CONTROL_PLANE_URL registers this host with a remote Beacon control plane (future).`)
			return
		case "agent-only":
			fmt.Println(`beacon agent-only: headless agent (API + WebSockets only, no bundled web UI).

Sets BEACON_MODE=agent and BEACON_AGENT_ONLY=true. Use when a remote UI or mobile app connects to this host.`)
			return
		case "reset-password":
			fmt.Println(`beacon reset-password <username> <new-password>

Reset a user password in the SQLite database. Invalidates existing sessions.
Requires DB_PATH to point at the on-disk database (not :memory:).`)
			return
		}
	}

	fmt.Print(`Beacon: self-hosted container logs, RBAC, and monitoring.

Usage:
  beacon [command]

Commands:
  server          Run full dashboard with embedded web UI (default)
  agent           Run as fleet agent (local UI + API)
  agent-only      Run headless agent (API/WebSockets only)
  reset-password  Reset a user password in SQLite
  config          Print effective non-secret configuration
  version         Print version
  help            Show this help

Examples:
  beacon
  beacon server
  beacon agent-only
  beacon reset-password admin 'NewSecurePass1'
  beacon config

Install globally:
  make install
  # or: go install .

Docker:
  docker exec beacon beacon reset-password admin 'NewSecurePass1'
`)
}
