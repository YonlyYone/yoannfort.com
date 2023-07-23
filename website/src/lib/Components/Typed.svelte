<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import Typed from 'typed.js';

	const dispatch = createEventDispatcher();

	export let strings :string[] = ['Hello World!'];
	export let stringsElement:string|Element|undefined = undefined;
	export let typeSpeed = 50;
	export let startDelay = 20;
	export let backSpeed = 50;
	export let smartBackspace = false;
	export let shuffle = false;
	export let backDelay = 700;
	export let fadeOut = false;
	export let fadeOutClass = 'typed-fade-out';
	export let fadeOutDelay = 500;
	export let loop = true;
	export let loopCount = Infinity;
	export let showCursor = false;
	export let cursorChar = '|';
	export let autoInsertCss = true;
	export let attr: string|undefined = undefined;
	export let bindInputFocusEvents = false;
	export let contentType = 'html';

	onMount(() => {
		let el = document.querySelector('.typing');

		const typed = new Typed(el, {
			strings,
			stringsElement,
			typeSpeed,
			startDelay,
			backSpeed,
			smartBackspace,
			shuffle,
			backDelay,
			fadeOut,
			fadeOutClass,
			fadeOutDelay,
			loop,
			loopCount,
			showCursor,
			cursorChar,
			autoInsertCss,
			attr,
			bindInputFocusEvents,
			contentType,
			onComplete: () => dispatch('complete'),
			onStringTyped: () => dispatch('stringTyped'),
			onLastStringBackspaced: () => dispatch('lastStringBackspaced'),
			onTypingPaused: () => dispatch('typingPaused'),
			onTypingResumed: () => dispatch('typingResumed'),
			onReset: () => dispatch('reset'),
			onStop: () => dispatch('stop'),
			onStart: () => dispatch('start'),
			onDestroy: () => dispatch('destroy')
		});
	});
</script>

<div>
	<slot />
</div>
