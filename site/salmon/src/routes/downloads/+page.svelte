<script lang="ts">
  import { onMount } from 'svelte';
  import { Sitename } from '$lib';

  interface Asset {
    name: string;
    browser_download_url: string;
    size: number;
    download_count: number;
  }

  interface Release {
    tag_name: string;
    published_at: string;
    assets: Asset[];
  }

  const GITHUB_REPO = 'joenikkai/salmon';
  const PLATFORMS = ['Windows', 'macOS', 'Linux'] as const;
  type Platform = (typeof PLATFORMS)[number];

  const navLinks = [
    { href: '/downloads', label: 'Download' },
    { href: '/license', label: 'Licence' },
    { href: '/terms-of-use', label: 'Terms of use' },
    { href: '/sponsor', label: 'Sponsor' }
  ];

  let release = $state<Release | null>(null);
  let error = $state<string | null>(null);
  let loading = $state(true);
  let menuOpen = $state(false);

  async function loadRelease() {
    loading = true;
    error = null;
    try {
      const res = await fetch(
        `https://api.github.com/repos/${GITHUB_REPO}/releases?per_page=30`
      );
      if (!res.ok) throw new Error(`GitHub responded with ${res.status}`);
      const releases: Release[] = await res.json();
      const appRelease = releases.find((r) => r.tag_name.startsWith('app-v'));
      if (!appRelease) throw new Error('No app release has been published yet.');
      release = appRelease;
    } catch (e) {
      error = e instanceof Error ? e.message : 'Something went wrong.';
    } finally {
      loading = false;
    }
  }

  onMount(loadRelease);

  const version = $derived(
    release ? release.tag_name.replace(/^app-v/, 'v') : ''
  );

  const published = $derived(
    release
      ? new Date(release.published_at).toLocaleDateString(undefined, {
          year: 'numeric',
          month: 'long',
          day: 'numeric'
        })
      : ''
  );

  const grouped = $derived(
    PLATFORMS.map((platform) => ({
      platform,
      assets: (release?.assets ?? [])
        .filter((a) => getPlatform(a.name) === platform)
        .sort((a, b) => a.name.localeCompare(b.name))
    }))
  );

  const totalDownloads = $derived(
    (release?.assets ?? []).reduce((sum, a) => sum + a.download_count, 0)
  );

  function getPlatform(name: string): Platform | 'Other' {
    const n = name.toLowerCase();
    if (n.endsWith('.exe') || n.endsWith('.msi')) return 'Windows';
    if (n.endsWith('.dmg') || n.endsWith('.pkg')) return 'macOS';
    if (n.endsWith('.deb') || n.endsWith('.rpm') || n.endsWith('.appimage'))
      return 'Linux';
    return 'Other';
  }

  function fileKind(name: string): string {
    const n = name.toLowerCase();
    if (n.endsWith('.exe')) return 'Windows installer';
    if (n.endsWith('.msi')) return 'MSI package';
    if (n.endsWith('.dmg')) return 'Disk image';
    if (n.endsWith('.pkg')) return 'Installer package';
    if (n.endsWith('.deb')) return 'Debian package';
    if (n.endsWith('.rpm')) return 'RPM package';
    if (n.endsWith('.appimage')) return 'AppImage';
    if (n.endsWith('.tar.gz')) return 'Tarball';
    if (n.endsWith('.zip')) return 'Zip archive';
    return 'Download';
  }

  function arch(name: string): string | null {
    const n = name.toLowerCase();
    if (n.includes('universal')) return 'Universal';
    if (n.includes('arm64') || n.includes('aarch64')) return 'ARM64';
    if (n.includes('x64') || n.includes('x86_64') || n.includes('amd64'))
      return 'x64';
    return null;
  }

  function formatSize(bytes: number): string {
    if (!bytes) return '—';
    const mb = bytes / 1024 / 1024;
    if (mb < 1) return `${Math.max(1, Math.round(bytes / 1024))} KB`;
    return `${mb.toFixed(1)} MB`;
  }
</script>

<svelte:head>
  <title>Download — {Sitename}</title>
  <meta
    name="description"
    content={`Download the latest version of ${Sitename} for Windows, macOS and Linux. Builds are published directly from GitHub Releases.`}
  />
</svelte:head>

<header class="site-header">
  <div class="container header-inner">
    <a class="brand" href="/" aria-label={Sitename}>
      <span class="brand-mark" aria-hidden="true">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          <path d="M3 10c2-3 4-3 6 0s4 3 6 0 4-3 6 0" />
          <path d="M3 16c2-3 4-3 6 0s4 3 6 0 4-3 6 0" />
        </svg>
      </span>
      <span class="brand-name">{Sitename}</span>
    </a>

    <nav class="nav-desktop" aria-label="Main">
      {#each navLinks as link}
        <a href={link.href} aria-current={link.href === '/downloads' ? 'page' : undefined}>
          {link.label}
        </a>
      {/each}
    </nav>

    <div class="header-actions">
      <a class="btn btn-ghost hide-sm" href="/sponsor">Sponsor</a>
      <a class="btn btn-primary" href="/license">Get a licence</a>
      <button
        class="menu-btn"
        aria-label="Toggle navigation"
        aria-expanded={menuOpen}
        onclick={() => (menuOpen = !menuOpen)}
      >
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          {#if menuOpen}
            <path d="M18 6 6 18M6 6l12 12" />
          {:else}
            <path d="M3 6h18M3 12h18M3 18h18" />
          {/if}
        </svg>
      </button>
    </div>
  </div>

  {#if menuOpen}
    <div class="mobile-menu">
      {#each navLinks as link}
        <a href={link.href} onclick={() => (menuOpen = false)}>{link.label}</a>
      {/each}
    </div>
  {/if}
</header>

<main>
  <!-- ── Page hero ─────────────────────────────────────── -->
  <section class="page-hero">
    <div class="container">
      <nav class="crumbs" aria-label="Breadcrumb">
        <a href="/">Home</a>
        <span aria-hidden="true">/</span>
        <span aria-current="page">Download</span>
      </nav>

      <span class="badge"><span class="dot"></span> Official builds · Verified releases</span>

      <h1>Download {Sitename}</h1>

      <p class="lede">
        Every build is published straight from GitHub Releases, so you always get the
        exact artifact that was shipped — nothing repackaged, nothing added.
      </p>

      {#if release}
        <div class="hero-meta">
          <div class="meta-item">
            <span class="meta-label">Latest version</span>
            <span class="meta-value">{version}</span>
          </div>
          <div class="meta-item">
            <span class="meta-label">Released</span>
            <span class="meta-value">{published}</span>
          </div>
          <div class="meta-item">
            <span class="meta-label">Total downloads</span>
            <span class="meta-value">{totalDownloads.toLocaleString()}</span>
          </div>
        </div>
      {/if}
    </div>
  </section>

  <!-- ── Downloads ─────────────────────────────────────── -->
  <section class="section">
    <div class="container">
      {#if loading}
        <div class="skeleton-grid" aria-hidden="true">
          {#each [1, 2, 3] as _}
            <div class="skeleton-card">
              <span class="sk sk-icon"></span>
              <span class="sk sk-line"></span>
              <span class="sk sk-line sk-short"></span>
            </div>
          {/each}
        </div>
        <p class="sr-only" role="status">Loading the latest release…</p>
      {:else if error}
        <div class="state-card" role="alert">
          <span class="state-icon" aria-hidden="true">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="9" />
              <path d="M12 7.5v5.5M12 16.4h.01" />
            </svg>
          </span>
          <h2>We couldn't load the downloads</h2>
          <p>{error}</p>
          <button class="btn btn-primary" onclick={loadRelease}>Try again</button>
          <p class="state-hint">
            You can also browse every release directly on
            <a
              href={`https://github.com/${GITHUB_REPO}/releases`}
              target="_blank"
              rel="noopener noreferrer"
            >
              GitHub
            </a>.
          </p>
        </div>
      {:else}
        <div class="platforms">
          {#each grouped as group (group.platform)}
            <section class="platform">
              <div class="platform-head">
                <span class="platform-icon" aria-hidden="true">
                  {@render platformIcon(group.platform)}
                </span>
                <h2>{group.platform}</h2>
                <span class="platform-count">
                  {group.assets.length}
                  {group.assets.length === 1 ? 'build' : 'builds'}
                </span>
              </div>

              {#if group.assets.length === 0}
                <p class="empty">No builds for this platform yet — check back soon.</p>
              {:else}
                <div class="assets-grid">
                  {#each group.assets as asset (asset.name)}
                    <a
                      class="download-card"
                      href={asset.browser_download_url}
                      target="_blank"
                      rel="noopener noreferrer"
                    >
                      <span class="asset-top">
                        <span class="asset-kind">{fileKind(asset.name)}</span>
                        {#if arch(asset.name)}
                          <span class="asset-arch">{arch(asset.name)}</span>
                        {/if}
                      </span>

                      <span class="asset-name">{asset.name}</span>

                      <span class="asset-meta">
                        <span>{formatSize(asset.size)}</span>
                        <span class="dot-sep" aria-hidden="true">·</span>
                        <span>{asset.download_count.toLocaleString()} downloads</span>
                      </span>

                      <span class="asset-cta" aria-hidden="true">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
                          <path d="M12 4v11" />
                          <path d="m7.5 11 4.5 4.5 4.5-4.5" />
                          <path d="M5 19.5h14" />
                        </svg>
                      </span>
                    </a>
                  {/each}
                </div>
              {/if}
            </section>
          {/each}
        </div>

        <aside class="assurance">
          <span class="assurance-icon" aria-hidden="true">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M12 3 5 6v5.4c0 4.3 2.9 7.8 7 9.1 4.1-1.3 7-4.8 7-9.1V6l-7-3Z" />
              <path d="m9.2 11.9 2.1 2.1 3.9-4" />
            </svg>
          </span>
          <p>
            Served directly from GitHub Releases. Source code, build logs and checksums are
            all public —
            <a
              href={`https://github.com/${GITHUB_REPO}`}
              target="_blank"
              rel="noopener noreferrer"
            >
              inspect the repository
            </a>.
          </p>
        </aside>
      {/if}
    </div>
  </section>

  <!-- ── CTA ───────────────────────────────────────────── -->
  <section class="cta">
    <div class="container">
      <div class="cta-card">
        <h2>Using {Sitename} commercially?</h2>
        <p>
          The AGPL covers most use cases for free. If you need to embed it in a
          closed-source product, a commercial licence keeps things simple.
        </p>
        <div class="cta-actions">
          <a class="btn btn-lg btn-white" href="/license">Get a licence</a>
          <a class="btn btn-lg btn-oncolor" href="/sponsor">Sponsor the project</a>
        </div>
      </div>
    </div>
  </section>
</main>

<footer class="footer">
  <div class="container">
    <div class="footer-grid">
      <div class="footer-brand">
        <a class="brand" href="/">
          <span class="brand-mark" aria-hidden="true">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
              <path d="M3 10c2-3 4-3 6 0s4 3 6 0 4-3 6 0" />
              <path d="M3 16c2-3 4-3 6 0s4 3 6 0 4-3 6 0" />
            </svg>
          </span>
          <span class="brand-name">{Sitename}</span>
        </a>
        <p>Open-source software, built in the open and licensed under AGPL-3.0.</p>
      </div>

      <div>
        <h4>Product</h4>
        <ul>
          <li><a href="/downloads" aria-current="page">Download app</a></li>
          <li><a href="/license">Get a licence</a></li>
          <li><a href="/sponsor">Sponsor</a></li>
        </ul>
      </div>

      <div>
        <h4>Legal</h4>
        <ul>
          <li><a href="/license">Licence</a></li>
          <li><a href="/terms-of-use">Terms of use</a></li>
        </ul>
      </div>

      <div>
        <h4>Community</h4>
        <ul>
          <li><a href={`https://github.com/${GITHUB_REPO}`} rel="noopener">GitHub</a></li>
          <li><a href="/sponsor">Support the project</a></li>
        </ul>
      </div>
    </div>

    <div class="footer-bottom">
      <span>© {new Date().getFullYear()} {Sitename}. All rights reserved.</span>
      <span>Released under the GNU AGPL-3.0.</span>
    </div>
  </div>
</footer>

{#snippet platformIcon(platform: Platform)}
  {#if platform === 'Windows'}
    <svg viewBox="0 0 24 24" fill="currentColor">
      <path d="M3 5.6 10.2 4.6v6.6H3V5.6Zm8.4-1.2L21 3v8.2h-9.6V4.4ZM3 12.4h7.2v6.6L3 18v-5.6Zm8.4 0H21V21l-9.6-1.3v-7.3Z" />
    </svg>
  {:else if platform === 'macOS'}
    <svg viewBox="0 0 24 24" fill="currentColor">
      <path d="M16.4 12.8c0-2.1 1.7-3.1 1.8-3.2-.9-1.4-2.4-1.6-2.9-1.7-1.3-.1-2.5.8-3.1.8-.6 0-1.6-.8-2.7-.7-1.4 0-2.6.8-3.3 2-1.4 2.5-.4 6.1 1 8.1.7 1 1.5 2.1 2.6 2 1 0 1.4-.6 2.7-.6s1.6.6 2.7.6c1.1 0 1.8-1 2.5-2 .8-1.2 1.1-2.3 1.1-2.4 0 0-2.2-.9-2.4-2.9Z" />
      <path d="M14.5 6.4c.6-.7 1-1.7.9-2.6-.9 0-1.9.6-2.5 1.3-.5.6-1 1.6-.9 2.5 1 .1 2-.5 2.5-1.2Z" />
    </svg>
  {:else}
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <rect x="2.5" y="4" width="19" height="16" rx="2.5" />
      <path d="m7 9.5 2.8 2.5L7 14.5M12.8 14.8h4.2" />
    </svg>
  {/if}
{/snippet}

<style>
  /* ── Design tokens (shared) ────────────────────────── */
  :global(:root) {
    --salmon-50: #fff4f1;
    --salmon-100: #ffe7e0;
    --salmon-200: #ffd0c4;
    --salmon-300: #ffb4a2;
    --salmon-400: #ff9581;
    --salmon-500: #fa8072;
    --salmon-600: #ea5f4d;
    --salmon-700: #c4442f;

    --ink: #0b1220;
    --ink-2: #334155;
    --muted: #64748b;
    --border: #e6e9ef;
    --surface: #ffffff;
    --surface-2: #f8fafc;

    --shadow-sm: 0 1px 2px rgba(15, 23, 42, 0.06);
    --shadow-md: 0 10px 30px -12px rgba(15, 23, 42, 0.18);
  }

  :global(*),
  :global(*::before),
  :global(*::after) {
    box-sizing: border-box;
  }

  :global(body) {
    margin: 0;
    font-family: "Inter", ui-sans-serif, system-ui, -apple-system, "Segoe UI",
      Roboto, Helvetica, Arial, sans-serif;
    color: var(--ink);
    background: var(--surface);
    -webkit-font-smoothing: antialiased;
    text-rendering: optimizeLegibility;
  }

  .container {
    width: 100%;
    max-width: 1140px;
    margin: 0 auto;
    padding: 0 24px;
  }

  .sr-only {
    position: absolute;
    width: 1px;
    height: 1px;
    padding: 0;
    margin: -1px;
    overflow: hidden;
    clip: rect(0 0 0 0);
    white-space: nowrap;
    border: 0;
  }

  /* ── Buttons ───────────────────────────────────────── */
  .btn {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    padding: 10px 18px;
    border: 1px solid transparent;
    border-radius: 10px;
    font: inherit;
    font-size: 0.9375rem;
    font-weight: 600;
    letter-spacing: -0.005em;
    text-decoration: none;
    white-space: nowrap;
    cursor: pointer;
    transition: transform 0.15s ease, background-color 0.15s ease,
      box-shadow 0.15s ease, border-color 0.15s ease, color 0.15s ease;
  }

  .btn-lg {
    padding: 13px 24px;
    font-size: 1rem;
    border-radius: 12px;
  }

  .btn-primary {
    color: #fff;
    background: linear-gradient(180deg, var(--salmon-500), var(--salmon-600));
    box-shadow: var(--shadow-sm), 0 10px 22px -10px rgba(234, 95, 77, 0.75);
  }

  .btn-primary:hover {
    transform: translateY(-1px);
    box-shadow: var(--shadow-sm), 0 14px 28px -12px rgba(234, 95, 77, 0.9);
  }

  .btn-ghost {
    color: var(--ink-2);
    background: transparent;
  }

  .btn-ghost:hover {
    background: var(--salmon-50);
    color: var(--salmon-700);
  }

  .btn:focus-visible,
  a:focus-visible,
  button:focus-visible {
    outline: 2px solid var(--salmon-500);
    outline-offset: 2px;
  }

  /* ── Header ────────────────────────────────────────── */
  .site-header {
    position: sticky;
    top: 0;
    z-index: 50;
    background: rgba(255, 255, 255, 0.82);
    backdrop-filter: saturate(180%) blur(14px);
    -webkit-backdrop-filter: saturate(180%) blur(14px);
    border-bottom: 1px solid var(--border);
  }

  .header-inner {
    display: flex;
    align-items: center;
    gap: 20px;
    height: 68px;
  }

  .brand {
    display: inline-flex;
    align-items: center;
    gap: 10px;
    font-size: 1.0625rem;
    font-weight: 700;
    letter-spacing: -0.02em;
    color: var(--ink);
    text-decoration: none;
  }

  .brand-mark {
    display: grid;
    place-items: center;
    width: 34px;
    height: 34px;
    border-radius: 10px;
    color: #fff;
    background: linear-gradient(135deg, var(--salmon-400), var(--salmon-600));
    box-shadow: 0 6px 16px -6px rgba(234, 95, 77, 0.85);
  }

  .brand-mark svg {
    width: 20px;
    height: 20px;
  }

  .nav-desktop {
    display: flex;
    gap: 2px;
    margin-left: 12px;
  }

  .nav-desktop a {
    padding: 8px 12px;
    border-radius: 8px;
    font-size: 0.9375rem;
    font-weight: 500;
    color: var(--ink-2);
    text-decoration: none;
    transition: background-color 0.15s ease, color 0.15s ease;
  }

  .nav-desktop a:hover {
    background: var(--salmon-50);
    color: var(--salmon-700);
  }

  .nav-desktop a[aria-current="page"] {
    background: var(--salmon-50);
    color: var(--salmon-700);
    font-weight: 600;
  }

  .header-actions {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-left: auto;
  }

  .menu-btn {
    display: none;
    place-items: center;
    width: 40px;
    height: 40px;
    padding: 0;
    border: 1px solid var(--border);
    border-radius: 10px;
    background: #fff;
    color: var(--ink);
    cursor: pointer;
  }

  .menu-btn svg {
    width: 20px;
    height: 20px;
  }

  .mobile-menu {
    display: none;
  }

  /* ── Page hero ─────────────────────────────────────── */
  .page-hero {
    position: relative;
    overflow: hidden;
    padding: 64px 0 56px;
    border-bottom: 1px solid var(--border);
  }

  .page-hero::before {
    content: "";
    position: absolute;
    inset: -40% -10% auto -10%;
    height: 520px;
    background: radial-gradient(
        620px 300px at 15% 0%,
        rgba(250, 128, 114, 0.16),
        transparent 70%
      ),
      radial-gradient(
        480px 280px at 85% 5%,
        rgba(255, 180, 162, 0.2),
        transparent 70%
      );
    pointer-events: none;
  }

  .page-hero .container {
    position: relative;
  }

  .crumbs {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 26px;
    font-size: 0.8125rem;
    color: var(--muted);
  }

  .crumbs a {
    color: var(--muted);
    text-decoration: none;
    transition: color 0.15s ease;
  }

  .crumbs a:hover {
    color: var(--salmon-600);
  }

  .badge {
    display: inline-flex;
    align-items: center;
    gap: 9px;
    padding: 6px 14px 6px 11px;
    border: 1px solid var(--salmon-200);
    border-radius: 999px;
    background: var(--salmon-50);
    color: var(--salmon-700);
    font-size: 0.8125rem;
    font-weight: 600;
  }

  .dot {
    width: 7px;
    height: 7px;
    border-radius: 50%;
    background: var(--salmon-500);
    box-shadow: 0 0 0 4px rgba(250, 128, 114, 0.2);
    animation: pulse 2.4s ease-out infinite;
  }

  @keyframes pulse {
    0% {
      box-shadow: 0 0 0 0 rgba(250, 128, 114, 0.55);
    }
    70% {
      box-shadow: 0 0 0 9px rgba(250, 128, 114, 0);
    }
    100% {
      box-shadow: 0 0 0 0 rgba(250, 128, 114, 0);
    }
  }

  .page-hero h1 {
    margin: 20px 0 0;
    font-size: clamp(2.25rem, 4.6vw, 3.25rem);
    line-height: 1.08;
    letter-spacing: -0.035em;
    font-weight: 800;
  }

  .lede {
    margin: 18px 0 0;
    max-width: 46rem;
    font-size: 1.125rem;
    line-height: 1.7;
    color: var(--muted);
  }

  .hero-meta {
    display: flex;
    flex-wrap: wrap;
    gap: 12px 40px;
    margin-top: 34px;
    padding-top: 26px;
    border-top: 1px solid var(--border);
  }

  .meta-item {
    display: grid;
    gap: 4px;
  }

  .meta-label {
    font-size: 0.75rem;
    font-weight: 700;
    letter-spacing: 0.1em;
    text-transform: uppercase;
    color: var(--muted);
  }

  .meta-value {
    font-size: 0.9375rem;
    font-weight: 600;
    color: var(--ink);
    font-variant-numeric: tabular-nums;
  }

  /* ── Sections ──────────────────────────────────────── */
  .section {
    padding: 72px 0 88px;
  }

  /* ── Platforms ─────────────────────────────────────── */
  .platforms {
    display: flex;
    flex-direction: column;
    gap: 48px;
  }

  .platform-head {
    display: flex;
    align-items: center;
    gap: 12px;
    padding-bottom: 16px;
    margin-bottom: 22px;
    border-bottom: 1px solid var(--border);
  }

  .platform-icon {
    display: grid;
    place-items: center;
    width: 38px;
    height: 38px;
    border-radius: 11px;
    background: var(--salmon-50);
    border: 1px solid var(--salmon-100);
    color: var(--salmon-600);
    flex: none;
  }

  .platform-icon svg {
    width: 18px;
    height: 18px;
  }

  .platform-head h2 {
    margin: 0;
    font-size: 1.125rem;
    font-weight: 700;
    letter-spacing: -0.02em;
  }

  .platform-count {
    margin-left: auto;
    font-size: 0.8125rem;
    font-weight: 600;
    color: var(--muted);
    padding: 4px 12px;
    border-radius: 999px;
    background: var(--surface-2);
    border: 1px solid var(--border);
  }

  .assets-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(268px, 1fr));
    gap: 14px;
  }

  .empty {
    margin: 0;
    font-size: 0.9375rem;
    color: var(--muted);
    padding: 18px 0;
  }

  /* ── Download card ─────────────────────────────────── */
  .download-card {
    position: relative;
    display: flex;
    flex-direction: column;
    gap: 6px;
    padding: 20px;
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 14px;
    box-shadow: var(--shadow-sm);
    text-decoration: none;
    color: inherit;
    overflow: hidden;
    transition: transform 0.2s ease, box-shadow 0.2s ease, border-color 0.2s ease;
  }

  .download-card::after {
    content: "";
    position: absolute;
    inset: 0 0 auto 0;
    height: 3px;
    background: linear-gradient(90deg, var(--salmon-500), var(--salmon-600));
    transform: scaleX(0);
    transform-origin: left;
    transition: transform 0.28s ease;
  }

  .download-card:hover {
    transform: translateY(-3px);
    border-color: var(--salmon-200);
    box-shadow: var(--shadow-md);
  }

  .download-card:hover::after,
  .download-card:focus-visible::after {
    transform: scaleX(1);
  }

  .asset-top {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 2px;
  }

  .asset-kind {
    font-size: 0.7rem;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    color: var(--salmon-700);
  }

  .asset-arch {
    font-size: 0.6875rem;
    font-weight: 700;
    letter-spacing: 0.04em;
    padding: 2px 8px;
    border-radius: 6px;
    background: var(--surface-2);
    border: 1px solid var(--border);
    color: var(--ink-2);
  }

  .asset-name {
    font-family: ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, monospace;
    font-size: 0.8125rem;
    line-height: 1.45;
    color: var(--ink-2);
    word-break: break-all;
  }

  .asset-meta {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    gap: 6px;
    margin-top: 4px;
    font-size: 0.8125rem;
    color: var(--muted);
    font-variant-numeric: tabular-nums;
  }

  .dot-sep {
    opacity: 0.55;
  }

  .asset-cta {
    position: absolute;
    top: 16px;
    right: 16px;
    display: grid;
    place-items: center;
    width: 32px;
    height: 32px;
    border-radius: 9px;
    background: var(--salmon-50);
    color: var(--salmon-600);
    opacity: 0;
    transform: translateY(-3px);
    transition: opacity 0.2s ease, transform 0.2s ease;
  }

  .asset-cta svg {
    width: 16px;
    height: 16px;
  }

  .download-card:hover .asset-cta,
  .download-card:focus-visible .asset-cta {
    opacity: 1;
    transform: translateY(0);
  }

  /* ── Skeleton ──────────────────────────────────────── */
  .skeleton-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(268px, 1fr));
    gap: 14px;
  }

  .skeleton-card {
    display: flex;
    flex-direction: column;
    gap: 12px;
    padding: 20px;
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 14px;
  }

  .sk {
    display: block;
    border-radius: 7px;
    background: linear-gradient(90deg, #f1eeeb 25%, #e7e3df 37%, #f1eeeb 63%);
    background-size: 400% 100%;
    animation: shimmer 1.4s ease infinite;
  }

  .sk-icon {
    width: 64px;
    height: 12px;
    border-radius: 999px;
  }

  .sk-line {
    height: 11px;
    width: 100%;
  }

  .sk-short {
    width: 55%;
  }

  @keyframes shimmer {
    0% {
      background-position: 100% 50%;
    }
    100% {
      background-position: 0 50%;
    }
  }

  /* ── State card ────────────────────────────────────── */
  .state-card {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
    padding: 34px 32px;
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 16px;
    box-shadow: var(--shadow-sm);
    max-width: 560px;
  }

  .state-icon {
    display: grid;
    place-items: center;
    width: 42px;
    height: 42px;
    border-radius: 12px;
    background: var(--salmon-50);
    border: 1px solid var(--salmon-100);
    color: var(--salmon-600);
    margin-bottom: 4px;
  }

  .state-icon svg {
    width: 20px;
    height: 20px;
  }

  .state-card h2 {
    margin: 0;
    font-size: 1.15rem;
    font-weight: 700;
    letter-spacing: -0.02em;
  }

  .state-card p {
    margin: 0;
    font-size: 0.9375rem;
    line-height: 1.65;
    color: var(--muted);
  }

  .state-card .btn {
    margin-top: 8px;
  }

  .state-hint {
    font-size: 0.875rem !important;
  }

  .state-hint a {
    color: var(--salmon-600);
    font-weight: 600;
    text-decoration: none;
  }

  .state-hint a:hover {
    text-decoration: underline;
  }

  /* ── Assurance ─────────────────────────────────────── */
  .assurance {
    display: flex;
    align-items: flex-start;
    gap: 14px;
    margin-top: 56px;
    padding: 20px 22px;
    background: var(--salmon-50);
    border: 1px solid var(--salmon-200);
    border-radius: 14px;
  }

  .assurance-icon {
    display: grid;
    place-items: center;
    width: 36px;
    height: 36px;
    flex: none;
    border-radius: 10px;
    background: #fff;
    border: 1px solid var(--salmon-100);
    color: var(--salmon-600);
  }

  .assurance-icon svg {
    width: 18px;
    height: 18px;
  }

  .assurance p {
    margin: 0;
    font-size: 0.9375rem;
    line-height: 1.65;
    color: var(--ink-2);
  }

  .assurance a {
    color: var(--salmon-700);
    font-weight: 600;
    text-decoration: none;
  }

  .assurance a:hover {
    text-decoration: underline;
  }

  /* ── CTA ───────────────────────────────────────────── */
  .cta {
    padding: 0 0 88px;
  }

  .cta-card {
    position: relative;
    overflow: hidden;
    padding: 56px 40px;
    border-radius: 24px;
    text-align: center;
    color: #fff;
    background: linear-gradient(
      135deg,
      var(--salmon-500),
      var(--salmon-600) 60%,
      #d94d38
    );
    box-shadow: 0 30px 60px -30px rgba(234, 95, 77, 0.9);
  }

  .cta-card::after {
    content: "";
    position: absolute;
    inset: 0;
    background: radial-gradient(
      520px 220px at 50% -20%,
      rgba(255, 255, 255, 0.35),
      transparent 70%
    );
    pointer-events: none;
  }

  .cta-card > * {
    position: relative;
  }

  .cta-card h2 {
    margin: 0;
    font-size: clamp(1.6rem, 3.2vw, 2.2rem);
    font-weight: 800;
    letter-spacing: -0.03em;
  }

  .cta-card p {
    max-width: 540px;
    margin: 14px auto 0;
    font-size: 1.0625rem;
    line-height: 1.7;
    color: rgba(255, 255, 255, 0.92);
  }

  .cta-actions {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    gap: 12px;
    margin-top: 30px;
  }

  .btn-white {
    color: var(--salmon-700);
    background: #fff;
    box-shadow: 0 10px 24px -12px rgba(0, 0, 0, 0.5);
  }

  .btn-white:hover {
    transform: translateY(-1px);
  }

  .btn-oncolor {
    color: #fff;
    background: rgba(255, 255, 255, 0.14);
    border-color: rgba(255, 255, 255, 0.4);
  }

  .btn-oncolor:hover {
    background: rgba(255, 255, 255, 0.22);
  }

  /* ── Footer ────────────────────────────────────────── */
  .footer {
    padding: 56px 0 32px;
    border-top: 1px solid var(--border);
    background: #fff;
  }

  .footer-grid {
    display: grid;
    grid-template-columns: 1.4fr repeat(3, 1fr);
    gap: 40px;
  }

  .footer-brand p {
    max-width: 280px;
    margin: 14px 0 0;
    font-size: 0.9375rem;
    line-height: 1.7;
    color: var(--muted);
  }

  .footer h4 {
    margin: 0 0 16px;
    font-size: 0.8125rem;
    font-weight: 700;
    letter-spacing: 0.12em;
    text-transform: uppercase;
    color: var(--ink);
  }

  .footer ul {
    display: grid;
    gap: 10px;
    margin: 0;
    padding: 0;
    list-style: none;
  }

  .footer ul a {
    font-size: 0.9375rem;
    color: var(--muted);
    text-decoration: none;
    transition: color 0.15s ease;
  }

  .footer ul a:hover {
    color: var(--salmon-600);
  }

  .footer-bottom {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    gap: 16px;
    margin-top: 48px;
    padding-top: 24px;
    border-top: 1px solid var(--border);
    font-size: 0.875rem;
    color: var(--muted);
  }

  /* ── Responsive ────────────────────────────────────── */
  @media (max-width: 980px) {
    .footer-grid {
      grid-template-columns: 1fr 1fr;
      gap: 32px;
    }
  }

  @media (max-width: 760px) {
    .nav-desktop,
    .hide-sm {
      display: none;
    }

    .menu-btn {
      display: grid;
    }

    .mobile-menu {
      display: grid;
      gap: 4px;
      padding: 8px 24px 20px;
      border-top: 1px solid var(--border);
      background: #fff;
    }

    .mobile-menu a {
      padding: 12px 4px;
      border-bottom: 1px solid var(--border);
      font-size: 1rem;
      font-weight: 500;
      color: var(--ink-2);
      text-decoration: none;
    }

    .mobile-menu a:last-child {
      border-bottom: 0;
    }

    .page-hero {
      padding: 48px 0 44px;
    }

    .section {
      padding: 56px 0 64px;
    }

    .assets-grid,
    .skeleton-grid {
      grid-template-columns: 1fr;
    }

    .state-card {
      padding: 26px 22px;
    }

    .cta {
      padding: 0 0 64px;
    }

    .cta-card {
      padding: 44px 24px;
    }

    .footer-grid {
      grid-template-columns: 1fr;
    }
  }

  @media (max-width: 520px) {
    .container {
      padding: 0 18px;
    }

    .footer-bottom {
      justify-content: center;
      text-align: center;
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .btn,
    .download-card,
    .download-card::after,
    .asset-cta,
    .dot {
      animation: none !important;
      transition: none !important;
    }

    .download-card:hover {
      transform: none;
    }
  }
</style>