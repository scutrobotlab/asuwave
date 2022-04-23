<template>
  <v-app id="inspire">
    <DrawerList ref="DrawerList" />

    <v-app-bar app dark color="primary">
      <v-app-bar-nav-icon @click.stop="switchDrawer()"></v-app-bar-nav-icon>
      <v-toolbar-title>坠好用的上位机</v-toolbar-title>
      <v-btn icon v-on:click="openDialog()">
        <v-icon>mdi-information-outline</v-icon>
      </v-btn>
      <v-spacer></v-spacer>
      <v-switch
        v-model="$vuetify.theme.dark"
        hide-details
        inset
        color="black"
        v-bind:label="$vuetify.theme.dark ? '深色模式' : '浅色模式'"
      ></v-switch>
    </v-app-bar>

    <v-main>
      <ChartCard />
    </v-main>

    <v-footer app color="primary">
      <span class="white--text">华工机器人实验室</span>
    </v-footer>

    <AboutDialog ref="AboutDialog" />
  </v-app>
</template>

<script>
import DrawerList from "@/components/DrawerList.vue";
import AboutDialog from "@/components/AboutDialog.vue";
import ChartCard from "@/components/ChartCard.vue";

export default {
  components: {
    DrawerList,
    AboutDialog,
    ChartCard,
  },
  props: {
    source: String,
  },
  async mounted() {
    this.$vuetify.theme.dark = window.matchMedia("(prefers-color-scheme: dark)").matches;
    await this.$store.dispatch("variables/getVType");
  },
  methods: {
    switchDrawer() {
      this.$refs.DrawerList.switchDrawer();
    },
    openDialog() {
      this.$refs.AboutDialog.openDialog();
    },
  },
};
</script>
