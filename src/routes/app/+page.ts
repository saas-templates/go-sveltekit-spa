import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
    await fetch("/api/ping")
    return {}
};
