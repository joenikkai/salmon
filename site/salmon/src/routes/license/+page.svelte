<script>
    import { Sitename } from "$lib";
    import { SALMON_MAINTAINERS_EMAIL } from "$lib/data";
    const CONTACT_EMAIL = SALMON_MAINTAINERS_EMAIL;
    // Loads your ./LICENSE file as a raw string (Vite `?raw` import).
    // Adjust the relative path if this route lives elsewhere.
    import licenseText from "../../../LICENSE?raw";

    const year = new Date().getFullYear();

    const canDo = [
        "Use it for any purpose — personal, academic or commercial.",
        "Read, study and audit every line of the source code.",
        "Modify it to fit your needs, or build your own version.",
        "Self-host it, on your own terms, with no vendor lock-in.",
        "Redistribute it, as long as the freedoms are preserved."
    ];

    const mustDo = [
        "Keep the copyright and licence notices intact.",
        "Release your modified source under the same AGPL-3.0 terms.",
        "Offer the source to anyone who interacts with your version over a network.",
        "State clearly that you changed the software, and when."
    ];

    const cannotDo = [
        "Turn a modified version into a closed-source product.",
        "Sublicense the code under different terms.",
        "Hold the authors liable for damages or lost data.",
        "Use it without accepting the terms of the licence."
    ];

    const commercial = [
        "Use Salmon in closed-source products",
        "Embed it in a proprietary service",
        "Remove AGPL network source-sharing obligations",
        "Priority support and long-term maintenance"
    ];
</script>

<svelte:head>
    <title>Licence — {Sitename}</title>
    <meta
        name="description"
        content="Salmon is released under the GNU Affero General Public License v3.0. Read what that means for you, and how to obtain a commercial licence."
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
            <a href="/downloads">Download</a>
            <a href="/license" aria-current="page">Licence</a>
            <a href="/terms-of-use">Terms of use</a>
            <a href="/sponsor">Sponsor</a>
        </nav>

        <div class="header-actions">
            <a class="btn btn-primary" href="/downloads">Download app</a>
        </div>
    </div>
</header>

<main>
    <!-- ── Page hero ─────────────────────────────────────── -->
    <section class="page-hero">
        <div class="container">
            <nav class="crumbs" aria-label="Breadcrumb">
                <a href="/">Home</a>
                <span aria-hidden="true">/</span>
                <span aria-current="page">Licence</span>
            </nav>

            <span class="badge"><span class="dot"></span> GNU AGPL-3.0</span>

            <h1>Licence</h1>

            <p class="lede">
                {Sitename} is free and open-source software, released under the
                <strong>GNU Affero General Public License, version 3.0</strong>. That means
                you're free to use, study, share and improve it — and anyone who builds on
                it passes those same freedoms on.
            </p>

            <div class="hero-meta">
                <div class="meta-item">
                    <span class="meta-label">Licence</span>
                    <span class="meta-value">AGPL-3.0-or-later</span>
                </div>
                <div class="meta-item">
                    <span class="meta-label">Copyright</span>
                    <span class="meta-value">© {year} Joseph Wangai Mwaniki</span>
                </div>
                <div class="meta-item">
                    <span class="meta-label">SPDX identifier</span>
                    <span class="meta-value mono">AGPL-3.0-or-later</span>
                </div>
            </div>

            <div class="hero-actions">
                <a class="btn btn-primary" href="#full-text">Read the full licence</a>
                <a class="btn btn-outline" href="/license#commercial">Get a commercial licence</a>
            </div>
        </div>
    </section>

    <!-- ── Plain-English summary ─────────────────────────── -->
    <section class="section">
        <div class="container">
            <div class="section-head">
                <span class="eyebrow">In plain English</span>
                <h2>What the AGPL means for you.</h2>
                <p>
                    The summary below is a friendly overview, not a substitute for the
                    licence itself. Where the two disagree, the licence text always wins.
                </p>
            </div>

            <div class="perm-grid">
                <article class="perm-card perm-yes">
                    <header>
                        <span class="perm-icon" aria-hidden="true">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.4" stroke-linecap="round" stroke-linejoin="round">
                                <path d="m5 12 4 4 10-10" />
                            </svg>
                        </span>
                        <h3>You can</h3>
                    </header>
                    <ul>
                        {#each canDo as item}
                            <li>{item}</li>
                        {/each}
                    </ul>
                </article>

                <article class="perm-card perm-must">
                    <header>
                        <span class="perm-icon" aria-hidden="true">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M12 9v4M12 17h.01" />
                                <path d="M10.29 3.86 1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z" />
                            </svg>
                        </span>
                        <h3>You must</h3>
                    </header>
                    <ul>
                        {#each mustDo as item}
                            <li>{item}</li>
                        {/each}
                    </ul>
                </article>

                <article class="perm-card perm-no">
                    <header>
                        <span class="perm-icon" aria-hidden="true">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.4" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M18 6 6 18M6 6l12 12" />
                            </svg>
                        </span>
                        <h3>You cannot</h3>
                    </header>
                    <ul>
                        {#each cannotDo as item}
                            <li>{item}</li>
                        {/each}
                    </ul>
                </article>
            </div>

            <p class="callout">
                <strong>The AGPL difference:</strong> if you run a modified version of
                {Sitename} on a network server, you must offer the source code of your
                modified version to the people using it. This is what keeps hosted software
                open too — not just distributed software.
            </p>
        </div>
    </section>

    <!-- ── Commercial licence ────────────────────────────── -->
    <section class="section section-alt" id="commercial">
        <div class="container commercial-inner">
            <div class="commercial-copy">
                <span class="eyebrow">Commercial licence</span>
                <h2>Need to keep your source closed?</h2>
                <p>
                    The AGPL is a great fit for open projects, but it isn't for everyone.
                    If your product can't meet the copyleft obligations, a commercial licence
                    gives you the exact same code under terms that work for you.
                </p>
                <p>
                    You keep full ownership of your own application. No source-sharing
                    requirement, no copyleft, no changes to your existing stack.
                </p>

                <ul class="commercial-list">
                    {#each commercial as item}
                        <li>
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.4" stroke-linecap="round" stroke-linejoin="round">
                                <path d="m5 12 4 4 10-10" />
                            </svg>
                            <span>{item}</span>
                        </li>
                    {/each}
                </ul>

                <div class="hero-actions">
                    <a class="btn btn-primary" href="/license/request">Request a licence</a>
                    <a class="btn btn-outline" href={CONTACT_EMAIL}>Email the maintainer</a>
                </div>
            </div>

            <aside class="commercial-card">
                <h3>Which licence is right for me?</h3>

                <div class="compare">
                    <div class="compare-row">
                        <span class="compare-label">Open-source project</span>
                        <span class="pill pill-yes">AGPL-3.0 — free</span>
                    </div>
                    <div class="compare-row">
                        <span class="compare-label">Self-hosting internally</span>
                        <span class="pill pill-yes">AGPL-3.0 — free</span>
                    </div>
                    <div class="compare-row">
                        <span class="compare-label">Publicly modified SaaS</span>
                        <span class="pill pill-mid">AGPL-3.0 — share source</span>
                    </div>
                    <div class="compare-row">
                        <span class="compare-label">Closed-source product</span>
                        <span class="pill pill-no">Commercial licence</span>
                    </div>
                    <div class="compare-row">
                        <span class="compare-label">Embedded in proprietary app</span>
                        <span class="pill pill-no">Commercial licence</span>
                    </div>
                </div>

                <p class="compare-note">
                    Not sure which applies to you? Get in touch — we'll help you figure it out.
                </p>
            </aside>
        </div>
    </section>

    <!-- ── Copyright notice ──────────────────────────────── -->
    <section class="section">
        <div class="container">
            <div class="notice-card">
                <div class="notice-mark" aria-hidden="true">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="12" cy="12" r="10" />
                        <path d="M14.83 14.83a4 4 0 1 1 0-5.66" />
                    </svg>
                </div>

                <div class="notice-body">
                    <h2>Copyright notice</h2>
                    <pre class="copyright-block">Salmon
Copyright (C) {year} Joseph Wangai Mwaniki

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.</pre>
                </div>
            </div>
        </div>
    </section>

    <!-- ── Full licence text ─────────────────────────────── -->
    <section class="section section-alt" id="full-text">
        <div class="container">
            <div class="section-head">
                <span class="eyebrow">The legal text</span>
                <h2>GNU Affero General Public License v3.0</h2>
                <p>
                    The complete, unmodified licence text as published by the Free Software
                    Foundation. This is the version that governs your use of {Sitename}.
                </p>
            </div>

            <details class="license-details" open>
                <summary>
                    <span class="summary-text">Full licence text</span>
                    <span class="summary-hint">Version 3, 19 November 2007</span>
                    <svg class="chev" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
                        <path d="m6 9 6 6 6-6" />
                    </svg>
                </summary>

                <div class="license-scroll" tabindex="0" aria-label="Full AGPL-3.0 licence text">
                    <pre>{licenseText}</pre>
                </div>

                <p class="license-footnote">
                    Can't find what you need? The canonical text lives at
                    <a href="https://www.gnu.org/licenses/agpl-3.0.html" rel="noopener noreferrer" target="_blank">
                        gnu.org/licenses/agpl-3.0.html
                    </a>.
                </p>
            </details>
        </div>
    </section>

    <!-- ── CTA ───────────────────────────────────────────── -->
    <section class="cta">
        <div class="container">
            <div class="cta-card">
                <h2>Questions about licensing?</h2>
                <p>
                    Whether you're shipping an open-source project or a commercial product,
                    we're happy to point you in the right direction.
                </p>
                <div class="cta-actions">
                    <a class="btn btn-lg btn-white" href="/license/request">Request a licence</a>
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
                    <li><a href="/downloads">Download app</a></li>
                    <li><a href="/license">Get a licence</a></li>
                    <li><a href="/sponsor">Sponsor</a></li>
                </ul>
            </div>

            <div>
                <h4>Legal</h4>
                <ul>
                    <li><a href="/license" aria-current="page">Licence</a></li>
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
            <span>© {year} {Sitename}. All rights reserved.</span>
            <span>Released under the GNU AGPL-3.0.</span>
        </div>
    </div>
</footer>

<style>
    /* ── Design tokens (same as landing page) ──────────── */
    :global(:root) {
        --salmon-50: #fff4f1;
        --salmon-100: #ffe7e0;
        --salmon-200: #ffd0c4;
        --salmon-300: #ffb4a2;
        --salmon-400: #ff9581;
        --salmon-500: #fa8072;
        --salmon-600: #ea5f4d;
        --salmon-700: #c4442f;

        --green-50: #f0fdf4;
        --green-600: #16a34a;

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

    .container {
        width: 100%;
        max-width: 1140px;
        margin: 0 auto;
        padding: 0 24px;
    }

    .mono {
        font-family: ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, monospace;
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

    .btn:focus-visible,
    a:focus-visible,
    summary:focus-visible {
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

    /* ── Page hero ─────────────────────────────────────── */
    .page-hero {
        position: relative;
        overflow: hidden;
        padding: 64px 0 72px;
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

    .lede strong {
        font-weight: 600;
        color: var(--ink-2);
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
    }

    .hero-actions {
        display: flex;
        flex-wrap: wrap;
        gap: 12px;
        margin-top: 30px;
    }

    /* ── Sections ──────────────────────────────────────── */
    .section {
        padding: 88px 0;
    }

    .section-alt {
        background: var(--surface-2);
        border-top: 1px solid var(--border);
        border-bottom: 1px solid var(--border);
    }

    .section-head {
        max-width: 660px;
        margin: 0 auto 48px;
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
    .commercial-copy h2 {
        margin: 14px 0 0;
        font-size: clamp(1.75rem, 3.2vw, 2.375rem);
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

    /* ── Permission cards ──────────────────────────────── */
    .perm-grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: 20px;
    }

    .perm-card {
        padding: 28px;
        border: 1px solid var(--border);
        border-radius: 18px;
        background: #fff;
        box-shadow: var(--shadow-sm);
    }

    .perm-card header {
        display: flex;
        align-items: center;
        gap: 12px;
        margin-bottom: 20px;
    }

    .perm-card h3 {
        margin: 0;
        font-size: 1.0625rem;
        font-weight: 700;
        letter-spacing: -0.015em;
    }

    .perm-icon {
        display: grid;
        place-items: center;
        width: 36px;
        height: 36px;
        flex: none;
        border-radius: 10px;
    }

    .perm-icon svg {
        width: 18px;
        height: 18px;
    }

    .perm-yes .perm-icon {
        background: var(--green-50);
        color: var(--green-600);
    }

    .perm-must .perm-icon {
        background: var(--salmon-50);
        color: var(--salmon-600);
    }

    .perm-no .perm-icon {
        background: #fef2f2;
        color: #dc2626;
    }

    .perm-card ul {
        display: grid;
        gap: 12px;
        margin: 0;
        padding: 0;
        list-style: none;
    }

    .perm-card li {
        position: relative;
        padding-left: 18px;
        font-size: 0.9375rem;
        line-height: 1.6;
        color: var(--ink-2);
    }

    .perm-card li::before {
        content: "";
        position: absolute;
        left: 0;
        top: 0.6em;
        width: 6px;
        height: 6px;
        border-radius: 50%;
        background: var(--salmon-300);
    }

    .callout {
        max-width: 760px;
        margin: 40px auto 0;
        padding: 20px 24px;
        border: 1px solid var(--salmon-200);
        border-left: 4px solid var(--salmon-500);
        border-radius: 12px;
        background: var(--salmon-50);
        font-size: 0.9375rem;
        line-height: 1.7;
        color: var(--ink-2);
        text-align: left;
    }

    .callout strong {
        color: var(--salmon-700);
        font-weight: 700;
    }

    /* ── Commercial ────────────────────────────────────── */
    .commercial-inner {
        display: grid;
        grid-template-columns: 1fr 0.9fr;
        align-items: center;
        gap: 64px;
    }

    .commercial-copy p {
        margin: 16px 0 0;
        font-size: 1rem;
        line-height: 1.75;
        color: var(--muted);
    }

    .commercial-list {
        display: grid;
        gap: 12px;
        margin: 28px 0 0;
        padding: 0;
        list-style: none;
    }

    .commercial-list li {
        display: flex;
        gap: 12px;
        font-size: 0.9375rem;
        line-height: 1.5;
        color: var(--ink-2);
    }

    .commercial-list svg {
        width: 18px;
        height: 18px;
        flex: none;
        margin-top: 2px;
        color: var(--salmon-600);
    }

    .commercial-card {
        padding: 28px;
        border: 1px solid var(--border);
        border-radius: 18px;
        background: #fff;
        box-shadow: var(--shadow-md);
    }

    .commercial-card h3 {
        margin: 0;
        font-size: 1rem;
        font-weight: 700;
        letter-spacing: -0.015em;
    }

    .compare {
        display: grid;
        gap: 2px;
        margin-top: 18px;
    }

    .compare-row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 14px;
        padding: 12px 0;
        border-bottom: 1px solid var(--border);
    }

    .compare-row:last-child {
        border-bottom: 0;
    }

    .compare-label {
        font-size: 0.9375rem;
        color: var(--ink-2);
    }

    .pill {
        flex: none;
        padding: 4px 10px;
        border-radius: 999px;
        font-size: 0.75rem;
        font-weight: 600;
        white-space: nowrap;
    }

    .pill-yes {
        background: var(--green-50);
        color: var(--green-600);
        border: 1px solid #bbf7d0;
    }

    .pill-mid {
        background: var(--salmon-50);
        color: var(--salmon-700);
        border: 1px solid var(--salmon-200);
    }

    .pill-no {
        background: #f1f5f9;
        color: var(--ink-2);
        border: 1px solid var(--border);
    }

    .compare-note {
        margin: 18px 0 0;
        font-size: 0.875rem;
        line-height: 1.6;
        color: var(--muted);
    }

    /* ── Copyright notice ──────────────────────────────── */
    .notice-card {
        display: grid;
        grid-template-columns: auto 1fr;
        gap: 28px;
        padding: 36px;
        border: 1px solid var(--border);
        border-radius: 20px;
        background: #fff;
        box-shadow: var(--shadow-sm);
    }

    .notice-mark {
        display: grid;
        place-items: center;
        width: 48px;
        height: 48px;
        border-radius: 14px;
        background: var(--salmon-50);
        border: 1px solid var(--salmon-100);
        color: var(--salmon-600);
    }

    .notice-mark svg {
        width: 24px;
        height: 24px;
    }

    .notice-body h2 {
        margin: 0;
        font-size: 1.25rem;
        font-weight: 700;
        letter-spacing: -0.02em;
    }

    .copyright-block {
        margin: 18px 0 0;
        padding: 20px;
        border: 1px solid var(--border);
        border-radius: 12px;
        background: var(--surface-2);
        font-family: ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, monospace;
        font-size: 0.8125rem;
        line-height: 1.7;
        color: var(--ink-2);
        white-space: pre-wrap;
        overflow-x: auto;
    }

    /* ── Full licence text ─────────────────────────────── */
    .license-details {
        border: 1px solid var(--border);
        border-radius: 18px;
        background: #fff;
        box-shadow: var(--shadow-sm);
        overflow: hidden;
    }

    .license-details summary {
        display: flex;
        align-items: center;
        gap: 16px;
        padding: 22px 26px;
        cursor: pointer;
        list-style: none;
        user-select: none;
    }

    .license-details summary::-webkit-details-marker {
        display: none;
    }

    .summary-text {
        font-size: 1rem;
        font-weight: 700;
        letter-spacing: -0.015em;
    }

    .summary-hint {
        margin-left: auto;
        font-size: 0.8125rem;
        color: var(--muted);
    }

    .chev {
        width: 20px;
        height: 20px;
        flex: none;
        color: var(--muted);
        transition: transform 0.2s ease;
    }

    .license-details[open] .chev {
        transform: rotate(180deg);
    }

    .license-scroll {
        max-height: 520px;
        overflow-y: auto;
        padding: 0 26px 26px;
        border-top: 1px solid var(--border);
    }

    .license-scroll pre {
        margin: 22px 0 0;
        font-family: ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, monospace;
        font-size: 0.75rem;
        line-height: 1.75;
        color: var(--ink-2);
        white-space: pre-wrap;
        word-break: break-word;
    }

    .license-footnote {
        margin: 0;
        padding: 18px 26px 22px;
        border-top: 1px solid var(--border);
        background: var(--surface-2);
        font-size: 0.875rem;
        color: var(--muted);
    }

    .license-footnote a {
        color: var(--salmon-600);
        font-weight: 600;
        text-decoration: none;
    }

    .license-footnote a:hover {
        text-decoration: underline;
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
        .perm-grid {
            grid-template-columns: 1fr;
        }

        .commercial-inner {
            grid-template-columns: 1fr;
            gap: 40px;
        }

        .footer-grid {
            grid-template-columns: 1fr 1fr;
            gap: 32px;
        }

        .nav-desktop {
            display: none;
        }
    }

    @media (max-width: 760px) {
        .section,
        .section-alt {
            padding: 64px 0;
        }

        .page-hero {
            padding: 48px 0 56px;
        }

        .notice-card {
            grid-template-columns: 1fr;
            gap: 20px;
            padding: 26px;
        }

        .summary-hint {
            display: none;
        }

        .license-scroll {
            max-height: 420px;
            padding: 0 18px 22px;
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
        .chev {
            transition: none;
        }
    }
</style>
