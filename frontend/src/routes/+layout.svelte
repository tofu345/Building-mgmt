<script>
    import { user } from "$lib/stores/user";
    import { getUser } from "$lib/utils/user.utils";
    import { onMount } from "svelte";

    onMount(async () => {
        console.log("fetching user data...");
        user.set(await getUser());
        console.log($user);
    });
</script>

<slot />
