import { fetchApi } from "./internal";

export async function getVersion() {
    const res = await fetchApi("https://api.github.com/repos/scutrobotlab/asuwave/releases/latest", "GET");
    return res;
}

export async function postUpdate() {
    const res = await fetchApi("/update", "POST");
    return res;
}
