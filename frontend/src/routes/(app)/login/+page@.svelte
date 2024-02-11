<script lang="ts">
    import { goto } from "$app/navigation";
    import axios from "$lib/axios";
    import { setCookie } from "$lib/cookies";
    import { user } from "$lib/stores/user";
    import { getUser } from "$lib/utils/user.utils";
    import { onMount } from "svelte";

    let errors: string | Object;

    async function onSubmit(e: SubmitEvent) {
        const formData = new FormData(e.target as HTMLFormElement);

        const data: { [key: string]: FormDataEntryValue } = {};
        for (let field of formData) {
            const [key, value] = field;
            data[key] = value;
        }

        const res = await axios
            .post("/token", data)
            .then((res) => res)
            .catch((err) => err.response);

        if (res.status !== 200) {
            errors = res.data.detail || res.data.errors;
            console.log(res);
            return;
        }

        if (!res.data.access || !res.data.refresh) {
            console.log(res.data);
            return;
        }

        setCookie("access", res.data.access);
        setCookie("refresh", res.data.refresh);

        goto("/");

        user.set(await getUser());
    }

    onMount(() => {
        if ($user) {
            goto("/");
        }
    });
</script>

{#if typeof errors === "object" && !Array.isArray(errors) && errors !== null}
    {#each Object.entries(errors) as [key, value]}
        <p>{key}: {value}</p>
    {/each}
{:else}
    <p>{errors || ""}</p>
{/if}

<form on:submit|preventDefault={onSubmit}>
    <label for="email">Email</label>
    <input class="border-2" type="text" name="email" id="email" required />
    <label for="password">Password</label>
    <input
        class="border-2"
        type="password"
        name="password"
        id="password"
        required
    />
    <button type="submit" class="border-2">Login</button>
</form>
