<script lang="ts">
	import MobileNavigation from "$lib/Components/MobileNavigation.svelte";

</script>

<div class="container">
    <MobileNavigation/>
    <div class="content-wrapper">
        <slot></slot>
    </div>
</div>