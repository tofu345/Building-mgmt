<script lang="ts">
    import axios from "$lib/axios";
    import { onMount } from "svelte";

    async function getLocations() {
        let response = await axios
            .get("/locations")
            .then((res) => res.data)
            .catch((err) => err);
        console.log(response);
    }

    onMount(() => {
        getLocations();
    });
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>
