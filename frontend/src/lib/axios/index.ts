import { getCookie } from "$lib/cookies";
import axios from "axios";
import { memoizedRefreshToken } from "./refreshToken";

const baseURL = "http://localhost:8000";

const instance = axios.create({
    baseURL,
    timeout: 15000,
});

instance.interceptors.request.use(
    async (config) => {
        const access = getCookie("access");

        if (access) {
            config.headers["Authorization"] = `Bearer ${access}`;
        }

        return config;
    },
    (error) => Promise.reject(error)
);

instance.interceptors.response.use(
    (response) => response,
    async (error) => {
        const config = error?.config;

        if (error?.response?.status === 401 && !config?.sent) {
            config.sent = true;

            const access = await memoizedRefreshToken();

            if (access) {
                config.headers["Authorization"] = `Bearer ${access}`;
            }

            return axios(config);
        }
        return Promise.reject(error);
    }
);

export default instance;
