<template>
  <div class="not-found-container flex-center">
    <div class="not-found-card glass">
      <div class="error-visual">
        <span class="error-code numeric">404</span>
      </div>

      <div class="error-content">
        <h1>Page not found</h1>
        <p>
          The resource you’re looking for doesn’t exist or may have been moved.
        </p>

        <router-link to="/dashboard" class="page-btn primary">
          Return to dashboard
        </router-link>
      </div>
    </div>

    <div class="background-decor" aria-hidden="true">
      <div class="grid-lines"></div>
    </div>
  </div>
</template>

<script setup>
</script>

<style scoped>
.not-found-container {
  height: 100vh;
  width: 100vw;
  background: var(--bg-main);
  position: relative;
  overflow: hidden;
}

.not-found-card {
  width: 100%;
  max-width: 420px;
  padding: 2rem 1.75rem;
  border-radius: var(--radius-xl);
  border: 1px solid var(--border);
  background: color-mix(in srgb, var(--bg-card) 88%, transparent);
  text-align: center;
  z-index: 10;
  box-shadow: var(--elevate);
}

.error-visual {
  margin-bottom: 1.25rem;
}

.error-code {
  display: inline-block;
  font-size: 2.5rem;
  font-weight: 700;
  letter-spacing: -0.04em;
  color: var(--accent);
  font-family: var(--font-mono);
}

.error-content h1 {
  font-size: 1.15rem;
  font-weight: 700;
  letter-spacing: -0.018em;
  margin: 0 0 0.4rem;
  color: var(--text-main);
}

.error-content p {
  font-size: 0.8125rem;
  color: var(--text-mute);
  margin: 0 0 1.5rem;
  line-height: 1.5;
}

.background-decor {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.grid-lines {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(to right, color-mix(in srgb, var(--border) 55%, transparent) 1px, transparent 1px),
    linear-gradient(to bottom, color-mix(in srgb, var(--border) 55%, transparent) 1px, transparent 1px);
  background-size: 40px 40px;
  opacity: 0.45;
  mask-image: radial-gradient(ellipse 55% 50% at 50% 45%, black, transparent 75%);
}
</style>
