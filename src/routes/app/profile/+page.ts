import type { PageLoad } from "./$types";

export const load: PageLoad = async () => {
    // TODO: this shouldn't be invoked when user is not authenticated.
    return {}
};
