<script lang="ts">
	import { createEventDispatcher, type SvelteComponent } from "svelte";

    const dispach = createEventDispatcher()

    const handleClick = (event: MouseEvent) => dispach('click', event) 

    interface $props extends SvelteComponent {}
</script>

<button on:click={handleClick} {...$$props}>
    <slot></slot>
</button>

<style>
    button {
        background: none;
        color: inherit;
        border: none;
        padding: 0;
        font: inherit;
        cursor: pointer;
        outline: inherit;
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        padding:0 8px;
    }
</style>