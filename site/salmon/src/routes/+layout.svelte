<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import image from '$lib/assets/hero.png'
    import {PUBLIC_SITE_VERSION} from '$env/static/public'
    const ofImage = `${image}?v=${PUBLIC_SITE_VERSION}`

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />

	<title>Salmon — A Super Awesome Ledger</title>
	<meta name="description" content="A Super Awesome ledger" />

	<!-- Open Graph tags for WhatsApp, Facebook, LinkedIn, etc. -->
	<meta property="og:title" content="Salmon — A Super Awesome Ledger" />
	<meta property="og:description" content="A Super Awesome ledger" />
	<meta property="og:image" content={ofImage} />
	<!-- <meta property="og:url" content="https://your-domain.com" /> -->
	<meta property="og:type" content="website" />
	<meta property="og:site_name" content="Salmon" />

	<!-- Optional: Twitter Card (also used by some platforms) -->
	<meta name="twitter:card" content="salmon logo" />
	<meta name="twitter:title" content="Salmon — A Super Awesome Ledger" />
	<meta name="twitter:description" content="A Super Awesome ledger" />
	<meta name="twitter:image" content={ofImage} />
</svelte:head>
{@render children()}
