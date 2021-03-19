<template>
  <v-list>
    <v-list-item>
      <v-list-item-title>变量·{{ showtext }}</v-list-item-title>
      <v-spacer></v-spacer>
      <v-list-item-icon>
        <v-btn icon v-on:click="openDialog()">
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </v-list-item-icon>
    </v-list-item>
    <ErrorAlert v-model="error" />
    <v-list-group v-for="i in variables" :key="i.Name">
      <template v-slot:activator>
        <v-list-item-icon>
          <v-icon>mdi-variable</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>{{ i.Name }}</v-list-item-title>
        </v-list-item-content>
      </template>

      <v-list-item>
        <v-list-item-icon>
          <v-icon>mdi-tag-multiple</v-icon>
        </v-list-item-icon>
        <v-list-item-content>{{ i.Type }}</v-list-item-content>
        <v-btn icon absolute small right v-on:click="delVariable(i)">
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </v-list-item>
      <v-list-item>
        <v-list-item-icon>
          <v-icon>mdi-view-list</v-icon>
        </v-list-item-icon>
        <v-list-item-content>{{ hexdsp(i.Addr) }}</v-list-item-content>
      </v-list-item>
    </v-list-group>
    <VariableNewDialog ref="VariableNewDialog" v-bind:opt="opt" />
  </v-list>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { deleteVariable } from "@/api/variable.js";
import VariableNewDialog from "@/components/VariableNewDialog.vue";
export default {
  mixins: [errorMixin],
  props: ["showtext", "opt"],
  components: {
    VariableNewDialog,
  },
  computed: {
    variables() {
      return this.$store.state.variables[this.opt];
    },
  },
  async mounted() {
    await this.getVariables();
  },
  methods: {
    openDialog() {
      this.$refs.VariableNewDialog.openDialog();
    },
    hexdsp(i) {
      var h = i.toString(16);
      var l = h.length;
      var z = new Array(9 - l).join("0");
      return "0x" + z + h;
    },
    async getVariables() {
      await this.$store.dispatch("getV", this.opt);
    },
    async delVariable(i) {
      await this.errorHandler(deleteVariable(this.opt, 1, i.Name, i.Type, i.Addr));
      await this.getVariables();
    },
  },
};
</script>
