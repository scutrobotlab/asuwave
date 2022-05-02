<template>
  <v-list>
    <v-list-item>
      <v-list-item-title>串口</v-list-item-title>
      <v-spacer />
      <v-switch v-model="status" inset @change="optSerial()" />
    </v-list-item>
    <ErrorAlert v-model="error" />
    <v-list-item>
      <v-list-item-content>
        <v-select
          v-model="serial"
          :items="serialList"
          :disabled="status"
          label="选择串口"
          @click="getSerialList()"
        />
        <v-select
          v-model="baud"
          :items="baudList"
          :disabled="status"
          label="波特率"
        />
      </v-list-item-content>
    </v-list-item>
  </v-list>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { getSerial, getSerialCur, postSerialCur, deleteSerialCur } from "@/api/serial.js";
import baudRate from "@/const/BaudRate.json";
export default {
  mixins: [errorMixin],
  data: () => ({
    serial: null,
    serialList: [],
    baud: null,
    baudList: baudRate
  }),
  computed: {
    status: {
      get() {
        return this.$store.state.serialPort.status;
      },
      set(val) {
        this.$store.commit("serialPort/setStatus", val);
      },
    },
  },
  mounted() {
    Promise.all([this.getSerialList(), this.getSerial()]);
  },
  methods: {
    async getSerialList() {
      this.serialList = await this.errorHandler(getSerial());
    },
    async getSerial() {
      this.serial = await this.errorHandler(getSerialCur());
      if (this.serial) {
        this.$store.commit("serialPort/setStatus", true);
      }
    },
    optSerial() {
      if (this.status) {
        this.errorHandler(postSerialCur(this.serial)).catch(() => {
          this.$store.commit("serialPort/setStatus", false);
        });
      } else {
        this.errorHandler(deleteSerialCur()).catch(() => {
          this.$store.commit("serialPort/setStatus", true);
        });
      }
    },
  },
};
</script>
