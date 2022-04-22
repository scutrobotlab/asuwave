<template>
  <v-list class="mb-8">
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
    <v-list-item dense v-for="i in variables" :key="i.Name">
        <v-list-item-avatar size="20" :color="i.Inputcolor"/>
      <v-list-item-content>
        <v-list-item-title>
          <span class="green--text">{{ i.Type }}</span>
          &nbsp;
          {{ i.Name }}
        </v-list-item-title>
        <v-list-item-subtitle>{{ hexdsp(i.Addr) }}</v-list-item-subtitle>
      </v-list-item-content>

      <v-list-item-action>
        <v-btn x-small icon v-on:click="delVariable(i)">
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </v-list-item-action>
    </v-list-item>
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
