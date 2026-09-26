<script lang="ts">
    import { onMount } from "svelte";
    import { Sitename } from "$lib";
    import { SALMON_LEGAL_EMAIL } from "$lib/data";

    const LAST_UPDATED = "15 September 2026";
    const EFFECTIVE_DATE = "15 September 2026";
    const CONTACT_EMAIL = SALMON_LEGAL_EMAIL;
    const GOVERNING_LAW = "the Republic of Kenya";

    const sections = [
        { id: "acceptance", n: "1", title: "Acceptance of these terms" },
        { id: "about", n: "2", title: "About Salmon" },
        { id: "software-licence", n: "3", title: "Licence to use the software" },
        { id: "accounts", n: "4", title: "Accounts" },
        { id: "acceptable-use", n: "5", title: "Acceptable use" },
        { id: "your-content", n: "6", title: "Your content" },
        { id: "intellectual-property", n: "7", title: "Intellectual property and branding" },
        { id: "third-parties", n: "8", title: "Third-party services" },
        { id: "self-hosted", n: "9", title: "Self-hosted instances" },
        { id: "availability", n: "10", title: "Availability and changes" },
        { id: "warranties", n: "11", title: "Disclaimer of warranties" },
        { id: "liability", n: "12", title: "Limitation of liability" },
        { id: "indemnity", n: "13", title: "Indemnification" },
        { id: "termination", n: "14", title: "Termination" },
        { id: "changes", n: "15", title: "Changes to these terms" },
        { id: "governing-law", n: "16", title: "Governing law and disputes" },
        { id: "general", n: "17", title: "General" },
        { id: "contact", n: "18", title: "Contact" }
    ];

    let active = "acceptance";

    onMount(() => {
        const elements = sections
            .map((s) => document.getElementById(s.id))
            .filter(Boolean);

        if (!elements.length || typeof IntersectionObserver === "undefined") {
            return;
        }

        const observer = new IntersectionObserver(
            (entries) => {
                for (const entry of entries) {
                    if (entry.isIntersecting) {
                        active = entry.target.id;
                    }
                }
            },
            { rootMargin: "-96px 0px -70% 0px", threshold: 0 }
        );

        elements.forEach((el) => observer.observe(el));

        return () => observer.disconnect();
    });
</script>

<svelte:head>
    <title>Terms of use — {Sitename}</title>
    <meta
        name="description"
        content="The terms governing your use of the Salmon website and any hosted services. The Salmon software itself is licensed separately under the GNU AGPL-3.0."
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
            <a href="/terms-of-use" aria-current="page">Terms of use</a>
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
                <span aria-current="page">Terms of use</span>
            </nav>

            <span class="badge"><span class="dot"></span> Legal</span>

            <h1>Terms of use</h1>

            <p class="lede">
                These terms govern your use of the {Sitename} website and any hosted
                services we operate. The {Sitename} software itself is licensed separately
                under the <a href="/license">GNU AGPL-3.0</a>.
            </p>

            <div class="hero-meta">
                <div class="meta-item">
                    <span class="meta-label">Last updated</span>
                    <span class="meta-value">{LAST_UPDATED}</span>
                </div>
                <div class="meta-item">
                    <span class="meta-label">Effective from</span>
                    <span class="meta-value">{EFFECTIVE_DATE}</span>
                </div>
                <div class="meta-item">
                    <span class="meta-label">Applies to</span>
                    <span class="meta-value">This website and hosted services</span>
                </div>
            </div>

            <p class="hero-note">
                We've written these terms in plain language wherever possible so they're
                actually readable. If anything is unclear, please
                <a href={`mailto:${CONTACT_EMAIL}`}>get in touch</a> — we'd rather answer a
                question than have you guess.
            </p>
        </div>
    </section>

    <!-- ── Body ──────────────────────────────────────────── -->
    <section class="section">
        <div class="container legal-layout">
            <!-- Table of contents -->
            <aside class="toc" aria-label="On this page">
                <div class="toc-inner">
                    <h2 class="toc-title">On this page</h2>
                    <ol class="toc-list">
                        {#each sections as s}
                            <li>
                                <a
                                    href={`#${s.id}`}
                                    class:active={active === s.id}
                                    aria-current={active === s.id ? "true" : undefined}
                                >
                                    <span class="toc-n">{s.n}</span>
                                    <span>{s.title}</span>
                                </a>
                            </li>
                        {/each}
                    </ol>

                    <a class="toc-top" href="#top">Back to top</a>
                </div>
            </aside>

            <!-- Content -->
            <div class="legal" id="top">
                <p class="lead">
                    Please read these Terms of Use ("Terms") carefully before using the
                    {Sitename} website or any hosted service we operate. By accessing or
                    using them, you agree to be bound by these Terms.
                </p>

                <!-- 1 -->
                <section id="acceptance" class="legal-section">
                    <h2><span class="n">1</span> Acceptance of these terms</h2>
                    <p>
                        By accessing this website, downloading material from it, or using any
                        hosted service we operate, you confirm that you have read, understood
                        and agree to be bound by these Terms. If you do not agree, please do
                        not use the website or those services.
                    </p>
                    <p>
                        If you are using them on behalf of an organisation, you represent
                        that you have the authority to bind that organisation, and "you"
                        refers to both you and that organisation.
                    </p>
                    <p>
                        These Terms apply only to the website and hosted services. They do
                        <strong>not</strong> apply to the {Sitename} software itself, which is
                        governed by its own licence — see section 3.
                    </p>
                </section>

                <!-- 2 -->
                <section id="about" class="legal-section">
                    <h2><span class="n">2</span> About {Sitename}</h2>
                    <p>
                        {Sitename} is an open-source software project created and maintained
                        by Joseph Wangai Mwaniki ("the maintainer", "we", "us" or "our").
                    </p>
                    <p>It is important to distinguish between two things:</p>
                    <ul>
                        <li>
                            <strong>The software</strong> — the source code, builds and
                            releases of {Sitename}, which you can download, run and modify.
                            Your rights to it come from the GNU AGPL-3.0.
                        </li>
                        <li>
                            <strong>This website and hosted services</strong> — the site you
                            are reading, our documentation, downloads page, and any hosted
                            instance we operate. These are governed by these Terms.
                        </li>
                    </ul>
                    <p>
                        Nothing in these Terms is intended to restrict, replace or take away
                        any right granted to you under the AGPL-3.0.
                    </p>
                </section>

                <!-- 3 -->
                <section id="software-licence" class="legal-section">
                    <h2><span class="n">3</span> Licence to use the software</h2>
                    <p>
                        The {Sitename} software is free and open-source software released
                        under the <strong>GNU Affero General Public License, version 3.0</strong>
                        (AGPL-3.0). It is <em>not</em> licensed under these Terms.
                    </p>
                    <p>
                        Your rights to use, study, modify and redistribute the software are
                        set out in full in that licence, which you should read at
                        <a href="/license">our licence page</a>. In summary, the AGPL-3.0
                        allows you to use the software for any purpose, provided you keep the
                        licence and copyright notices intact and release any modified version
                        under the same terms — including when you make it available to others
                        over a network.
                    </p>

                    <div class="callout callout-info">
                        <strong>If these Terms and the AGPL-3.0 conflict on anything relating
                        to the software, the AGPL-3.0 prevails.</strong>
                        These Terms cannot and do not impose additional restrictions on the
                        rights that licence grants you.
                    </div>
                </section>

                <!-- 4 -->
                <section id="accounts" class="legal-section">
                    <h2><span class="n">4</span> Accounts</h2>
                    <p>
                        Some hosted services may allow or require you to create an account. If
                        you do, you agree to:
                    </p>
                    <ul>
                        <li>provide accurate and current information;</li>
                        <li>keep your credentials secure and confidential;</li>
                        <li>
                            accept responsibility for all activity that occurs under your
                            account; and
                        </li>
                        <li>
                            notify us promptly if you suspect unauthorised access or any
                            security breach.
                        </li>
                    </ul>
                    <p>
                        You must be at least the age of majority in your jurisdiction to
                        create an account. We may suspend or close accounts that breach these
                        Terms, that we reasonably believe are being used unlawfully, or that
                        place the service or other users at risk.
                    </p>
                </section>

                <!-- 5 -->
                <section id="acceptable-use" class="legal-section">
                    <h2><span class="n">5</span> Acceptable use</h2>
                    <p>
                        You may use this website and any hosted services for lawful purposes
                        only. In particular, you agree <strong>not</strong> to:
                    </p>
                    <ul>
                        <li>break any applicable law or regulation;</li>
                        <li>
                            infringe the intellectual property, privacy or other rights of
                            anyone else;
                        </li>
                        <li>
                            upload, transmit or distribute malware, or anything designed to
                            disrupt or damage systems;
                        </li>
                        <li>
                            attempt to gain unauthorised access to any part of the service,
                            other accounts, or the systems or networks connected to it;
                        </li>
                        <li>
                            interfere with, overload, or launch automated attacks against the
                            service, including denial-of-service attempts and excessive
                            scraping;
                        </li>
                        <li>
                            harass, threaten, impersonate or abuse other users or the
                            maintainer;
                        </li>
                        <li>
                            misrepresent your identity or your affiliation with the project;
                            or
                        </li>
                        <li>use the service to send spam or unsolicited advertising.</li>
                    </ul>

                    <div class="callout callout-warn">
                        <strong>The AGPL grants you rights to the software — not to our
                        infrastructure.</strong>
                        You are free to fork, modify and self-host {Sitename}. That freedom
                        does not extend to abusing, overloading or attacking the services we
                        operate.
                    </div>
                </section>

                <!-- 6 -->
                <section id="your-content" class="legal-section">
                    <h2><span class="n">6</span> Your content</h2>
                    <p>
                        You retain all ownership rights in any content you create, upload or
                        submit through the service. We do not claim ownership over it.
                    </p>
                    <p>
                        To operate the service, you grant us a limited, non-exclusive,
                        worldwide, royalty-free licence to host, store, process, transmit and
                        back up your content — solely for the purpose of running and
                        maintaining the service, and for no other purpose. This licence ends
                        when you delete your content or your account, except where we are
                        required to retain it by law or where it persists in routine backups
                        for a reasonable period.
                    </p>
                    <p>
                        You are responsible for your content and confirm that you have all
                        necessary rights to submit it, and that it does not infringe anyone
                        else's rights or violate any law.
                    </p>
                </section>

                <!-- 7 -->
                <section id="intellectual-property" class="legal-section">
                    <h2><span class="n">7</span> Intellectual property and branding</h2>
                    <p>
                        The {Sitename} name, logo, visual identity, website design, and the
                        written documentation on this site are the property of Joseph Wangai
                        Mwaniki and are protected by copyright and other intellectual property
                        laws. All rights not expressly granted are reserved.
                    </p>
                    <p>
                        <strong>The AGPL-3.0 covers the software source code.</strong> It does
                        not grant you rights to use the project's name, logo or branding. In
                        particular, you may not:
                    </p>
                    <ul>
                        <li>
                            use the {Sitename} name or logo in a way that suggests your fork,
                            product or service is official, endorsed by, or affiliated with
                            the project; or
                        </li>
                        <li>
                            use our branding in a misleading way or in connection with
                            unlawful or harmful activity.
                        </li>
                    </ul>
                    <p>
                        If you distribute a modified version, please give it a different name
                        and its own visual identity, so users are not confused about its
                        origin. You are, of course, welcome to state truthfully that your
                        project is based on {Sitename}.
                    </p>
                </section>

                <!-- 8 -->
                <section id="third-parties" class="legal-section">
                    <h2><span class="n">8</span> Third-party services</h2>
                    <p>
                        The website may link to, integrate with, or rely on third-party
                        services — for example code hosting, payment processors, package
                        registries or analytics-free hosting providers. We do not control
                        those services and are not responsible for their content, policies or
                        practices.
                    </p>
                    <p>
                        Your use of any third-party service is governed by that provider's own
                        terms and privacy policy. We encourage you to review them. Links from
                        our site do not imply endorsement.
                    </p>
                </section>

                <!-- 9 -->
                <section id="self-hosted" class="legal-section">
                    <h2><span class="n">9</span> Self-hosted instances</h2>
                    <p>
                        One of the main benefits of open-source software is that you can run
                        it yourself. If you deploy your own instance of {Sitename}, you are
                        its operator and you are solely responsible for it.
                    </p>
                    <p>That responsibility includes, without limitation:</p>
                    <ul>
                        <li>compliance with all laws applicable to your instance;</li>
                        <li>
                            data protection and privacy obligations in relation to your users;
                        </li>
                        <li>security, availability, backups and disaster recovery;</li>
                        <li>user support and moderation; and</li>
                        <li>any content hosted or processed on your instance.</li>
                    </ul>
                    <p>
                        These Terms govern this website and the hosted services we operate.
                        They do not govern your instance, and we accept no responsibility or
                        liability for it.
                    </p>
                </section>

                <!-- 10 -->
                <section id="availability" class="legal-section">
                    <h2><span class="n">10</span> Availability and changes</h2>
                    <p>
                        We aim to keep the website and any hosted services running smoothly,
                        but we provide them on a best-effort basis. We do not guarantee any
                        particular level of uptime, availability or performance.
                    </p>
                    <p>
                        We may add, change, suspend or discontinue any part of the website or
                        hosted services at any time, with or without notice. This does not
                        affect the software itself: the source code remains available to you
                        under the AGPL-3.0, and you are always free to self-host.
                    </p>
                </section>

                <!-- 11 -->
                <section id="warranties" class="legal-section">
                    <h2><span class="n">11</span> Disclaimer of warranties</h2>
                    <p class="caps">
                        The website and any hosted services are provided "as is" and "as
                        available", without warranties of any kind, whether express, implied
                        or statutory.
                    </p>
                    <p class="caps">
                        To the fullest extent permitted by applicable law, we disclaim all
                        implied warranties, including any implied warranties of
                        merchantability, fitness for a particular purpose, title and
                        non-infringement.
                    </p>
                    <p class="caps">
                        We do not warrant that the website or hosted services will be
                        uninterrupted, timely, secure, error-free, or free of viruses or other
                        harmful components, or that any defects will be corrected.
                    </p>
                    <p>
                        You use the website and hosted services at your own discretion and
                        risk, and you are solely responsible for any damage to your systems or
                        loss of data that results from doing so.
                    </p>
                </section>

                <!-- 12 -->
                <section id="liability" class="legal-section">
                    <h2><span class="n">12</span> Limitation of liability</h2>
                    <p class="caps">
                        To the maximum extent permitted by applicable law, the maintainer
                        shall not be liable for any indirect, incidental, special,
                        consequential, exemplary or punitive damages, or for any loss of
                        profits, revenue, data, goodwill or business opportunity, arising out
                        of or in connection with your use of, or inability to use, the website
                        or any hosted service — whether based in contract, tort, negligence,
                        strict liability or any other theory, and even if advised of the
                        possibility of such damages.
                    </p>
                    <p class="caps">
                        To the maximum extent permitted by applicable law, the total aggregate
                        liability of the maintainer arising out of or relating to these Terms
                        or the website and hosted services shall not exceed the greater of
                        (a) the total amount you have paid us in the twelve (12) months
                        immediately preceding the event giving rise to the claim, or
                        (b) fifty United States dollars (USD 50).
                    </p>
                    <p>
                        Some jurisdictions do not allow the exclusion or limitation of certain
                        warranties or liabilities. Where that is the case, the exclusions and
                        limitations above apply to the fullest extent permitted by law, and
                        nothing in these Terms excludes or limits any liability that cannot
                        lawfully be excluded or limited.
                    </p>
                </section>

                <!-- 13 -->
                <section id="indemnity" class="legal-section">
                    <h2><span class="n">13</span> Indemnification</h2>
                    <p>
                        You agree to indemnify, defend and hold harmless the maintainer from
                        and against any claims, liabilities, damages, losses, costs and
                        expenses — including reasonable legal fees — arising out of or in any
                        way connected with:
                    </p>
                    <ul>
                        <li>your access to or use of the website or hosted services;</li>
                        <li>your content;</li>
                        <li>your operation of a self-hosted instance; or</li>
                        <li>your breach of these Terms or of any law or third-party right.</li>
                    </ul>
                    <p>
                        We reserve the right to assume the exclusive defence and control of
                        any matter otherwise subject to indemnification by you, in which case
                        you agree to cooperate with us in asserting any available defences.
                    </p>
                </section>

                <!-- 14 -->
                <section id="termination" class="legal-section">
                    <h2><span class="n">14</span> Termination</h2>
                    <p>
                        You may stop using the website and hosted services at any time. We may
                        suspend or terminate your access to them at any time, with or without
                        notice, if we reasonably believe you have breached these Terms or if
                        we discontinue the relevant service.
                    </p>
                    <p>
                        Termination does not affect your rights to the software under the
                        AGPL-3.0. Because the software is open source, it cannot be taken away
                        from you — a copy you have lawfully obtained remains yours to use under
                        that licence.
                    </p>
                    <p>
                        Any provision of these Terms that by its nature should survive
                        termination will do so, including sections 7 (intellectual property),
                        11 (disclaimer of warranties), 12 (limitation of liability),
                        13 (indemnification) and 16 (governing law).
                    </p>
                </section>

                <!-- 15 -->
                <section id="changes" class="legal-section">
                    <h2><span class="n">15</span> Changes to these terms</h2>
                    <p>
                        We may update these Terms from time to time — for example to reflect
                        changes to the service or to the law. When we do, we will update the
                        "Last updated" date at the top of this page, and where the changes are
                        material we will take reasonable steps to bring them to your
                        attention, such as a notice on the site.
                    </p>
                    <p>
                        Your continued use of the website or hosted services after any change
                        takes effect constitutes acceptance of the revised Terms. If you do
                        not agree with the changes, please stop using those services.
                    </p>
                    <p>
                        Changes to these Terms do not alter the licence terms of the software,
                        which can only be changed in accordance with the AGPL-3.0 itself.
                    </p>
                </section>

                <!-- 16 -->
                <section id="governing-law" class="legal-section">
                    <h2><span class="n">16</span> Governing law and disputes</h2>
                    <p>
                        These Terms and any dispute or claim arising out of or in connection
                        with them are governed by the laws of {GOVERNING_LAW}, without regard
                        to its conflict-of-law rules.
                    </p>
                    <p>
                        The courts of {GOVERNING_LAW} shall have exclusive jurisdiction to
                        settle any dispute arising out of or in connection with these Terms,
                        subject to any right you may have to bring proceedings in your local
                        courts where mandatory consumer protection law applies.
                    </p>
                    <p>
                        Before starting formal proceedings, we ask that you contact us first at
                        <a href={`mailto:${CONTACT_EMAIL}`}>{CONTACT_EMAIL}</a>. Most issues can
                        be resolved quickly and informally, and we would genuinely prefer that.
                    </p>
                </section>

                <!-- 17 -->
                <section id="general" class="legal-section">
                    <h2><span class="n">17</span> General</h2>
                    <p>
                        <strong>Severability.</strong> If any provision of these Terms is found
                        to be invalid, unlawful or unenforceable, that provision will be
                        limited or removed to the minimum extent necessary, and the remaining
                        provisions will continue in full force and effect.
                    </p>
                    <p>
                        <strong>No waiver.</strong> Our failure to enforce any right or
                        provision of these Terms is not a waiver of that right or provision.
                    </p>
                    <p>
                        <strong>Assignment.</strong> You may not assign or transfer your rights
                        under these Terms without our prior written consent. We may assign our
                        rights without restriction.
                    </p>
                    <p>
                        <strong>Entire agreement.</strong> These Terms, together with the
                        AGPL-3.0 and any other policies we publish on this site, constitute
                        the entire agreement between you and us regarding the website and
                        hosted services, and supersede any prior agreements.
                    </p>
                </section>

                <!-- 18 -->
                <section id="contact" class="legal-section">
                    <h2><span class="n">18</span> Contact</h2>
                    <p>
                        If you have any questions about these Terms, want to report a problem,
                        or need to get in touch about a legal matter, please contact:
                    </p>

                    <div class="contact-card">
                        <div class="contact-row">
                            <span class="contact-label">Maintainer</span>
                            <span class="contact-value">Joseph Wangai Mwaniki</span>
                        </div>
                        <div class="contact-row">
                            <span class="contact-label">Email</span>
                            <span class="contact-value">
                                <a href={`mailto:${CONTACT_EMAIL}`}>{CONTACT_EMAIL}</a>
                            </span>
                        </div>
                        <div class="contact-row">
                            <span class="contact-label">Repository</span>
                            <span class="contact-value">
                                <a href="https://github.com/joenikkai/salmon" rel="noopener">
                                    github.com/joenikkai/salmon
                                </a>
                            </span>
                        </div>
                    </div>
                </section>

                <div class="legal-footnote">
                    <p>
                        This document is a template provided for convenience and does not
                        constitute legal advice. Laws vary by jurisdiction, and the enforceability
                        of certain clauses — particularly the limitation of liability — depends on
                        local law. If you rely on these Terms commercially, please have them
                        reviewed by a qualified lawyer in your jurisdiction.
                    </p>
                </div>
            </div>
        </div>
    </section>

    <!-- ── CTA ───────────────────────────────────────────── -->
    <section class="cta">
        <div class="container">
            <div class="cta-card">
                <h2>Questions about these terms?</h2>
                <p>
                    We'd rather answer a question than have you guess. Reach out any time, and
                    take a look at the licence page if you want to know what the AGPL lets you
                    do.
                </p>
                <div class="cta-actions">
                    <a class="btn btn-lg btn-white" href={`mailto:${CONTACT_EMAIL}`}>Email us</a>
                    <a class="btn btn-lg btn-oncolor" href="/license">Read the licence</a>
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
                    <li><a href="/license">Licence</a></li>
                    <li><a href="/terms-of-use" aria-current="page">Terms of use</a></li>
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

    :global(html) {
        scroll-behavior: smooth;
        scroll-padding-top: 90px;
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

    .btn:focus-visible,
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

    .lede a {
        color: var(--salmon-600);
        font-weight: 600;
        text-decoration: none;
    }

    .lede a:hover {
        text-decoration: underline;
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

    .hero-note {
        margin: 26px 0 0;
        padding: 16px 20px;
        max-width: 46rem;
        border: 1px solid var(--border);
        border-left: 4px solid var(--salmon-400);
        border-radius: 10px;
        background: var(--surface-2);
        font-size: 0.9375rem;
        line-height: 1.65;
        color: var(--ink-2);
    }

    .hero-note a {
        color: var(--salmon-600);
        font-weight: 600;
        text-decoration: none;
    }

    .hero-note a:hover {
        text-decoration: underline;
    }

    /* ── Layout ────────────────────────────────────────── */
    .section {
        padding: 72px 0 88px;
    }

    .legal-layout {
        display: grid;
        grid-template-columns: 260px 1fr;
        gap: 64px;
        align-items: start;
    }

    /* ── TOC ───────────────────────────────────────────── */
    .toc {
        position: sticky;
        top: 92px;
    }

    .toc-inner {
        padding: 22px;
        border: 1px solid var(--border);
        border-radius: 16px;
        background: #fff;
        box-shadow: var(--shadow-sm);
        max-height: calc(100vh - 120px);
        overflow-y: auto;
    }

    .toc-title {
        margin: 0 0 14px;
        font-size: 0.75rem;
        font-weight: 700;
        letter-spacing: 0.12em;
        text-transform: uppercase;
        color: var(--muted);
    }

    .toc-list {
        display: grid;
        gap: 2px;
        margin: 0;
        padding: 0;
        list-style: none;
    }

    .toc-list a {
        display: flex;
        gap: 10px;
        padding: 7px 10px;
        border-radius: 8px;
        font-size: 0.8125rem;
        line-height: 1.4;
        color: var(--ink-2);
        text-decoration: none;
        transition: background-color 0.15s ease, color 0.15s ease;
    }

    .toc-list a:hover {
        background: var(--salmon-50);
        color: var(--salmon-700);
    }

    .toc-list a.active {
        background: var(--salmon-50);
        color: var(--salmon-700);
        font-weight: 600;
    }

    .toc-n {
        flex: none;
        width: 1.1em;
        font-variant-numeric: tabular-nums;
        color: var(--muted);
        font-weight: 600;
    }

    .toc-list a.active .toc-n {
        color: var(--salmon-500);
    }

    .toc-top {
        display: inline-block;
        margin-top: 16px;
        padding-top: 14px;
        border-top: 1px solid var(--border);
        width: 100%;
        font-size: 0.8125rem;
        font-weight: 600;
        color: var(--salmon-600);
        text-decoration: none;
    }

    .toc-top:hover {
        text-decoration: underline;
    }

    /* ── Legal content ─────────────────────────────────── */
    .legal {
        max-width: 720px;
        scroll-margin-top: 90px;
    }

    .legal .lead {
        margin: 0 0 8px;
        padding-bottom: 32px;
        border-bottom: 1px solid var(--border);
        font-size: 1.0625rem;
        line-height: 1.75;
        color: var(--ink-2);
    }

    .legal-section {
        padding-top: 40px;
        scroll-margin-top: 90px;
    }

    .legal-section h2 {
        display: flex;
        align-items: baseline;
        gap: 12px;
        margin: 0 0 18px;
        font-size: 1.25rem;
        font-weight: 700;
        letter-spacing: -0.02em;
        line-height: 1.3;
    }

    .legal-section h2 .n {
        flex: none;
        display: grid;
        place-items: center;
        width: 26px;
        height: 26px;
        border-radius: 8px;
        background: var(--salmon-50);
        border: 1px solid var(--salmon-100);
        color: var(--salmon-600);
        font-size: 0.75rem;
        font-weight: 800;
        font-variant-numeric: tabular-nums;
        transform: translateY(1px);
    }

    .legal-section p {
        margin: 0 0 16px;
        font-size: 0.9375rem;
        line-height: 1.8;
        color: var(--ink-2);
    }

    .legal-section p:last-child {
        margin-bottom: 0;
    }

    .legal-section strong {
        color: var(--ink);
        font-weight: 600;
    }

    .legal-section a {
        color: var(--salmon-600);
        font-weight: 600;
        text-decoration: none;
    }

    .legal-section a:hover {
        text-decoration: underline;
    }

    .legal-section ul {
        display: grid;
        gap: 10px;
        margin: 0 0 16px;
        padding: 0;
        list-style: none;
    }

    .legal-section li {
        position: relative;
        padding-left: 22px;
        font-size: 0.9375rem;
        line-height: 1.75;
        color: var(--ink-2);
    }

    .legal-section li::before {
        content: "";
        position: absolute;
        left: 4px;
        top: 0.68em;
        width: 6px;
        height: 6px;
        border-radius: 50%;
        background: var(--salmon-300);
    }

    .caps {
        text-transform: none;
        letter-spacing: 0.005em;
        font-size: 0.875rem !important;
        color: var(--ink-2);
    }

    /* ── Callouts ──────────────────────────────────────── */
    .callout {
        margin: 22px 0;
        padding: 18px 22px;
        border-radius: 12px;
        font-size: 0.9375rem;
        line-height: 1.7;
        color: var(--ink-2);
    }

    .callout strong {
        display: block;
        margin-bottom: 4px;
        color: var(--ink);
    }

    .callout-info {
        border: 1px solid var(--salmon-200);
        border-left: 4px solid var(--salmon-500);
        background: var(--salmon-50);
    }

    .callout-info strong {
        color: var(--salmon-700);
    }

    .callout-warn {
        border: 1px solid var(--border);
        border-left: 4px solid var(--salmon-400);
        background: var(--surface-2);
    }

    /* ── Contact card ──────────────────────────────────── */
    .contact-card {
        margin-top: 22px;
        padding: 8px 24px;
        border: 1px solid var(--border);
        border-radius: 14px;
        background: var(--surface-2);
    }

    .contact-row {
        display: flex;
        flex-wrap: wrap;
        gap: 6px 20px;
        padding: 14px 0;
        border-bottom: 1px solid var(--border);
        font-size: 0.9375rem;
    }

    .contact-row:last-child {
        border-bottom: 0;
    }

    .contact-label {
        flex: none;
        width: 110px;
        font-weight: 600;
        color: var(--muted);
    }

    .contact-value {
        color: var(--ink);
        font-weight: 500;
    }

    .contact-value a {
        color: var(--salmon-600);
        text-decoration: none;
    }

    .contact-value a:hover {
        text-decoration: underline;
    }

    /* ── Footnote ──────────────────────────────────────── */
    .legal-footnote {
        margin-top: 48px;
        padding: 22px 24px;
        border: 1px dashed var(--border);
        border-radius: 14px;
        background: var(--surface-2);
    }

    .legal-footnote p {
        margin: 0;
        font-size: 0.8125rem;
        line-height: 1.7;
        color: var(--muted);
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
        .legal-layout {
            grid-template-columns: 1fr;
            gap: 40px;
        }

        .toc {
            position: static;
        }

        .toc-inner {
            max-height: none;
        }

        .toc-list {
            grid-template-columns: repeat(2, 1fr);
        }

        .nav-desktop {
            display: none;
        }

        .footer-grid {
            grid-template-columns: 1fr 1fr;
            gap: 32px;
        }
    }

    @media (max-width: 760px) {
        .section {
            padding: 56px 0 64px;
        }

        .page-hero {
            padding: 48px 0 56px;
        }

        .toc-list {
            grid-template-columns: 1fr;
        }

        .contact-label {
            width: 100%;
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

    @media (prefers-reduced-motion: reduce) {
        :global(html) {
            scroll-behavior: auto;
        }

        .btn {
            transition: none;
        }
    }
</style>
