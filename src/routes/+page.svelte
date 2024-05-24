<script lang="ts">
	import '../app.css';
	import SideBar from '$lib/SideBar.svelte';
	import { onMount } from 'svelte';

	let time = new Date();

	$: [greeting_text, greeting_color] = getGreeting(time);

	function getGreeting(time: Date): [string, string] {
		if (time.getHours() < 12) {
			return ["Good morning", "text-yellow-400"];
		} else if (time.getHours() < 18) {
			return ["Good afternoon", "text-green-400"];
		} else {
			return ["Good evening", "text-blue-400"];
		}
	}

	onMount(() => {
		const interval = setInterval(() => {
			time = new Date();
		}, 1_000);

		return () => {
			clearInterval(interval);
		};
	});
</script>

<SideBar />

<main class="flex flex-col items-center justify-center">
	<h1 class="text-5xl font-extrabold text-white p-10 text-center">
		<span class="{greeting_color}">{greeting_text}</span>, I'm Bastian Asmussen!
	</h1>
	<p class="text-lg text-gray-300 p-5 text-center">
		I'm a software engineer who loves to build things with code and share my knowledge with others.
	</p>

	<p class="text-lg text-gray-300 p-5 text-center">
		Contact me <a href="mailto:contact@asmussen.tech" class="text-blue-500 hover:underline">here</a> if you have any questions or just want to chat!
	</p>
</main>

<style lang="postcss">
	:global(html) {
		background-color: theme(colors.gray.900);
	}
</style>
