import { getVariable, getVariableType } from "@/api/variable.js";

export default {
  namespaced: true,
  state: {
    proj: [],
    read: {},
    write: {},
    vTypes: [],
  },
  getters: {},
  mutations: {
    setV(state, i) {
      if (i.t == "proj") {
        state[i.t] = Object.values(i.v);
        return;
      }
      state[i.t] = i.v;
    },
    setVType(state, i) {
      state.vTypes = i;
    },
  },
  actions: {
    getV({ commit }, t) {
      getVariable(t).then((v) => {
        commit("setV", { t, v });
      });
    },
    getVType({ commit }) {
      getVariableType().then((v) => {
        commit("setVType", v);
      });
    },
  },
  modules: {},
};