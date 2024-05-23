<form
	method="POST"
	class="flex flex-col items-center justify-center min-h-screen bg-gray-100 text-gray-900"
>
	<h1 class="text-6xl font-bold text-blue-500">Contact Me</h1>
	<p class="mt-4 text-xl">Please fill out the form below to get in touch with me.</p>
	<div class="mt-8 w-full max-w-md">
		<label for="name" class="block">Name</label>
		<input
			type="text"
			id="name"
			name="name"
			class="w-full px-4 py-2 mt-2 border border-gray-300 rounded"
			required
		/>
	</div>
	<div class="mt-4 w-full max-w-md">
		<label for="email" class="block">Email</label>
		<input
			type="email"
			id="email"
			name="email"
			class="w-full px-4 py-2 mt-2 border border-gray-300 rounded"
			required
		/>
	</div>
	<div class="mt-4 w-full max-w-md">
		<label for="message" class="block">Message</label>
		<textarea
			id="message"
			name="message"
			class="w-full px-4 py-2 mt-2 border border-gray-300 rounded"
			rows="4"
			required
		></textarea>
	</div>
	<div class="mt-8">
		<button
			type="submit"
			class="px-4 py-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700"
			>Send Message</button
		>
	</div>
</form>
