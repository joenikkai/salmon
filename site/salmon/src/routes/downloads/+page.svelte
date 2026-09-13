<script lang="ts">
  import { onMount } from 'svelte';

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

  let release: Release | null = $state(null);
  let error: string | null = $state(null);

  const GITHUB_REPO = 'joenikkai/salmon';

  onMount(async () => {
    try {
      const res = await fetch(`https://api.github.com/repos/${GITHUB_REPO}/releases/latest`);
      if (!res.ok) throw new Error('Failed to fetch release');
      const releases = await res.json();
      release = releases.find((r: Release) => r.tag_name.startsWith('app-v'));
    } catch (e) {
      error = e instanceof Error ? e.message : 'Unknown error';
    }
  });

  // Categorize assets by platform based on filename
  function getPlatform(assetName: string): string {
    if (assetName.endsWith('.exe')) return 'Windows';
    if (assetName.endsWith('.dmg') || assetName.endsWith('.pkg')) return 'macOS';
    if (assetName.endsWith('.deb') || assetName.endsWith('.AppImage')) return 'Linux';
    return 'Other';
  }

  function formatSize(bytes: number): string {
    return (bytes / 1024 / 1024).toFixed(2) + ' MB';
  }
</script>

<svelte:head>
  <title>Download - salmon</title>
  <meta name="description" content="Download the latest version of salmon for Windows, macOS, and Linux." />
</svelte:head>

<div class="downloads-container">
  <h1>Download salmon</h1>

  {#if error}
    <p class="error">Could not load downloads: {error}</p>
  {:else if !release}
    <p>Loading latest release...</p>
  {:else}
    <p class="version-info">
      Latest version: <strong>{release.tag_name}</strong>
      (published {new Date(release.published_at).toLocaleDateString()})
    </p>

    {#each ['Windows', 'macOS', 'Linux'] as platform}
      <section class="platform-section">
        <h2>{platform}</h2>
        <div class="assets-grid">
          {#each release.assets.filter(a => getPlatform(a.name) === platform) as asset}
            <a href={asset.browser_download_url} class="download-card" target="_blank" rel="noopener">
              <span class="asset-name">{asset.name}</span>
              <span class="asset-meta">
                {formatSize(asset.size)} &middot; {asset.download_count} downloads
              </span>
            </a>
          {/each}
        </div>
      </section>
    {/each}
  {/if}
</div>

<style>
  .downloads-container {
    max-width: 900px;
    margin: 2rem auto;
    padding: 0 1rem;
  }
  .version-info {
    color: #555;
    margin-bottom: 2rem;
  }
  .platform-section {
    margin-bottom: 2.5rem;
  }
  .assets-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 1rem;
  }
  .download-card {
    display: flex;
    flex-direction: column;
    padding: 1rem;
    border: 1px solid #ddd;
    border-radius: 8px;
    text-decoration: none;
    color: inherit;
    transition: border-color 0.2s, box-shadow 0.2s;
  }
  .download-card:hover {
    border-color: #0070f3;
    box-shadow: 0 2px 8px rgba(0, 112, 243, 0.15);
  }
  .asset-name {
    font-weight: 600;
    margin-bottom: 0.25rem;
  }
  .asset-meta {
    font-size: 0.85rem;
    color: #666;
  }
  .error {
    color: #d32f2f;
  }
</style>