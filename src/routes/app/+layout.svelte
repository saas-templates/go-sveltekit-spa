<script lang="ts">
	import { curUser, whoAmI } from '$lib/store';
	import { onMount } from 'svelte';

	onMount(() => {
		whoAmI();
	});
</script>

{#if $curUser === undefined}
	<div class="w-full min-h-screen grid place-items-center">
		<progress class="progress w-56" />
	</div>
{:else if $curUser === null}
	<!-- show login component -->
	<h1>Please Login</h1>
{:else}
	<slot />
{/if}
