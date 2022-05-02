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
        <v-alert type="success" :value="alert">
          添加成功！
        </v-alert>
        <!-- <v-tabs vertical>
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
          <v-tab-item> -->
        <v-file-input v-model="file" label="上传elf或者axf文件" />
        <!-- </v-tab-item>
          <v-tab-item>
            <v-file-input v-model="file" label="上传elf或者axf文件" />
          </v-tab-item>
        </v-tabs> -->
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
        <v-data-table :headers="headers" :items="searchData">
          <template #item.isRead="{ item }">
            <v-btn
              icon color="green"
              @click="openVariableDialog(item.Name, item.Type, null, item.Addr, 'read')"
            >
              <v-icon>mdi-eye</v-icon>
            </v-btn>
          </template>
          <template #item.isModi="{ item }">
            <v-btn
              icon color="green"
              @click="openVariableDialog(item.Name, item.Type, null, item.Addr, 'modi')"
            >
              <v-icon>mdi-pen</v-icon>
            </v-btn>
          </template>
        </v-data-table>
        <VariableNewDialog ref="VariableNewDialog" :opt="opt" />
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { postVariableToProj, deleteVariableAll } from "@/api/variable.js"; //postVariable,
import VariableNewDialog from "@/components/VariableNewDialog.vue";

export default {
  components: {
    VariableNewDialog,
  },
  mixins: [errorMixin],
  data: () => ({
    dialog: false,
    file: null,
    keyword: "",
    reaction: "",
    alert: false,
    opt: "",
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
        value: "isModi"
      },
    ]
  }),
  computed: {
    lists() {
      return this.$store.state.variables.variables.proj;
    },
    searchData() {
      return this.$store.getters['variables/searchVToProj'](this.keyword)
    },
  },
  watch: {
    file: async function () {
      await this.errorHandler(postVariableToProj(this.file));
      await this.getVariableList();
    },
  },
  async mounted() {
    await this.getVariableList();
    this.$bus.$on("sendalert", (data) => {
      this.alert = data;
      setTimeout(this.realert, 1000);
    });
  },
  methods: {
    openDialog() {
      this.dialog = true;
    },
    openVariableDialog(name, type, board, addr, opt) {
      this.opt = opt;
      this.$refs.VariableNewDialog.openDialogFromList(name, type, board, addr);
    },
    async closeDialog() {
      await this.$store.dispatch("variables/getV", "read");
      await this.$store.dispatch("variables/getV", "modi");
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
    realert() {
      this.alert = false;
    },
  },
};
</script>
