 <script>
    import { Sitename } from "$lib";
    import { SALMON_MAINTAINERS_EMAIL } from "$lib/data";

    const year = new Date().getFullYear();

    const licenceTypes = [
        { value: "commercial", label: "Commercial licence (closed-source use)" },
        { value: "oem", label: "OEM / embedding in a proprietary product" },
        { value: "saas", label: "Hosted SaaS that can't meet AGPL obligations" },
        { value: "clarification", label: "I just have a licensing question" },
        { value: "other", label: "Something else" }
    ];

    const deployments = [
        { value: "saas", label: "Hosted service / SaaS" },
        { value: "on-prem", label: "On-premise installation" },
        { value: "desktop", label: "Desktop application" },
        { value: "mobile", label: "Mobile application" },
        { value: "embedded", label: "Embedded / shipped with hardware" },
        { value: "internal", label: "Internal tooling only" },
        { value: "other", label: "Other" }
    ];

    let fullName = "";
    let email = "";
    let company = "";
    let website = "";
    let licenceType = "commercial";
    let deployment = "saas";
    let message = "";
    let agree = false;

    let sending = false;
    let submitted = false;
    let error = "";

    async function handleSubmit() {
        error = "";

        if (!fullName.trim() || !email.trim() || !message.trim()) {
            error = "Please fill in your name, email address and a short description of your use case.";
            return;
        }

        if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.trim())) {
            error = "That email address doesn't look quite right. Please check it and try again.";
            return;
        }

        if (!agree) {
            error = "Please confirm you're happy for us to contact you about this request.";
            return;
        }

        sending = true;

        try {
            // Replace this with your own endpoint or a SvelteKit form action.
            await new Promise((resolve) => setTimeout(resolve, 700));
            submitted = true;
        } catch (err) {
            error = `Something went wrong sending your request. Please email ${SALMON_MAINTAINERS_EMAIL} instead.`;
        } finally {
            sending = false;
        }
    }

    function resetForm() {
        submitted = false;
        error = "";
        fullName = "";
        email = "";
        company = "";
        website = "";
        licenceType = "commercial";
        deployment = "saas";
        message = "";
        agree = false;
    }
</script>

<svelte:head>
    <title>Request a licence — {Sitename}</title>
    <meta
        name="description"
        content="Request a commercial licence for Salmon. Tell us about your project and we'll get back to you with terms that work for your team."
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
                <a href="/license">Licence</a>
                <span aria-hidden="true">/</span>
                <span aria-current="page">Request a licence</span>
            </nav>

            <span class="badge"><span class="dot"></span> Commercial licence</span>

            <h1>Request a licence</h1>

            <p class="lede">
                Tell us a little about what you're building and how you plan to use
                {Sitename}. We'll come back to you with terms that fit — no copyleft, no
                source-sharing obligation, and no changes to your existing stack.
            </p>

            <div class="hero-meta">
                <div class="meta-item">
                    <span class="meta-label">Typical reply</span>
                    <span class="meta-value">1–2 business days</span>
                </div>
                <div class="meta-item">
                    <span class="meta-label">Licence</span>
                    <span class="meta-value">Per project, perpetual or annual</span>
                </div>
                <div class="meta-item">
                    <span class="meta-label">Prefer email?</span>
                    <span class="meta-value mono">{SALMON_MAINTAINERS_EMAIL}</span>
                </div>
            </div>
        </div>
    </section>

    <!-- ── Request form ──────────────────────────────────── -->
    <section class="section">
        <div class="container request-inner">
            <div>
                {#if submitted}
                    <!-- ── Success state ─────────────────────── -->
                    <div class="form-card success-card">
                        <span class="success-mark" aria-hidden="true">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.4" stroke-linecap="round" stroke-linejoin="round">
                                <path d="m5 12 4 4 10-10" />
                            </svg>
                        </span>

                        <h2>Request received</h2>

                        <p>
                            Thanks, {fullName.split(" ")[0] || "there"} — we've got your request.
                            You'll hear back from a maintainer at
                            <strong>{email}</strong> within one to two business days.
                        </p>

                        <p class="success-note">
                            Need it sooner, or want to add something? Reply to
                            <a href="mailto:{SALMON_MAINTAINERS_EMAIL}">{SALMON_MAINTAINERS_EMAIL}</a>
                            and mention your project name.
                        </p>

                        <div class="hero-actions success-actions">
                            <a class="btn btn-primary" href="/downloads">Download {Sitename}</a>
                            <button class="btn btn-outline" type="button" on:click={resetForm}>
                                Send another request
                            </button>
                        </div>
                    </div>
                {:else}
                    <!-- ── Form ──────────────────────────────── -->
                    <form class="form-card" on:submit|preventDefault={handleSubmit} novalidate>
                        <div class="form-head">
                            <h2>Tell us about your use case</h2>
                            <p>Fields marked with <span class="req">*</span> are required.</p>
                        </div>

                        <div class="form-grid">
                            <div class="field">
                                <label for="fullName">Full name <span class="req">*</span></label>
                                <input
                                    id="fullName"
                                    class="control"
                                    type="text"
                                    name="fullName"
                                    autocomplete="name"
                                    placeholder="Ada Lovelace"
                                    bind:value={fullName}
                                />
                            </div>

                            <div class="field">
                                <label for="email">Work email <span class="req">*</span></label>
                                <input
                                    id="email"
                                    class="control"
                                    type="email"
                                    name="email"
                                    autocomplete="email"
                                    placeholder="ada@company.com"
                                    bind:value={email}
                                />
                            </div>

                            <div class="field">
                                <label for="company">Company / organisation</label>
                                <input
                                    id="company"
                                    class="control"
                                    type="text"
                                    name="company"
                                    autocomplete="organization"
                                    placeholder="Analytical Engines Ltd."
                                    bind:value={company}
                                />
                            </div>

                            <div class="field">
                                <label for="website">Website</label>
                                <input
                                    id="website"
                                    class="control"
                                    type="url"
                                    name="website"
                                    autocomplete="url"
                                    placeholder="https://example.com"
                                    bind:value={website}
                                />
                                <span class="field-hint">Optional — helps us understand your product.</span>
                            </div>

                            <div class="field">
                                <label for="licenceType">What do you need? <span class="req">*</span></label>
                                <div class="select-wrap">
                                    <select id="licenceType" class="control" name="licenceType" bind:value={licenceType}>
                                        {#each licenceTypes as option}
                                            <option value={option.value}>{option.label}</option>
                                        {/each}
                                    </select>
                                    <svg class="chev-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
                                        <path d="m6 9 6 6 6-6" />
                                    </svg>
                                </div>
                            </div>

                            <div class="field">
                                <label for="deployment">How will it be deployed? <span class="req">*</span></label>
                                <div class="select-wrap">
                                    <select id="deployment" class="control" name="deployment" bind:value={deployment}>
                                        {#each deployments as option}
                                            <option value={option.value}>{option.label}</option>
                                        {/each}
                                    </select>
                                    <svg class="chev-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
                                        <path d="m6 9 6 6 6-6" />
                                    </svg>
                                </div>
                            </div>

                            <div class="field field-full">
                                <label for="message">Describe your project <span class="req">*</span></label>
                                <textarea
                                    id="message"
                                    class="control"
                                    name="message"
                                    rows="6"
                                    placeholder="What are you building, who uses it, and why does the AGPL not fit your model?"
                                    bind:value={message}
                                ></textarea>
                                <span class="field-hint">
                                    A few sentences is plenty. If you already know your expected
                                    distribution volume or team size, mention it here.
                                </span>
                            </div>

                            <div class="field field-full">
                                <div class="check-row">
                                    <input id="agree" type="checkbox" name="agree" bind:checked={agree} />
                                    <label for="agree">
                                        I'm happy for {Sitename}'s maintainers to contact me about this
                                        request. We'll only use your details to answer it.
                                    </label>
                                </div>
                            </div>
                        </div>

                        {#if error}
                            <div class="form-error" role="alert">
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
                                    <circle cx="12" cy="12" r="10" />
                                    <path d="M12 8v4M12 16h.01" />
                                </svg>
                                <span>{error}</span>
                            </div>
                        {/if}

                        <div class="form-foot">
                            <button class="btn btn-primary btn-lg" type="submit" disabled={sending}>
                                {sending ? "Sending…" : "Send request"}
                            </button>
                            <p class="note">
                                No payment now — we'll send a quote and agreement first.
                            </p>
                        </div>
                    </form>
                {/if}
            </div>

            <!-- ── Sidebar ───────────────────────────────────── -->
            <aside class="request-aside">
                <div class="aside-card">
                    <h3>What happens next</h3>
                    <ol class="steps">
                        <li>
                            <span class="step-num" aria-hidden="true">1</span>
                            <span>We read your request — usually within one to two business days.</span>
                        </li>
                        <li>
                            <span class="step-num" aria-hidden="true">2</span>
                            <span>We may ask a couple of clarifying questions about your deployment.</span>
                        </li>
                        <li>
                            <span class="step-num" aria-hidden="true">3</span>
                            <span>You receive a written quote and a draft licence agreement.</span>
                        </li>
                        <li>
                            <span class="step-num" aria-hidden="true">4</span>
                            <span>Once signed, you're covered — and free to ship.</span>
                        </li>
                    </ol>
                </div>

                <div class="aside-card aside-card-alt">
                    <h3>Not sure you need one?</h3>
                    <p>
                        Plenty of teams use {Sitename} under the AGPL without ever needing a
                        commercial licence — self-hosting internally, open-source projects and
                        public forks all stay free.
                    </p>
                    <p class="aside-note">
                        If your product is closed-source, embeds {Sitename} in something
                        proprietary, or runs a modified version as a hosted service, a
                        commercial licence is the right route.
                    </p>
                    <a class="btn btn-outline" href="/license#commercial">Compare the options</a>
                </div>

                <div class="aside-card">
                    <h3>Rather just email us?</h3>
                    <p class="aside-note">
                        Skip the form and write to us directly — a human reads every message.
                    </p>
                    <a class="btn btn-outline" href="mailto:{SALMON_MAINTAINERS_EMAIL}">
                        {SALMON_MAINTAINERS_EMAIL}
                    </a>
                </div>
            </aside>
        </div>
    </section>

    <!-- ── CTA ───────────────────────────────────────────── -->
    <section class="cta">
        <div class="container">
            <div class="cta-card">
                <h2>Still weighing it up?</h2>
                <p>
                    Read the plain-English breakdown of what the AGPL lets you do, or support
                    the project if you're sticking with the open-source licence.
                </p>
                <div class="cta-actions">
                    <a class="btn btn-lg btn-white" href="/license">Read the licence guide</a>
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
                    <li><a href="/license/request">Get a licence</a></li>
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

    .btn-primary:hover:not(:disabled) {
        transform: translateY(-1px);
        box-shadow: var(--shadow-sm), 0 14px 28px -12px rgba(234, 95, 77, 0.9);
    }

    .btn-primary:disabled {
        opacity: 0.65;
        cursor: progress;
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
    input:focus-visible,
    select:focus-visible,
    textarea:focus-visible {
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
        flex-wrap: wrap;
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

    /* ── Request layout ────────────────────────────────── */
    .request-inner {
        display: grid;
        grid-template-columns: 1.35fr 0.75fr;
        align-items: start;
        gap: 48px;
    }

    .request-aside {
        position: sticky;
        top: 92px;
    }

    /* ── Form card ─────────────────────────────────────── */
    .form-card {
        padding: 36px;
        border: 1px solid var(--border);
        border-radius: 20px;
        background: #fff;
        box-shadow: var(--shadow-sm);
    }

    .form-head {
        margin-bottom: 28px;
    }

    .form-head h2 {
        margin: 0;
        font-size: 1.375rem;
        font-weight: 700;
        letter-spacing: -0.025em;
    }

    .form-head p {
        margin: 8px 0 0;
        font-size: 0.875rem;
        color: var(--muted);
    }

    .req {
        color: var(--salmon-600);
        margin-left: 1px;
    }

    .form-grid {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 20px;
    }

    .field {
        display: grid;
        gap: 8px;
        align-content: start;
    }

    .field-full {
        grid-column: 1 / -1;
    }

    .field label {
        font-size: 0.875rem;
        font-weight: 600;
        letter-spacing: -0.005em;
        color: var(--ink);
    }

    .field-hint {
        font-size: 0.8125rem;
        line-height: 1.55;
        color: var(--muted);
    }

    .control {
        width: 100%;
        padding: 11px 14px;
        border: 1px solid var(--border);
        border-radius: 10px;
        background: #fff;
        font: inherit;
        font-size: 0.9375rem;
        color: var(--ink);
        transition: border-color 0.15s ease, box-shadow 0.15s ease;
    }

    .control::placeholder {
        color: #94a3b8;
    }

    .control:hover {
        border-color: #d7dce5;
    }

    .control:focus {
        outline: none;
        border-color: var(--salmon-400);
        box-shadow: 0 0 0 4px rgba(250, 128, 114, 0.18);
    }

    textarea.control {
        min-height: 148px;
        resize: vertical;
        line-height: 1.65;
    }

    .select-wrap {
        position: relative;
    }

    .select-wrap .control {
        appearance: none;
        -webkit-appearance: none;
        padding-right: 42px;
        cursor: pointer;
    }

    .chev-icon {
        position: absolute;
        top: 50%;
        right: 14px;
        width: 18px;
        height: 18px;
        transform: translateY(-50%);
        color: var(--muted);
        pointer-events: none;
    }

    /* ── Checkbox row ──────────────────────────────────── */
    .check-row {
        display: flex;
        align-items: flex-start;
        gap: 12px;
        padding: 15px 16px;
        border: 1px solid var(--border);
        border-radius: 12px;
        background: var(--surface-2);
    }

    .check-row input {
        width: 18px;
        height: 18px;
        margin: 1px 0 0;
        flex: none;
        accent-color: var(--salmon-600);
        cursor: pointer;
    }

    .check-row label {
        font-size: 0.875rem;
        font-weight: 500;
        line-height: 1.6;
        color: var(--ink-2);
        cursor: pointer;
    }

    /* ── Error ─────────────────────────────────────────── */
    .form-error {
        display: flex;
        align-items: flex-start;
        gap: 10px;
        margin-top: 22px;
        padding: 14px 16px;
        border: 1px solid #fecaca;
        border-radius: 12px;
        background: #fef2f2;
        color: #b91c1c;
        font-size: 0.875rem;
        line-height: 1.6;
    }

    .form-error svg {
        width: 18px;
        height: 18px;
        flex: none;
        margin-top: 2px;
    }

    /* ── Form footer ───────────────────────────────────── */
    .form-foot {
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        gap: 14px 20px;
        margin-top: 28px;
        padding-top: 24px;
        border-top: 1px solid var(--border);
    }

    .form-foot .note {
        margin: 0;
        font-size: 0.8125rem;
        line-height: 1.6;
        color: var(--muted);
    }

    /* ── Success state ─────────────────────────────────── */
    .success-card {
        text-align: center;
        padding: 56px 36px;
    }

    .success-mark {
        display: grid;
        place-items: center;
        width: 56px;
        height: 56px;
        margin: 0 auto;
        border-radius: 16px;
        background: var(--green-50);
        border: 1px solid #bbf7d0;
        color: var(--green-600);
    }

    .success-mark svg {
        width: 26px;
        height: 26px;
    }

    .success-card h2 {
        margin: 22px 0 0;
        font-size: 1.5rem;
        font-weight: 700;
        letter-spacing: -0.025em;
    }

    .success-card p {
        max-width: 34rem;
        margin: 14px auto 0;
        font-size: 1rem;
        line-height: 1.7;
        color: var(--muted);
    }

    .success-card p strong {
        color: var(--ink-2);
        font-weight: 600;
    }

    .success-note a {
        color: var(--salmon-600);
        font-weight: 600;
        text-decoration: none;
        word-break: break-word;
    }

    .success-note a:hover {
        text-decoration: underline;
    }

    .success-actions {
        justify-content: center;
        margin-top: 32px;
    }

    /* ── Aside cards ───────────────────────────────────── */
    .aside-card {
        padding: 26px;
        border: 1px solid var(--border);
        border-radius: 18px;
        background: #fff;
        box-shadow: var(--shadow-sm);
    }

    .aside-card + .aside-card {
        margin-top: 20px;
    }

    .aside-card h3 {
        margin: 0 0 18px;
        font-size: 1rem;
        font-weight: 700;
        letter-spacing: -0.015em;
    }

    .steps {
        display: grid;
        gap: 18px;
        margin: 0;
        padding: 0;
        list-style: none;
    }

    .steps li {
        display: flex;
        gap: 14px;
        font-size: 0.9375rem;
        line-height: 1.6;
        color: var(--ink-2);
    }

    .step-num {
        display: grid;
        place-items: center;
        width: 26px;
        height: 26px;
        flex: none;
        border-radius: 8px;
        background: var(--salmon-50);
        border: 1px solid var(--salmon-100);
        color: var(--salmon-700);
        font-size: 0.75rem;
        font-weight: 700;
    }

    .aside-card-alt {
        background: var(--surface-2);
    }

    .aside-card p {
        margin: 0;
        font-size: 0.9375rem;
        line-height: 1.7;
        color: var(--ink-2);
    }

    .aside-note {
        margin-top: 12px !important;
        font-size: 0.875rem !important;
        color: var(--muted) !important;
    }

    .aside-card .btn {
        margin-top: 18px;
        width: 100%;
        white-space: normal;
        text-align: center;
        word-break: break-word;
    }

    /* ── CTA ───────────────────────────────────────────── */
    .cta {
        padding: 0 0 88px;
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
        .request-inner {
            grid-template-columns: 1fr;
            gap: 40px;
        }

        .request-aside {
            position: static;
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
        .section {
            padding: 64px 0;
        }

        .page-hero {
            padding: 48px 0 56px;
        }

        .form-card {
            padding: 26px;
        }

        .form-grid {
            grid-template-columns: 1fr;
        }

        .success-card {
            padding: 44px 24px;
        }

        .cta {
            padding-bottom: 64px;
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
        .control {
            transition: none;
        }
    }
</style>
