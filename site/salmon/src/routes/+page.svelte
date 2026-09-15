<script>
    import { Sitename } from "$lib";
    import Hero from "$lib/assets/hero.png";

    let menuOpen = false;

    const navLinks = [
        { href: "/downloads", label: "Download" },
        { href: "/license", label: "Licence" },
        { href: "/terms-of-use", label: "Terms of use" },
        { href: "/sponsor", label: "Sponsor" }
    ];

    const checks = ["No tracking", "Auditable source", "Free to self-host"];

    const trust = [
        { label: "AGPL-3.0 licensed", icon: "M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" },
        { label: "Fully auditable source", icon: "m16 18 6-6-6-6M8 6l-6 6 6 6" },
        { label: "Self-host friendly", icon: "M4 4h16v6H4zM4 14h16v6H4zM8 7h.01M8 17h.01" },
        { label: "No vendor lock-in", icon: "M12 2a10 10 0 1 0 0 20 10 10 0 0 0 0-20zM2 12h20M12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 0-20" }
    ];

    const features = [
        {
            title: "Built for speed",
            body: "A lean core and zero bloat, so everything stays instant — even on modest hardware.",
            icon: "M13 2 3 14h9l-1 8 10-12h-9l1-8z"
        },
        {
            title: "Private by default",
            body: "Your data stays on your machine. No telemetry, no tracking, no quiet uploads.",
            icon: "M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"
        },
        {
            title: "Open source",
            body: "Every line is public and licensed under AGPL-3.0. Read it, audit it, fork it.",
            icon: "m16 18 6-6-6-6M8 6l-6 6 6 6"
        },
        {
            title: "Self-host ready",
            body: "Run it on your own server in minutes, or use a hosted instance if you prefer.",
            icon: "M4 4h16v6H4zM4 14h16v6H4zM8 7h.01M8 17h.01"
        },
        {
            title: "Works everywhere",
            body: "A consistent experience across desktop and the web, with no platform lock-in.",
            icon: "M12 2a10 10 0 1 0 0 20 10 10 0 0 0 0-20zM2 12h20M12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 0-20"
        },
        {
            title: "Community driven",
            body: "The roadmap is shaped in the open by the people who actually use it.",
            icon: "M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2M9 11a4 4 0 1 0 0-8 4 4 0 0 0 0 8zM22 21v-2a4 4 0 0 0-3-3.87M16 3.13a4 4 0 0 1 0 7.75"
        }
    ];

    const licensePoints = [
        "Free to use, modify and distribute under AGPL-3.0",
        "Full source available — no closed components",
        "Copyleft keeps every fork open for everyone",
        "Commercial licence available for closed-source use"
    ];
</script>

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
                <a href={link.href}>{link.label}</a>
            {/each}
        </nav>

        <div class="header-actions">
            <a class="btn btn-ghost hide-sm" href="/sponsor">Sponsor</a>
            <a class="btn btn-primary" href="/downloads">Download app</a>
            <button
                class="menu-btn"
                aria-label="Toggle navigation"
                aria-expanded={menuOpen}
                on:click={() => (menuOpen = !menuOpen)}
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
                <a href={link.href} on:click={() => (menuOpen = false)}>{link.label}</a>
            {/each}
        </div>
    {/if}
</header>

<main>
    <!-- ── Hero ─────────────────────────────────────────── -->
    <section class="hero">
        <div class="container hero-inner">
            <div class="hero-copy">
                <span class="badge"><span class="dot"></span> Open source · AGPL-3.0</span>

                <h1>Software you can <span class="accent">actually trust</span>.</h1>

                <p class="lede">
                    {Sitename} is built in the open — fast, private, and free to inspect,
                    self-host and improve. No black boxes, no tracking, no lock-in.
                </p>

                <div class="hero-cta">
                    <a class="btn btn-primary btn-lg" href="/downloads">Download app</a>
                    <a class="btn btn-outline btn-lg" href="/license">Get a licence</a>
                </div>

                <ul class="hero-checks">
                    {#each checks as item}
                        <li>
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                                <path d="m5 12 4 4 10-10" />
                            </svg>
                            {item}
                        </li>
                    {/each}
                </ul>
            </div>

            <div class="hero-visual">
                <div class="glow" aria-hidden="true"></div>
                <div class="shot">
                    <img src={Hero} alt={`${Sitename} interface preview`} />
                </div>
            </div>
        </div>
    </section>

    <!-- ── Trust bar ────────────────────────────────────── -->
    <section class="trust">
        <div class="container trust-inner">
            {#each trust as item}
                <div class="trust-item">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                        <path d={item.icon} />
                    </svg>
                    {item.label}
                </div>
            {/each}
        </div>
    </section>

    <!-- ── Features ─────────────────────────────────────── -->
    <section class="section">
        <div class="container">
            <div class="section-head">
                <span class="eyebrow">Why {Sitename}</span>
                <h2>Everything you'd expect. Nothing you wouldn't.</h2>
                <p>
                    Built to be dependable, transparent and yours to control — from the first
                    line of code to the licence it ships under.
                </p>
            </div>

            <div class="features">
                {#each features as feature}
                    <article class="feature">
                        <div class="feature-icon">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                                <path d={feature.icon} />
                            </svg>
                        </div>
                        <h3>{feature.title}</h3>
                        <p>{feature.body}</p>
                    </article>
                {/each}
            </div>
        </div>
    </section>

    <!-- ── Licence ──────────────────────────────────────── -->
    <section class="license">
        <div class="container license-inner">
            <div class="license-copy">
                <span class="eyebrow">Licensing</span>
                <h2>Open source, with clear terms.</h2>
                <p>
                    {Sitename} is released under the <strong>GNU Affero General Public
                    License v3.0</strong>. You're free to run, study, share and improve it —
                    and anyone who builds on it shares those freedoms back.
                </p>
                <p>
                    Need to embed it in a closed-source product? A commercial licence is
                    available with the same code and none of the copyleft obligations.
                </p>

                <div class="license-actions">
                    <a class="btn btn-primary" href="/license">Get a licence</a>
                    <a class="btn btn-outline" href="/terms-of-use">Read the terms of use</a>
                </div>
            </div>

            <aside class="license-card">
                <h3>What that means for you</h3>
                <ul class="license-list">
                    {#each licensePoints as point}
                        <li>
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                                <path d="m5 12 4 4 10-10" />
                            </svg>
                            <span>{point}</span>
                        </li>
                    {/each}
                </ul>
            </aside>
        </div>
    </section>

    <!-- ── CTA ──────────────────────────────────────────── -->
    <section class="cta">
        <div class="container">
            <div class="cta-card">
                <h2>Start using {Sitename} today.</h2>
                <p>
                    Download it, read the source, or help fund the work that keeps it
                    independent and free for everyone.
                </p>
                <div class="cta-actions">
                    <a class="btn btn-lg btn-white" href="/downloads">Download app</a>
                    <a class="btn btn-lg btn-oncolor" href="/sponsor">Sponsor the project</a>
                </div>
            </div>
        </div>
    </section>
</main>

<!-- ── Footer ───────────────────────────────────────────── -->
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
                    <li><a href="/downloads">Download app</a></li>
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
                    <li><a href="https://github.com/joenikkai/salmon" rel="noopener">GitHub</a></li>
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

<style>
    /* ── Design tokens ─────────────────────────────────── */
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
        --shadow-lg: 0 30px 60px -25px rgba(15, 23, 42, 0.28);
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

    :global(img) {
        max-width: 100%;
    }

    .container {
        width: 100%;
        max-width: 1140px;
        margin: 0 auto;
        padding: 0 24px;
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

    .btn-outline {
        color: var(--ink);
        background: #fff;
        border-color: var(--border);
        box-shadow: var(--shadow-sm);
    }

    .btn-outline:hover {
        transform: translateY(-1px);
        border-color: var(--salmon-300);
        color: var(--salmon-700);
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
    .menu-btn:focus-visible,
    a:focus-visible {
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

    /* ── Hero ──────────────────────────────────────────── */
    .hero {
        position: relative;
        overflow: hidden;
        padding: 84px 0 96px;
    }

    .hero::before {
        content: "";
        position: absolute;
        inset: -30% -10% auto -10%;
        height: 640px;
        background: radial-gradient(
                620px 320px at 12% 0%,
                rgba(250, 128, 114, 0.18),
                transparent 70%
            ),
            radial-gradient(
                520px 320px at 88% 8%,
                rgba(255, 180, 162, 0.22),
                transparent 70%
            );
        pointer-events: none;
    }

    .hero-inner {
        position: relative;
        display: grid;
        grid-template-columns: 1.05fr 0.95fr;
        align-items: center;
        gap: 64px;
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
    }

    .hero h1 {
        margin: 22px 0 0;
        font-size: clamp(2.4rem, 5.2vw, 3.5rem);
        line-height: 1.06;
        letter-spacing: -0.035em;
        font-weight: 800;
    }

    .accent {
        color: var(--salmon-600);
        background: linear-gradient(transparent 66%, var(--salmon-200) 66%);
        border-radius: 2px;
    }

    .lede {
        margin: 20px 0 0;
        max-width: 34rem;
        font-size: 1.125rem;
        line-height: 1.7;
        color: var(--muted);
    }

    .hero-cta {
        display: flex;
        flex-wrap: wrap;
        gap: 12px;
        margin-top: 34px;
    }

    .hero-checks {
        display: flex;
        flex-wrap: wrap;
        gap: 10px 24px;
        margin: 30px 0 0;
        padding: 0;
        list-style: none;
        font-size: 0.875rem;
        font-weight: 500;
        color: var(--ink-2);
    }

    .hero-checks li {
        display: inline-flex;
        align-items: center;
        gap: 8px;
    }

    .hero-checks svg {
        width: 16px;
        height: 16px;
        flex: none;
        color: var(--salmon-600);
    }

    .hero-visual {
        position: relative;
    }

    .glow {
        position: absolute;
        inset: 10% 4% 6%;
        border-radius: 40px;
        background: linear-gradient(
            135deg,
            rgba(250, 128, 114, 0.45),
            rgba(255, 180, 162, 0.25)
        );
        filter: blur(48px);
    }

    .shot {
        position: relative;
        padding: 10px;
        border: 1px solid var(--border);
        border-radius: 18px;
        background: #fff;
        box-shadow: var(--shadow-lg);
        overflow: hidden;
    }

    .shot img {
        display: block;
        width: 100%;
        height: auto;
        border-radius: 10px;
    }

    /* ── Trust bar ─────────────────────────────────────── */
    .trust {
        padding: 22px 0;
        border-top: 1px solid var(--border);
        border-bottom: 1px solid var(--border);
        background: var(--surface-2);
    }

    .trust-inner {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 14px 44px;
    }

    .trust-item {
        display: inline-flex;
        align-items: center;
        gap: 10px;
        font-size: 0.875rem;
        font-weight: 600;
        color: var(--ink-2);
    }

    .trust-item svg {
        width: 18px;
        height: 18px;
        flex: none;
        color: var(--salmon-600);
    }

    /* ── Sections ──────────────────────────────────────── */
    .section {
        padding: 96px 0;
    }

    .section-head {
        max-width: 640px;
        margin: 0 auto 56px;
        text-align: center;
    }

    .eyebrow {
        display: inline-block;
        font-size: 0.8125rem;
        font-weight: 700;
        letter-spacing: 0.12em;
        text-transform: uppercase;
        color: var(--salmon-600);
    }

    .section-head h2,
    .license-copy h2 {
        margin: 14px 0 0;
        font-size: clamp(1.75rem, 3.4vw, 2.5rem);
        line-height: 1.15;
        letter-spacing: -0.03em;
        font-weight: 800;
    }

    .section-head p {
        margin: 16px 0 0;
        font-size: 1.0625rem;
        line-height: 1.7;
        color: var(--muted);
    }

    /* ── Features ──────────────────────────────────────── */
    .features {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: 20px;
    }

    .feature {
        padding: 26px;
        border: 1px solid var(--border);
        border-radius: 16px;
        background: #fff;
        transition: transform 0.2s ease, box-shadow 0.2s ease,
            border-color 0.2s ease;
    }

    .feature:hover {
        transform: translateY(-3px);
        border-color: var(--salmon-200);
        box-shadow: var(--shadow-md);
    }

    .feature-icon {
        display: grid;
        place-items: center;
        width: 42px;
        height: 42px;
        border: 1px solid var(--salmon-100);
        border-radius: 12px;
        background: var(--salmon-50);
        color: var(--salmon-600);
    }

    .feature-icon svg {
        width: 20px;
        height: 20px;
    }

    .feature h3 {
        margin: 18px 0 0;
        font-size: 1.0625rem;
        font-weight: 700;
        letter-spacing: -0.015em;
    }

    .feature p {
        margin: 8px 0 0;
        font-size: 0.9375rem;
        line-height: 1.65;
        color: var(--muted);
    }

    /* ── Licence ───────────────────────────────────────── */
    .license {
        padding: 96px 0;
        border-top: 1px solid var(--border);
        border-bottom: 1px solid var(--border);
        background: var(--surface-2);
    }

    .license-inner {
        display: grid;
        grid-template-columns: 1fr 0.85fr;
        align-items: center;
        gap: 64px;
    }

    .license-copy p {
        margin: 16px 0 0;
        font-size: 1rem;
        line-height: 1.75;
        color: var(--muted);
    }

    .license-copy strong {
        font-weight: 600;
        color: var(--ink);
    }

    .license-actions {
        display: flex;
        flex-wrap: wrap;
        gap: 12px;
        margin-top: 30px;
    }

    .license-card {
        padding: 28px;
        border: 1px solid var(--border);
        border-radius: 18px;
        background: #fff;
        box-shadow: var(--shadow-md);
    }

    .license-card h3 {
        margin: 0;
        font-size: 0.8125rem;
        font-weight: 700;
        letter-spacing: 0.12em;
        text-transform: uppercase;
        color: var(--muted);
    }

    .license-list {
        display: grid;
        gap: 14px;
        margin: 18px 0 0;
        padding: 0;
        list-style: none;
    }

    .license-list li {
        display: flex;
        gap: 12px;
        font-size: 0.9375rem;
        line-height: 1.5;
        color: var(--ink-2);
    }

    .license-list svg {
        width: 18px;
        height: 18px;
        flex: none;
        margin-top: 2px;
        color: var(--salmon-600);
    }

    /* ── CTA ───────────────────────────────────────────── */
    .cta {
        padding: 88px 0;
    }

    .cta-card {
        position: relative;
        overflow: hidden;
        padding: 64px 40px;
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
        font-size: clamp(1.75rem, 3.4vw, 2.4rem);
        font-weight: 800;
        letter-spacing: -0.03em;
    }

    .cta-card p {
        max-width: 520px;
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
        margin-top: 32px;
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
        .hero {
            padding: 60px 0 72px;
        }

        .hero-inner {
            grid-template-columns: 1fr;
            gap: 48px;
        }

        .hero-visual {
            max-width: 620px;
        }

        .features {
            grid-template-columns: repeat(2, 1fr);
        }

        .license-inner {
            grid-template-columns: 1fr;
            gap: 40px;
        }

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

        .section,
        .license {
            padding: 64px 0;
        }

        .features {
            grid-template-columns: 1fr;
        }

        .cta-card {
            padding: 48px 24px;
        }

        .footer-grid {
            grid-template-columns: 1fr;
        }
    }

    @media (prefers-reduced-motion: reduce) {
        .btn,
        .feature {
            transition: none;
        }
    }
</style>