<template>
  <v-card style="overflow-y: hidden">
    <div :class="themeClasses" ref="chart" style="width: 100%; height: 77vh"></div>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn text @click="exportData">导出</v-btn>
      <v-btn text color="primaryText" @click="follow">{{ showFollow }}</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import timechart from "timechart";
import Themeable from "vuetify/lib/mixins/themeable";
import { getVariable } from "@/api/variable.js";

export default {
  name: "ChartCard",
  mixins: [Themeable],
  data: () => ({
    chart: null,
    ws: null,
    indexColor: -2,
    inputColor: "",
    showFollow: "跟随",
    putColors: [],
  }),
  created() {
    this.initWS();
  },
  destroyed() {
    this.ws.close();
  },
  async mounted() {
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
    this.$bus.$on("sendcolor", (data) => {
      this.inputColor = data;
    });
    getVariable("read").then((variable) => {
      this.putColors = variable;
      this.indexColor = -1;
    });
  },
  watch: {
    isDark: function() {
      // for (const s of this.chart.options.series) {
      //   this.updateColor(s);
      // }
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
      if (this.inputColor != "") {
        series.color = this.inputColor;
      } else {
        series.color = this.putColors[index].Inputcolor;
      }
    },
    praseWS(data) {
      if (!data) {
        return;
      }
      const jsonWS = JSON.parse(data);
      const seriesArray = this.chart.options.series;
      if (this.indexColor != -2) {
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
      }
      this.chart.update();
    },
    follow() {
      switch (this.showFollow) {
        case "跟随":
          this.chart.options.realTime = true;
          this.showFollow = "取消跟随";
          break;
        case "取消跟随":
          this.chart.options.realTime = false;
          this.showFollow = "跟随";
          break;
      }
    },
    exportData() {
      console.log(this.chart.options.series);
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
