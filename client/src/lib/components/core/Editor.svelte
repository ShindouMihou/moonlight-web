<script lang="ts">
    import {Quotes} from "../../quotes/quotes";
    import {Random} from "../../math/random";
    import {contents, title} from "$lib/stores/editor";
    import {onDestroy, onMount} from "svelte";
    import axios from "axios";
    import {ServerAddr} from "../../config";
    import {currentChapter} from "../../stores/editor";
    import {AUTH_TOKEN} from "../../store";
    import {ExclamationTriangle} from "@steeze-ui/radix-icons";
    import {Icon} from "@steeze-ui/svelte-icon";

    const placeholder = Quotes[Random(0, Quotes.length)]

    let autoSaveInterval
    let lastSavedContent = $contents

    let error: string | null = null

    onMount(() => { autoSaveInterval = setInterval(save, 250); console.log('respawned editor')})
    onDestroy(() => { clearInterval(autoSaveInterval); autoSaveInterval = null; console.log('despawned editor'); })

    // Reset auto-save since that may lead to problems.
    currentChapter.subscribe(onChanged, onChanged)

    function onChanged() {
        console.info('change state detected, respawning autosave')
        clearInterval(autoSaveInterval);
        lastSavedContent = $contents

        autoSaveInterval = setInterval(save, 250)
    }

    async function save() {
        if ($contents !== lastSavedContent) {
            try {
                // TODO: Support new chapters somehow.
                let response = await axios.patch(ServerAddr + "/chapters/" + $currentChapter.id, { title: $title, contents: $contents }, {
                    headers: {
                        Authorization: 'Bearer ' + $AUTH_TOKEN
                    } as any
                })
                if (response.status === 401) {
                    localStorage.removeItem("auth.token")
                    error = "Your session has expired, and auto-save couldn't work. Please manually backup before refreshing!"
                    return
                }
                if (response.status === 404) {
                    error = "The chapter cannot be found, and auto-save couldn't work. Please manually backup before refreshing!"
                    return
                }
                if (response.status === 204) {
                    lastSavedContent = $contents
                    error = null
                    return
                }
                error = "Auto-save encountered a mysterious error: Status " + response.status
            } catch (e) {
                error = "Auto-save couldn't save for some reason, we'll keep retrying!"
            }
        }
    }
</script>
<div class="flex flex-col py-8 px-12 gap-4">
    {#if error != null}
        <div class="p-3 bg-red-700 text-white rounded-lg hover:opacity-60 transition duration-500 flex flex-row items-center gap-4">
            <Icon src={ExclamationTriangle} size="16" class="text-red-400 flex-nowrap"></Icon>
            <p class="text-sm font-light">{error}</p>
        </div>
    {/if}
    <input
            bind:value={$title}
            class="font-bold text-2xl leading-none placeholder:text-[#151515] bg-black outline-none selection:text-black selection:bg-white"
            placeholder="Chapter Title"
    />
    <textarea
            class="text-neutral-50 bg-black outline-none text-base placeholder:text-neutral-600 min-h-screen resize-none selection:text-black selection:bg-white"
            bind:value={$contents}
            placeholder={placeholder}
    ></textarea>
</div>