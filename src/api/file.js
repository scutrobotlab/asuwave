import { fetchApi, uploadApi } from "./internal";

export async function uploadFile(file) {
  return await uploadApi("/file/upload", "PUT", file);
}

export async function getFilePath() {
  const res = await fetchApi("/file/path", "GET");
  return res;
}

export async function setFilePath(Path) {
  const res = await fetchApi("/file/path", "PUT", { Path });
  return res;
}

export async function deleteFilePath() {
  const res = await fetchApi("/file/path", "DELETE");
  return res;
}
  