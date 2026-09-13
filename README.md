# salmon

> super awesome ledger.

<img src="https://upload.wikimedia.org/wikipedia/commons/3/39/Salmo_salar.jpg" alt="A picture of salmon gotten from wikipedia">


## WHY?

> I noticed my dad having a problem managing his inventory. This was so for many local businesses. I think it will be really helpful.

## GOALS

> salmon should be able to help a business to efficiently manage stock and boost profits by:

- The Kenya Revenue Authority Standard
  - Automatic eTIMS invoices.
  - VAT exemption.
  - Tax returns.
  - ETR receipts.
- Money Flow
  - Mpesa C2B, STK, Till and Paybill.
  - Cash, Card and credit sales.
  - Daily Z report ( Closing report for the casier).
  - Debtors and creditors
- Core features.
    - A plug and play binary that is able to run on. Windows, Linux and Mac.
    - A fast solution.
    - Offline usage/ A local first that it becomes hardened against power and network outages. (Assuming the business is working on a laptop or has power generators and Uninterruptible power supply ).
    - Fast Barcode and search. returns/refunds, discounts,customer display.
    - Hardware: printer(receipts), scanner and cash drawer.
    - Inventory
      - Expiry.
      - Real stock
      - Multistore transfer.
      - Stock take
    - Purchasing
      - LPO
      - Supplier bills
    - Role Based access Control

> Some of this terms are jargons I will document everything in a down to earth manner.

Backups up to 5 years are a requirement by [Kenyan Law](https://www.kra.go.ke/helping-tax-payers/facts-about-kra/category/9).

We don't bundle cloud backup because we can't run that infrastructure for free. Local backups and export tools will be provided; businesses are responsible for their own offsite retention per KRA requirements.


## The stack

- `Sveltekit` for the Website.
- `Wails Svelte` for the Desktop app.

I previously thought on adding a mobile version of the app. Handling product records on a phones keyboard was a hustle for me.
But if salmon's users want a mobile version later. 


## the scope

> the initial scope is a point of sale terminal but as we go on, we might cover other areas for the business.


**WARNING:** Avoid using a single profile for all business needs.

Salmon is elastic at installation time: you can define multiple profiles. Each profile must be scoped to a single concern. Point of Sale, for instance, should never be combined with Enterprise Resource Management in the same profile.

<!--
## install 
&lt;!-- will be added Later --&gt;

## run
&lt;!-- will be added Later --&gt;
-->