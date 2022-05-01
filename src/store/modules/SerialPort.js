export default {
  namespaced: true,
  state: {
    status: false,
  },
  getters: {},

  mutations: {
    setStatus(state, i) {
      state.status = i;
    },
  },
  actions: {},
  modules: {},
};