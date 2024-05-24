<script>
	import '../app.css';
	import SideBar from '$lib/SideBar.svelte';
</script>

<SideBar />

<main class="mx-auto p-8 my-8 max-w-4xl">
	<h1 class="text-4xl lg:text-6xl font-extrabold dark:text-white p-5 text-center mt-20">
		Hello, my name is <span class="text-blue-500">Bastian</span>!
	</h1>

	<p class="text-lg dark:text-gray-300 p-5 text-center">
		I'm a software engineer who loves to build things with code and share my knowledge with others.
	</p>

	<p class="text-lg dark:text-gray-300 p-5 text-center">
		Contact me <a href="mailto:contact@asmussen.tech" class="text-blue-500">here</a>.
</main>

<style lang="postcss">
	:global(html) {
		background-color: theme(colors.gray.900);
	}
</style>
