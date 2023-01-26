<script lang="ts">
    import {Icon} from "@steeze-ui/svelte-icon";
    import {ChevronRight} from "@steeze-ui/radix-icons";
    import {onMount} from "svelte";
    import type {Book} from "../../../types/Book";
    import axios from "axios";
    import {ServerAddr} from "../../../config";

    let books: Book[]
    onMount(async () => {
        const token = localStorage.getItem("auth.token")
        if (token == null) {
            window.location.replace("/auth")
            return
        }

        let response = await axios.get(ServerAddr + "/books", {
            headers: {
                Authorization: 'Bearer ' + token
            } as any
        })
        if (response.status === 401) {
            localStorage.removeItem("auth.token")
            window.location.replace("/auth")
            return
        }
        books = (response.data as { books: Book[], length: number }).books as Book[]
    })
</script>
<div class="flex flex-row flex-grow flex-wrap gap-4 p-4 px-12 gap-4">
    {#if books != null}
        {#each books as book}
            <a href="editor/{book.id}" class="border border-[#151515] p-5 h-fit hover:opacity-60 transition duration-500 items-center flex flex-row gap-2">
                <h4 class="font-medium text-xl leading-none">{book.name}</h4>
                <Icon src={ChevronRight} size="16"></Icon>
            </a>
        {/each}
    {/if}
</div>