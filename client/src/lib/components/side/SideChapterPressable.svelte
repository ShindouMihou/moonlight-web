<script lang="ts">
    export let index: number;
    export let title: string;
    // In reality, this would be a UUID that addresses to the chapter.
    export let addressable: string;
</script>
<a href="/chapter/{addressable}" class="border-b border-b-[#151515] p-6 hover:opacity-40 transition duration-500">
    <h3 class="text-lg font-light items-center flex flex-row gap-3"><span class="text-xs font-light text-[#fff]">{index}.</span>{title}</h3>
</a>