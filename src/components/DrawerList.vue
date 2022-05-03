<template>
  <v-navigation-drawer v-model="drawer" width="310" app>
    <SerialPort />
    <v-divider />
    <VariableReadList />
    <v-divider />
    <VariableModiList />
    <v-divider />
    <SaveConfig />
  </v-navigation-drawer>
</template>

<script>
import SerialPort from "@/components/SerialPort.vue";
import SaveConfig from "@/components/SaveConfig.vue";
import VariableModiList from "@/components/VariableModiList.vue";
import VariableReadList from "@/components/VariableReadList.vue";
export default {
  components: {
    SerialPort,
    SaveConfig,
    VariableModiList,
    VariableReadList,
  },
  data: () => ({
    drawer: null,
  }),
  methods: {
    switchDrawer() {
      this.drawer = !this.drawer;
    },
  },
};
</script>
