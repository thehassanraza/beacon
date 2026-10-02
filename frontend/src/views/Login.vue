<template>
  <div class="login-page">
    <aside class="login-visual" aria-hidden="true">
      <div class="visual-grid"></div>
      <div class="visual-glow"></div>
      <div class="visual-content">
        <div class="visual-brand">
          <BrandMark :size="32" />
          <span class="visual-wordmark">Bea<span class="brand-accent">con</span></span>
        </div>
        <h2 class="visual-headline">
          Observe every container.<br />
          Act with clarity.
        </h2>
        <p class="visual-sub">
          Unified logs, workloads, and health signals in one ops console.
        </p>
        <div class="visual-panel">
          <div class="panel-chrome">
            <span class="dot"></span><span class="dot"></span><span class="dot"></span>
            <span class="panel-title">stream · prod-api</span>
          </div>
          <pre class="panel-log"><code><span class="t-mute">12:41:08</span> <span class="t-info">INFO</span>  listener ready :8080
<span class="t-mute">12:41:09</span> <span class="t-ok">OK</span>    health check passed
<span class="t-mute">12:41:11</span> <span class="t-warn">WARN</span>  retry queue depth=2
<span class="t-mute">12:41:12</span> <span class="t-ok">OK</span>    drain complete</code></pre>
        </div>
      </div>
      <div class="visual-footer">
        <span>Infrastructure observability</span>
      </div>
    </aside>

    <main class="login-panel">
      <div class="login-panel-inner animate-slide-up">
        <div class="login-mobile-brand">
          <BrandMark :size="28" />
          <span class="logo-text">Bea<span class="brand-accent">con</span></span>
        </div>

        <header class="login-header">
          <p class="login-eyebrow">Sign in</p>
          <h1>Welcome back</h1>
          <p class="login-sub">Enter your credentials to access the console.</p>
        </header>

        <form @submit.prevent="handleLogin" class="login-form">
          <div class="input-group">
            <label for="login-username">Username</label>
            <div class="premium-input-wrapper">
              <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
              <input
                id="login-username"
                v-model="username"
                type="text"
                placeholder="admin"
                autocomplete="username"
                required
              />
            </div>
          </div>

          <div class="input-group">
            <label for="login-password">Password</label>
            <div class="premium-input-wrapper">
              <svg viewBox="0 0 24 24" width="16" height="16" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="3" y="11" width="18" height="11" rx="0" ry="0"></rect>
                <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
              </svg>
              <input
                id="login-password"
                v-model="password"
                type="password"
                placeholder="••••••••"
                autocomplete="current-password"
                required
              />
            </div>
          </div>

          <button type="submit" :disabled="loading" class="premium-btn primary full-width login-btn">
            {{ loading ? "Signing in…" : "Sign in" }}
          </button>

          <Transition name="fade">
            <p v-if="error" class="error-msg" role="alert">{{ error }}</p>
          </Transition>
        </form>

        <footer class="login-footer">
          <span>v{{ appVersion }}</span>
          <div class="dot-sep"></div>
          <span>Encrypted session</span>
        </footer>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { secureStorage } from "../utils/storage";
import { apiFetch } from "../utils/apiFetch";
import { sharedState } from "../utils/sharedState";
import BrandMark from "../components/BrandMark.vue";
import pkg from "../../package.json";

const appVersion = pkg.version;
const router = useRouter();
const username = ref("");
const password = ref("");
const loading = ref(false);
const error = ref("");

const handleLogin = async () => {
  loading.value = true;
  error.value = "";

  try {
    const formData = new FormData();
    formData.append("username", username.value.trim());
    formData.append("password", password.value);

    const res = await apiFetch("/api/token", {
      method: "POST",
      body: formData,
    });

    if (res.ok) {
      const data = await res.json();
      secureStorage.setItem("token", data.access_token);
      sharedState.forcePasswordChange = data.password_changed === false;
      sharedState.showPasswordModal = data.password_changed === false;
      router.push("/dashboard");
    } else {
      let message = "Invalid credentials";
      try {
        const data = await res.json();
        if (data?.error) message = data.error;
      } catch {
        /* keep default */
      }
      if (res.status === 429) {
        message = "Too many login attempts. Try again in 15 minutes.";
      }
      error.value = message;
    }
  } catch {
    error.value = "Unable to reach the server";
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.login-page {
  position: fixed;
  inset: 0;
  display: grid;
  grid-template-columns: minmax(0, 1.05fr) minmax(360px, 0.95fr);
  background: var(--bg-main);
  color: var(--text-main);
}

.login-visual {
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 2rem 2.5rem;
  background: #030303;
  border-right: 1px solid var(--border);
  overflow: hidden;
}

[data-theme="light"] .login-visual {
  background: #0a0a0c;
  color: #ececec;
}

.visual-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(to right, rgba(255, 255, 255, 0.04) 1px, transparent 1px),
    linear-gradient(to bottom, rgba(255, 255, 255, 0.04) 1px, transparent 1px);
  background-size: 40px 40px;
  mask-image: radial-gradient(ellipse 70% 60% at 40% 40%, black, transparent 75%);
  pointer-events: none;
}

.visual-glow {
  position: absolute;
  width: 55%;
  height: 45%;
  left: -10%;
  top: 15%;
  background: radial-gradient(ellipse, rgba(var(--accent-rgb), 0.18), transparent 70%);
  pointer-events: none;
}

.visual-content {
  position: relative;
  z-index: 1;
  max-width: 440px;
  margin-top: 8vh;
}

.visual-brand {
  display: flex;
  align-items: center;
  gap: 0.65rem;
  margin-bottom: 2rem;
}

.visual-wordmark {
  font-size: 1.05rem;
  font-weight: 700;
  letter-spacing: -0.018em;
  color: #ececec;
}

.visual-headline {
  font-size: clamp(1.6rem, 2.4vw, 2.1rem);
  font-weight: 700;
  letter-spacing: -0.03em;
  line-height: 1.2;
  margin: 0 0 0.85rem;
  color: #f2f2f2;
}

.visual-sub {
  margin: 0 0 2rem;
  font-size: 0.9rem;
  line-height: 1.55;
  color: #8a8a8a;
  max-width: 34ch;
}

.visual-panel {
  border: 1px solid #1a1a1a;
  background: #07080c;
  overflow: hidden;
  box-shadow: 0 20px 50px -24px rgba(0, 0, 0, 0.8);
}

.panel-chrome {
  display: flex;
  align-items: center;
  gap: 0.35rem;
  padding: 0.55rem 0.75rem;
  border-bottom: 1px solid #1a1a1a;
  background: #0a0a0a;
}

.panel-chrome .dot {
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background: #2a2a2a;
}

.panel-title {
  margin-left: 0.5rem;
  font-family: var(--font-mono);
  font-size: 0.625rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: #6e6e6e;
}

.panel-log {
  margin: 0;
  padding: 0.9rem 1rem 1.1rem;
  font-family: var(--font-mono);
  font-size: 0.72rem;
  line-height: 1.65;
  color: #c8ecd4;
  white-space: pre-wrap;
}

.t-mute { color: #5c6670; }
.t-info { color: #43b2e1; }
.t-ok { color: #4dbc92; }
.t-warn { color: #fcaa38; }

.visual-footer {
  position: relative;
  z-index: 1;
  font-family: var(--font-mono);
  font-size: 0.625rem;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: #5a5a5a;
}

.login-panel {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  background: var(--bg-main);
}

.login-panel-inner {
  width: 100%;
  max-width: 360px;
}

.login-mobile-brand {
  display: none;
  align-items: center;
  gap: 0.6rem;
  margin-bottom: 2rem;
}

.login-header {
  margin-bottom: 1.75rem;
}

.login-eyebrow {
  margin: 0 0 0.5rem;
  font-family: var(--font-mono);
  font-size: 0.625rem;
  font-weight: 600;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: var(--accent);
}

.login-header h1 {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 700;
  letter-spacing: -0.02em;
  color: var(--text-main);
}

.login-sub {
  margin: 0.45rem 0 0;
  font-size: 0.8125rem;
  color: var(--text-mute);
  line-height: 1.45;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.input-group label {
  display: block;
  font-size: 0.625rem;
  font-weight: 600;
  color: var(--text-mute);
  text-transform: uppercase;
  letter-spacing: 0.12em;
  font-family: var(--font-mono);
  margin-bottom: 0.45rem;
}

.premium-input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.premium-input-wrapper svg {
  position: absolute;
  left: 0.75rem;
  color: var(--text-mute);
  transition: color 0.15s;
  pointer-events: none;
}

.premium-input-wrapper input {
  width: 100%;
  height: 40px;
  padding: 0 0.75rem 0 2.35rem;
  background: var(--bg-card);
  border: 1px solid var(--border);
  border-radius: 0;
  color: var(--text-main);
  font-size: 0.875rem;
  font-weight: 500;
  transition: border-color 0.15s, box-shadow 0.15s;
}

.premium-input-wrapper input:focus {
  outline: none;
  border-color: var(--accent);
  box-shadow: 0 0 0 1px var(--accent);
}

.premium-input-wrapper:focus-within svg {
  color: var(--accent);
}

.login-btn {
  width: 100%;
  height: 40px;
  margin-top: 0.35rem;
  border-radius: 0;
}

.error-msg {
  text-align: left;
  margin: 0;
}

.login-footer {
  margin-top: 2rem;
  display: flex;
  align-items: center;
  gap: 0.65rem;
  font-size: 0.625rem;
  font-weight: 600;
  color: var(--text-mute);
  text-transform: uppercase;
  letter-spacing: 0.1em;
  font-family: var(--font-mono);
}

.dot-sep {
  width: 3px;
  height: 3px;
  border-radius: 50%;
  background: var(--border-strong);
}

@media (max-width: 900px) {
  .login-page {
    grid-template-columns: 1fr;
  }

  .login-visual {
    display: none;
  }

  .login-mobile-brand {
    display: flex;
  }

  .login-panel {
    padding: 1.5rem;
  }
}
</style>
