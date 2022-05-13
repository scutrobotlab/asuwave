<template>
  <v-dialog v-model="dialog" max-width="50%">
    <v-card>
      <v-list-item>
        <v-list-item-avatar tile>
          <img src="@/assets/logo.png" alt="Logo">
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-title>坠好用的上位机</v-list-item-title>
          <v-list-item-subtitle>{{ current_tag }}</v-list-item-subtitle>
        </v-list-item-content>
        <v-list-item-icon>
          <v-btn icon @click="dialog = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-list-item-icon>
      </v-list-item>
      <AboutVersion />
      <v-divider />
      <AboutOption />
      <v-card-actions class="caption text--secondary">
        <span>&copy; {{ new Date().getFullYear() }} 华工机器人实验室</span>
        <v-spacer />
        <v-tooltip bottom>
          <template #activator="{ on, attrs }">
            <v-btn
              icon
              href="https://github.com/scutrobotlab/asuwave"
              target="_blank"
              v-bind="attrs"
              v-on="on"
            >
              <v-icon>mdi-github</v-icon>
            </v-btn>
          </template>
          <span>GitHub反馈</span>
        </v-tooltip>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import AboutOption from './AboutOption.vue';
import AboutVersion from './AboutVersion.vue';
export default {
  components: { AboutOption, AboutVersion },
  mixins: [errorMixin],
  data: () => ({
    dialog: false,
  }),
  computed: {
    current_tag() {
      return this.$store.state.version.current_tag;
    },
  },
  methods: {
    openDialog() {
      this.dialog = true;
    },
  },
};
</script>
