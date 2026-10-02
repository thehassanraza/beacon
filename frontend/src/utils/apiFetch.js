export async function apiFetch(input, init = {}) {
  const headers = new Headers(init.headers || {});
  if (!headers.has("X-Beacon-Client")) {
    headers.set("X-Beacon-Client", "web");
  }
  return fetch(input, { ...init, headers });
}
