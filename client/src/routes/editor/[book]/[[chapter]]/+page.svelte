<script lang="ts">
    import TopHeader from "$lib/components/core/TopHeader.svelte";
    import {Icon} from "@steeze-ui/svelte-icon";
    import {ChevronLeft, Gear, Plus} from "@steeze-ui/radix-icons";
    import {onMount} from "svelte";
    import axios from "axios";
    import {ServerAddr} from "$lib/config";
    import {AUTH_TOKEN} from "$lib/store";
    import {page} from "$app/stores";
    import EditorPage from "$lib/components/pages/editor/EditorPage.svelte";
    import {chapters, contents, currentBook, currentChapter, title} from "$lib/stores/editor";
    import EditorSettings from "$lib/components/pages/editor/EditorSettings.svelte";
    import {redirect} from "$lib/utils/window";

    let currentPage = 'editor'

    onMount(async () => {
        const token = localStorage.getItem("auth.token")
        if (token == null) {
            redirect("/auth")
            return
        }
        $AUTH_TOKEN = token
        let response = await axios.get(ServerAddr + "/books/" + $page.params["book"], {
            headers: {
                Authorization: 'Bearer ' + token
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
        $currentBook = response.data
        response = await axios.get(ServerAddr + "/books/" + $currentBook.id + "/chapters", {
            headers: {
                Authorization: 'Bearer ' + token
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
        $chapters = response.data.chapters
        let currentChapterId = $page.params["chapter"]
        if (currentChapterId != null) {
            await open(currentChapterId)
        }
    })

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

    const back = () => window.location.replace('/dashboard')

    const move = (to: string, otherwise: string) => {
        if (currentPage === to) {
            currentPage = otherwise
            return
        }
        currentPage = to
    }
</script>

<div>
    <TopHeader>
        <div class="flex flex-row justify-between my-auto px-6 items-center w-full">
            <button on:click={back} class="flex flex-row items-center gap-4 hover:opacity-60 transition duration-500">
                <Icon src={ChevronLeft} size="21"></Icon>
                <p class="font-bold text-xl leading-none">{$currentBook?.name}</p>
            </button>
            <div class="flex flex-row gap-4">
                <button><Icon src={Plus} size="21"></Icon></button>
                <button on:click={() => move("settings", "editor")}><Icon src={Gear} size="21"></Icon></button>
            </div>
        </div>
    </TopHeader>
    {#if currentPage === "editor"}
        <EditorPage/>
    {/if}
    {#if currentPage === "settings"}
        <EditorSettings/>
    {/if}
</div>