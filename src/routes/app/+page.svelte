<script lang="ts">
	import { curUser } from '$lib/store';
</script>

{#if $curUser}
	<div class="flex flex-row items-center gap-2 p-2">
		<img class="rounded-full w-12" src={$curUser.picture} alt="" />
		<div class="flex flex-col">
			<h1 class="text-2xl">{$curUser.name}</h1>
			<h3>{$curUser.email}</h3>
		</div>
	</div>
{/if}

<a href="/profile">Profile</a>
