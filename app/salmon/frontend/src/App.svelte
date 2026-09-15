<!--/**
 * Copyright (C) 2026 Joseph Wangai Mwaniki
 * 
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 * 
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 **/ -->
 <script lang="ts">
  import {onMount} from 'svelte';
  import {Events, WML} from "@wailsio/runtime";
  import {GreetService} from "../bindings/changeme";

  const wailsVersion = "v3.0.0-beta.20";

  let name: string = $state('');
  let time: string = $state('Listening for Time event...');

  let titleNameEl: HTMLElement;
  let toastEl: HTMLElement;
  let resultEl: HTMLElement;
  let toastTimer: ReturnType<typeof setTimeout>;

  onMount(() => {
    Events.On('time', (v: any) => {
      // On a narrow screen the full RFC1123 stamp is too wide for the footer, so
      // show just the clock time there (matching the CSS breakpoint).
      const full = v.data;
      const compact = (full.match(/\d{1,2}:\d{2}:\d{2}/) || [full])[0];
      time = window.matchMedia('(max-width: 640px)').matches ? compact : full;
    });
    // Wire up data-wml-openURL links (logos + footer "Docs" link).
    WML.Reload();
  });

  // Crossfade the framework word in the heading ("Wails + Svelte") to the name
  // the user entered ("Wails + <name>"): the old word fades out while the new one
  // fades in over the same spot.
  function swapTitleName(name: string): void {
    const current = titleNameEl.querySelector('.title-name-text:not(.is-outgoing)');
    if (!current || current.textContent === name) {
      return;
    }
    const incoming = document.createElement('span');
    incoming.className = 'title-name-text is-entering';
    incoming.textContent = name;
    current.classList.add('is-outgoing');
    titleNameEl.appendChild(incoming);
    // Force a reflow so the transitions run from the starting state.
    void incoming.offsetWidth;
    incoming.classList.remove('is-entering');
    current.classList.add('is-leaving');
    current.addEventListener('transitionend', () => current.remove(), {once: true});
  }

  // Pop the toast with the message Go returned, then auto-dismiss it.
  function showToast(message: string): void {
    resultEl.innerText = message;
    toastEl.classList.add('is-visible');
    clearTimeout(toastTimer);
    toastTimer = setTimeout(() => toastEl.classList.remove('is-visible'), 4000);
  }

  const doGreet = (): void => {
    let n = name || 'anonymous';
    swapTitleName(n);
    GreetService.Greet(n).then(showToast).catch(console.error);
  }
</script>
<main>
  <h1>Salmon</h1>
</main>
<style>
  /* Put your standard CSS here */
</style>
