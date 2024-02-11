<script lang="ts">
    export let title = "Log in";
    export let method = "POST";
</script>

<form {method} class="text-black text-sm p-3 border">
    <div class="pb-4">
        <h1 class="text-3xl">{title}</h1>
    </div>

    <slot />
</form>
