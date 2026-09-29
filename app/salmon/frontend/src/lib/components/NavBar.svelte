<script>
    const ctxs = [
      {title:"track sales"},
      {title:"customers"},
      {title:"stocks"},
      {title:"suppliers"},
      {title:"summary"},
      {title:"settings"},
    ];
</script>
<div class="navbar">
    {#each ctxs as ctx,idx}
        <div class="ctx" title={ctx.title}><img src="/settings.svg" alt={ctx.title}></div>
    {/each}
</div>
<style>

    .navbar {
        height: 100%;
        display: flex;
        flex-direction: column;
    }

    .ctx {
        width: 200px;
        height: 200px;
        background:var(--salmon-accent-1);
        margin:0;

        display:flex;
        justify-content: center;
        align-items:center;

        transition:750ms;
    }

    .ctx img {
        width:100px;
        height:100px;
    }

    .ctx:hover {
        background-color: transparent;
    }
</style>
