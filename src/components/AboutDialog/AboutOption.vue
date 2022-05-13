<template>
  <div>
    <v-card-title>设置</v-card-title>
    <v-card-text>
      <v-slider
        v-model="LogLevel" label="日志等级"
        ticks="always" min="0" max="5"
        dense
      />
      <v-checkbox 
        v-model="SaveVarList" dense label="保存变量列表"
      />
      <v-checkbox
        v-model="SaveFilePath" dense label="保存监控文件路径"
      />
      <v-checkbox 
        v-model="UpdateByProj" dense label="监控文件更新时，根据变量名更新列表"
      />
    </v-card-text>
  </div>
</template>

<script>
export default {
  computed: {
    LogLevel: {
      get(){
        return this.$store.state.option.LogLevel;
      },
      set(v){
        this.$store.dispatch("option/set", {k: "LogLevel", v})
      }
    },
    SaveVarList: {
      get(){
        return this.$store.state.option.LogLevel;
      },
      set(v){
        this.$store.dispatch("option/set", {k: "SaveVarList", v})
      }
    },
    SaveFilePath: {
      get(){
        return this.$store.state.option.LogLevel;
      },
      set(v){
        this.$store.dispatch("option/set", {k: "SaveFilePath", v})
      }
    },
    UpdateByProj: {
      get(){
        return this.$store.state.option.LogLevel;
      },
      set(v){
        this.$store.dispatch("option/set", {k: "UpdateByProj", v})
      }
    },
  }
}
</script>

<style>

</style>