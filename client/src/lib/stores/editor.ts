import {writable} from "svelte/store";
import type {Book, Chapter} from "../types/Book";

export const chapters = writable<Chapter[]>(undefined)
export const currentBook = writable<Book>(undefined)
export const currentChapter = writable<Chapter>(undefined)
export const title = writable<string>('')
export const contents = writable<string>('')