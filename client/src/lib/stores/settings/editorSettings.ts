import {browser} from "$app/environment";
import {writable} from "svelte/store";

export const AutoSave = writable(stored('autosave.enabled', 'true') === 'true')
// :note(v) in seconds
export const AutoSaveInterval = writable(Number.parseInt(stored('autosave.interval', '5')))

if (browser) {
    AutoSave.subscribe((value) => localStorage.setItem('autosave.enabled', String(value)))
    AutoSaveInterval.subscribe((value) => localStorage.setItem('autosave.interval', String(value)))
}

function stored(key: string, def: string) {
    if (browser) {
        return (localStorage.getItem(key) ?? def)
    }

    return def;
}