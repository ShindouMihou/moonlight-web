<script lang="ts">
    import TopHeader from "$lib/components/core/TopHeader.svelte";
    import SideHeader from "$lib/components/core/SideHeader.svelte";
    import {Icon} from "@steeze-ui/svelte-icon";
    import {Bookmark, Plus} from "@steeze-ui/radix-icons";
    import SideMenuPressable from "$lib/components/side/SideMenuPressable.svelte";
    import Books from "$lib/components/pages/dashboard/Books.svelte";

    let page = 'books'
</script>

<div>
    <TopHeader>
        <div class="flex flex-row justify-between my-auto px-6 items-center w-full">
            <div>
                <h1 class="font-bold text-xl leading-none">Dashboard</h1>
            </div>
        </div>
    </TopHeader>
    <div class="flex flex-row justify-between">
        <SideHeader>
            <SideMenuPressable bind:page href="books">
                <h3 class="text-sm font-light items-center flex flex-row gap-3"><Icon src={Bookmark} size="16"/>Books</h3>
            </SideMenuPressable>
            <SideMenuPressable bind:page href="create_new_book">
                <h3 class="text-sm font-light items-center flex flex-row gap-3"><Icon src={Plus} size="16"/>Create new book</h3>
            </SideMenuPressable>
        </SideHeader>
        {#if page === "books"}
            <Books/>
        {/if}
    </div>
</div>