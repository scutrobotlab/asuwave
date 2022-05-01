import Vue from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";
import store from "./store";

import "@mdi/font/css/materialdesignicons.css";
import "@fontsource/roboto";

Vue.config.productionTip = false;

new Vue({
  vuetify,
  store,
  beforeCreate() {
    Vue.prototype.$bus = this;
  },
  render: (h) => h(App),
}).$mount("#app");
