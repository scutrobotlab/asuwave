import { getOption, setOption } from "@/api/option.js"; //postVariable,

export default {
  namespaced: true,
  state: {
    LogLevel: 0,
    SaveVarList: false,
    SaveFilePath: false,
    UpdateByProj: false
  },
  getters: {},
  mutations: {
    all(state, s) {
      state.LogLevel = s.LogLevel
      state.SaveVarList = s.SaveVarList
      state.SaveFilePath = s.SaveFilePath
      state.UpdateByProj = s.UpdateByProj
    },
    kv(state, {k, v}) {
      state[k] = v
    }
  },
  actions: {
    async get ({ commit }) {
      await getOption().then((r)=>{
        commit("all", r)
      });
    },
    async set ({ commit }, {k, v}) {
      window.console.log(k, v);
      await setOption(k, v).then(()=>{
        commit("kv", {k, v});
      });
    },
  },
  modules: {},
};