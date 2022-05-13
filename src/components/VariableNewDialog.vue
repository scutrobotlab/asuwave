<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" persistent max-width="600px">
      <v-card>
        <v-toolbar dense color="primary">
          <v-toolbar-title>添加变量</v-toolbar-title>
          <v-spacer />
          <v-btn icon dark @click="dialog = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-toolbar>
        <v-card-text>
          <v-container>
            <ErrorAlert v-model="error" />
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-row>
                <v-col cols="12" sm="6" md="6">
                  <v-color-picker
                    v-model="Inputcolor"
                    dot-size="25"
                    swatches-max-height="100"
                    show-swatches
                  />
                </v-col>
                <v-col cols="12" sm="6" md="6">
                  <v-select
                    v-model="Board"
                    :items="[1, 2, 3]"
                    label="板子代号"
                    hint="保持默认为1即可"
                    :rules="[(v) => !!v || '板子代号是必要的']"
                    required
                  />
                  <v-select
                    v-model="Type"
                    :items="types"
                    label="变量类型"
                    :rules="[(v) => !!v || '变量类型是必要的']"
                    required
                    :disabled="disable"
                  />
                  <v-text-field
                    v-model="Name"
                    label="变量名称"
                    type="text"
                    required
                    :rules="NameRules"
                  />
                  <v-text-field
                    v-model="Addr"
                    label="变量地址"
                    type="text"
                    :rules="AddrRules"
                    :disabled="disable"
                    hint="[20000000, 7fffffff]区间的16进制数。"
                    required
                  />
                  <v-row dense>
                    <v-col cols="12" sm="6" dense>
                      <v-text-field
                        v-model="SignalGain"
                        label="信号增益（Gain）"
                        type="number"
                        :hint="signalHint"
                      />
                    </v-col>
                    <v-col cols="12" sm="6" dense>
                      <v-text-field
                        v-model="SignalBias"
                        label="信号偏置（Bias）"
                        type="number"
                        :hint="signalHint"
                      />
                    </v-col>
                  </v-row>
                  <v-text-field
                    v-model="Inputcolor"
                    label="变量颜色"
                    type="text"
                    disabled
                    required
                  />
                </v-col>
              </v-row>
            </v-form>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="error" @click="reset()">
            清空
          </v-btn>
          <v-btn color="primary" :disabled="!valid" @click="addVariable()">
            添加
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import errorMixin from "@/mixins/errorMixin.js";
import { postVariable } from "@/api/variable.js";

export default {
  mixins: [errorMixin],
  props: {
    mod: {
      type: String,
      default: "",
    },
  },
  data: () => ({
    dialog: false,
    disable: false,
    colorcard: false,
    valid: true,
    Board: 1,
    Name: "",
    Type: "",
    Addr: "",
    SignalGain: 1,
    SignalBias: 0,
    Inputcolor: "#F3CB09FF",
    showcolorcard: "",
    AddrRules: [
      (v) => !!v || "变量地址是必要的",
      (v) => /^(0x)?[2-7][0-9a-fA-F]{7}$/.test(v) || "请输入[2000 0000, 7fff ffff]区间的16进制数。",
    ],
    NameRules: [
      (v) => !!v || "变量名称是必要的",
      (v) => /^[_a-zA-z]/.test(v) || "变量名应以字母或下划线开始。",
    ],
  }),
  computed: {
    types() {
      return this.$store.state.variables.vTypes;
    },
    signalHint() {
      let hint = "";
      if (this.SignalGain != 1) {
        hint += this.SignalGain + " * ";
      }
      if (this.Name != "") {
        hint += this.Name;
      } else {
        hint += "var";
      }
      if (this.SignalBias != 0) {
        hint += " + " + this.SignalBias;
      }
      return hint;
    },
  },
  methods: {
    reset() {
      this.Board = 1;
      this.Name = "";
      this.Type = "";
      this.Addr = "";
      this.SignalGain = 1;
      this.SignalBias = 0;
      this.Inputcolor = this.randomColor();
    },
    openDialog() {
      this.dialog = true;
      this.Inputcolor = this.randomColor();
    },
    openDialogFromList(name, type, board, addr) {
      this.dialog = true;
      this.disable = true;
      this.Board = board;
      this.Name = name;
      this.Type = type;
      this.Addr = addr;
      this.SignalGain = 1;
      this.SignalBias = 0;
      this.Inputcolor = this.randomColor();
    },
    addVariable() {
      if (this.$refs.form.validate()) {
        this.errorHandler(
          postVariable(
            this.mod,
            this.Board,
            this.Name,
            this.Type,
            parseInt(this.Addr, 16),
            this.Inputcolor,
            parseFloat(this.SignalGain),
            parseFloat(this.SignalBias)
          )
        ).then(async () => {
          this.dialog = false;
          await this.$store.dispatch("variables/getV", this.mod);
          this.$bus.$emit("sendalert", true);
          console.log(this.mod);
        });
      }
    },
    randomColor() {
      let color = "#";
      for (let i = 0; i < 6; i++) {
        let seed = "0123456789abcdef";
        if (Math.random() < 0.5) {
          seed = this.$vuetify.theme.dark ? "89abcdef" : "01234567";
        }
        color += seed[Math.floor(Math.random() * seed.length)];
      }
      color += "FF";
      return color;
    },
  },
};
</script>
