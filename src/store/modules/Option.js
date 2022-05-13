import { getOption, setOption } from "@/api/option.js"; //postVariable,

export default {
  namespaced: true,
  state: {
    LogLevel: 3,
    SaveVarList: false,
    SaveFilePath: true,
    UpdateByProj: true
  },
  getters: {},
  mutations: {},
  actions: {
    async get ({ state }) {
      await getOption().then((r)=>{
        state = r
      });
    },
    async set ({ state }, k, v) {
      await setOption(k, v).then(()=>{
        state.k = v;
      });
    },
  },
  modules: {},
};