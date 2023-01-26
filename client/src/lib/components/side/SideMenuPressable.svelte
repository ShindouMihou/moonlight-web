<script lang="ts">

    export let page: string;
    export let href: string;

    const go = () => { page = href }
</script>
<button on:click={go} class="border-b border-b-[#151515] p-6 hover:opacity-40 transition duration-500">
    <slot/>
</button>