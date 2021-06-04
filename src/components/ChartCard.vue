<template>
  <v-card style="overflow-y: hidden">
    <div :class="themeClasses" ref="chart" style="width: 100%; height: 77vh"></div>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn text @click="exportData">导出</v-btn>
      <v-btn text color="primaryText" @click="follow">跟随</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import timechart from "timechart";
import colors from "vuetify/lib/util/colors";
import Themeable from "vuetify/lib/mixins/themeable";

const lineColors = {
  light: [
    colors.red.base,
    colors.green.base,
    colors.orange.base,
    colors.purple.base,
    colors.indigo.base,
    colors.teal.base,
    colors.pink.base,
  ],
  dark: [
    colors.red.lighten2,
    colors.green.lighten2,
    colors.orange.lighten2,
    colors.purple.lighten2,
    colors.indigo.lighten2,
    colors.teal.lighten2,
    colors.pink.lighten2,
  ],
};

export default {
  name: "ChartCard",
  mixins: [Themeable],
  data: () => ({
    chart: null,
    ws: null,
    indexColor: -1,
  }),
  created() {
    this.initWS();
  },
  destroyed() {
    this.ws.close();
  },
  mounted() {
    this.chart = new timechart(this.$refs.chart, {
      baseTime: Date.now(),
      series: [],
      xRange: { min: 0, max: 20 * 1000 },
      realTime: true,
      zoom: {
        x: {
          autoRange: true,
          minDomainExtent: 50,
        },
        y: {
          autoRange: true,
        },
      },
    });
    this.$nextTick(() => {
      this.chart.onResize();
    });
  },
  watch: {
    isDark: function() {
      for (const s of this.chart.options.series) {
        this.updateColor(s);
      }
      this.chart.update();
    },
  },
  methods: {
    initWS() {
      let url = "";
      if (process.env.NODE_ENV === "production") {
        url =
          (document.location.protocol == "https:" ? "wss" : "ws") +
          "://" +
          window.location.host +
          "/ws";
      } else {
        url = "ws://localhost:8000/ws";
      }
      this.ws = new WebSocket(url);
      this.ws.onopen = this.WSonopen;
      this.ws.onclose = this.WSclose;
      this.ws.onmessage = this.WSonmessage;
      this.ws.onerror = this.WSonerror;
    },
    WSonopen() {
      console.log("连接成功");
    },
    WSclose() {
      console.log("连接断开");
    },
    WSonmessage(evt) {
      this.praseWS(evt.data);
    },
    WSonerror(evt) {
      console.log("ERROR: " + evt.data);
    },
    updateColor(series) {
      const index = series.colorIndex;
      const colorArray = this.isDark ? lineColors.dark : lineColors.light;
      series.color = colorArray[index % colorArray.length];
    },
    praseWS(data) {
      if (!data) {
        return;
      }

      const jsonWS = JSON.parse(data);
      const seriesArray = this.chart.options.series;
      for (const dp of jsonWS.Variables) {
        let series = seriesArray.find((a) => a.name == dp.Name);
        if (!series) {
          this.indexColor++;
          series = {
            name: dp.Name,
            colorIndex: this.indexColor,
            data: [],
          };
          this.updateColor(series);
          seriesArray.push(series);
        }
        series.data.push({
          x: dp.Tick,
          y: dp.Data,
        });
      }
      this.chart.update();
    },
    follow() {
      this.chart.options.realTime = true;
    },
    exportData() {
      const json = JSON.stringify(this.chart.options);
      const blob = new Blob([json], { type: "application/json" });
      const anchor = document.createElement("a");
      anchor.href = URL.createObjectURL(blob);
      anchor.download = "上位机导出数据.json";
      anchor.style.display = "none";
      document.body.appendChild(anchor);
      anchor.click();
      document.body.removeChild(anchor);
      URL.revokeObjectURL(anchor.href);
    },
  },
};
</script>

<style scoped>
.theme--light {
  --background-overlay: white;
}
.theme--dark {
  --background-overlay: black;
}
</style>
