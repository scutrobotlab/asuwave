<template>
  <v-dialog v-model="dialog" max-width="400">
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
      <div>
        <v-card-text v-if="update.checking">
          <v-progress-circular
            indeterminate :size="20" class="mr-2"
            color="primary"
          />
          正在检查更新...
        </v-card-text>
        <v-card-text v-else-if="current_tag === update.response.tag_name">
          <v-icon color="success" left>
            mdi-check-circle
          </v-icon>已是最新版本。
          <a class="text--secondary " @click="checkUpdate">重新检查</a>
        </v-card-text>
        <div v-else>
          <v-card-title>
            {{ update.response.tag_name }}
            <v-chip
              color="warning" outlined small
              class="mx-2"
            >
              New
            </v-chip>
          </v-card-title>
          <v-card-text class="text--secondary">
            更新日志：
            <div>{{ update.response.body }}</div>
            <div class="text-center text--disabled my-3">
              <v-btn
                v-if="download_link"
                rounded
                color="success"
                :href="download_link.browser_download_url"
              >
                <v-icon left>
                  mdi-download
                </v-icon>下载
              </v-btn>
              <div v-if="download_link" class="my-1">
                {{ download_link.name }} ({{ ByteUnitConvert(download_link.size) }})
              </div>
              <div class="my-1">
                <a target="_blank" href="https://github.com/scutrobotlab/asuwave/releases">
                  查看所有版本
                </a>
              </div>
            </div>
          </v-card-text>
        </div>
      </div>
      <v-divider />
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
import { getVersion } from "@/api/version.js";
export default {
  mixins: [errorMixin],
  data: () => ({
    dialog: false,
    os: "",
    arch: "",
    current_tag: "",
    update: {
      error: false,
      checking: true,
      response: {
        tag_name: "",
        body: "",
        assets: [],
      },
    },
  }),
  computed: {
    download_link() {
      return this.update.response.assets.find((asset) => {
        return (
          asset.browser_download_url.includes(this.os) &&
          asset.browser_download_url.includes(this.arch)
        );
      });
    },
  },
  mounted() {
    this.current_tag = process.env.VUE_APP_GITTAG;
    window.console.log(this.current_tag)

    if (/Win|win/i.test(navigator.userAgent)) this.os = "windows";
    else if (/Mac|mac|darwin/i.test(navigator.userAgent)) this.os = "darwin";
    else if (/linux|gnu/i.test(navigator.userAgent)) this.os = "linux";

    if (/(?:(amd|x(?:(?:86|64)[-_])?|wow|win)64)[;)]/i.test(navigator.userAgent))
      this.arch = "amd64";
    else if (/\b(aarch64|arm(v?8e?l?|_?64))\b/i.test(navigator.userAgent)) this.arch = "arm64";

    window.console.log(navigator.userAgent);
    this.checkUpdate();
  },
  methods: {
    openDialog() {
      this.dialog = true;
    },
    ByteUnitConvert(val) {
      return (
        Math.floor(Math.log2(val) / 10) +
        ["B", "KB", "MB", "GB", "TB"][Math.floor(Math.log2(val) / 10)]
      );
    },
    async checkUpdate() {
      this.update.checking = true;
      this.update.response = await this.errorHandler(getVersion());
      this.update.checking = false;
    },
  },
};
</script>
