<script type="module">
	import SideBar from '$lib/SideBar.svelte';

	async function fetchProjects() {
		const res = await fetch('https://api.github.com/users/BastianAsmussen/repos');
		const data = await res.json();

		let projects = data.map((project) => {
			return {
				name: project.name,
				description: project.description || 'No description provided...',
				url: project.html_url,
				language: project.language || 'N/A',
				date: project.created_at,
				archived: project.archived
			};
		});

		// Sort by updated date.
		projects.sort((a, b) => Date.parse(b.date) - Date.parse(a.date));

		return projects;
	}
</script>

<SideBar />

<h1 class="text-5xl font-bold text-white p-10 text-center">Projects</h1>

{#await fetchProjects()}
	<div class="p-10 text-center">
		<p class="mt-4 text-lg text-gray-300">Fetching projects from GitHub...</p>
	</div>
{:then projects}
	<div class="grid grid-cols-1 gap-4 md:grid-cols-3 md:gap-4 lg:grid-cols-3 lg:gap-6 p-10">
		{#each projects as project}
			<a
				href={project.url}
				target="_blank"
				rel="noopener"
				class="p-4 rounded shadow-md bg-gray-800 hover:shadow-lg hover:scale-105 transition duration-300"
			>
				<h2 class="text-2xl font-bold text-white">
					{project.name}
				</h2>
				<p class="mt-2 text-lg text-gray-300">
					{project.description}
				</p>
				<div class="flex justify-between mt-4">
					<p class="text-sm text-gray-500">{project.language}</p>
					<p class="text-sm text-gray-500">{new Date(project.date).toLocaleDateString()}</p>
				</div>
			</a>
		{/each}
	</div>
{:catch someError}
	<div class="p-10 text-center">
		<h1 class="text-4xl font-bold text-white">Error</h1>
		<p class="mt-4 text-lg text-gray-300">
			{someError.message}
		</p>
	</div>
{/await}

<style lang="postcss">
	:global(html) {
		background-color: theme(colors.gray.900);
	}
</style>
