import { getCookie } from "$lib/cookies";
import axios from "axios";
import { memoizedRefreshToken } from "./refreshToken";

const baseURL = "http://localhost:8000";

const instance = axios.create({
    baseURL,
    timeout: 15000,
});

// Attach headers
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

// Try refresh token
instance.interceptors.response.use(
    (response) => response,
    async (error) => {
        const config = error?.config;
        const url = config.url;

        if (error?.response?.status !== 200) {
            const access = await memoizedRefreshToken();
            if (access) {
                config.headers["Authorization"] = `Bearer ${access}`;
            }

            return axios(config);
        }
        // TODO: Redirect to login

        return Promise.reject(error);
    }
);

export default instance;
