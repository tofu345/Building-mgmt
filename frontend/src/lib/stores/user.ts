import type { User } from "$lib/types";
import { writable, type Writable } from "svelte/store";

export const isAuthenticated = writable(false);

export const user: Writable<User | null> = writable(null);
