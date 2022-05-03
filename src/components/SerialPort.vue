<template>
  <v-list>
    <v-list-item>
      <v-list-item-title>串口</v-list-item-title>
      <v-spacer />
      <v-switch 
        v-model="status" inset 
        :disabled="!valid"
        @change="optSerial()"
      />
    </v-list-item>
    <ErrorAlert v-model="error" />
    <v-list-item dense>
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
    serial: "",
    serialList: [],
    baud: "",
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
    valid() {
      return this.serial!="" && this.serial!=null &&
            this.baud!="" && this.baud!=null 
    }
  },
  mounted() {
    Promise.all([this.getSerialList(), this.getSerial()]);
  },
  methods: {
    async getSerialList() {
      this.serialList = await this.errorHandler(getSerial());
    },
    async getSerial() {
      let serialCur = await this.errorHandler(getSerialCur());
      this.serial = serialCur.Serial;
      this.baud = serialCur.Baud;
      if (serialCur.Serial) {
        this.$store.commit("serialPort/setStatus", true);
      }
    },
    optSerial() {
      this.errorHandler(
        this.status?
          postSerialCur(this.serial, this.baud):
          deleteSerialCur()
      ).catch(() => {
        this.$store.commit("serialPort/setStatus", false);
      })
    },
  },
};
</script>
