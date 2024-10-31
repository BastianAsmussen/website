<script lang="ts">
	import '../app.css';
	import SideBar from '$lib/SideBar.svelte';
	import { onMount } from 'svelte';

	let time = new Date();

	$: [greeting_text, greeting_color] = getGreeting(time);

	function getGreeting(time: Date): [string, string] {
		if (time.getHours() < 9) {
			return ['God morgen', 'text-sky-500'];
		} else if (time.getHours() < 12) {
			return ['God formiddag', 'text-amber-500'];
		} else if (time.getHours() < 13) {
			return ['God middag', 'text-yellow-500'];
		} else if (time.getHours() < 17) {
			return ['God eftermiddag', 'text-lime-500'];
		} else {
			return ['God aften', 'text-indigo-500'];
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
		<span class={greeting_color}>{greeting_text}</span>, mit navn er Bastian Asmussen!
	</h1>
	<p class="text-lg text-gray-300 p-5 text-center">
		Jeg er en programmør som elsker at kode og dele min viden med andre.
	</p>

	<p class="text-lg text-gray-300 p-5 text-center">
		Kontakt mig <a href="mailto:contact@asmussen.tech" class="text-blue-500 hover:underline">her</a>
		hvis du har spørgsmål, eller bare vil snakke!
	</p>
</main>

<style lang="postcss">
	:global(html) {
		background-color: theme(colors.gray.900);
	}
</style>
