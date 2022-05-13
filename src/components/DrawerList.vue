<template>
  <v-navigation-drawer v-model="drawer" width="310" app>
    <SerialPort />
    <v-divider />
    <VariableReadList />
    <v-divider />
    <VariableWriteList />
  </v-navigation-drawer>
</template>

<script>
import SerialPort from "@/components/SerialPort.vue";
import VariableWriteList from "@/components/VariableWriteList.vue";
import VariableReadList from "@/components/VariableReadList.vue";
export default {
  components: {
    SerialPort,
    VariableWriteList,
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
