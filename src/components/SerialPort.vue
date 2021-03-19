<template>
  <v-list>
    <v-list-item>
      <v-list-item-title>串口</v-list-item-title>
    </v-list-item>
    <ErrorAlert v-model="error" />
    <v-list-item>
      <v-list-item-action>
        <v-switch v-model="status" v-on:change="optSerial()" inset></v-switch>
      </v-list-item-action>
      <v-select
        :items="serialList"
        v-model="serial"
        v-on:click="getSerialList()"
        v-bind:disabled="status"
        label="选择串口"
      ></v-select>
    </v-list-item>
  </v-list>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { getSerial, getSerialCur, postSerialCur, deleteSerialCur } from "@/api/serial.js";
export default {
  mixins: [errorMixin],
  data: () => ({
    status: false,
    serial: null,
    serialList: [],
  }),
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
        this.status = true;
      }
    },
    optSerial() {
      if (this.status) {
        this.errorHandler(postSerialCur(this.serial)).catch(() => {
          this.status = false;
        });
      } else {
        this.errorHandler(deleteSerialCur()).catch(() => {
          this.status = true;
        });
      }
    },
  },
};
</script>
