<script>

import {ClientName} from "$lib/config";
</script>
<div class="flex flex-row justify-between border-b border-b-[#151515]">
    <div class="flex flex-col border-r border-r-[#151515] p-6 max-w-xs w-full">
        <h1 class="font-bold text-2xl leading-none">Moonlight</h1>
        <p class="font-light text-xs text-gray-500 leading-none">You are connected to <span class="font-bold">{ClientName}</span></p>
    </div>
    <div class="flex flex-grow p-6">
        <slot/>
    </div>
</div>