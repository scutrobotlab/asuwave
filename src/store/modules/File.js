import { uploadFile, setFilePath } from "@/api/file.js"; //postVariable,

export default {
  namespaced: true,
  state: {
    path: "",
    upload: ""
  },
  getters: {
    fileStatus (state) {
      if (state.path != "") {
        return "正在监听：" + state.path
      } else if (state.upload != "") {
        return "已导入：" + state.upload
      } else {
        return "未设置文件"
      }
    }
  },
  
  mutations: {},
  actions: {
    async setUpload ({ state }, f) {
      await uploadFile(f).then(()=>{
        state.upload = f.name;
        state.path = "";
      });
    },
    async setPath ({ state }, f) {
      await setFilePath(f).then(()=>{
        state.upload = "";
        state.path = f;
      });
    },
  },
  modules: {},
};