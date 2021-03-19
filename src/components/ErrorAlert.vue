<template>
  <v-expand-transition>
    <v-alert v-if="error" text dense :prominent="critical" type="error" :icon="alert_icon">
      <v-row align="center">
        <v-col class="grow" :class="critical || 'py-0'">
          <div :class="!critical || 'title'">{{ title }}</div>
          <div v-if="critical">错误代码：{{ error.status }}</div>
        </v-col>
        <v-col class="shrink" @click="action" :class="critical || 'py-0'">
          <v-btn outlined color="error" v-if="!critical && action_text">{{ action_text }}</v-btn>
          <v-btn icon color="error" v-else>
            <v-icon @click="close">mdi-close-circle</v-icon>
          </v-btn>
        </v-col>
      </v-row>
    </v-alert>
  </v-expand-transition>
</template>

<script>
export default {
  model: {
    prop: "error",
    event: "close",
  },
  props: {
    error: { Type: Object },
    critical: Boolean,
  },
  computed: {
    title() {
      if (this.error.data) return this.error.data;
      switch (this.error.status) {
        case 401:
          return "身份信息已过期。";
        case 404:
          return "找不到链接。";
        case 503:
          return "服务暂时不可用。";
        default:
          return "出现错误。";
      }
    },
    alert_icon() {
      if (!this.critical) return false;
      switch (this.error.status) {
        case 401:
          return "mdi-account-alert";
        case 404:
          return "mdi-link-variant-off";
        case 503:
          return "mdi-server-network-off";
        default:
          return "mdi-alert-circle";
      }
    },
    action_text() {
      switch (this.error.status) {
        case 401:
          return "登录";
        case 404:
          return "返回";
        case 503:
          return "刷新";
        default:
          return null;
      }
    },
  },
  methods: {
    action() {
      switch (this.error.status) {
        case 401:
          this.$router.push("/");
          return;
        case 404:
          this.$router.go(-1);
          return;
        case 503:
          this.$router.go();
          return;
        default:
          this.close();
          return;
      }
    },
    close() {
      this.$emit("close", null);
    },
  },
};
</script>

<style></style>
