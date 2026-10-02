<template>
  <div class="page-view settings-view animate-fade-in">
    <section class="page-hero">
      <div class="page-hero-body">
        <div class="page-hero-copy">
          <span class="page-hero-eyebrow">Account</span>
          <h1>Settings</h1>
          <p class="page-hero-sub">Manage your profile, security, and appearance preferences.</p>
        </div>
        <div class="page-hero-stats">
          <div class="page-hero-stat">
            <span class="page-hero-stat-val">{{ sharedState.currentUser?.username?.[0]?.toUpperCase() || "?" }}</span>
            <span class="page-hero-stat-lbl">User</span>
          </div>
          <div class="page-hero-stat" :class="sharedState.currentUser?.is_admin ? 'warning' : 'muted'">
            <span class="page-hero-stat-val">{{ sharedState.currentUser?.is_admin ? "Admin" : "Staff" }}</span>
            <span class="page-hero-stat-lbl">Role</span>
          </div>
        </div>
      </div>
      <div class="page-hero-mesh" aria-hidden="true"></div>
    </section>

    <div class="settings-grid">
      <div class="settings-card">
        <div class="card-header">
          <div class="card-icon">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </div>
          <div>
            <h3>Profile</h3>
            <p class="card-desc">Your account identity</p>
          </div>
        </div>
        <div class="card-body">
          <div class="info-row">
            <label>Username</label>
            <div class="value-box">{{ sharedState.currentUser?.username }}</div>
          </div>
          <div class="info-row">
            <label>Role</label>
            <div class="value-box">
              <span :class="['badge', sharedState.currentUser?.is_admin ? 'badge-warning' : 'badge-dim']">
                {{ sharedState.currentUser?.is_admin ? "Administrator" : "Staff member" }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <div class="settings-card">
        <div class="card-header">
          <div class="card-icon">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
              <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
              <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
            </svg>
          </div>
          <div>
            <h3>Security</h3>
            <p class="card-desc">Update your password</p>
          </div>
        </div>
        <div class="card-body">
          <form @submit.prevent="handlePasswordUpdate" class="settings-form">
            <div class="input-group" v-if="sharedState.currentUser?.password_changed !== false">
              <label>Current password</label>
              <input type="password" v-model="currentPassword" placeholder="Enter current password" class="premium-input" required />
            </div>
            <div class="input-group">
              <label>New password</label>
              <input type="password" v-model="newPassword" placeholder="At least 8 characters" class="premium-input" required />
            </div>
            <div class="input-group">
              <label>Confirm password</label>
              <input type="password" v-model="confirmPassword" placeholder="Confirm new password" class="premium-input" required />
            </div>
            <button type="submit" :disabled="loading" class="page-btn primary full-width">
              {{ loading ? "Updating..." : "Update password" }}
            </button>
            <p v-if="error" class="error-msg">{{ error }}</p>
          </form>
        </div>
      </div>

      <div class="settings-card">
        <div class="card-header">
          <div class="card-icon">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
              <circle cx="12" cy="12" r="3"></circle>
              <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
            </svg>
          </div>
          <div>
            <h3>Appearance</h3>
            <p class="card-desc">{{ themeDescription }}</p>
          </div>
        </div>
        <div class="card-body">
          <div class="theme-options">
            <button
              v-for="option in themeOptions"
              :key="option.value"
              type="button"
              :class="['page-filter-pill', { active: sharedState.themePreference === option.value }]"
              @click="applyTheme(option.value)"
            >
              {{ option.label }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { sharedState, showToast, applyTheme } from "../utils/sharedState";
import { apiFetch } from "../utils/apiFetch";

const themeOptions = [
  { value: "auto", label: "Auto" },
  { value: "light", label: "Light" },
  { value: "dark", label: "Dark" },
];

const themeDescription = computed(() => {
  if (sharedState.themePreference === "auto") {
    return `Following system (${sharedState.theme})`;
  }
  return `Using ${sharedState.themePreference} theme`;
});

const newPassword = ref("");
const confirmPassword = ref("");
const currentPassword = ref("");
const loading = ref(false);
const error = ref("");

const handlePasswordUpdate = async () => {
  if (newPassword.value !== confirmPassword.value) {
    error.value = "Passwords do not match";
    return;
  }
  if (newPassword.value.length < 8) {
    error.value = "Password must be at least 8 characters";
    return;
  }

  const mustProvideCurrent =
    sharedState.currentUser?.password_changed !== false &&
    !sharedState.forcePasswordChange;

  if (mustProvideCurrent && !currentPassword.value) {
    error.value = "Current password is required";
    return;
  }

  loading.value = true;
  error.value = "";

  try {
    const formData = new FormData();
    formData.append("password", newPassword.value);
    if (mustProvideCurrent) {
      formData.append("current_password", currentPassword.value);
    }

    const res = await apiFetch("/api/user/change-password", {
      method: "POST",
      body: formData,
    });

    if (res.ok) {
      showToast("Success", "Password updated successfully", "success");
      newPassword.value = "";
      confirmPassword.value = "";
      currentPassword.value = "";
      sharedState.forcePasswordChange = false;
      sharedState.showPasswordModal = false;
      if (sharedState.currentUser) {
        sharedState.currentUser.password_changed = true;
      }
    } else {
      const data = await res.json().catch(() => ({}));
      error.value = data.error || "Failed to update password";
    }
  } catch (err) {
    error.value = "System connection failed";
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.settings-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 1rem;
}

.settings-card {
  padding: 1.15rem;
  border-radius: var(--radius-xl);
  border: 1px solid var(--border);
  background: color-mix(in srgb, var(--bg-card) 90%, transparent);
  box-shadow: var(--elevate);
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.card-header {
  display: flex;
  align-items: flex-start;
  gap: 0.75rem;
}

.card-icon {
  width: 32px;
  height: 32px;
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(var(--accent-rgb), 0.12);
  color: var(--accent);
  border: 1px solid rgba(var(--accent-rgb), 0.2);
  flex-shrink: 0;
}

.card-icon svg {
  width: 16px;
  height: 16px;
}

.card-header h3 {
  margin: 0;
  font-size: 0.9375rem;
  font-weight: 700;
  color: var(--text-main);
  letter-spacing: -0.018em;
}

.card-desc {
  margin: 0.15rem 0 0;
  font-size: 0.75rem;
  color: var(--text-mute);
}

.card-body {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.info-row {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.info-row label {
  font-size: 0.625rem;
  font-weight: 600;
  color: var(--text-mute);
  text-transform: uppercase;
  letter-spacing: 0.12em;
  font-family: var(--font-mono);
}

.value-box {
  padding: 0 0.75rem;
  height: 36px;
  display: flex;
  align-items: center;
  background: transparent;
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
  color: var(--text-main);
  font-weight: 600;
}

.settings-form {
  display: flex;
  flex-direction: column;
  gap: 0.85rem;
}

.input-group label {
  display: block;
  font-size: 0.68rem;
  font-weight: 800;
  color: var(--text-mute);
  text-transform: uppercase;
  letter-spacing: 0.06em;
  margin-bottom: 0.4rem;
}

.full-width {
  width: 100%;
  justify-content: center;
  margin-top: 0.25rem;
}

.theme-options {
  display: flex;
  gap: 0.35rem;
  padding: 0.25rem;
  background: var(--bg-input);
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
}

.theme-options .page-filter-pill {
  flex: 1;
  justify-content: center;
}

@media (max-width: 768px) {
  .settings-grid {
    grid-template-columns: 1fr;
  }
}
</style>
