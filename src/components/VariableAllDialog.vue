<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
      <v-card>
        <v-toolbar dark color="primary">
          <v-toolbar-items>
            <v-btn icon dark v-on:click="closeDialog">
              <v-icon>mdi-close</v-icon>
            </v-btn>
          </v-toolbar-items>
          <v-toolbar-title>变量列表</v-toolbar-title>
        </v-toolbar>
        <v-alert type="success" :value="alert">添加成功！ </v-alert>
        <v-row>
          <v-col cols="4">
            <v-file-input label="上传elf或者axf文件" v-model="file"></v-file-input>
          </v-col>
          <v-col cols="7">
            <v-text-field
              clearable
              placeholder="搜索变量"
              prepend-icon="mdi-magnify"
              v-model="keyword"
            ></v-text-field>
          </v-col>
          <v-col cols="1">
            <v-btn class="ma-2" outlined color="primary" @click="deleteVariable()">
              一键删除
            </v-btn>
          </v-col>
        </v-row>

        <ErrorAlert v-model="error" />
        <v-simple-table dense fixed-header height="780px">
          <template v-slot:default>
            <thead>
              <tr>
                <th>名称</th>
                <th>类型</th>
                <th>地址</th>
                <th>观察</th>
                <th>修改</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="list in searchData" :key="list.Name">
                <td>{{ list.Name }}</td>
                <td>{{ list.Type }}</td>
                <td>{{ list.Addr }}</td>
                <td>
                  <v-btn
                    icon
                    v-on:click="openVariableDialog(list.Name, list.Type, 1, list.Addr, 'read')"
                  >
                    <v-icon>mdi-plus</v-icon>
                    {{ reaction }}
                  </v-btn>
                </td>
                <td>
                  <v-btn
                    icon
                    v-on:click="openVariableDialog(list.Name, list.Type, 1, list.Addr, 'modi')"
                  >
                    <v-icon>mdi-plus</v-icon>
                  </v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
        <VariableNewDialog ref="VariableNewDialog" v-bind:opt="opt" />
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { postVariableToProj, deleteVariableAll } from "@/api/variable.js"; //postVariable,
import VariableNewDialog from "@/components/VariableNewDialog.vue";

export default {
  mixins: [errorMixin],
  data: () => ({
    dialog: false,
    file: null,
    keyword: "",
    reaction: "",
    alert: false,
    opt: "",
  }),
  components: {
    VariableNewDialog,
  },
  computed: {
    lists() {
      return this.$store.state.variables.variables["proj"];
    },
    searchData() {
      if (this.$store.getters.variables.searchVToProj(this.keyword) == null) return null;
      return this.$store.getters.variables.searchVToProj(this.keyword).slice(0, 200);
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
