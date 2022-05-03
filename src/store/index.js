import Vue from "vue";
import Vuex from "vuex";
import File from "./modules/File";
import Variables from "./modules/Variables";
import SerialPort from "./modules/SerialPort";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    file: File,
    variables: Variables,
    serialPort: SerialPort,
  },
});