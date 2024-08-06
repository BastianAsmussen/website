<script>
	import { page } from '$app/stores';
	import '../app.css';
	import SideBar from '$lib/SideBar.svelte';
</script>

<svelte:head>
	<title>{$page.status} - {$page.error.message}</title>
	<meta name="description" content="Oops! Noget gik galt, prøv igen senere." />
</svelte:head>

<SideBar />

<div class="flex flex-col items-center justify-center min-h-screen bg-gray-100 text-gray-900">
	<h1 class="text-6xl font-bold text-red-500">{$page.status} - {$page.error.message}</h1>
	<p class="mt-4 text-xl">Ups! Noget gik galt, prøv igen senere.</p>
	<div class="mt-8">
		<button
			on:click={() => (location.href = '/')}
			class="px-4 py-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700">Go Home</button
		>
	</div>
</div>

<style lang="postcss">
	:global(html) {
		background-color: theme(colors.gray.900);
	}
</style>
