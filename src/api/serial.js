import { fetchApi } from "./internal";

export async function getSerial() {
  const res = await fetchApi("/serial");
  return res.Serials;
}

export async function getSerialCur() {
  const res = await fetchApi("/serial_cur");
  return res;
}

export async function postSerialCur(Serial, Baud) {
  const res = await fetchApi("/serial_cur", "POST", { Serial, Baud });
  return res;
}

export async function deleteSerialCur() {
  return await fetchApi("/serial_cur", "DELETE");
}
