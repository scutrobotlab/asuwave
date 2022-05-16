<template>
  <v-card style="overflow-y: hidden">
    <div
      ref="chart"
      style="width: 100%; height: calc(100vh - 56px - 36px - 52px)"
    />
    <v-card-actions>
      <v-file-input
        v-model="file" dense 
        label="导入json文件" accept=".json"
      />
      <v-spacer />
      <v-btn text @click="exportData">
        导出
      </v-btn>
      <v-btn text color="primaryText" @click="toggleFollow">
        {{ follow?"取消跟随":"跟随" }}
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
    file: null,
    chart: null,
    ws: null,
    follow: true,
    putColors: [],
  }),
  computed: {
    variables() {
      return this.$store.state.variables.read;
    },
  },
  watch: {
    isDark: function () {
      this.chart.update();
    },
    file: function(v) {
      window.console.log(v)
      var reader = new FileReader();
      reader.onload = (event)=>{
        console.log(event.target.result);
        this.chart.options.series = JSON.parse(event.target.result).series;
        this.chart.update();
      };
      reader.readAsText(v);
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
          "/dataws";
      } else {
        url = "ws://localhost:8000/dataws";
      }
      this.ws = new WebSocket(url);
      this.ws.onopen = ()=>{console.log("连接成功")};
      this.ws.onclose = ()=>{console.log("连接断开")};
      this.ws.onmessage = (evt) => {this.parseWS(evt.data)};
      this.ws.onerror = (evt)=>{console.log("ERROR: " + evt.data)}
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
    toggleFollow() {
      this.follow = !this.follow;
      this.chart.options.realTime = this.follow;
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
