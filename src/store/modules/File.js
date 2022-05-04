import { uploadFile, setFilePath, getFilePath} from "@/api/file.js"; //postVariable,

export default {
  namespaced: true,
  state: {
    path: "",
    upload: "",
    error: null
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
  
  mutations: {
    setError (state, err) {
      state.error = err
    }
  },
  actions: {
    async refreshPath ({ state }) {
      window.console.log("!!!")
      await getFilePath().then((r)=>{
        window.console.log(r);
        if (r.length != 0) {
          state.upload = "";
          state.path = r;
        }
      });
    },
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