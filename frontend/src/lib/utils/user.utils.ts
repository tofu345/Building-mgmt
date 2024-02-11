import { browser } from "$app/environment";
import axios from "$lib/axios";
import { getCookie } from "$lib/cookies";
import type { User } from "$lib/types";

export const getUser = async (): Promise<User | null> => {
    if (!browser) {
        console.error("Cannot get user when not in browser");
        return null;
    }

    const token = getCookie("access");
    if (!token) {
        return null;
    }

    const res = await axios
        .get("/users/info")
        .then((res) => res)
        .catch((err) => err.response);

    if (res.status != 200) {
        return null;
    }

    return res.data;
};
