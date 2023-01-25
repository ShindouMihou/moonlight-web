<script>
    import TopHeader from "$lib/components/core/TopHeader.svelte";
    import SideHeader from "$lib/components/core/SideHeader.svelte";
    import SideChapterPressable from "$lib/components/side/SideChapterPressable.svelte";
    import Editor from "$lib/components/core/Editor.svelte";
    import {Icon} from "@steeze-ui/svelte-icon";
    import {Gear, Plus} from "@steeze-ui/radix-icons";
</script>

<div>
    <TopHeader>
        <div class="flex flex-row justify-between my-auto px-6 items-center w-full">
            <div>
                <h1 class="font-bold text-xl leading-none">Book Name</h1>
            </div>
            <div class="flex flex-row gap-4">
                <Icon src={Plus} size="21"></Icon>
                <Icon src={Gear} size="21"></Icon>
            </div>
        </div>
    </TopHeader>
    <div class="flex flex-row justify-between">
        <SideHeader>
            <SideChapterPressable index={1} title="Chapter 1" addressable="chapter-1"/>
        </SideHeader>
        <div class="flex flex-col flex-grow">
            <Editor title="Chapter 1" value=""/>
        </div>
    </div>
</div>