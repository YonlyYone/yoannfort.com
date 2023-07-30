<script lang="ts">
	import { afterUpdate, onMount } from 'svelte'; 

	import links from '$lib/Constants/Navigation.json';
    import Profil from '$lib/Constants/Profil.json'

	import DiscreteButton from './Common/DiscreteButton.svelte';
	import Typed from './Common/Typed.svelte';

	const handleDarkModeToogle = () => {
		console.log('darkmode')
	}

	const handleCartToogle = () => {
		console.log('show cart')
	}

	const handlePanelToogle = () =>{
		console.log('show panel')
	}

	const getLocationHash = () => {
		const { hash } = location
		currentHash = hash === '' ? "/#" : hash
		console.log({currentHash, location})
	}

	let currentHash = "/#"

	onMount(() => {
		getLocationHash()
	})

	afterUpdate(() => {
		getLocationHash()
	});
</script>

<header class="mobile-nav">
	<div class="profile">
		<div>
			<p>Yoann Fort</p>
			<Typed className="mobile-typed" strings={Profil.typed}>
				<p class="typing" />
			</Typed>
		</div>
		<div class="button-container">
			<div>
				<DiscreteButton on:click={handleDarkModeToogle}>
					<i class="fa-regular fa-moon" />
				</DiscreteButton>
			</div>
			<div>
				<DiscreteButton on:click={handleCartToogle}>
					<i class="fa-solid fa-cart-shopping" />
				</DiscreteButton>
			</div>
			<div>
				<DiscreteButton on:click={handlePanelToogle }>
					<i class="fa-solid fa-bars" />
				</DiscreteButton>
			</div>
		</div>
	</div>

	<div class="top-menu">
		<ul>
			{#each links.links as {url, icon, name}}
				<li>
					<a href={`/${url}`}>
						<i class={icon}/>
						<p>{name.toUpperCase()}</p>
					</a>
				</li>
			{/each}
		</ul>
	</div>
</header>

<style>
	i {
		font-size: 22px;
	}

	.mobile-nav {
		background-color: white;
        position: static;
	}

	.profile {
		padding: 4px 18px;
		height: 42px;
		text-align: space;
		display: flex;
		flex-direction: row;
		flex-wrap: nowrap;
		justify-content: space-between;
	}

    .button-container > div {
        width: 100%;
        height: 100%;
        min-width: 32px;
		margin-right: 12px;
    }

	.button-container > div:last-child {
		margin-right: 0;
    }

	.button-container  {
		justify-content: flex-end;
	}

	.profile > div {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: flex-start;
	}

	.profile > div:last-child {
		flex-direction: row;
	}

	.profile p {
		margin: 0;
		padding: 0;
	}

	.top-menu {
		font-size: 16px;
		font-weight: bold;
	}

	.top-menu > ul {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100%;
		list-style: none;
		border-top: 1px solid grey;
		border-bottom: 1px solid grey;
		padding-left: 0px;
		margin: 0;
		overflow: scroll;
	}

	.top-menu > ul > li:first-child {
		border-left: 1px solid grey;
	}

	.top-menu > ul > li {
		padding-top: 8px;
		border-right: 1px solid grey;
		margin: 10px 0;
		display: block;
		min-width: 86px;
		width: 20%;

	}

	.top-menu > ul > li > a {
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: column;
		flex-wrap: wrap;
	}

	.top-menu > ul > li > a > p {
		font-size: 12px;
		margin-top: 6px;
	}

	.top-menu > ul > li > a {
		text-decoration: none;
        color: inherit;
	}
	.top-menu > ul > li > a:visited {
        color: inherit;
		text-decoration: none;
	}
</style>
