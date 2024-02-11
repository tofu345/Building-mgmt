export function setCookie(
    name: string,
    value: string | null,
    daysToLive: number | null = 7
) {
    const date = new Date();
    if (!daysToLive) {
        daysToLive = 0;
    }

    date.setTime(date.getTime() + daysToLive * 24 * 60 * 60 * 1000); // Convert date to milliseconds
    let expires = date.toUTCString();
    document.cookie = `${name}=${value}; expires=${expires}; path=/; SameSite=Lax`;
}

export function deleteCookie(name: string) {
    setCookie(name, null, null);
}

export function getCookie(name: string): string | null {
    const cookieString = decodeURIComponent(document.cookie);
    const cookies = cookieString.split("; ");
    let result = null;

    cookies.forEach((element) => {
        if (element.indexOf(name) == 0) {
            result = element.substring(name.length + 1);
        }
    });

    return result;
}
