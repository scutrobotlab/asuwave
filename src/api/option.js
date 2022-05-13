import { fetchApi } from "./internal";

export async function getOption() {
  const res = await fetchApi("/option");
  return res;
}

export async function putOption(Key, Value) {
  return await fetchApi("/option", "PUT", { Key, Value });
}
