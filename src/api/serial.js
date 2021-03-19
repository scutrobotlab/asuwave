import { fetchApi } from "./internal";

export async function getSerial() {
  const res = await fetchApi("/serial");
  return res.Serials;
}

export async function getSerialCur() {
  const res = await fetchApi("/serial_cur");
  return res.Serial;
}

export async function postSerialCur(Serial) {
  const res = await fetchApi("/serial_cur", "POST", { Serial });
  return res.Serial;
}

export async function deleteSerialCur() {
  return await fetchApi("/serial_cur", "DELETE");
}
