package cli

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureOutput(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestDispatchVersion(t *testing.T) {
	exit, code, _ := Dispatch([]string{"beacon", "version"})
	if !exit || code != 0 {
		t.Fatalf("expected exit 0, got exit=%v code=%d", exit, code)
	}
}

func TestDispatchUnknownCommand(t *testing.T) {
	exit, code, _ := Dispatch([]string{"beacon", "not-a-command"})
	if !exit || code != 1 {
		t.Fatalf("expected exit 1, got exit=%v code=%d", exit, code)
	}
}

func TestApplyRunModes(t *testing.T) {
	t.Setenv("BEACON_MODE", "")
	t.Setenv("BEACON_AGENT_ONLY", "")

	var rt Runtime
	applyRunMode(&rt, "agent-only")
	if rt.ServeFrontend {
		t.Fatal("agent-only should not serve frontend")
	}
	if os.Getenv("BEACON_AGENT_ONLY") != "true" {
		t.Fatal("expected BEACON_AGENT_ONLY=true")
	}

	applyRunMode(&rt, "agent")
	if !rt.ServeFrontend {
		t.Fatal("agent should serve frontend")
	}

	applyRunMode(&rt, "server")
	if !rt.ServeFrontend {
		t.Fatal("server should serve frontend")
	}
}

func TestPrintConfig(t *testing.T) {
	t.Setenv("DB_PATH", t.TempDir()+"/missing.db")
	out := captureOutput(func() { PrintConfig(Runtime{Mode: "server"}) })
	if !bytes.Contains([]byte(out), []byte("Beacon configuration")) {
		t.Fatalf("unexpected config output: %q", out)
	}
}

func TestDispatchServerReturnsRuntime(t *testing.T) {
	exit, code, rt := Dispatch([]string{"beacon", "server"})
	if exit || code != 0 {
		t.Fatalf("expected to continue server, got exit=%v code=%d", exit, code)
	}
	if !rt.ServeFrontend || rt.Mode != "server" {
		t.Fatalf("unexpected runtime: %+v", rt)
	}
}
