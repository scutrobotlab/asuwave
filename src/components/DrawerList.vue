<template>
  <v-navigation-drawer v-model="drawer" width="310" app>
    <SerialPort />
    <v-divider />

    <v-list>
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title>读取变量列表</v-list-item-title>
          <!-- <v-list-item-subtitle>TODO文件路径</v-list-item-subtitle> -->
        </v-list-item-content>
        <v-spacer />
        <v-list-item-icon>
          <v-btn icon @click="openDialog()">
            <v-icon>mdi-file-upload</v-icon>
          </v-btn>
        </v-list-item-icon>
      </v-list-item>
    </v-list>
    <v-divider />

    <VariableReadList />
    <v-divider />
    <VariableModiList />
    <v-divider />
    <SaveConfig />
    <VariableAllDialog ref="VariableAllDialog" />
  </v-navigation-drawer>
</template>

<script>
import SerialPort from "@/components/SerialPort.vue";
import SaveConfig from "@/components/SaveConfig.vue";
import VariableModiList from "@/components/VariableModiList.vue";
import VariableReadList from "@/components/VariableReadList.vue";
import VariableAllDialog from "@/components/VariableAllDialog.vue";
export default {
  components: {
    SerialPort,
    SaveConfig,
    VariableModiList,
    VariableReadList,
    VariableAllDialog,
  },
  data: () => ({
    drawer: null,
  }),
  methods: {
    switchDrawer() {
      this.drawer = !this.drawer;
    },
    openDialog() {
      this.$refs.VariableAllDialog.openDialog();
    },
  },
};
</script>
