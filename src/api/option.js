import { fetchApi } from "./internal";

export async function getOption() {
  const res = await fetchApi("/option");
  return res.Save;
}

export async function putOption(Save) {
  return await fetchApi("/option", "PUT", { Save });
}
