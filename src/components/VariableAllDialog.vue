<template>
  <v-row justify="center">
    <v-dialog
      v-model="dialog" fullscreen hide-overlay
      transition="dialog-bottom-transition"
    >
      <v-card>
        <v-toolbar dark color="primary">
          <v-toolbar-items>
            <v-btn icon dark @click="closeDialog">
              <v-icon>mdi-close</v-icon>
            </v-btn>
          </v-toolbar-items>
          <v-toolbar-title>变量列表</v-toolbar-title>
        </v-toolbar>
        <v-alert type="success" :value="alert!=null">
          {{ alert }}
        </v-alert>
        <v-tabs vertical background-color="primary">
          <v-tab>
            <v-icon left>
              mdi-file-upload-outline
            </v-icon>
            上传文件
          </v-tab>
          <v-tab>
            <v-icon left>
              mdi-book-sync-outline 
            </v-icon>
            监控文件
          </v-tab>
          <v-tab-item>
            <v-card flat style="background:#88F2" class="px-3">
              <v-card-text class="pb-0">
                通过文件资源管理器查找并上传elf或者axf文件，解析返回工程变量。上传文件会导致监控的文件路径被清除。
              </v-card-text>
              <v-file-input
                v-model="file" label="上传elf或者axf文件"
                append-outer-icon="mdi-send" @click:append-outer="uploadFile"
              />
            </v-card>
          </v-tab-item>
          <v-tab-item>
            <v-card flat style="background:#88F2" class="px-3">
              <v-card-text class="pb-0">
                指定后端设备上的文件路径，监控文件变化，并在每次文件变化时更新工程变量。
              </v-card-text>
              <v-text-field 
                v-model="filepath" 
                prepend-icon="mdi-file-link-outline" 
                label="输入elf或者axf文件路径"
                append-outer-icon="mdi-send"
                @click:append-outer="watchFile"
              />
            </v-card>
          </v-tab-item>
        </v-tabs>
        <v-row class="mx-1">
          <v-col cols="8" lg="10">
            <v-text-field
              v-model="keyword"
              clearable
              placeholder="搜索变量"
              prepend-icon="mdi-magnify"
            />
          </v-col>
          <v-col cols="4" lg="2">
            <v-btn
              class="my-3" block outlined
              @click="deleteVariable()"
            >
              清空
            </v-btn>
          </v-col>
        </v-row>

        <ErrorAlert v-model="error" />
        <v-data-table
          :headers="headers"
          :items="proj_vars"
          :search="keyword"
          group-by="group"
        >
          <template #[`item.isRead`]="{ item }">
            <v-btn
              icon :color="item.isRead?'green':'grey'"
              @click="toggleVar(item.isRead, item, 'read')"
            >
              <v-icon>mdi-eye{{ item.isRead?"":"-plus-outline" }}</v-icon>
            </v-btn>
          </template>
          <template #[`item.isWrite`]="{ item }">
            <v-btn
              icon :color="item.isWrite?'green':'grey'"
              @click="toggleVar(item.isWrite, item, 'write')"
            >
              <v-icon>mdi-pencil{{ item.isWrite?"":"-plus-outline" }}</v-icon>
            </v-btn>
          </template>
        </v-data-table>
        <VariableNewDialog ref="VariableNewDialog" :mod="mod" />
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { getVariable, deleteVariable, deleteVariableAll } from "@/api/variable.js"; //postVariable,
import VariableNewDialog from "@/components/VariableNewDialog.vue";

export default {
  components: {
    VariableNewDialog,
  },
  mixins: [errorMixin],
  data: () => ({
    dialog: false,
    file: null,
    filepath: null,
    keyword: "",
    reaction: "",
    alert: null,
    mod: "",
    ws: null,
    headers:[
      {
        text:"变量名称",
        value: "Name"
      },
      {
        text:"变量地址",
        value: "Addr"
      },
      {
        text:"变量类型",
        value: "Type"
      },
      {
        text:"只读变量",
        value: "isRead"
      },
      {
        text:"可写变量",
        value: "isWrite"
      },
    ]
  }),
  computed: {
    proj_vars() {
      this.alert; //触发更新
      return this.$store.state.variables.proj.map((p)=>{
        let addr = parseInt(p.Addr, 16)
        p.isRead = (this.$store.state.variables.read[addr] != null);
        p.isWrite = (this.$store.state.variables.write[addr] != null);
        p.group = p.Name.split(".")[0];
        return p;
      })
    },
  },
  async mounted() {
    this.$store.dispatch("file/refreshPath");
    this.initWS();
    await this.getVariableList();
    this.$bus.$on("sendalert", (data) => {
      this.alert = "添加成功";
      setTimeout(this.realert, 1000);
    });
  },
  destroyed() {
    this.ws.close();
  },
  methods: {
    async toggleVar(state, item, mod) {
      if (state) {
        await deleteVariable(mod, 1, item.Name, item.Type, parseInt(item.Addr, 16));
        await this.$store.dispatch("variables/getV", mod);
        this.alert = "删除成功";
        setTimeout(this.realert, 1000);
      }else {
        this.openVariableDialog(item.Name, item.Type, 1, item.Addr, mod);
      }
    },
    async uploadFile() { 
      await this.$store.dispatch('file/setUpload', this.file)
      await this.getVariableList();
    },
    async watchFile() { 
      await this.$store.dispatch('file/setPath', this.filepath)
      await this.getVariableList();
    },
    openDialog() {
      this.dialog = true;
    },
    openVariableDialog(name, type, board, addr, mod) {
      this.mod = mod;
      this.$refs.VariableNewDialog.openDialogFromList(name, type, board, addr);
    },
    async closeDialog() {
      await this.$store.dispatch("variables/getV", "read");
      await this.$store.dispatch("variables/getV", "write");
      this.dialog = false;
    },
    async getVariableList() {
      await this.$store.dispatch("variables/getV", "proj");
    },
    // async variableAdd(mode, i) {
    //   await this.errorHandler(postVariable(mode, 1, i.Name, i.Type, parseInt(i.Addr, 16)));
    //   this.alert = true;
    //   setTimeout(this.realert, 1000);
    // },
    async deleteVariable() {
      await this.errorHandler(deleteVariableAll());
      this.getVariableList();
    },
    initWS() {
      let url = "";
      let path = "/filews"
      if (process.env.NODE_ENV === "production") {
        url =
        (document.location.protocol == "https:" ? "wss" : "ws") +
        "://" +
        window.location.host + path;
      } else {
        url = "ws://localhost:8000" + path;
      }
      this.ws = new WebSocket(url);
      this.ws.onopen = (() => {window.console.log("file: 连接成功")});
      this.ws.onclose = (() => {window.console.log("file: 连接断开")});
      this.ws.onmessage = ((e) => {
        window.console.log(e.data)
        this.$store.dispatch("variables/getV", "read");
        this.$store.dispatch("variables/getV", "write");
        this.$store.dispatch("variables/getV", "proj");
      });
      this.ws.onerror = ((e) => { 
        this.$store.commit("file/setError", e.data);
      });
    },
    realert() {
      this.alert = null;
    },
  },
};
</script>
