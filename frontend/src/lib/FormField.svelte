<script lang="ts">
    export let label: string;
    export let name: string;
    export let placeholder: string;
    export let type: string;
</script>

<div>
    <label for="email" class="block mb-2 text-gray-900"> {label} </label>
    <input
        {type}
        {name}
        class="bg-gray-50 border text-gray-900 rounded-lg focus:border-[#3398e1] w-full p-2.5 mb-2"
        {placeholder}
    />
</div>
