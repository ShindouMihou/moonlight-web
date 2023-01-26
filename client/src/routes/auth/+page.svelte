<script lang="ts">
    import TopHeader from "$lib/components/core/TopHeader.svelte";
    import SideHeader from "$lib/components/core/SideHeader.svelte";
    import {Icon} from "@steeze-ui/svelte-icon";
    import {ExclamationTriangle, LockClosed, Person} from "@steeze-ui/radix-icons";
    import {ClientName, ServerAddr} from "$lib/config";
    import type {MidnightError} from "../../lib/types/Error";
    import {AUTH_TOKEN} from "../../lib/store";

    let username: string = ""
    let password: string = ""

    let error: string | null = null

    type AuthenticationResponse = { token: string, expiresIn: Date }

    async function authenticate() {
        if (username.length === 0 || password.length === 0) {
            error = "The username, or password cannot be empty."
            return
        }
        try {
            const response = await fetch(ServerAddr + "/token", {
                method: 'PUT',
                body: JSON.stringify({ username: username, password: password })
            })
            if (response.status === 401) {
                const errorResponse: MidnightError = await response.json()
                error = errorResponse.error
                return
            }
            if (!response.ok) {
                error = "The server is having a mushroom day today."
                return
            }
            const payload: AuthenticationResponse = await response.json()

            // IMPORTANT: Local storage is unsafe, but unless you do some sketchy stuff that'll make
            // the site vulnerable to XSS, this should be good enough.
            localStorage.setItem("auth.token", payload.token)
            $AUTH_TOKEN = payload.token

            window.location.replace("/dashboard")
        } catch (e) {
            error = "An internal error occurred: " + e
        }
    }
</script>

<div>
    <TopHeader>
        <div class="flex flex-row justify-between my-auto px-6 items-center w-full">
            <div>
                <h1 class="font-bold text-xl leading-none">Authentication</h1>
            </div>
        </div>
    </TopHeader>
    <div class="flex flex-row">
        <SideHeader/>
        <div class="flex flex-col flex-grow-0 py-8 px-12 gap-4 max-w-xl">
            {#if error != null}
                <div class="p-3 bg-red-700 text-white rounded-lg hover:opacity-60 transition duration-500 flex flex-row items-center gap-4">
                    <Icon src={ExclamationTriangle} size="16" class="text-red-400 flex-nowrap"></Icon>
                    <p class="text-sm font-light">{error}</p>
                </div>
            {/if}
            <p class="text-sm">
                You are now about to connect into {ClientName}, but before you can continue, you have to authenticate into
                the server by entering a valid account in the server.
            </p>
            <div class="p-2 border-b border-b-[#151515] flex flex-row items-center gap-2">
                <Icon src={Person} size="16"></Icon>
                <input name="username"
                       type="text"
                       bind:value={username}
                       autocomplete="false"
                       placeholder="Username"
                       class="outline-none bg-black p-2 placeholder:text-gray-500 w-full"
                >
            </div>
            <div class="p-2 border-b border-b-[#151515] flex flex-row items-center gap-2">
                <Icon src={LockClosed} size="16"></Icon>
                <input name="password"
                       type="password"
                       bind:value={password}
                       autocomplete="false"
                       placeholder="Password"
                       class="outline-none bg-black p-2 placeholder:text-gray-500 w-full"
                >
            </div>
            <button class="p-2 bg-gray-50 text-black rounded-lg hover:opacity-60 transition duration-500" on:click={authenticate}>Login</button>
        </div>
    </div>
</div>