<script lang="ts">
    import axios from "$lib/axios";
    import { setCookie } from "$lib/cookies";

    async function onSubmit(e: SubmitEvent) {
        const formData = new FormData(e.target as HTMLFormElement);

        const data: { [key: string]: FormDataEntryValue } = {};
        for (let field of formData) {
            const [key, value] = field;
            data[key] = value;
        }

        let res = await axios
            .post("/token", data)
            .then((res) => res)
            .catch((err) => err);

        if (res.status !== 200) {
            console.error(res.data);
            return;
        }

        if (!res.data.access || !res.data.refresh) {
            console.error(res.data);
            return;
        }

        setCookie("access", res.data.access);
        setCookie("refresh", res.data.refresh);

        console.log(res.data);
    }
</script>

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
