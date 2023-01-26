<script lang="ts">
    import SideHeader from "$lib/components/core/SideHeader.svelte";
    import SideChapterPressable from "$lib/components/side/SideChapterPressable.svelte";
    import Editor from "$lib/components/core/Editor.svelte";
    import axios from "axios";
    import {ServerAddr} from "$lib/config";
    import {AUTH_TOKEN} from "$lib/store";
    import {contents, currentChapter, title, chapters, currentBook} from "$lib/stores/editor";
    import {redirect} from "$lib/utils/window";
    
    const back = () => window.location.replace('/dashboard')
    
    async function move(event: CustomEvent) {
        let addressable = event.detail.address
        await open(addressable)
    }

    async function open(chapter: string) {
        let response = await axios.get(ServerAddr + "/chapters/" + chapter, {
            headers: {
                Authorization: 'Bearer ' + $AUTH_TOKEN
            } as any
        })
        if (response.status === 401) {
            localStorage.removeItem("auth.token")
            redirect("/auth")
            return
        }
        if (response.status === 404) {
            back()
            return
        }
        $currentChapter = response.data
        $title = $currentChapter.title
        $contents = $currentChapter.contents
    }
</script>

<div>
    <div class="flex flex-row justify-between">
        <SideHeader>
            {#if $chapters != null}
                {#each $chapters as chapter, index}
                    <SideChapterPressable id={chapter.id} index={index + 1} title={chapter.title} on:click={move}/>
                {/each}
            {/if}
        </SideHeader>
        <div class="flex flex-col flex-grow">
            {#if $currentChapter != null && $currentBook != null}
                <Editor/>
            {/if}
        </div>
    </div>
</div>