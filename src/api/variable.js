import { fetchApi, uploadFile } from "./internal";

export async function getVariable(mode) {
  const res = await fetchApi("/variable_" + mode);
  return res;
}

export async function getVariableType() {
  const res = await fetchApi("/variable_type");
  return res.Types;
}

export async function postVariable(mode, Board, Name, Type, Addr, Inputcolor, SignalGain, SignalBias) {
  return await fetchApi("/variable_" + mode, "POST", { Board, Name, Type, Addr, Inputcolor, SignalGain, SignalBias });
}

export async function postVariableToProj(file) {
  return await uploadFile("/variable_proj", "POST", file);
}

export async function putVariable(mode, Board, Name, Type, Addr, Data, Inputcolor) {
  return await fetchApi("/variable_" + mode, "PUT", { Board, Name, Type, Addr, Data, Inputcolor });
}

export async function deleteVariable(mode, Board, Name, Type, Addr, Inputcolor) {
  return await fetchApi("/variable_" + mode, "DELETE", { Board, Name, Type, Addr, Inputcolor });
}
export async function deleteVariableAll() {
  return await fetchApi("/variable_proj", "DELETE", "vToProj.json");
}
