<template>
  <v-dialog v-model="dialog" max-width="400">
    <v-card>
      <v-list-item>
        <v-list-item-avatar tile>
          <img src="@/assets/logo.png" alt="Logo" />
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-title>坠好用的上位机</v-list-item-title>
          <v-list-item-subtitle>{{ current.version.tag_name }}</v-list-item-subtitle>
        </v-list-item-content>
        <v-list-item-icon>
          <v-btn icon v-on:click="dialog = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-list-item-icon>
      </v-list-item>
      <div>
        <v-card-text v-if="update.checking">
          <v-progress-circular indeterminate :size="20" class="mr-2" color="primary" />
          正在检查更新...
        </v-card-text>
        <v-card-text v-else-if="current.version.tag_name === update.version.tag_name">
          <v-icon color="success" class="mr-2">mdi-check-circle</v-icon>已是最新版本。
        </v-card-text>
        <div v-else>
          <v-card-title>
            {{ update.version.tag_name
            }}<v-chip color="warning" outlined small class="mx-2">New</v-chip>
          </v-card-title>
          <v-card-text class="text--secondary">
            更新日志：
            <div>{{ update.version.body }}</div>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn small text color="success" @click="doUpdate"
              ><v-icon left>mdi-arrow-up-circle</v-icon>立即更新</v-btn
            >
          </v-card-actions>
        </div>
      </div>
      <v-divider></v-divider>
      <v-card-actions class="caption text--secondary">
        <span>&copy; {{ new Date().getFullYear() }} 华工机器人实验室</span>
        <v-spacer></v-spacer>
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <v-btn icon href="https://github.com/scutrobotlab/asuwave" v-bind="attrs" v-on="on">
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
import { getVersion, postUpdate } from "@/api/version.js";
export default {
  mixins: [errorMixin],
  data: () => ({
    dialog: false,
    hc: "",
    current: {
      version: {
        tag_name: "v 1.6.9"
      }
    },
    update: {
      error: false,
      checking: true,
      version: {
        tag_name: "",
        body: "",
      },
    },
  }),
  methods: {
    openDialog() {
      this.dialog = true;
    },
    async checkUpdate() {
      this.update.checking = true;
      this.update.version = await this.errorHandler(getVersion());
      this.update.checking = false;
    },
    doUpdate() {
      this.errorHandler(postUpdate());
    },
  },
  mounted() {
    console.log("version", process.env.VUE_APP_GITTAG)
    this.current.version = {
      tag_name: process.env.VUE_APP_GITTAG,
    }
    this.checkUpdate();
  },
};
</script>
