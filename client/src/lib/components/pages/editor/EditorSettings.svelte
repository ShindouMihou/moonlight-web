<script lang="ts">
    import SideHeader from "$lib/components/core/SideHeader.svelte";
    import SideMenuPressable from "$lib/components/side/SideMenuPressable.svelte";
    import {Bookmark} from "@steeze-ui/radix-icons";
    import {Icon} from "@steeze-ui/svelte-icon";
    import {currentBook} from "$lib/stores/editor";
    import axios from "axios";
    import {ServerAddr} from "$lib/config";
    import {AUTH_TOKEN} from "$lib/store";
    import {redirect} from "$lib/utils/window";


    let page = 'book';
    let name = $currentBook.name

    const back = () => window.location.replace('/dashboard')

    const save = async () => {
        let response = await axios.patch(ServerAddr + "/books/" + $currentBook.id, {name}, {
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
        if (response.status === 204) {
            const book = $currentBook
            book.name = name
            $currentBook = book
        }
    }
</script>

<div>
    <div class="flex flex-row justify-between">
        <SideHeader>
            <SideMenuPressable bind:page href="book.settings">
                <h3 class="text-sm font-light items-center flex flex-row gap-3"><Icon src={Bookmark} size="16"/>Book Settings</h3>
            </SideMenuPressable>
        </SideHeader>
        <div class="flex flex-col flex-grow py-8 px-12 gap-4">
            <div>
                <p class="font-bold text-2xl leading-none">Book Name</p>
                <p class="font-light text-xs text-gray-200">A book needs to have a good name, after all.</p>
                <div class="p-2 border-b border-b-[#151515] flex flex-row items-center gap-2 pt-4 px-0">
                    <Icon src={Bookmark} size="16"></Icon>
                    <input name="name"
                           type="text"
                           bind:value={name}
                           autocomplete="false"
                           placeholder="Book Name"
                           class="outline-none bg-black p-2 placeholder:text-gray-500 w-full"
                    >
                </div>
            </div>
            <button class="p-2 bg-gray-50 text-black rounded-lg hover:opacity-60 transition duration-500" on:click={save}>Save</button>
        </div>
    </div>
</div>