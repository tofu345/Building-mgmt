import axios from "$lib/axios";
import { getCookie } from "$lib/cookies";
import type { User } from "$lib/types";

export const getUser = async (): Promise<User | null> => {
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
