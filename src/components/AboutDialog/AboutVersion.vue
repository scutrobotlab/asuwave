<template>
  <div>
    <v-card-text v-if="update.checking">
      <v-progress-circular
        indeterminate :size="20" class="mr-2"
        color="primary"
      />
      正在检查更新...
    </v-card-text>
    <v-card-text v-else-if="!NewVersion">
      <v-icon color="success" left>
        mdi-check-circle
      </v-icon>已是最新版本。
      <a class="text--secondary" @click="checkUpdate">重新检查</a>
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
        <div>{{ update.response.body }}</div>
        <div class="text-center text--disabled my-3">
          <v-btn
            v-if="asset!=null"
            rounded
            color="success"
            :href="asset.browser_download_url"
          >
            <v-icon left>
              mdi-download
            </v-icon>下载
          </v-btn>
          <div v-if="asset" class="my-1">
            {{ asset.name }} ({{ ByteUnitConvert(asset.size) }})
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
</template>

<script>
export default {
  computed: {
    asset() {
      return this.$store.getters["version/asset"];
    },
    update() {
      return this.$store.state.version.update;
    },
    current_tag() {
      return this.$store.state.version.current_tag;
    },
    NewVersion() {
      return this.$store.getters["version/NewVersion"];
    }
  },
  methods: {
    checkUpdate() {
      this.$store.dispatch("version/CheckUpdate");
    },
    ByteUnitConvert(val) {
      return (
        Math.floor(Math.log2(val) / 10) +
        ["B", "KB", "MB", "GB", "TB"][Math.floor(Math.log2(val) / 10)]
      );
    },
  },
}
</script>

<style>

</style>