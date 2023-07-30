<script lang="ts">
	import type { SvelteComponent } from "svelte";

    interface $props extends SvelteComponent {}

    export let title: string | null = null;
</script>

<section 
    {...$$props} 
    style:padding={title? '30px' : '0'} 
    style:width={title ? '420px' : '480px' }
>
    {#if title}
        <h2>title</h2>
    {/if}
    <slot></slot>
</section>

<style>
    section{
        display: flex;
        justify-content: flex-start;
        align-items: flex-start;
        flex-direction: column;
        margin-top: 12px;
        min-height: 300px;
        max-width: 100vw;
        background-color: white;
        margin: 15px 0;
        opacity: 1;
        border-radius: 4px;
    }
    section > h2 {
        margin-bottom: 30px;
    }
</style>