import { getVariable, getVariableType } from "@/api/variable.js";

export default {
  namespaced: true,
  state: {
    proj: [],
    read: {},
    modi: {},
    vTypes: [],
  },
  getters: {
    searchVToProj: (state) => (keyword) => {
      if (keyword) {
        let kw = keyword.toLowerCase()
        state.proj.filter(
          (obj) => obj.Name.toLowerCase().includes(kw)
        )
      } else {
        return state.proj;
      }
    },
  },
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