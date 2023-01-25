<script lang="ts">
    import {Quotes} from "../../quotes/quotes";
    import {Random} from "../../math/random";

    export let title: string;
    export let value: string;

    const placeholder = Quotes[Random(0, Quotes.length)]
</script>
<div class="flex flex-col py-8 px-12 gap-4">
    <h2 class="font-bold text-2xl leading-none">{title}</h2>
    <textarea
            class="text-neutral-50 bg-black outline-none text-base placeholder:text-neutral-600 min-h-screen resize-none selection:text-black selection:bg-white"
            bind:value
            placeholder={placeholder}
    ></textarea>
</div>