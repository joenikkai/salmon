# Salmon

> A super awesome ledger.

<img src="https://upload.wikimedia.org/wikipedia/commons/3/39/Salmo_salar.jpg" alt="A picture of a salmon from Wikipedia">

## Why?

I noticed my dad struggling to manage his inventory. Many local businesses face the same problem. Salmon is my attempt to help.

## Goals

Salmon should help a business efficiently manage stock and boost profits by providing:

### Kenya Revenue Authority (KRA) Compliance
- Automatic eTIMS invoices
- VAT exemption handling
- Tax returns
- ETR receipts

### Money Flow
- M-Pesa C2B, STK, Till, and Paybill integration
- Cash, card, and credit sales
- Daily Z report (closing report for the cashier)
- Debtors and creditors tracking

### Core Features
- **Plug and play binary** — runs on Windows, Linux, and Mac
- **Fast** — built for speed
- **Offline-first** — hardened against power and network outages (assuming the business uses a laptop with a power generator or UPS)
- **Fast barcode scanning and search** — returns/refunds, discounts, customer display
- **Hardware support** — receipt printer, scanner, cash drawer
- **Inventory**
  - Expiry tracking
  - Real stock levels
  - Multi-store transfers
  - Stock take
- **Purchasing**
  - LPO (Local Purchase Orders)
  - Supplier bills
- **Role-Based Access Control**

> Some of these terms are jargon. I will document everything in a down-to-earth manner.

Backups up to 5 years are required by [Kenyan Law](https://www.kra.go.ke/helping-tax-payers/facts-about-kra/category/9).

We don't bundle cloud backup because we can't run that infrastructure for free. Local backups and export tools will be provided. Businesses are responsible for their own offsite retention per KRA requirements.

## The Stack

- **SvelteKit** — for the website
- **Wails + Svelte** — for the desktop app

I previously considered adding a mobile version. Handling product records on a phone keyboard was a hassle for me. But if Salmon's users want a mobile version later, it will be added in version 3.

## The Scope

> The initial scope is a point of sale terminal. As we go on, we might cover other areas for the business.

**WARNING:** Avoid using a single profile for all business needs.

Salmon is elastic at installation time: you can define multiple profiles. Each profile must be scoped to a single concern. Point of Sale, for instance, should never be combined with Enterprise Resource Management in the same profile.

<!--
## Install
will be added later

## Run
will be added later
-->