<template>
  <v-card style="overflow-y: hidden">
    <div
      ref="chart"
      style="width: 100%; height: calc(100vh - 56px - 36px - 52px)"
    />
    <v-card-actions>
      <v-spacer />
      <v-btn text @click="exportData">
        导出
      </v-btn>
      <v-btn text color="primaryText" @click="follow">
        {{ showFollow }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import timechart from "timechart";
import { lineChart } from 'timechart/plugins/lineChart';
import { d3Axis } from 'timechart/plugins/d3Axis';
import { legend } from 'timechart/plugins/legend';
import { crosshair } from 'timechart/plugins/crosshair';
import { nearestPoint } from 'timechart/plugins/nearestPoint';

export default {
  name: "ChartCard",
  data: () => ({
    chart: null,
    ws: null,
    showFollow: "跟随",
    putColors: [],
  }),
  computed: {
    variables() {
      return this.$store.state.variables.variables.read;
    },
  },
  watch: {
    isDark: function () {
      this.chart.update();
    },
  },
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
      tooltip: true,
      zoom: {
        x: {
          autoRange: true,
          minDomainExtent: 50,
        },
        y: {
          autoRange: true,
        },
      },
      plugins: {
        lineChart,
        d3Axis,
        legend,
        crosshair,
        nearestPoint,
      },
    });
    this.$nextTick(() => {
      this.chart.onResize();
    });
    this.$store.dispatch("variables/getV", "read");
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
      this.parseWS(evt.data);
    },
    WSonerror(evt) {
      console.log("ERROR: " + evt.data);
    },
    parseWS(data) {
      if (!data) {
        return;
      }
      const jsonWS = JSON.parse(data);
      let seriesArray = this.chart.options.series;
      seriesArray.forEach((s) => {
        if (Object.entries(this.variables).find(([, v]) => s.name == v.Name) === undefined) {
          s.visible = false;
        }
      });
      for (const [, variable] of Object.entries(this.variables)) {
        let chart_var = jsonWS.find((c) => c.Name == variable.Name);
        let series = seriesArray.find((s) => s.name == variable.Name);
        if (!series) {
          series = {
            name: variable.Name,
            color: variable.Inputcolor,
            data: [],
          };
          seriesArray.push(series);
        }
        if (!series.visible) {
          series.visible = true;
        }
        if (chart_var) {
          series.data.push({
            x: chart_var.Tick,
            y: chart_var.Data,
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
