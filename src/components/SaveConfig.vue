<template>
  <v-list>
    <v-list-item>
      <v-list-item-title>保存选项</v-list-item-title>
    </v-list-item>
    <ErrorAlert v-model="error" />
    <v-list-item v-for="i in config" :key="i.t" dense>
      <v-switch
        v-model="i.v" dense :label="i.t"
        inset @change="updateConfig(i)"
      />
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
        t: "监控文件",
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
