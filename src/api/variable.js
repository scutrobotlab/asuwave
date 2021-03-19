import { fetchApi, uploadFile } from "./internal";

export async function getVariable(mode) {
  const res = await fetchApi("/variable_" + mode);
  return res.Variables;
}

export async function getVariableType() {
  const res = await fetchApi("/variable_type");
  return res.Types;
}

export async function postVariable(mode, Board, Name, Type, Addr) {
  return await fetchApi("/variable_" + mode, "POST", { Board, Name, Type, Addr });
}

export async function postVariableToProj(file) {
  return await uploadFile("/variable_proj", "POST", file);
}

export async function putVariable(mode, Board, Name, Type, Addr, Data) {
  return await fetchApi("/variable_" + mode, "PUT", { Board, Name, Type, Addr, Data });
}

export async function deleteVariable(mode, Board, Name, Type, Addr) {
  return await fetchApi("/variable_" + mode, "DELETE", { Board, Name, Type, Addr });
}
