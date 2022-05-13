<template>
  <v-app id="inspire">
    <DrawerList ref="DrawerList" />

    <v-app-bar app dark color="primary">
      <v-app-bar-nav-icon @click.stop="switchDrawer()" />
      <v-toolbar-title>坠好用的上位机</v-toolbar-title>
      <v-btn icon @click="openAboutDialog()">
        <v-icon v-if="NewVersion" color="warning">
          mdi-alert-decagram
        </v-icon>
        <v-icon v-else>
          mdi-information-outline
        </v-icon>
      </v-btn>
      <v-spacer />
      <v-switch
        v-model="$vuetify.theme.dark"
        hide-details
        inset
        color="black"
        :label="$vuetify.theme.dark ? '深色模式' : '浅色模式'"
      />
    </v-app-bar>

    <v-main>
      <ChartCard />
    </v-main>

    <v-footer app color="primary" class="pa-0">
      <v-btn
        tile class="ma-0"
        dark color="green darken-3"
        @click="openAllDialog"
      >
        <v-icon>
          mdi-file-outline
        </v-icon>
      </v-btn>
      <span class="white--text mx-3">{{ fileStatus }}</span>
    </v-footer>

    <VariableAllDialog ref="VariableAllDialog" />
    <AboutDialog ref="AboutDialog" />
  </v-app>
</template>

<script>
import VariableAllDialog from "@/components/VariableAllDialog.vue";
import AboutDialog from "@/components/AboutDialog/AboutDialog.vue";
import DrawerList from "@/components/DrawerList.vue";
import ChartCard from "@/components/ChartCard.vue";

export default {
  components: {
    VariableAllDialog,
    AboutDialog,
    DrawerList,
    ChartCard,
  },
  computed: {
    fileStatus () {
      return this.$store.getters["file/fileStatus"]
    }, 
    NewVersion() {
      return this.$store.getters["version/NewVersion"];
    }
  },
  mounted() {
    this.$vuetify.theme.dark = window.matchMedia("(prefers-color-scheme: dark)").matches;
    this.$store.dispatch("variables/getVType");
    this.$store.dispatch("version/Init");
    this.$store.dispatch("option/get");
  },
  methods: {
    switchDrawer() {
      this.$refs.DrawerList.switchDrawer();
    },
    openAboutDialog() {
      this.$refs.AboutDialog.openDialog();
    },
    openAllDialog() {
      this.$refs.VariableAllDialog.openDialog();
    },
  },
};
</script>
