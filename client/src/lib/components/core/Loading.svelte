<script lang="ts">
    import {onMount} from "svelte";
    import {AUTH_TOKEN} from "$lib/store";

    export let authenticatedOnly: boolean = true
    export let defaultRoute: string | null = null

    onMount(() => {
        const token = localStorage.getItem("auth.token")
        if (token == null && authenticatedOnly) {
            redirect("/auth")
            return
        }
        $AUTH_TOKEN = token
        if (defaultRoute) {
            redirect(defaultRoute)
        }
    })

    function redirect(location: string) {
        window.location.replace(location)
    }
</script>

<div class="p-4 m-auto items-center flex flex-col justify-center min-h-screen">
    <h1 class="font-bold text-2xl leading-none animate-bounce">Moonlight</h1>
</div>