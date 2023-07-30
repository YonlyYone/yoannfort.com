<script lang="ts">
	import Card from "$lib/Components/Common/Card.svelte";
	import ProfileCard from "$lib/Components/ProfileCard.svelte";
	import ResumeCard from "$lib/Components/ResumeCard.svelte";
</script>

<ProfileCard/>
<div class="content-container">
	<ResumeCard/>
	<Card id="works"></Card>
	<Card id="blog"></Card>
	<Card id="products"></Card>
	<Card id="contact"></Card>
</div>

<style>
	.content-container{
		max-height: 75vh;
        overflow: scroll;
	}
</style>