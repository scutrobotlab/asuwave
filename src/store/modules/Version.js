import { getVersion } from "@/api/version.js";

export default {
  namespaced: true,
  state: {
    os: "",
    arch: "",
    current_tag: "",
    update: {
      error: false,
      checking: true,
      response: {
        tag_name: "",
        body: "",
        assets: [],
      },
    },
  },
  getters: {
    asset(state) {
      return state.update.response.assets.find((asset) => {
        return (
          asset.browser_download_url.includes(state.os) &&
          asset.browser_download_url.includes(state.arch)
        );
      });
    },
    NewVersion(state) {
      return (state.update.response.tag_name != "" && state.current_tag != state.update.response.tag_name);
    }
  },
  mutations: {
    Init(state) {
      state.current_tag = process.env.VUE_APP_GITTAG;
      window.console.log(state.current_tag);
  
      if (/Win|win/i.test(navigator.userAgent)) state.os = "windows";
      else if (/Mac|mac|darwin/i.test(navigator.userAgent)) state.os = "darwin";
      else if (/linux|gnu/i.test(navigator.userAgent)) state.os = "linux";
  
      if (/(?:(amd|x(?:(?:86|64)[-_])?|wow|win)64)[;)]/i.test(navigator.userAgent))
        state.arch = "amd64";
      else if (/\b(aarch64|arm(v?8e?l?|_?64))\b/i.test(navigator.userAgent)) state.arch = "arm64";
  
      window.console.log(navigator.userAgent);
    },
  },
  actions: {
    async Init({dispatch, commit}) {
      commit("Init");
      dispatch("CheckUpdate")
    },
    async CheckUpdate({state}) {
      state.update.checking = true;
      state.update.response = await getVersion();
      state.update.checking = false;
    },
  },
  modules: {},
};