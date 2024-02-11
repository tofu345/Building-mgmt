<script>
    import { user } from "./stores/user";
</script>

<nav class="w-full h-14 flex justify-between bg-gray-100">
    <div class="h-full w-full mx-3 flex justify-start items-center gap-2">
        <div
            class="font-sans rounded-md bg-gray-300 px-2 py-1 text-lg text-gray-600 font-bold"
        >
            BUILDING MGMT
        </div>
    </div>
    <div class="h-full w-full mx-3 flex justify-end items-center gap-2">
        {#if $user}
            <div
                class="cursor-pointer rounded-full w-10 h-10 border-2 border-gray-500 bg-gray-200"
            />
        {:else}
            <a href="/login">Login</a>
        {/if}
    </div>
</nav>
