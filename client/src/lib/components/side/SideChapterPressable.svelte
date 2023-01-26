<script lang="ts">
    import {createEventDispatcher} from "svelte";

    export let id: string;
    export let index: number;
    export let title: string;

    const dispatch = createEventDispatcher()
    const click = () => dispatch('click', { address: id })
</script>
<button on:click={click} class="border-b border-b-[#151515] p-6 hover:opacity-40 transition duration-500">
    <p class="text-lg font-light items-center flex flex-row gap-3"><span class="text-xs font-light text-[#fff]">{index}.</span>{title}</p>
</button>