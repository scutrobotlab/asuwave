<template>
  <v-list>
    <v-list-item>
      <v-list-item-title>保存选项</v-list-item-title>
    </v-list-item>
    <ErrorAlert v-model="error" />
    <v-list-item dense v-for="i in config" :key="i.t">
      <v-switch dense v-model="i.v" :label="i.t" v-on:change="updateConfig(i)" inset></v-switch>
    </v-list-item>
  </v-list>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { getOption, putOption } from "@/api/option.js";
export default {
  mixins: [errorMixin],
  data: () => ({
    save: 7,
    config: [
      {
        t: "变量列表",
        i: 1,
        v: true,
      },
      {
        t: "观察变量",
        i: 2,
        v: true,
      },
      {
        t: "修改变量",
        i: 4,
        v: true,
      },
    ],
  }),
  mounted() {
    this.getConfig();
  },
  methods: {
    async getConfig() {
      this.save = await this.errorHandler(getOption());
      for (var i = 0; i < 3; i++) {
        this.config[i].v = this.checkCanSave(this.config[i].i);
      }
    },
    async updateConfig(i) {
      i.v ? (this.save += i.i) : (this.save -= i.i);
      await this.errorHandler(putOption(this.save));
    },
    checkCanSave(i) {
      return (this.save & i) == i;
    },
  },
};
</script>
