import { writable } from "svelte/store"

export type User = {
    id: string,
    name: string,
    picture: string,
    email: string,
    username: string,
    created_at: string,
    updated_at: string,
}

export const curUser = writable<User | null | undefined>(undefined);

export const whoAmI = async () => {
    try {
        const resp = await fetch("/api/me?simulate_err=false")
        if (resp.status === 200) {
            const user = (await resp.json()) as User
            curUser.set(user);
        } else {
            curUser.set(null)
        }
    } catch (error) {
        console.log("refresh failed: ", error)
        curUser.set(null)
    }
}
