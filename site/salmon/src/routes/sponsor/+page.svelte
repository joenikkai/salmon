<script>
    import { Sitename } from "$lib";

    const year = new Date().getFullYear();

    const impact = [
        {
            title: "Development time",
            body: "Direct funding means more hours on features, bug fixes and releases — not just evenings and weekends.",
            icon: "M12 6v6l4 2M12 2a10 10 0 1 0 0 20 10 10 0 0 0 0-20z"
        },
        {
            title: "Infrastructure",
            body: "Build servers, CI minutes, test devices and hosting for docs and downloads.",
            icon: "M4 4h16v6H4zM4 14h16v6H4zM8 7h.01M8 17h.01"
        },
        {
            title: "Independence",
            body: "No investors, no ads, no data selling. Funding comes from the people who use it.",
            icon: "M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"
        },
        {
            title: "Long-term care",
            body: "Security patches, dependency updates and keeping the project healthy for years to come.",
            icon: "M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l8.84 8.84 8.84-8.84a5.5 5.5 0 0 0 0-7.78z"
        }
    ];

    const tiers = [
        {
            name: "Supporter",
            price: "5",
            period: "/month",
            tagline: "For individuals who want to say thanks.",
            perks: [
                "Name in the SUPPORTERS file",
                "Supporter badge on GitHub",
                "Early access to release notes"
            ],
            cta: "Become a supporter",
            featured: false
        },
        {
            name: "Backer",
            price: "25",
            period: "/month",
            tagline: "For regular users who rely on the project.",
            perks: [
                "Everything in Supporter",
                "Your name or handle in the README",
                "Vote on the public roadmap",
                "Priority responses on issues"
            ],
            cta: "Become a backer",
            featured: true
        },
        {
            name: "Sponsor",
            price: "100",
            period: "/month",
            tagline: "For teams building on top of Salmon.",
            perks: [
                "Everything in Backer",
                "Logo in the README and on the site",
                "Acknowledgement in release posts",
                "Direct line for questions"
            ],
            cta: "Become a sponsor",
            featured: false
        }
    ];

    const oneTime = [
        { label: "Buy me a coffee", price: "$5", note: "One-off thanks" },
        { label: "Support a release", price: "$50", note: "Funds a focused day" },
        { label: "Sponsor a feature", price: "$500", note: "Commission work" }
    ];

    const sponsors = [
        { name: "Your name here", tier: "Backer" },
        { name: "Your name here", tier: "Backer" },
        { name: "Your team here", tier: "Sponsor" },
        { name: "Your name here", tier: "Supporter" }
    ];

    const faqs = [
        {
            q: "What does my money actually pay for?",
            a: "Mainly time. Development, code review, documentation and releases all take hours. The rest covers infrastructure like build servers and hosting. Nothing goes to ads or third parties."
        },
        {
            q: "Is this a donation or a purchase?",
            a: "It's a sponsorship, not a purchase. You're supporting the work — you don't receive goods or services in return, and sponsorship doesn't change your rights under the AGPL."
        },
        {
            q: "Does sponsoring give me a commercial licence?",
            a: "No. Sponsorship and licensing are separate. If you need to use Salmon in a closed-source product, request a commercial licence. Sponsoring is never required to use the software."
        },
        {
            q: "Can I sponsor as a company?",
            a: "Absolutely, and it's appreciated. Company sponsors get their logo in the README and on this page. Get in touch if you'd like an invoice or a custom arrangement."
        },
        {
            q: "Can I cancel any time?",
            a: "Yes. Monthly sponsorships can be cancelled whenever you like, and you keep all the recognition you've already earned."
        },
        {
            q: "I can't afford to sponsor — can I still help?",
            a: "Definitely. Reporting bugs, improving docs, answering questions, sharing the project and starring the repo all make a real difference."
        }
    ];

    let openFaq = 0;

    function toggleFaq(i) {
        openFaq = openFaq === i ? -1 : i;
    }
</script>

<svelte:head>
    <title>Sponsor — {Sitename}</title>
    <meta
        name="description"
        content="Sponsor Salmon and keep it independent, maintained and free for everyone. Monthly and one-time options available."
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
            <a href="/license">Licence</a>
            <a href="/terms-of-use">Terms of use</a>
            <a href="/sponsor" aria-current="page">Sponsor</a>
        </nav>

        <div class="header-actions">
            <a class="btn btn-primary" href="#tiers">Sponsor now</a>
        </div>
    </div>
</header>

<main>
    <!-- ── Hero ──────────────────────────────────────────── -->
    <section class="page-hero">
        <div class="container">
            <nav class="crumbs" aria-label="Breadcrumb">
                <a href="/">Home</a>
                <span aria-hidden="true">/</span>
                <span aria-current="page">Sponsor</span>
            </nav>

            <span class="badge"><span class="dot"></span> Funded by the community</span>

            <h1>Keep {Sitename} <span class="accent">independent</span>.</h1>

            <p class="lede">
                {Sitename} is free, open source and has no investors, no ads and no interest
                in your data. Sponsorship is what keeps it that way — paying for the time
                and infrastructure that keep the project healthy.
            </p>

            <div class="hero-actions">
                <a class="btn btn-primary btn-lg" href="#tiers">Sponsor the project</a>
                <a class="btn btn-outline btn-lg" href="https://github.com/joenikkai/salmon" rel="noopener">
                    Star on GitHub
                </a>
            </div>

            <div class="hero-stats">
                <div class="stat">
                    <span class="stat-value">AGPL-3.0</span>
                    <span class="stat-label">Free forever</span>
                </div>
                <div class="stat">
                    <span class="stat-value">0</span>
                    <span class="stat-label">Trackers or ads</span>
                </div>
                <div class="stat">
                    <span class="stat-value">100%</span>
                    <span class="stat-label">Goes to the project</span>
                </div>
            </div>
        </div>
    </section>

    <!-- ── Why sponsor ───────────────────────────────────── -->
    <section class="section">
        <div class="container">
            <div class="section-head">
                <span class="eyebrow">Where it goes</span>
                <h2>Every contribution funds real work.</h2>
                <p>
                    No salaries, no offices, no marketing budget. Sponsorship goes directly
                    into building and maintaining the software you use.
                </p>
            </div>

            <div class="impact-grid">
                {#each impact as item}
                    <article class="impact-card">
                        <div class="impact-icon">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                                <path d={item.icon} />
                            </svg>
                        </div>
                        <h3>{item.title}</h3>
                        <p>{item.body}</p>
                    </article>
                {/each}
            </div>
        </div>
    </section>

    <!-- ── Tiers ─────────────────────────────────────────── -->
    <section class="section section-alt" id="tiers">
        <div class="container">
            <div class="section-head">
                <span class="eyebrow">Monthly sponsorship</span>
                <h2>Pick a tier that suits you.</h2>
                <p>
                    Recurring support is the most valuable kind — it makes long-term planning
                    possible. Cancel any time, no questions asked.
                </p>
            </div>

            <div class="tiers">
                {#each tiers as tier}
                    <article class="tier" class:featured={tier.featured}>
                        {#if tier.featured}
                            <span class="tier-ribbon">Most popular</span>
                        {/if}

                        <h3>{tier.name}</h3>
                        <p class="tier-tagline">{tier.tagline}</p>

                        <div class="tier-price">
                            <span class="currency">$</span>
                            <span class="amount">{tier.price}</span>
                            <span class="period">{tier.period}</span>
                        </div>

                        <ul class="tier-perks">
                            {#each tier.perks as perk}
                                <li>
                                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                                        <path d="m5 12 4 4 10-10" />
                                    </svg>
                                    <span>{perk}</span>
                                </li>
                            {/each}
                        </ul>

                        <a
                            class="btn {tier.featured ? 'btn-primary' : 'btn-outline'} tier-cta"
                            href="https://github.com/sponsors/joenikkai"
                            rel="noopener"
                        >
                            {tier.cta}
                        </a>
                    </article>
                {/each}
            </div>

            <!-- One-time -->
            <div class="onetime">
                <div class="onetime-head">
                    <h3>Prefer a one-time contribution?</h3>
                    <p>No commitment, same impact. Perfect if you just want to say thanks.</p>
                </div>

                <div class="onetime-grid">
                    {#each oneTime as option}
                        <a
                            class="onetime-card"
                            href="https://github.com/sponsors/joenikkai"
                            rel="noopener"
                        >
                            <span class="onetime-price">{option.price}</span>
                            <span class="onetime-label">{option.label}</span>
                            <span class="onetime-note">{option.note}</span>
                        </a>
                    {/each}
                </div>
            </div>
        </div>
    </section>

    <!-- ── Sponsors wall ─────────────────────────────────── -->
    <section class="section">
        <div class="container">
            <div class="section-head">
                <span class="eyebrow">With thanks to</span>
                <h2>Our supporters.</h2>
                <p>
                    These people and teams make {Sitename} possible. Sponsors at the Backer
                    tier and above are listed here and in the README.
                </p>
            </div>

            <div class="wall">
                {#each sponsors as s}
                    <div class="wall-item" class:placeholder={s.name === "Your name here" || s.name === "Your team here"}>
                        <span class="wall-name">{s.name}</span>
                        <span class="wall-tier">{s.tier}</span>
                    </div>
                {/each}
            </div>

            <p class="wall-note">
                Want your name here? Pick any tier above — recognition is applied within a day.
            </p>
        </div>
    </section>

    <!-- ── FAQ ───────────────────────────────────────────── -->
    <section class="section section-alt">
        <div class="container">
            <div class="section-head">
                <span class="eyebrow">Questions</span>
                <h2>Before you sponsor.</h2>
                <p>The honest answers to what people usually ask.</p>
            </div>

            <div class="faq">
                {#each faqs as faq, i}
                    <div class="faq-item" class:open={openFaq === i}>
                        <button
                            class="faq-q"
                            aria-expanded={openFaq === i}
                            on:click={() => toggleFaq(i)}
                        >
                            <span>{faq.q}</span>
                            <svg class="chev" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
                                <path d="m6 9 6 6 6-6" />
                            </svg>
                        </button>

                        {#if openFaq === i}
                            <div class="faq-a">
                                <p>{faq.a}</p>
                            </div>
                        {/if}
                    </div>
                {/each}
            </div>
        </div>
    </section>

    <!-- ── Other ways to help ────────────────────────────── -->
    <section class="section">
        <div class="container">
            <div class="help-card">
                <div class="help-copy">
                    <span class="eyebrow">Not able to sponsor?</span>
                    <h2>There are other ways to help.</h2>
                    <p>
                        Money isn't the only contribution that matters. Every issue reported,
                        typo fixed and question answered makes the project better for
                        everyone.
                    </p>
                </div>

                <ul class="help-list">
                    <li>
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                            <path d="m12 3 2.09 6.26L20.5 9.5l-5 4.02 1.81 6.48L12 16.77 6.69 20l1.81-6.48-5-4.02 6.41-.24z" />
                        </svg>
                        <div>
                            <strong>Star the repository</strong>
                            <span>Helps others discover the project.</span>
                        </div>
                    </li>
                    <li>
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M12 20h9M16.5 3.5a2.12 2.12 0 0 1 3 3L7 19l-4 1 1-4z" />
                        </svg>
                        <div>
                            <strong>Improve the docs</strong>
                            <span>Clear documentation is a gift.</span>
                        </div>
                    </li>
                    <li>
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M12 9v4M12 17h.01M10.29 3.86 1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z" />
                        </svg>
                        <div>
                            <strong>Report bugs clearly</strong>
                            <span>Good reports save hours of guesswork.</span>
                        </div>
                    </li>
                    <li>
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2M9 11a4 4 0 1 0 0-8 4 4 0 0 0 0 8zM23 21v-2a4 4 0 0 0-3-3.87M16 3.13a4 4 0 0 1 0 7.75" />
                        </svg>
                        <div>
                            <strong>Tell someone about it</strong>
                            <span>Word of mouth is the best marketing.</span>
                        </div>
                    </li>
                </ul>

                <div class="help-actions">
                    <a class="btn btn-outline" href="https://github.com/joenikkai/salmon" rel="noopener">
                        Contribute on GitHub
                    </a>
                </div>
            </div>
        </div>
    </section>

    <!-- ── CTA ───────────────────────────────────────────── -->
    <section class="cta">
        <div class="container">
            <div class="cta-card">
                <h2>Thank you for considering it.</h2>
                <p>
                    Every sponsor, at every tier, keeps {Sitename} free and independent for
                    the people who need it most.
                </p>
                <div class="cta-actions">
                    <a class="btn btn-lg btn-white" href="#tiers">Sponsor the project</a>
                    <a class="btn btn-lg btn-oncolor" href="/downloads">Download app</a>
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
                    <li><a href="/sponsor" aria-current="page">Sponsor</a></li>
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
            <span>© {year} {Sitename}. All rights reserved.</span>
            <span>Released under the GNU AGPL-3.0.</span>
        </div>
    </div>
</footer>

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

    .accent {
        color: var(--salmon-600);
        background: linear-gradient(transparent 66%, var(--salmon-200) 66%);
        border-radius: 2px;
    }

    .lede {
        margin: 18px 0 0;
        max-width: 46rem;
        font-size: 1.125rem;
        line-height: 1.7;
        color: var(--muted);
    }

    .hero-actions {
        display: flex;
        flex-wrap: wrap;
        gap: 12px;
        margin-top: 30px;
    }

    .hero-stats {
        display: flex;
        flex-wrap: wrap;
        gap: 16px 48px;
        margin-top: 40px;
        padding-top: 26px;
        border-top: 1px solid var(--border);
    }

    .stat {
        display: grid;
        gap: 2px;
    }

    .stat-value {
        font-size: 1.25rem;
        font-weight: 800;
        letter-spacing: -0.02em;
        color: var(--salmon-600);
    }

    .stat-label {
        font-size: 0.8125rem;
        color: var(--muted);
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
        margin: 0 auto 52px;
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
    .help-copy h2 {
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

    /* ── Impact ────────────────────────────────────────── */
    .impact-grid {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        gap: 20px;
    }

    .impact-card {
        padding: 26px;
        border: 1px solid var(--border);
        border-radius: 16px;
        background: #fff;
        transition: transform 0.2s ease, box-shadow 0.2s ease,
            border-color 0.2s ease;
    }

    .impact-card:hover {
        transform: translateY(-3px);
        border-color: var(--salmon-200);
        box-shadow: var(--shadow-md);
    }

    .impact-icon {
        display: grid;
        place-items: center;
        width: 42px;
        height: 42px;
        border: 1px solid var(--salmon-100);
        border-radius: 12px;
        background: var(--salmon-50);
        color: var(--salmon-600);
    }

    .impact-icon svg {
        width: 20px;
        height: 20px;
    }

    .impact-card h3 {
        margin: 18px 0 0;
        font-size: 1.0625rem;
        font-weight: 700;
        letter-spacing: -0.015em;
    }

    .impact-card p {
        margin: 8px 0 0;
        font-size: 0.9375rem;
        line-height: 1.65;
        color: var(--muted);
    }

    /* ── Tiers ─────────────────────────────────────────── */
    .tiers {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: 22px;
        align-items: start;
    }

    .tier {
        position: relative;
        display: flex;
        flex-direction: column;
        padding: 32px 28px;
        border: 1px solid var(--border);
        border-radius: 20px;
        background: #fff;
        box-shadow: var(--shadow-sm);
        transition: transform 0.2s ease, box-shadow 0.2s ease,
            border-color 0.2s ease;
    }

    .tier:hover {
        transform: translateY(-3px);
        box-shadow: var(--shadow-md);
    }

    .tier.featured {
        border-color: var(--salmon-300);
        box-shadow: var(--shadow-md), 0 0 0 1px var(--salmon-200);
    }

    .tier-ribbon {
        position: absolute;
        top: -12px;
        left: 28px;
        padding: 4px 12px;
        border-radius: 999px;
        background: linear-gradient(180deg, var(--salmon-500), var(--salmon-600));
        color: #fff;
        font-size: 0.6875rem;
        font-weight: 700;
        letter-spacing: 0.06em;
        text-transform: uppercase;
        box-shadow: 0 8px 18px -10px rgba(234, 95, 77, 0.9);
    }

    .tier h3 {
        margin: 0;
        font-size: 1.125rem;
        font-weight: 700;
        letter-spacing: -0.02em;
    }

    .tier-tagline {
        margin: 8px 0 0;
        font-size: 0.875rem;
        line-height: 1.55;
        color: var(--muted);
        min-height: 2.6em;
    }

    .tier-price {
        display: flex;
        align-items: baseline;
        gap: 2px;
        margin: 22px 0 0;
        padding-bottom: 22px;
        border-bottom: 1px solid var(--border);
    }

    .currency {
        font-size: 1.25rem;
        font-weight: 700;
        color: var(--ink-2);
        align-self: flex-start;
        margin-top: 6px;
    }

    .amount {
        font-size: 2.75rem;
        font-weight: 800;
        letter-spacing: -0.04em;
        line-height: 1;
    }

    .period {
        font-size: 0.9375rem;
        font-weight: 500;
        color: var(--muted);
    }

    .tier-perks {
        display: grid;
        gap: 12px;
        margin: 22px 0 0;
        padding: 0;
        list-style: none;
        flex: 1;
    }

    .tier-perks li {
        display: flex;
        gap: 10px;
        font-size: 0.9375rem;
        line-height: 1.5;
        color: var(--ink-2);
    }

    .tier-perks svg {
        width: 17px;
        height: 17px;
        flex: none;
        margin-top: 2px;
        color: var(--salmon-600);
    }

    .tier-cta {
        width: 100%;
        margin-top: 26px;
    }

    /* ── One-time ──────────────────────────────────────── */
    .onetime {
        margin-top: 56px;
        padding: 32px;
        border: 1px solid var(--border);
        border-radius: 20px;
        background: #fff;
    }

    .onetime-head h3 {
        margin: 0;
        font-size: 1.125rem;
        font-weight: 700;
        letter-spacing: -0.02em;
    }

    .onetime-head p {
        margin: 8px 0 0;
        font-size: 0.9375rem;
        color: var(--muted);
    }

    .onetime-grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: 14px;
        margin-top: 22px;
    }

    .onetime-card {
        display: grid;
        gap: 4px;
        padding: 20px;
        border: 1px solid var(--border);
        border-radius: 14px;
        background: var(--surface-2);
        text-decoration: none;
        color: inherit;
        transition: border-color 0.15s ease, background-color 0.15s ease,
            transform 0.15s ease;
    }

    .onetime-card:hover {
        transform: translateY(-2px);
        border-color: var(--salmon-300);
        background: var(--salmon-50);
    }

    .onetime-price {
        font-size: 1.375rem;
        font-weight: 800;
        letter-spacing: -0.03em;
        color: var(--salmon-600);
    }

    .onetime-label {
        font-size: 0.9375rem;
        font-weight: 600;
        color: var(--ink);
    }

    .onetime-note {
        font-size: 0.8125rem;
        color: var(--muted);
    }

    /* ── Sponsors wall ─────────────────────────────────── */
    .wall {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        gap: 14px;
    }

    .wall-item {
        display: grid;
        gap: 4px;
        padding: 22px 20px;
        border: 1px solid var(--border);
        border-radius: 14px;
        background: #fff;
        text-align: center;
    }

    .wall-item.placeholder {
        border-style: dashed;
        background: var(--surface-2);
    }

    .wall-name {
        font-size: 0.9375rem;
        font-weight: 600;
        color: var(--ink);
    }

    .wall-item.placeholder .wall-name {
        color: var(--muted);
        font-weight: 500;
    }

    .wall-tier {
        font-size: 0.75rem;
        font-weight: 700;
        letter-spacing: 0.08em;
        text-transform: uppercase;
        color: var(--salmon-600);
    }

    .wall-item.placeholder .wall-tier {
        color: #94a3b8;
    }

    .wall-note {
        margin: 26px 0 0;
        text-align: center;
        font-size: 0.875rem;
        color: var(--muted);
    }

    /* ── FAQ ───────────────────────────────────────────── */
    .faq {
        max-width: 780px;
        margin: 0 auto;
        border: 1px solid var(--border);
        border-radius: 18px;
        background: #fff;
        overflow: hidden;
        box-shadow: var(--shadow-sm);
    }

    .faq-item + .faq-item {
        border-top: 1px solid var(--border);
    }

    .faq-q {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 20px;
        width: 100%;
        padding: 22px 26px;
        border: 0;
        background: transparent;
        font: inherit;
        font-size: 1rem;
        font-weight: 600;
        letter-spacing: -0.01em;
        color: var(--ink);
        text-align: left;
        cursor: pointer;
        transition: background-color 0.15s ease;
    }

    .faq-q:hover {
        background: var(--salmon-50);
    }

    .chev {
        width: 20px;
        height: 20px;
        flex: none;
        color: var(--muted);
        transition: transform 0.2s ease;
    }

    .faq-item.open .chev {
        transform: rotate(180deg);
        color: var(--salmon-600);
    }

    .faq-a {
        padding: 0 26px 24px;
    }

    .faq-a p {
        margin: 0;
        max-width: 62ch;
        font-size: 0.9375rem;
        line-height: 1.75;
        color: var(--muted);
    }

    /* ── Other ways to help ────────────────────────────── */
    .help-card {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 48px;
        padding: 44px;
        border: 1px solid var(--border);
        border-radius: 22px;
        background: #fff;
        box-shadow: var(--shadow-sm);
    }

    .help-copy p {
        margin: 16px 0 0;
        font-size: 1rem;
        line-height: 1.75;
        color: var(--muted);
    }

    .help-actions {
        margin-top: 28px;
    }

    .help-list {
        display: grid;
        gap: 18px;
        margin: 0;
        padding: 0;
        list-style: none;
    }

    .help-list li {
        display: flex;
        gap: 14px;
    }

    .help-list svg {
        width: 20px;
        height: 20px;
        flex: none;
        margin-top: 2px;
        color: var(--salmon-600);
    }

    .help-list strong {
        display: block;
        font-size: 0.9375rem;
        font-weight: 600;
        color: var(--ink);
    }

    .help-list span {
        display: block;
        margin-top: 2px;
        font-size: 0.875rem;
        line-height: 1.55;
        color: var(--muted);
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
        .impact-grid {
            grid-template-columns: repeat(2, 1fr);
        }

        .tiers {
            grid-template-columns: 1fr;
            max-width: 460px;
            margin: 0 auto;
        }

        .onetime-grid {
            grid-template-columns: 1fr;
        }

        .wall {
            grid-template-columns: repeat(2, 1fr);
        }

        .help-card {
            grid-template-columns: 1fr;
            gap: 32px;
            padding: 32px;
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

        .impact-grid {
            grid-template-columns: 1fr;
        }

        .wall {
            grid-template-columns: 1fr;
        }

        .faq-q {
            padding: 18px 20px;
            font-size: 0.9375rem;
        }

        .faq-a {
            padding: 0 20px 20px;
        }

        .onetime {
            padding: 24px;
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
        .chev,
        .tier,
        .impact-card,
        .onetime-card {
            transition: none;
        }
    }
</style>