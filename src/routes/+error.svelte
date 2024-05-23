<script>
    import { page } from '$app/stores';
    import '../app.css';
</script>

<div class="flex flex-col items-center justify-center h-screen bg-white dark:bg-gray-900 text-black dark:text-white">
    <h1 class="text-9xl font-bold text-red-500">{$page.status} - {$page.error.message}</h1>
    <button on:click={() => location.href = '/'} class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-10">Go Home</button>
</div>

<style lang="postcss">
    :global(html) {
        background-color: theme(colors.gray.100);
    }
</style>
