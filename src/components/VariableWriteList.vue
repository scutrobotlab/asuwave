<template>
  <v-list class="mb-8">
    <v-list-item>
      <v-list-item-title>可写变量</v-list-item-title>
      <v-spacer />
      <v-list-item-icon>
        <v-btn icon @click="openDialog()">
          <v-icon>mdi-plus-circle</v-icon>
        </v-btn>
      </v-list-item-icon>
    </v-list-item>
    <ErrorAlert v-model="error" />
    <v-list-item v-for="(i, Addr) in variables" :key="Addr" class="mb-2">
      <v-text-field
        v-model="i.Data"
        style="font-family: monospace"
        dense
        :label="i.Type + ' ' + i.Name + ' ='"
        :hint="hexdsp(i.Addr)"
        append-icon="mdi-send"
        type="number"
        :disabled="!serial_status"
        @click:append="writeVariable(i)"
      />
      <v-list-item-action>
        <v-btn small icon @click="delVariable(i)">
          <v-icon small>
            mdi-close
          </v-icon>
        </v-btn>
      </v-list-item-action>
    </v-list-item>
    <VariableNewDialog ref="VariableNewDialog" mod="write" />
  </v-list>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { putVariable, deleteVariable } from "@/api/variable.js";
import VariableNewDialog from "@/components/VariableNewDialog.vue";
export default {
  components: {
    VariableNewDialog,
  },
  mixins: [errorMixin],
  computed: {
    variables() {
      return this.$store.state.variables.write;
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
      await this.$store.dispatch("variables/getV", "write");
    },
    async delVariable(i) {
      await this.errorHandler(deleteVariable("write", 1, i.Name, i.Type, i.Addr));
      await this.getVariables();
    },

    async writeVariable(i) {
      await this.errorHandler(putVariable("write", 1, i.Name, i.Type, i.Addr, parseFloat(i.Data)));
    },
  },
};
</script>
