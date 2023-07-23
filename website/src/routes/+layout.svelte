<script>
	import MobileNavigation from "$lib/Components/MobileNavigation.svelte";

</script>

<div class="container">
    <MobileNavigation/>
    <div class="content-wrapper">
        <slot></slot>
    </div>
</div>

<style>
    .container {
        width: 100vw;
        height: 100vh;
        position: fixed;
        top: 0;
        left: 0;
        background-color: lightgray;
    }
    .content-wrapper {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        width: 85%;
        height: 90%;
    }
</style>