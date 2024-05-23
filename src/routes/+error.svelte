<script>
	import { page } from '$app/stores';
	import '../app.css';
</script>

<svelte:head>
	<title>{$page.status} - {$page.error.message}</title>
	<meta
		name="description"
		content="Oops, something went wrong. Please try again or contact me if the problem persists."
	/>
</svelte:head>

<div class="flex flex-col items-center justify-center min-h-screen bg-gray-100 text-gray-900">
	<h1 class="text-6xl font-bold text-red-500">{$page.status} - {$page.error.message}</h1>
	<p class="mt-4 text-xl">
		Oops, something went wrong. Please try again or contact me if the problem persists.
	</p>
	<div class="mt-8">
		<button
			on:click={() => (location.href = '/')}
			class="px-4 py-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700">Go Home</button
		>
		<button
			on:click={() => (location.href = '/contact')}
			class="ml-4 px-4 py-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700"
			>Contact Me</button
		>
	</div>
</div>

<style lang="postcss">
	:global(html) {
		background-color: theme(colors.gray.100);
	}
</style>
