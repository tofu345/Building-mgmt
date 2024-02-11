<script>
    import { goto } from "$app/navigation";
    import { deleteCookie } from "$lib/cookies";
    import { user } from "$lib/stores/user";
    import { onMount } from "svelte";

    onMount(async () => {
        deleteCookie("access");
        deleteCookie("refresh");
        user.set(null);

        goto("/login");
    });
</script>
