<template>
  <v-navigation-drawer width="300" v-model="drawer" app>
    <SerialPort />
    <v-divider />

    <v-list>
      <v-list-item>
        <v-list-item-title>读取变量列表</v-list-item-title>
        <v-spacer></v-spacer>
        <v-list-item-icon>
          <v-btn icon v-on:click="openDialog()">
            <v-icon>mdi-file-upload</v-icon>
          </v-btn>
        </v-list-item-icon>
      </v-list-item>
    </v-list>
    <v-divider />

    <VariableList showtext="观察" opt="read" />
    <v-divider />
    <VariableList showtext="修改" opt="modi" />
    <v-divider />
    <SaveConfig />
    <VariableAllDialog ref="VariableAllDialog" />
  </v-navigation-drawer>
</template>

<script>
import SerialPort from "@/components/SerialPort.vue";
import SaveConfig from "@/components/SaveConfig.vue";
import VariableList from "@/components/VariableList.vue";
import VariableAllDialog from "@/components/VariableAllDialog.vue";
export default {
  components: {
    SerialPort,
    SaveConfig,
    VariableList,
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
