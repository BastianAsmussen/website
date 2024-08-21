<script type="module">
	import SideBar from '$lib/SideBar.svelte';
	import { writable } from 'svelte/store';

	const repositoryBlacklist = ['BastianAsmussen'];
	const itemsPerPage = 9; // Number of projects per page
	let currentPage = writable(1);

	async function fetchProjects() {
		const res = await fetch('https://api.github.com/users/BastianAsmussen/repos');
		const data = await res.json();

		let projects = data
			.map((project) => {
				return {
					name: project.name,
					description: project.description || 'Ikke udfyldt.',
					url: project.html_url,
					language: project.language || 'Intet',
					createdAt: project.created_at,
					updatedAt: project.updated_at,
					isArchived: project.archived
				};
			})
			.filter((project) => !repositoryBlacklist.includes(project.name))
			.sort((a, b) => Date.parse(b.createdAt) - Date.parse(a.createdAt));

		return projects;
	}

	function getPaginatedProjects(projects, page, itemsPerPage) {
		const start = (page - 1) * itemsPerPage;
		const end = start + itemsPerPage;
		return projects.slice(start, end);
	}

	function totalPages(projects, itemsPerPage) {
		return Math.ceil(projects.length / itemsPerPage);
	}

	function changePage(newPage, projects, itemsPerPage) {
		if (newPage > 0 && newPage <= totalPages(projects, itemsPerPage)) {
			currentPage.set(newPage);
		}
	}
</script>

<SideBar />

<h1 class="text-5xl font-bold text-white p-10 text-center">Projekter</h1>

{#await fetchProjects()}
	<div class="p-10 text-center">
		<p class="mt-4 text-lg text-gray-300">Henter projekter fra GitHub...</p>
	</div>
{:then projects}
	{#key $currentPage}
		{#if projects.length > 0}
			<div class="grid grid-cols-1 gap-4 md:grid-cols-3 md:gap-4 lg:grid-cols-3 lg:gap-6 p-10">
				{#each getPaginatedProjects(projects, $currentPage, itemsPerPage) as project}
					<a
						href={project.url}
						target="_blank"
						rel="noopener"
						class="p-4 rounded shadow-md bg-gray-800 hover:shadow-lg hover:scale-105 transition duration-300 flex flex-col justify-between"
					>
						<div>
							<h2 class="text-2xl font-bold text-white">
								{project.name}
							</h2>
							<p class="mt-2 text-lg text-gray-300">
								{project.description}
							</p>
						</div>
						<div class="flex justify-between mt-4">
							<p class="text-sm text-gray-500">Sprog: {project.language}</p>
							<p class="text-sm text-gray-500">
								Oprettet: {new Date(project.createdAt).toLocaleDateString()}
							</p>
						</div>
					</a>
				{/each}
			</div>

			<!-- Centered Pagination controls and text -->
			<div class="flex flex-col items-center mt-8 space-y-4">
				<p class="text-lg text-white font-semibold">
					{`Side ${$currentPage} af ${totalPages(projects, itemsPerPage)}`}
				</p>
				<div class="flex space-x-4">
					<button
						class="px-5 py-2 bg-blue-600 text-white font-semibold rounded-full shadow-md hover:bg-blue-700 transition duration-300 disabled:bg-gray-500 disabled:cursor-not-allowed"
						on:click={() => changePage($currentPage - 1, projects, itemsPerPage)}
						disabled={$currentPage === 1}
					>
						&larr; Forrige
					</button>
					<button
						class="px-5 py-2 bg-blue-600 text-white font-semibold rounded-full shadow-md hover:bg-blue-700 transition duration-300 disabled:bg-gray-500 disabled:cursor-not-allowed"
						on:click={() => changePage($currentPage + 1, projects, itemsPerPage)}
						disabled={$currentPage === totalPages(projects, itemsPerPage)}
					>
						Næste &rarr;
					</button>
				</div>
			</div>
		{:else}
			<p class="text-lg text-gray-300">Ingen projekter at vise...</p>
		{/if}
	{/key}
{:catch someError}
	<div class="p-10 text-center">
		<h1 class="text-4xl font-bold text-white">Fejl</h1>
		<p class="mt-4 text-lg text-gray-300">
			{someError.message}
		</p>
	</div>
{/await}

<style lang="postcss">
	:global(html) {
		background-color: theme(colors.gray.900);
	}

	/* Ensure all cards have the same height */
	a {
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		min-height: 160px; /* Adjust this value as needed */
	}
</style>
