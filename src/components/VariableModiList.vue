<template>
  <v-list class="mb-8">
    <v-list-item>
      <v-list-item-title>可写变量</v-list-item-title>
      <v-spacer></v-spacer>
      <v-list-item-icon>
        <v-btn icon v-on:click="openDialog()">
          <v-icon>mdi-plus-circle</v-icon>
        </v-btn>
      </v-list-item-icon>
    </v-list-item>
    <ErrorAlert v-model="error" />
    <v-list-item v-for="i in variables" :key="i.Addr" class="mb-2">
      <v-text-field
        style="font-family: monospace"
        dense
        :label="i.Type + ' ' + i.Name + ' ='"
        :hint="hexdsp(i.Addr)"
        append-icon="mdi-send"
        v-model="i.Data"
        type="number"
        :disabled="!serial_status"
        @click:append="modiVariable(i)"
      >
      </v-text-field>
      <v-list-item-action>
        <v-btn small icon v-on:click="delVariable(i)">
          <v-icon small>mdi-close</v-icon>
        </v-btn>
      </v-list-item-action>
    </v-list-item>
    <VariableNewDialog ref="VariableNewDialog" opt="modi" />
  </v-list>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { putVariable, deleteVariable } from "@/api/variable.js";
import VariableNewDialog from "@/components/VariableNewDialog.vue";
export default {
  mixins: [errorMixin],
  components: {
    VariableNewDialog,
  },
  computed: {
    variables() {
      return this.$store.state.variables.variables.modi;
    },
    serial_status() {
      return this.$store.state.serialPort.status;
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
      await this.$store.dispatch("variables/getV", "modi");
    },
    async delVariable(i) {
      await this.errorHandler(deleteVariable("modi", 1, i.Name, i.Type, i.Addr));
      await this.getVariables();
    },

    async modiVariable(i) {
      await this.errorHandler(putVariable("modi", 1, i.Name, i.Type, i.Addr, parseFloat(i.Data)));
    },
  },
};
</script>
