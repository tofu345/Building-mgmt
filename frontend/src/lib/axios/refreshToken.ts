import { deleteCookie, getCookie, setCookie } from "$lib/cookies";
import mem from "mem";
import axios from "./index";

const refreshTokenFn = async () => {
    const refresh = getCookie("refresh");
    if (!refresh) {
        return null;
    }

    try {
        const res = await axios
            .post("/token/refresh", { token: refresh })
            .then((res) => res)
            .catch((err) => err);

        if (res.status != 200) {
            deleteCookie("access");
            deleteCookie("refresh");
        }

        setCookie("access", res.data.access, 7);
    } catch (error) {
        deleteCookie("access");
        deleteCookie("refresh");
    }
};

const maxAge = 10000; // 10 seconds

// Cache function for 10 seconds
export const memoizedRefreshToken = mem(refreshTokenFn, { maxAge });
