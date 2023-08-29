<script>
    import Nav from "$lib/Nav.svelte";
    import { user } from "$lib/stores/user";
    import { getUser } from "$lib/utils/user.utils";
    import { onMount } from "svelte";
    import "../../app.css";

    onMount(async () => {
        user.set(await getUser());
    });
</script>

<Nav />

<slot />
