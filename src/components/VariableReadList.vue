<template>
  <v-list class="mb-8">
    <v-list-item>
      <v-list-item-title>只读变量</v-list-item-title>
      <v-spacer></v-spacer>
      <v-list-item-icon>
        <v-btn icon v-on:click="openDialog()">
          <v-icon>mdi-plus-circle</v-icon>
        </v-btn>
      </v-list-item-icon>
    </v-list-item>
    <ErrorAlert v-model="error" />
    <v-list-item dense v-for="(i, Addr) in variables" :key="Addr" style="font-family: monospace">
      <v-list-item-avatar size="20" :color="i.Inputcolor" />
      <v-list-item-content>
        <v-list-item-title>
          <span class="green--text">{{ i.Type }}</span> {{ i.Name }}
          <span class="text--disabled"> -> {{ hexdsp(i.Addr) }}</span
          >;
        </v-list-item-title>
        <v-list-item-subtitle>
          <span style="color: #ff08;">return </span>
          <span v-if="i.SignalGain != 1" class="text--disabled"> {{ i.SignalGain }} * </span>
          <span>{{ i.Name }}</span>
          <span v-if="i.SignalBias != 0" class="text--disabled"> + {{ i.SignalBias }} </span>;
        </v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-action>
        <v-btn small icon v-on:click="delVariable(i)">
          <v-icon small>mdi-close</v-icon>
        </v-btn>
      </v-list-item-action>
    </v-list-item>
    <VariableNewDialog ref="VariableNewDialog" opt="read" />
  </v-list>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { deleteVariable } from "@/api/variable.js";
import VariableNewDialog from "@/components/VariableNewDialog.vue";
export default {
  mixins: [errorMixin],
  components: {
    VariableNewDialog,
  },
  computed: {
    variables() {
      return this.$store.state.variables.variables.read;
    },
    serial_status() {
      return this.$store.state.serialPort.status;
    },
  },
  async mounted() {
    this.getVariables();
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
      await this.$store.dispatch("variables/getV", "read");
    },
    async delVariable(i) {
      await this.errorHandler(deleteVariable("read", 1, i.Name, i.Type, i.Addr));
      await this.getVariables();
    },
  },
};
</script>
